package client

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmic/api"
	"github.com/openconfig/gnmic/target"
)

const (
	DefaultMaxRetries         int     = 2
	DefaultBackoffMinDelay    int     = 4
	DefaultBackoffMaxDelay    int     = 60
	DefaultBackoffDelayFactor float64 = 3
)

type SetOperation string

const (
	Update  SetOperation = "update"
	Replace SetOperation = "replace"
	Delete  SetOperation = "delete"
)

type Client struct {
	Devices map[string]*Device
	// Maximum number of retries
	MaxRetries int
	// Minimum delay between two retries
	BackoffMinDelay int
	// Maximum delay between two retries
	BackoffMaxDelay int
	// Backoff delay factor
	BackoffDelayFactor float64
}

type Device struct {
	SetMutex *sync.Mutex
	Target   *target.Target
}

func NewClient() Client {
	devices := make(map[string]*Device)
	return Client{
		Devices:            devices,
		MaxRetries:         DefaultMaxRetries,
		BackoffMinDelay:    DefaultBackoffMinDelay,
		BackoffMaxDelay:    DefaultBackoffMaxDelay,
		BackoffDelayFactor: DefaultBackoffDelayFactor,
	}
}

func (c *Client) AddTarget(ctx context.Context, device, host, username, password, certificate, key, caCertificate string, verifyCertificate, Tls bool) diag.Diagnostics {
	var diags diag.Diagnostics

	if !strings.Contains(host, ":") {
		host = host + ":57400"
	}

	t, err := api.NewTarget(
		api.Name(device),
		api.Address(host),
		api.Username(username),
		api.Password(password),
		api.TLSCert(certificate),
		api.TLSKey(key),
		api.TLSCA(caCertificate),
		api.SkipVerify(!verifyCertificate),
		api.Insecure(!Tls),
	)
	if err != nil {
		diags.AddError(
			"Unable to create target",
			"Unable to create target:\n\n"+err.Error(),
		)
		return diags
	}
	err = t.CreateGNMIClient(ctx)
	if err != nil {
		diags.AddError(
			"Unable to create gNMI client",
			"Unable to create gNMI client:\n\n"+err.Error(),
		)
		return diags
	}

	if device == "" {
		c.Devices[""] = &Device{}
		c.Devices[""].Target = t
		c.Devices[""].SetMutex = &sync.Mutex{}
	} else {
		c.Devices[device] = &Device{}
		c.Devices[device].Target = t
		c.Devices[device].SetMutex = &sync.Mutex{}
	}

	return nil
}

func (c *Client) Set(ctx context.Context, device, path, body string, operation SetOperation) (*gnmi.SetResponse, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := c.Devices[device].Target

	var setReq *gnmi.SetRequest
	var err error
	if operation == Update {
		setReq, err = api.NewSetRequest(
			api.Update(
				api.Path(path),
				api.Value(body, "json_ietf"),
			),
		)
	} else if operation == Replace {
		setReq, err = api.NewSetRequest(
			api.Replace(
				api.Path(path),
				api.Value(body, "json_ietf"),
			),
		)
	} else if operation == Delete {
		setReq, err = api.NewSetRequest(
			api.Delete(
				path,
			),
		)
	}
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to create set request, got error: %s", err))
		return nil, diags
	}

	tflog.Debug(ctx, fmt.Sprintf("gNMI set request: %s", setReq.String()))

	var setResp *gnmi.SetResponse
	for attempts := 0; ; attempts++ {
		c.Devices[device].SetMutex.Lock()
		setResp, err = target.Set(ctx, setReq)
		c.Devices[device].SetMutex.Unlock()
		if err != nil {
			if ok := c.Backoff(ctx, attempts); !ok {
				diags.AddError("Client Error", fmt.Sprintf("Set request failed, got error: %s", err))
				return nil, diags
			} else {
				tflog.Error(ctx, fmt.Sprintf("gNMI set request failed: %s, retries: %v", err, attempts))
				continue
			}
		}
		break
	}

	tflog.Debug(ctx, fmt.Sprintf("gNMI set response: %s", setResp.String()))

	return setResp, nil
}

func (c *Client) Get(ctx context.Context, device, path string) (*gnmi.GetResponse, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := c.Devices[device].Target

	getReq, err := api.NewGetRequest(api.Path(path), api.Encoding("json_ietf"))
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to create get request, got error: %s", err))
		return nil, diags
	}

	tflog.Debug(ctx, fmt.Sprintf("gNMI get request: %s", getReq.String()))

	var getResp *gnmi.GetResponse
	for attempts := 0; ; attempts++ {
		getResp, err = target.Get(ctx, getReq)
		if err != nil {
			if ok := c.Backoff(ctx, attempts); !ok {
				diags.AddError("Client Error", fmt.Sprintf("Get request failed, got error: %s", err))
				return nil, diags
			} else {
				tflog.Error(ctx, fmt.Sprintf("gNMI get request failed: %s, retries: %v", err, attempts))
				continue
			}
		}
		break
	}
	tflog.Debug(ctx, fmt.Sprintf("gNMI get response: %s", getResp.String()))

	return getResp, nil
}

// Backoff waits following an exponential backoff algorithm
func (c *Client) Backoff(ctx context.Context, attempts int) bool {
	tflog.Debug(ctx, fmt.Sprintf("Begining backoff method: attempts %v on %v", attempts, c.MaxRetries))
	if attempts >= c.MaxRetries {
		tflog.Debug(ctx, "Exit from backoff method with return value false")
		return false
	}

	minDelay := time.Duration(c.BackoffMinDelay) * time.Second
	maxDelay := time.Duration(c.BackoffMaxDelay) * time.Second

	min := float64(minDelay)
	backoff := min * math.Pow(c.BackoffDelayFactor, float64(attempts))
	if backoff > float64(maxDelay) {
		backoff = float64(maxDelay)
	}
	backoff = (rand.Float64()/2+0.5)*(backoff-min) + min
	backoffDuration := time.Duration(backoff)
	tflog.Debug(ctx, fmt.Sprintf("Starting sleeping for %v", backoffDuration.Round(time.Second)))
	time.Sleep(backoffDuration)
	tflog.Debug(ctx, "Exit from backoff method with return value true")
	return true
}

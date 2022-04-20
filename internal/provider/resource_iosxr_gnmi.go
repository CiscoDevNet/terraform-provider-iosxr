package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/tidwall/gjson"
)

type resourceGnmiType struct{}

func (t resourceGnmiType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Manages IOS-XR objects via gNMI calls. This resource can only manage a single object. It is able to read the state and therefore reconcile configuration drift.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"path": {
				MarkdownDescription: "A gNMI path, e.g. `openconfig-interfaces:/interfaces/interface`.",
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"attributes": {
				Type:                types.MapType{ElemType: types.StringType},
				MarkdownDescription: "Map of key-value pairs which represents the attributes and its values.",
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (t resourceGnmiType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceGnmi{
		provider: provider,
	}, diags
}

type resourceGnmi struct {
	provider provider
}

func (r resourceGnmi) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan Gnmi

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.Value))

	if !plan.Attributes.Unknown {
		body := plan.toBody(ctx)

		_, diags = r.provider.client.Set(ctx, plan.Device.Value, plan.Path.Value, body, client.Update)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	plan.Id = plan.Path
	plan.Attributes.Unknown = false

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourceGnmi) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state Gnmi

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	getResp, diags := r.provider.client.Get(ctx, state.Device.Value, state.Path.Value)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ea map[string]string
	state.Attributes.ElementsAs(ctx, &ea, false)
	existingAttr := make([]string, len(ea))
	for k := range ea {
		existingAttr = append(existingAttr, k)
	}

	attributes := make(map[string]attr.Value)

	for attr, value := range gjson.ParseBytes(getResp.Notification[0].Update[0].Val.GetJsonIetfVal()).Map() {
		if helpers.Contains(existingAttr, attr) {
			attributes[attr] = types.String{Value: value.String()}
		}
		// if len(existingAttr) > 0 && !helpers.Contains(existingAttr, attr) {
		// 	continue
		// }
		// attributes[attr] = types.String{Value: value.String()}
	}
	state.Attributes.Elems = attributes

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceGnmi) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan Gnmi

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.Value))

	if !plan.Attributes.Unknown {
		body := plan.toBody(ctx)

		_, diags = r.provider.client.Set(ctx, plan.Device.Value, plan.Path.Value, body, client.Update)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r resourceGnmi) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state Gnmi

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.Value))

	_, diags = r.provider.client.Set(ctx, state.Device.Value, state.Path.Value, "", client.Delete)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.Value))

	resp.State.RemoveResource(ctx)
}

func (r resourceGnmi) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Import", req.ID))

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("path"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req.ID)...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Import finished successfully", req.ID))
}

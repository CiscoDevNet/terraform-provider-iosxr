package helpers

import (
	"context"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	stringPatternValidatorErr = "Minimum length: %v. Maximum length: %v. Regexes: [%s]."
)

type stringPatternValidator struct {
	regexes              []string
	minLength, maxLength int
}

func (v stringPatternValidator) Description(context.Context) string {
	return fmt.Sprintf(stringPatternValidatorErr, v.minLength, v.maxLength, strings.Join(v.regexes, ", "))
}

func (v stringPatternValidator) MarkdownDescription(context.Context) string {
	return fmt.Sprintf(stringPatternValidatorErr, v.minLength, v.maxLength, strings.Join(v.regexes, ", "))
}

func (v stringPatternValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var str types.String
	{
		diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &str)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() {
			return
		}
	}

	if str.IsUnknown() || str.IsNull() {
		return
	}

	maxLength := v.maxLength
	if maxLength == 0 {
		maxLength = math.MaxInt
	}

	if len(str.ValueString()) >= v.minLength && len(str.ValueString()) <= maxLength {
		match := true
		for _, regex := range v.regexes {
			if ok, _ := regexp.MatchString(regex, str.ValueString()); !ok {
				match = false
			}
		}
		if match {
			return
		}
	}

	resp.Diagnostics.AddAttributeError(
		req.AttributePath,
		"Invalid String",
		fmt.Sprintf(stringPatternValidatorErr, v.minLength, v.maxLength, strings.Join(v.regexes, ", ")),
	)
}

func StringPatternValidator(minLength, maxLength int, regexes ...string) tfsdk.AttributeValidator {
	return stringPatternValidator{
		minLength: minLength,
		maxLength: maxLength,
		regexes:   regexes,
	}
}

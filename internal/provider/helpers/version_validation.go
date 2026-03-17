// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package helpers

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VersionAtLeast checks if the current provider version is at least the required version.
// Both versions should be in internal format (e.g., "2442", "2522").
func VersionAtLeast(currentVersion, requiredVersion string) bool {
	if currentVersion == "" || requiredVersion == "" {
		return true // If version is not set, allow all features
	}
	current, err1 := strconv.Atoi(currentVersion)
	required, err2 := strconv.Atoi(requiredVersion)
	if err1 != nil || err2 != nil {
		return true // If conversion fails, allow the feature
	}
	return current >= required
}

// FieldVersionConstraint represents version requirements for a field
type FieldVersionConstraint struct {
	FieldPath        string // e.g., "ca_trustpoints.method_est_credential_certificate"
	AddedInVersion   string // Version when field was introduced (empty = available from start)
	RemovedInVersion string // Version when field was removed (empty = still available)
}

// VersionRange represents min/max range for a specific version
type VersionRange struct {
	Min int64
	Max int64
}

// FieldRangeConstraint represents version-specific range constraints for an integer field
type FieldRangeConstraint struct {
	FieldPath     string                  // e.g., "endpoint_default_probe_tx_interval"
	VersionRanges map[string]VersionRange // version -> range (e.g., "2442" -> {Min: 30000, Max: 15000000})
}

// Validatable is an interface for models that support version validation
type Validatable interface {
	GetVersionConstraints() []FieldVersionConstraint
	GetRangeConstraints() []FieldRangeConstraint
}

// Validate performs all version-related validations on a model
// This includes: version constraints (added/removed fields) and range constraints
// Returns true if validation passes, false otherwise
func Validate(providerVersion string, model Validatable, diagnostics *diag.Diagnostics) bool {
	if providerVersion == "" {
		// If no version configured, skip validation (backwards compatibility)
		return true
	}

	// Get constraints from model
	versionConstraints := model.GetVersionConstraints()
	rangeConstraints := model.GetRangeConstraints()

	// Skip validation if no constraints defined
	if len(versionConstraints) == 0 && len(rangeConstraints) == 0 {
		return true
	}

	// Validate version constraints (field support)
	ValidateVersionConstraints(providerVersion, model, versionConstraints, diagnostics)
	if diagnostics.HasError() {
		return false
	}

	// Validate version-specific ranges for integer fields
	ValidateVersionRanges(providerVersion, model, rangeConstraints, diagnostics)
	if diagnostics.HasError() {
		return false
	}

	return true
}

// ValidateVersionConstraints checks if fields set in the plan are valid for the provider version
func ValidateVersionConstraints(
	providerVersion string,
	planValue interface{},
	constraints []FieldVersionConstraint,
	diagnostics *diag.Diagnostics,
) {
	if providerVersion == "" {
		// If no version configured, skip validation (backwards compatibility)
		return
	}

	// For each constraint, check if the field is set and if it's valid for this version
	for _, constraint := range constraints {
		// Handle resource-level constraint (empty FieldPath means entire resource)
		if constraint.FieldPath == "" {
			// Check if entire resource was removed
			if constraint.RemovedInVersion != "" {
				if VersionAtLeast(providerVersion, constraint.RemovedInVersion) {
					diagnostics.AddError(
						fmt.Sprintf("Resource Not Supported in IOS-XR Version %s", FormatVersion(providerVersion)),
						fmt.Sprintf(
							"This resource is not supported in IOS-XR version %s. "+
								"It was removed/deprecated in version %s.\n\n"+
								"Please use an earlier IOS-XR version that supports this resource, "+
								"or remove this resource from your configuration.",
							FormatVersion(providerVersion),
							FormatVersion(constraint.RemovedInVersion),
						),
					)
					return // Don't check individual fields if entire resource is invalid
				}
			}
			// Check if entire resource requires newer version
			if constraint.AddedInVersion != "" {
				if !VersionAtLeast(providerVersion, constraint.AddedInVersion) {
					diagnostics.AddError(
						fmt.Sprintf("Resource Not Supported in IOS-XR Version %s", FormatVersion(providerVersion)),
						fmt.Sprintf(
							"This resource is only available in IOS-XR version %s and later. "+
								"Your configured version is %s.\n\n"+
								"To use this resource, please upgrade to IOS-XR version %s or later, "+
								"or remove this resource from your configuration.",
							FormatVersion(constraint.AddedInVersion),
							FormatVersion(providerVersion),
							FormatVersion(constraint.AddedInVersion),
						),
					)
					return // Don't check individual fields if entire resource is invalid
				}
			}
			continue
		}

		// Check if field is set
		if !isFieldSet(planValue, constraint.FieldPath) {
			continue
		}

		// Check if field was added after the current version
		if constraint.AddedInVersion != "" {
			if !VersionAtLeast(providerVersion, constraint.AddedInVersion) {
				diagnostics.AddError(
					fmt.Sprintf("Field Not Supported in IOS-XR Version %s", FormatVersion(providerVersion)),
					fmt.Sprintf(
						"The field '%s' is only available in IOS-XR version %s and later. "+
							"Your configured version is %s.\n\n"+
							"To use this field, please upgrade to IOS-XR version %s or later, "+
							"or remove this field from your configuration.",
						constraint.FieldPath,
						FormatVersion(constraint.AddedInVersion),
						FormatVersion(providerVersion),
						FormatVersion(constraint.AddedInVersion),
					),
				)
				continue
			}
		}

		// Check if field was removed before or at the current version
		if constraint.RemovedInVersion != "" {
			if VersionAtLeast(providerVersion, constraint.RemovedInVersion) {
				// Provide helpful message about when the field was available
				availabilityMsg := "earlier versions"
				if constraint.AddedInVersion != "" {
					availabilityMsg = fmt.Sprintf("versions %s through %s",
						FormatVersion(constraint.AddedInVersion),
						formatVersionPrevious(constraint.RemovedInVersion))
				} else {
					availabilityMsg = fmt.Sprintf("versions earlier than %s", FormatVersion(constraint.RemovedInVersion))
				}

				diagnostics.AddError(
					fmt.Sprintf("Field Not Supported in IOS-XR Version %s", FormatVersion(providerVersion)),
					fmt.Sprintf(
						"The field '%s' is not supported in IOS-XR version %s. "+
							"This field was removed in version %s and is only available in %s.\n\n"+
							"Please remove this field from your configuration or use an earlier IOS-XR version.",
						constraint.FieldPath,
						FormatVersion(providerVersion),
						FormatVersion(constraint.RemovedInVersion),
						availabilityMsg,
					),
				)
			}
		}
	}
}

// isFieldSet checks if a field at the given path is set (not null/unknown) in the plan value
func isFieldSet(planValue interface{}, fieldPath string) bool {
	parts := strings.Split(fieldPath, ".")
	current := reflect.ValueOf(planValue)

	for i, part := range parts {
		// If it's not a struct, we can't navigate further
		if current.Kind() == reflect.Ptr {
			if current.IsNil() {
				return false
			}
			current = current.Elem()
		}

		if current.Kind() != reflect.Struct {
			return false
		}

		// Find the field
		field := current.FieldByName(toCamelCase(part))
		if !field.IsValid() {
			return false
		}

		// Handle Go slices (e.g., []CryptoCaTrustpoints)
		if field.Kind() == reflect.Slice {
			// If this is the last part, check if slice has elements
			if i == len(parts)-1 {
				return field.Len() > 0
			}

			// Check nested fields in slice items
			remainingPath := strings.Join(parts[i+1:], ".")
			for j := 0; j < field.Len(); j++ {
				sliceElem := field.Index(j)
				if isFieldSetInValue(sliceElem, remainingPath) {
					return true
				}
			}
			return false
		}

		// Check if it's a terraform types value (String, Int64, Bool, etc.)
		switch v := field.Interface().(type) {
		case types.String:
			if v.IsNull() || v.IsUnknown() {
				return false
			}
			// If this is the last part, field is set
			return true
		case types.Int64:
			if v.IsNull() || v.IsUnknown() {
				return false
			}
			return true
		case types.Bool:
			if v.IsNull() || v.IsUnknown() {
				return false
			}
			return true
		case types.List:
			if v.IsNull() || v.IsUnknown() {
				return false
			}
			// If this is the last part, just checking if list is set
			if i == len(parts)-1 {
				return true
			}
			// Need to check nested fields in list items
			// Get the list elements and check if any item has the nested field set
			var listElements []types.Object
			diags := v.ElementsAs(context.Background(), &listElements, false)
			if diags.HasError() {
				return false
			}

			// Check each list item for the remaining path
			remainingPath := strings.Join(parts[i+1:], ".")
			for _, elem := range listElements {
				if elem.IsNull() || elem.IsUnknown() {
					continue
				}
				// Convert types.Object to a struct by checking its attributes
				attrs := elem.Attributes()
				if checkObjectAttributes(attrs, remainingPath) {
					return true
				}
			}
			return false
		default:
			// For nested structs, continue navigating
			current = field
		}
	}

	return false
}

// checkObjectAttributes checks if a nested field path is set in a types.Object's attributes
func checkObjectAttributes(attrs map[string]attr.Value, fieldPath string) bool {
	parts := strings.Split(fieldPath, ".")
	if len(parts) == 0 {
		return false
	}

	// Convert to snake_case for attribute lookup
	attrName := parts[0]
	attrValue, ok := attrs[attrName]
	if !ok {
		return false
	}

	// Check if the attribute is set (not null/unknown)
	if attrValue == nil {
		return false
	}

	switch v := attrValue.(type) {
	case types.String:
		if v.IsNull() || v.IsUnknown() {
			return false
		}
		return true
	case types.Int64:
		if v.IsNull() || v.IsUnknown() {
			return false
		}
		return true
	case types.Bool:
		if v.IsNull() || v.IsUnknown() {
			return false
		}
		return true
	case types.List:
		if v.IsNull() || v.IsUnknown() {
			return false
		}
		// If there are more parts, need to recurse into list
		if len(parts) > 1 {
			var listElements []types.Object
			diags := v.ElementsAs(context.Background(), &listElements, false)
			if diags.HasError() {
				return false
			}
			remainingPath := strings.Join(parts[1:], ".")
			for _, elem := range listElements {
				if elem.IsNull() || elem.IsUnknown() {
					continue
				}
				if checkObjectAttributes(elem.Attributes(), remainingPath) {
					return true
				}
			}
		}
		return true
	case types.Object:
		if v.IsNull() || v.IsUnknown() {
			return false
		}
		// If there are more parts, recurse into the object
		if len(parts) > 1 {
			remainingPath := strings.Join(parts[1:], ".")
			return checkObjectAttributes(v.Attributes(), remainingPath)
		}
		return true
	default:
		return false
	}
}

// isFieldSetInValue checks if a field path is set in a reflect.Value (used for slice elements)
func isFieldSetInValue(value reflect.Value, fieldPath string) bool {
	parts := strings.Split(fieldPath, ".")
	current := value

	for _, part := range parts {
		// Dereference pointers
		if current.Kind() == reflect.Ptr {
			if current.IsNil() {
				return false
			}
			current = current.Elem()
		}

		if current.Kind() != reflect.Struct {
			return false
		}

		// Find the field
		field := current.FieldByName(toCamelCase(part))
		if !field.IsValid() {
			return false
		}

		// Check if it's a terraform types value
		switch v := field.Interface().(type) {
		case types.String:
			return !v.IsNull() && !v.IsUnknown()
		case types.Int64:
			return !v.IsNull() && !v.IsUnknown()
		case types.Bool:
			return !v.IsNull() && !v.IsUnknown()
		case types.List:
			return !v.IsNull() && !v.IsUnknown()
		case types.Set:
			return !v.IsNull() && !v.IsUnknown()
		case types.Object:
			return !v.IsNull() && !v.IsUnknown()
		default:
			// Continue navigating for nested structs
			current = field
		}
	}

	return false
}

// toCamelCase converts snake_case to CamelCase for field name lookup
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][0:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// FormatVersion converts internal version format to display format
// e.g., "2442" -> "24.4.2", "2522" -> "25.2.2"
func FormatVersion(version string) string {
	if len(version) == 4 {
		// Format: XYZZ -> X.Y.Z (e.g., 2442 -> 24.4.2)
		return fmt.Sprintf("%s.%s.%s", version[0:2], version[2:3], version[3:4])
	}
	return version
}

// formatVersionPrevious returns a user-friendly description of the previous version
func formatVersionPrevious(version string) string {
	return FormatVersion(version) + " (or earlier)"
}

// ValidateVersionRanges validates integer fields against version-specific ranges
func ValidateVersionRanges(
	providerVersion string,
	planValue interface{},
	rangeConstraints []FieldRangeConstraint,
	diagnostics *diag.Diagnostics,
) {
	if providerVersion == "" {
		// If no version configured, skip validation (backwards compatibility)
		return
	}

	// Note: This function is called AFTER ValidateVersionConstraints, so fields that are
	// removed/not-yet-added in the current version will have already been caught.
	// Here we only validate the range of fields that ARE valid for this version.

	// For each range constraint, check if the field is set and validate its range
	for _, constraint := range rangeConstraints {
		// Get the field value
		fieldValue, ok := getInt64FieldValue(planValue, constraint.FieldPath)
		if !ok {
			// Field not set or not an int64, skip
			continue
		}

		// Find the applicable range for the provider version
		applicableRange, applicableVersion := getRangeForVersion(constraint.VersionRanges, providerVersion)
		if applicableRange == nil {
			// No range constraint for this version, use widest range
			applicableRange, applicableVersion = getWidestRange(constraint.VersionRanges)
		}

		if applicableRange != nil {
			// Validate the value against the range
			if fieldValue < applicableRange.Min || fieldValue > applicableRange.Max {
				// Build error message showing only the current version's range
				diagnostics.AddError(
					fmt.Sprintf("Value Out of Range for IOS-XR Version %s", FormatVersion(providerVersion)),
					fmt.Sprintf(
						"The field '%s' value %d is outside the valid range for IOS-XR version %s.\n"+
							"Valid range for version %s: %d-%d",
						constraint.FieldPath,
						fieldValue,
						FormatVersion(providerVersion),
						FormatVersion(applicableVersion),
						applicableRange.Min,
						applicableRange.Max,
					),
				)
			}
		}
	}
}

// getInt64FieldValue retrieves the int64 value of a field at the given path
func getInt64FieldValue(planValue interface{}, fieldPath string) (int64, bool) {
	parts := strings.Split(fieldPath, ".")
	current := reflect.ValueOf(planValue)

	for i, part := range parts {
		// Dereference pointers
		if current.Kind() == reflect.Ptr {
			if current.IsNil() {
				return 0, false
			}
			current = current.Elem()
		}

		if current.Kind() != reflect.Struct {
			return 0, false
		}

		// Find the field
		field := current.FieldByName(toCamelCase(part))
		if !field.IsValid() {
			return 0, false
		}

		// Handle slices - check all elements
		if field.Kind() == reflect.Slice {
			if i == len(parts)-1 {
				// Last part is a slice, not an int64
				return 0, false
			}

			// Check nested fields in slice items
			remainingPath := strings.Join(parts[i+1:], ".")
			for j := 0; j < field.Len(); j++ {
				sliceElem := field.Index(j)
				if val, ok := getInt64FromValue(sliceElem, remainingPath); ok {
					return val, true
				}
			}
			return 0, false
		}

		// If this is the last part, check if it's an Int64
		if i == len(parts)-1 {
			if v, ok := field.Interface().(types.Int64); ok {
				if !v.IsNull() && !v.IsUnknown() {
					return v.ValueInt64(), true
				}
			}
			return 0, false
		}

		// Continue navigating
		current = field
	}

	return 0, false
}

// getInt64FromValue extracts int64 from a reflect.Value by navigating a field path
func getInt64FromValue(value reflect.Value, fieldPath string) (int64, bool) {
	parts := strings.Split(fieldPath, ".")
	current := value

	for i, part := range parts {
		if current.Kind() == reflect.Ptr {
			if current.IsNil() {
				return 0, false
			}
			current = current.Elem()
		}

		if current.Kind() != reflect.Struct {
			return 0, false
		}

		field := current.FieldByName(toCamelCase(part))
		if !field.IsValid() {
			return 0, false
		}

		// Handle slices
		if field.Kind() == reflect.Slice {
			if i == len(parts)-1 {
				return 0, false
			}
			remainingPath := strings.Join(parts[i+1:], ".")
			for j := 0; j < field.Len(); j++ {
				if val, ok := getInt64FromValue(field.Index(j), remainingPath); ok {
					return val, true
				}
			}
			return 0, false
		}

		// Check if it's the target Int64 field
		if i == len(parts)-1 {
			if v, ok := field.Interface().(types.Int64); ok {
				if !v.IsNull() && !v.IsUnknown() {
					return v.ValueInt64(), true
				}
			}
			return 0, false
		}

		current = field
	}

	return 0, false
}

// getRangeForVersion finds the applicable range for a given version
func getRangeForVersion(versionRanges map[string]VersionRange, version string) (*VersionRange, string) {
	// Check exact match first
	if r, ok := versionRanges[version]; ok {
		return &r, version
	}

	// Find the highest version that is <= the requested version
	var bestVersion string
	var bestRange *VersionRange

	for ver, r := range versionRanges {
		if VersionAtLeast(version, ver) {
			if bestVersion == "" || VersionAtLeast(ver, bestVersion) {
				bestVersion = ver
				rCopy := r
				bestRange = &rCopy
			}
		}
	}

	return bestRange, bestVersion
}

// getWidestRange returns the union of all ranges (min of all mins, max of all maxs)
func getWidestRange(versionRanges map[string]VersionRange) (*VersionRange, string) {
	if len(versionRanges) == 0 {
		return nil, ""
	}

	var minRange, maxRange int64
	var firstVersion string
	first := true

	for ver, r := range versionRanges {
		if first {
			minRange = r.Min
			maxRange = r.Max
			firstVersion = ver
			first = false
		} else {
			if r.Min < minRange {
				minRange = r.Min
			}
			if r.Max > maxRange {
				maxRange = r.Max
			}
		}
	}

	return &VersionRange{Min: minRange, Max: maxRange}, firstVersion
}

---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_snmp_vrf_host Data Source - terraform-provider-iosxr"
subcategory: "Management"
description: |-
  This data source can read the SNMP VRF Host configuration.
---

# iosxr_snmp_vrf_host (Data Source)

This data source can read the SNMP VRF Host configuration.

## Example Usage

```terraform
data "iosxr_snmp_vrf_host" "example" {
  vrf_name = "11.11.11.11"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vrf_name` (String) VRF name

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.
- `traps_unencrypted_unencrypted_string_version_v3_security_level` (String)


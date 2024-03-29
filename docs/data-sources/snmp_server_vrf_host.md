---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_snmp_server_vrf_host Data Source - terraform-provider-iosxr"
subcategory: "Management"
description: |-
  This data source can read the SNMP Server VRF Host configuration.
---

# iosxr_snmp_server_vrf_host (Data Source)

This data source can read the SNMP Server VRF Host configuration.

## Example Usage

```terraform
data "iosxr_snmp_server_vrf_host" "example" {
  vrf_name = "VRF1"
  address  = "11.11.11.11"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) Specify hosts to receive SNMP notifications
- `vrf_name` (String) VRF name

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.
- `unencrypted_strings` (Attributes List) The UNENCRYPTED (cleartext) community string (see [below for nested schema](#nestedatt--unencrypted_strings))

<a id="nestedatt--unencrypted_strings"></a>
### Nested Schema for `unencrypted_strings`

Read-Only:

- `community_string` (String) The UNENCRYPTED (cleartext) community string
- `udp_port` (String) udp port to which notifications should be sent
- `version_v3_security_level` (String)

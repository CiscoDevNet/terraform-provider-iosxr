---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_router_vrrp_interface Data Source - terraform-provider-iosxr"
subcategory: "VRRP"
description: |-
  This data source can read the Router VRRP Interface configuration.
---

# iosxr_router_vrrp_interface (Data Source)

This data source can read the Router VRRP Interface configuration.

## Example Usage

```terraform
data "iosxr_router_vrrp_interface" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_name` (String) VRRP interface configuration subcommands

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `bfd_minimum_interval` (Number) Hello interval
- `bfd_multiplier` (Number) Detect multiplier
- `delay_minimum` (Number) Set minimum delay on every interface up event
- `delay_reload` (Number) Set reload delay for first interface up event
- `id` (String) The path of the retrieved object.
- `mac_refresh` (Number) Set the Subordinate MAC-refresh rate for this interface

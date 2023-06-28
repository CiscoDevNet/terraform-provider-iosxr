---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_router_vrrp_interface_address_family_ipv4 Resource - terraform-provider-iosxr"
subcategory: "VRRP"
description: |-
  This resource can manage the Router VRRP Interface Address Family IPv4 configuration.
---

# iosxr_router_vrrp_interface_address_family_ipv4 (Resource)

This resource can manage the Router VRRP Interface Address Family IPv4 configuration.

## Example Usage

```terraform
resource "iosxr_router_vrrp_interface_address_family_ipv4" "example" {
  interface_name                      = "GigabitEthernet0/0/0/1"
  vrrp_id                             = 123
  version                             = 2
  address                             = "1.1.1.1"
  priority                            = 250
  name                                = "TEST"
  text_authentication                 = "7"
  timer_advertisement_time_in_seconds = 123
  timer_force                         = false
  preempt_disable                     = false
  preempt_delay                       = 255
  accept_mode_disable                 = false
  track_interfaces = [
    {
      interface_name     = "GigabitEthernet0/0/0/1"
      priority_decrement = 12
    }
  ]
  track_objects = [
    {
      object_name        = "OBJECT"
      priority_decrement = 22
    }
  ]
  bfd_fast_detect_peer_ipv4 = "33.33.33.3"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_name` (String) VRRP interface configuration subcommands
- `version` (Number) VRRP version
  - Range: `2`-`3`
- `vrrp_id` (Number) VRRP configuration
  - Range: `1`-`255`

### Optional

- `accept_mode_disable` (Boolean) Disable accept mode
- `address` (String) Enable VRRP and specify IP address(es)
- `bfd_fast_detect_peer_ipv4` (String) BFD peer interface IPv4 address
- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `name` (String) Configure VRRP Session name
- `preempt_delay` (Number) Wait before preempting
  - Range: `1`-`3600`
- `preempt_disable` (Boolean) Disable preemption
- `priority` (Number) Set priority level
  - Range: `1`-`254`
- `secondary_addresses` (Attributes List) VRRP IPv4 address (see [below for nested schema](#nestedatt--secondary_addresses))
- `text_authentication` (String) Set plain text authentication string
- `timer_advertisement_time_in_milliseconds` (Number) Configure in milliseconds
  - Range: `100`-`40950`
- `timer_advertisement_time_in_seconds` (Number) Advertisement time in seconds
  - Range: `1`-`255`
- `timer_force` (Boolean) Force the configured values to be used
- `track_interfaces` (Attributes List) Track an interface (see [below for nested schema](#nestedatt--track_interfaces))
- `track_objects` (Attributes List) Object Tracking (see [below for nested schema](#nestedatt--track_objects))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--secondary_addresses"></a>
### Nested Schema for `secondary_addresses`

Required:

- `address` (String) VRRP IPv4 address


<a id="nestedatt--track_interfaces"></a>
### Nested Schema for `track_interfaces`

Required:

- `interface_name` (String) Track an interface

Optional:

- `priority_decrement` (Number) Priority decrement
  - Range: `1`-`254`


<a id="nestedatt--track_objects"></a>
### Nested Schema for `track_objects`

Required:

- `object_name` (String) Object to be tracked
- `priority_decrement` (Number) Priority decrement
  - Range: `1`-`254`

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_router_vrrp_interface_address_family_ipv4.example "Cisco-IOS-XR-um-router-vrrp-cfg:/router/vrrp/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]/address-family/ipv4/vrrps/vrrp[vrrp-id=%!d(string=123)][version=%!d(string=2)]"
```
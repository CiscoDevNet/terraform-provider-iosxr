---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_router_static_vrf Data Source - terraform-provider-iosxr"
subcategory: "Routing"
description: |-
  This data source can read the Router Static VRF configuration.
---

# iosxr_router_static_vrf (Data Source)

This data source can read the Router Static VRF configuration.

## Example Usage

```terraform
data "iosxr_router_static_vrf" "example" {
  vrf_name = "VRF2"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vrf_name` (String) VRF Static route configuration subcommands

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.
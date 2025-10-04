---
subcategory: "Guides"
page_title: "Manage Multiple Devices"
description: |-
    Howto manage multiple devices.
---

# Manage Multiple Devices

When it comes to managing multiple IOS-XR devices, one can create multiple provider configurations and distinguish them by `alias` ([documentation](https://www.terraform.io/language/providers/configuration#alias-multiple-provider-configurations)).

```terraform
provider "iosxr" {
  alias = "ROUTER-1"
  host  = "10.1.1.1:57400"
}

provider "iosxr" {
  alias = "ROUTER-2"
  host  = "10.1.1.2:57400"
}

resource "iosxr_hostname" "ROUTER-1" {
  provider            = iosxr.ROUTER-1
  system_network_name = "ROUTER-1"
}

resource "iosxr_hostname" "ROUTER-2" {
  provider            = iosxr.ROUTER-2
  system_network_name = "ROUTER-2"
}
```

The disadvantages here is that the `provider` attribute of resources cannot be dynamic and therefore cannot be used in combination with `for_each` as an example. The issue is being tracked [here](https://github.com/hashicorp/terraform/issues/24476).

This provider offers an alternative approach where mutliple devices can be managed by a single provider configuration and the optional `device` attribute, which is available in every resource and data source, can then be used to select the respective device. This assumes that every device uses the same credentials.

```terraform
locals {
  routers = [
    {
      name = "ROUTER-1"
      host = "10.1.1.1:57400"
    },
    {
      name = "ROUTER-2"
      host = "10.1.1.2:57400"
    },
  ]
}

provider "iosxr" {
  devices = local.routers
}

resource "iosxr_hostname" "hostname" {
  for_each            = toset([for router in local.routers : router.name])
  device              = each.key
  system_network_name = each.key
}
```

## Device-Level Management Control

### The "managed" Flag

Each device in the `devices` list supports an optional `managed` attribute that controls whether the device is actively managed by Terraform:

- **`managed = true`** (default): Device is actively managed - configurations are applied
- **`managed = false`**: Device is "frozen" - state preserved but no changes applied
- **Not specified**: Defaults to `true`

### Basic Managed Flag Usage

```hcl
locals {
  devices = [
    {
      name    = "production-rtr01"
      host    = "10.1.1.10:57400"
      managed = true   # Actively managed
    },
    {
      name    = "maintenance-rtr02"
      host    = "10.1.1.20:57400"
      managed = false  # Temporarily frozen for maintenance
    },
    {
      name    = "active-rtr03"
      host    = "10.1.1.30:57400"
      # managed defaults to true when not specified
    }
  ]
}

provider "iosxr" {
  devices = local.devices
}

resource "iosxr_banner" "login_banner" {
  for_each = toset([for device in local.devices : device.name])
  device   = each.key
  login    = "Authorized Access Only - ${each.key}"
}
```

**Result**:
- `production-rtr01` and `active-rtr03`: Banner configuration applied
- `maintenance-rtr02`: Configuration skipped, existing state preserved

### Relationship with `selected_devices`

**Important**: The [`selected_devices` provider attribute](selective_deploy.md) takes precedence over individual `managed` flags:

- **When `selected_devices` is specified**: Individual `managed` flags are ignored
- **When `selected_devices` is not specified**: Individual `managed` flags are respected

#### Example: selected_devices Override
```hcl
provider "iosxr" {
  selected_devices = ["router-01", "router-03"]  # Only these devices managed
  devices = [
    { name = "router-01", host = "10.1.1.10:57400", managed = false },  # Overridden to managed=true
    { name = "router-02", host = "10.1.1.20:57400", managed = true },   # Overridden to managed=false
    { name = "router-03", host = "10.1.1.30:57400", managed = true }    # Remains managed=true
  ]
}
```

**Result**: Only `router-01` and `router-03` are managed, regardless of their individual `managed` flags.

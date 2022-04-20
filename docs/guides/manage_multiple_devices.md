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
  alias    = "ROUTER-1"
  username = "admin"
  password = "Cisco123"
  host     = "10.1.1.1"
}

provider "iosxr" {
  alias    = "ROUTER-2"
  username = "admin"
  password = "Cisco123"
  host     = "10.1.1.2"
}

resource "iosxr_gnmi" "ROUTER-1" {
  provider   = iosxr.ROUTER-1
  path       = "openconfig-system:system/config"
  attributes = {
    hostname = "ROUTER-1"
  }
}

resource "iosxr_gnmi" "ROUTER-2" {
  provider   = iosxr.ROUTER-2
  path       = "openconfig-system:system/config"
  attributes = {
    hostname = "ROUTER-2"
  }
}
```

The disadvantages here is that the `provider` attribute of resources cannot be dynamic and therefore cannot be used in combination with `for_each` as an example. The issue is being tracked [here](https://github.com/hashicorp/terraform/issues/24476).

This provider offers an alternative approach where mutliple devices can be managed by a single provider configuration and the optional `device` attribute, which is available in every resource and data source, can then be used to select the respective device. This assumes that every device uses the same credentials.

```terraform
locals {
  routers = [
    {
      name = "ROUTER-1"
      host  = "10.1.1.1"
    },
    {
      name = "ROUTER-2"
      host  = "10.1.1.2"
    },
  ]
}

provider "iosxr" {
  username = "admin"
  password = "Cisco123"
  devices  = local.routers
}

resource "iosxr_gnmi" "hostname" {
  for_each   = toset([for router in local.routers : router.name])
  device     = each.key
  path       = "openconfig-system:system/config"
  attributes = {
    hostname = each.key
  }
}
```

---
subcategory: "Guides"
page_title: "Importing Resources"
description: |-
    Importing Resources
---

# Importing Resources

A resource can be imported by using the `terraform import` command or by using an `import` block in the configuration. The resource documentation has an example for the `terraform import` command. An example for importing a resource using the `import` block is shown below. Assuming we have the following resource in our configuration:

```terraform
resource "iosxr_router_bgp" "bgp" {
  as_number = "65000"
}
```

We could add an import block to import the resource as shown below:

```terraform
import {
  to = iosxr_router_bgp.bgp
  id = "65000"
}
```

This will populate the state for the `iosxr_router_bgp` resource for the "default" device. When managing multiple devices, the device name can be appended to the import identifier using a comma as a separator:

```terraform
resource "iosxr_router_bgp" "bgp" {
  device    = "device1"
  as_number = "65000"
}

import {
  to = iosxr_router_bgp.bgp
  id = "65000,device1"
}
```

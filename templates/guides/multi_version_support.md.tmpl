---
subcategory: "Guides"
page_title: "Multi-Version Support"
description: |-
    Howto manage devices running different IOS-XR releases.
---

# Multi-Version Support

## Overview

A single version of this provider can manage devices running different IOS-XR releases. The provider determines the release of each device and uses the matching YANG paths, attributes and value constraints, so one configuration can target a mixed fleet.

### Supported Releases

| Release | Version used by the provider |
|---------|------------------------------|
| 24.4.2  | `24.4`                       |
| 25.4.2  | `25.4`                       |
| 26.2.1  | `26.2`                       |

The provider works with `major.minor` and ignores the patch, so `24.4.1` and `24.4.2` are both `24.4`. Only the releases listed above are tested. The behavior on other releases is not guaranteed.

### How the Version Is Determined

1. **Explicit version**: the `iosxr_version` provider attribute or the `IOSXR_VERSION` environment variable. The value applies to every device of the provider.
2. **Auto-detection**: if no version is set, the provider reads the release of every managed device when the provider is configured. It queries `Cisco-IOS-XR-install-oper:install/version`, normalizes the label to `major.minor` and caches the result per device.

Keep the following in mind:

- Auto-detection uses gNMI only. When the provider uses `protocol = "netconf"`, set `iosxr_version` explicitly. NETCONF support for version auto-detection will be added at a later date.
- Auto-detection happens when the provider is configured, so every managed device must be reachable. If a device cannot be queried, the provider fails with `Unable to Auto-Detect IOS-XR Version`. Set `iosxr_version` to skip the detection, or set `managed = false` for the device.
- An explicit `iosxr_version` applies to all devices. To manage devices with different releases in one provider configuration, use auto-detection.

## Configuration Reference

### Auto-Detection

This is the default. No version needs to be configured:

```hcl
provider "iosxr" {
  devices = [
    { name = "router-01", host = "10.1.1.10:57400" },  # 24.4.2
    { name = "router-02", host = "10.1.1.20:57400" }   # 26.2.1
  ]
}

```

### Explicit Version

Set the version with the `iosxr_version` attribute. The patch version in `major.minor.patch` is optional and ignored:

```hcl
provider "iosxr" {
  iosxr_version = "25.4"
  host          = "10.1.1.10:57400"
}

```

Alternatively, use the `IOSXR_VERSION` environment variable:

```bash
export IOSXR_VERSION="25.4.2"
```

### Reading the Detected Version

The `iosxr_device_info` data source returns the version that the provider uses for a device. It does not open an additional connection, the value is read from the provider. The value is an empty string if the detection was skipped, for example for a device with `managed = false`.

```hcl
data "iosxr_device_info" "router_02" {
  device = "router-02"
}

locals {
  router_02_version = data.iosxr_device_info.router_02.version  # for example "26.2"
}
```

## Version-Specific Attributes and Validation

Between releases, attributes can be added or removed, and value ranges, enum values, string lengths and patterns can change. The provider checks the configuration against the release of the device. A check fails with an error such as:

- `Field Not Supported in IOS-XR Version`: the attribute is not available in the device release. It was added in a later release, or removed in that release or an earlier one.
- `Value Out of Range for IOS-XR Version`: the number is outside the range of the device release.
- `Invalid Enum Value for IOS-XR Version`: the value is not valid for the device release.
- `Invalid String Length for IOS-XR Version` and `Invalid Value for IOS-XR Version`: the length or the pattern is not valid for the device release.

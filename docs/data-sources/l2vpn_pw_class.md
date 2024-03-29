---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_l2vpn_pw_class Data Source - terraform-provider-iosxr"
subcategory: "L2VPN"
description: |-
  This data source can read the L2VPN PW Class configuration.
---

# iosxr_l2vpn_pw_class (Data Source)

This data source can read the L2VPN PW Class configuration.

## Example Usage

```terraform
data "iosxr_l2vpn_pw_class" "example" {
  name = "PWC1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Pseudowire class template

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `encapsulation_mpls` (Boolean) Set pseudowire encapsulation to MPLS
- `encapsulation_mpls_load_balancing_flow_label_both` (Boolean) Insert/Discard Flow label on transmit/recceive
- `encapsulation_mpls_load_balancing_flow_label_both_static` (Boolean) Set Flow label parameters statically
- `encapsulation_mpls_load_balancing_flow_label_code_one7` (Boolean) Legacy code value
- `encapsulation_mpls_load_balancing_flow_label_code_one7_disable` (Boolean) Disables sending code 17 TLV
- `encapsulation_mpls_load_balancing_flow_label_receive` (Boolean) Discard Flow label on receive
- `encapsulation_mpls_load_balancing_flow_label_receive_static` (Boolean) Set Flow label parameters statically
- `encapsulation_mpls_load_balancing_flow_label_transmit` (Boolean) Insert Flow label on transmit
- `encapsulation_mpls_load_balancing_flow_label_transmit_static` (Boolean) Set Flow label parameters statically
- `encapsulation_mpls_load_balancing_pw_label` (Boolean) Enable PW VC label based load balancing
- `encapsulation_mpls_transport_mode_ethernet` (Boolean) Ethernet port mode
- `encapsulation_mpls_transport_mode_passthrough` (Boolean) passthrough incoming tags
- `encapsulation_mpls_transport_mode_vlan` (Boolean) Vlan tagged mode
- `id` (String) The path of the retrieved object.

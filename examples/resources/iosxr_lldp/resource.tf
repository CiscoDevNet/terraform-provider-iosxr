resource "iosxr_lldp" "example" {
  holdtime                               = 50
  timer                                  = 6
  reinit                                 = 3
  subinterfaces_enable                   = true
  priorityaddr_enable                    = true
  extended_show_width_enable             = true
  tlv_select_management_address_disable  = true
  tlv_select_port_description_disable    = true
  tlv_select_system_capabilities_disable = true
  tlv_select_system_description_disable  = true
  tlv_select_system_name_disable         = true
}

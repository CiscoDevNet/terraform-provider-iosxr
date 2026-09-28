resource "iosxr_lldp" "example" {
  chassis_id                             = "FOC22439P72"
  extended_show_width_enable             = true
  holdtime                               = 50
  management_enable                      = true
  priorityaddr_enable                    = true
  reinit                                 = 3
  subinterfaces_enable                   = true
  subinterfaces_tagged                   = true
  system_description                     = "Router1-Description"
  system_name                            = "Router1"
  timer                                  = 6
  tlv_select_management_address_disable  = true
  tlv_select_port_description_disable    = true
  tlv_select_system_capabilities_disable = true
  tlv_select_system_description_disable  = true
  tlv_select_system_name_disable         = true
}

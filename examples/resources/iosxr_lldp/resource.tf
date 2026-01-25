resource "iosxr_lldp" "example" {
  holdtime                               = 50
  timer                                  = 6
  reinit                                 = 3
  system_name                            = "Router1"
  system_description                     = "Router1-Description"
  chassis_id                             = "FOC22439P72"
  chassis_id_type_local                  = true
  subinterfaces_enable                   = true
  subinterfaces_tagged                   = true
  management_enable                      = true
  priorityaddr_enable                    = true
  extended_show_width_enable             = true
  tlv_select_management_address_disable  = true
  tlv_select_port_description_disable    = true
  tlv_select_system_capabilities_disable = true
  tlv_select_system_description_disable  = true
  tlv_select_system_name_disable         = true
}

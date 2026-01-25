resource "iosxr_router_hsrp_interface" "example" {
  interface_name            = "GigabitEthernet0/0/0/1"
  hsrp_use_bia              = true
  hsrp_redirects_disable    = true
  hsrp_delay_minimum        = 500
  hsrp_delay_reload         = 700
  hsrp_bfd_minimum_interval = 20000
  hsrp_bfd_multiplier       = 40
  hsrp_mac_refresh          = 5000
}

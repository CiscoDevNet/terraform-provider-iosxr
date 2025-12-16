resource "iosxr_performance_measurement_interface" "example" {
  interface_name                    = "GigabitEthernet0/0/0/1"
  delay_measurement                 = true
  delay_measurement_advertise_delay = 1000
  delay_measurement_profile_name    = "DELAY_PROFILE_1"
  delay_measurement_static_delay    = 1000000
  next_hop_ipv4                     = "10.0.0.1"
}

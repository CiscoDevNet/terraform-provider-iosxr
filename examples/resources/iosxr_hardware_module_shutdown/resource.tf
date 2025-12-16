resource "iosxr_hardware_module_shutdown" "example" {
  location_name = "0/0/CPU0"
  unshut        = true
}

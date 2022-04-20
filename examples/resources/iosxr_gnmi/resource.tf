resource "iosxr_gnmi" "example" {
  path = "openconfig-system:/system/config"
  attributes = {
    hostname = "ROUTER-1"
  }
}

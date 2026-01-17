resource "iosxr_cli" "example" {
  cli = "interface Loopback0 description configured-via-gnmi-cli"
}

resource "iosxr_ftp" "example" {
  client_vrfs = [
    {
      vrf_name = "VRF1"
      passive = true
      source_interface = "Loopback0"
      anonymous_password = "mypassword"
      username = "ftpuser"
      password = "myencryptedpassword"
    }
  ]
}

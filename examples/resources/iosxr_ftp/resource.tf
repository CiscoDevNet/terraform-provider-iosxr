resource "iosxr_ftp" "example" {
  client_vrfs = [
    {
      anonymous_password = "mypassword"
      passive            = true
      password           = "myencryptedpassword"
      source_interface   = "Loopback0"
      username           = "ftpuser"
      vrf_name           = "VRF1"
    }
  ]
}

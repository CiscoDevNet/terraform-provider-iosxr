resource "iosxr_pce" "example" {
  address_ipv4 = "77.77.77.1"
  state_sync_ipv4s = [
    {
      address = "100.100.100.11"
    }
  ]
  peer_filter_ipv4_access_list = "Accesslist1"
  api_authentication_digest    = true
  api_sibling_ipv4             = "100.100.100.2"
  api_users = [
    {
      user_name          = "rest-user"
      password_encrypted = "00141215174C04140B"
    }
  ]
}

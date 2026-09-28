resource "iosxr_crypto_client_authentication" "example" {
  # NOTE: This resource is only supported from IOS-XR version 25.4 and above
  profile = [
    {
      password_six = "Cisco123!"
      profile_name = "EAP_PROFILE"
      username     = "dot1x_client"
    }
  ]
}

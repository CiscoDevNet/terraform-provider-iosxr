resource "iosxr_crypto_client_authentication" "example" {
  # NOTE: This resource is only supported from IOS-XR version 25.1 and above
  profile = [
    {
      profile_name = "EAP_PROFILE"
      password_six = "Cisco123!"
      username     = "dot1x_client"
    }
  ]
}

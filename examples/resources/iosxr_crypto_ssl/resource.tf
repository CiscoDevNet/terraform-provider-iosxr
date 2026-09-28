resource "iosxr_crypto_ssl" "example" {
  # NOTE: This resource is only supported from IOS-XR version 25.4 and above
  profile = [
    {
      certificate  = "CORP_PKI_CA"
      profile_name = "MTLS_PROFILE"
    }
  ]
}

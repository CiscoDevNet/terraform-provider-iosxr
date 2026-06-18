resource "iosxr_crypto_ssl" "example" {
  # NOTE: This resource is only supported from IOS-XR version 25.1 and above
  profile = [
    {
      profile_name = "MTLS_PROFILE"
      certificate  = "CORP_PKI_CA"
    }
  ]
}

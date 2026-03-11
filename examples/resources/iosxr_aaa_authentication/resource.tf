resource "iosxr_aaa_authentication" "example" {
  login = [
    {
      list = "AAA-LIST"
      a1_tacacs = true
      a2_radius = true
      a3_group = "AAA3"
      a4_local = true
    }
  ]
}

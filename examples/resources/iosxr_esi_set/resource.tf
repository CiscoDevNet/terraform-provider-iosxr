resource "iosxr_esi_set" "example" {
  rpl      = "esi-set POLICYSET\n  1234.1234.1234.1234.1234\nend-set\n"
  set_name = "POLICYSET"
}

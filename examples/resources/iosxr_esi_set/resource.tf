resource "iosxr_esi_set" "example" {
  set_name = "POLICYSET"
  rpl      = "esi-set POLICYSET\n  1234.1234.1234.1234.1234\nend-set\n"
}

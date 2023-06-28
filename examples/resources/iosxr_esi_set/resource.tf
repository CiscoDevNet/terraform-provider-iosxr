resource "iosxr_esi_set" "example" {
  set_name        = "POLICYSET"
  esi_set_as_text = "esi-set POLICYSET\n  1234.1234.1234.1234.1234\nend-set\n"
}

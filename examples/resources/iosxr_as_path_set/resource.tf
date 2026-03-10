resource "iosxr_as_path_set" "example" {
  set_name = "TEST1"
  rpl      = "as-path-set TEST1\n  length ge 10\nend-set\n"
}

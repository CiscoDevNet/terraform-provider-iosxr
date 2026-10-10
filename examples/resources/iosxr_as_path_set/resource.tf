resource "iosxr_as_path_set" "example" {
  rpl      = "as-path-set TEST1\n  length ge 10\nend-set\n"
  set_name = "TEST1"
}

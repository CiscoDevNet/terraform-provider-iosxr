resource "iosxr_as_path_set" "example" {
  set_name       = "TEST1"
  rplas_path_set = "as-path-set TEST1\n  length ge 10\nend-set\n"
}

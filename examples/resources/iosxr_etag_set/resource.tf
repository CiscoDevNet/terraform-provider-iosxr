resource "iosxr_etag_set" "example" {
  rpl      = "etag-set ETAG1\n  65536,\n  401001\nend-set\n"
  set_name = "ETAG1"
}

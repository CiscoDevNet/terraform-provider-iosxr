resource "iosxr_large_community_set" "example" {
  set_name = "LARGE1"
  rpl = "large-community-set LARGE1\n  0:65001:1,\n  0:65002:*,\n  [65536..40001]:201001:301001,\n  ios-regex '^65000:.*:.*'\nend-set\n"
}

resource "iosxr_policy_global_set" "example" {
  rpl = "policy-global\n  GLOBAL1 'global1',\n  GLOBAL2 'global2'\nend-global\n"
}

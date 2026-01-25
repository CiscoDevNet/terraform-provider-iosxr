resource "iosxr_vty_pool" "example" {
  default_first_vty = 0
  default_last_vty = 10
  default_line_template = "default"
  eem_first_vty = 100
  eem_last_vty = 105
  eem_line_template = "EEM_TEMPLATE"
  pools = [
    {
      pool_name = "USER_POOL"
      first_vty = "20"
      last_vty = "30"
      line_template = "USER_TEMPLATE"
    }
  ]
}

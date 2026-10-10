resource "iosxr_aaa_accounting" "example" {
  commands = [
    {
      a1_tacacs  = true
      a2_group   = "AAA2"
      a3_local   = true
      a4_none    = true
      list       = "AAA-COMMANDS"
      start_stop = true
    }
  ]
  exec = [
    {
      a1_tacacs  = true
      a2_radius  = true
      a3_group   = "AAA3"
      a4_none    = true
      list       = "AAA-EXEC"
      start_stop = true
    }
  ]
  network = [
    {
      a1_tacacs  = true
      a2_radius  = true
      a3_group   = "AAA3"
      a4_none    = true
      list       = "AAA-NETWORK"
      start_stop = true
    }
  ]
  system = [
    {
      a1_tacacs  = true
      a2_radius  = true
      a3_group   = "AAA3"
      a4_none    = true
      list       = "AAA-SYSTEM"
      start_stop = true
    }
  ]
  update_newinfo = true
}

resource "iosxr_aaa_accounting" "example" {
  update_newinfo = true
  exec = [
    {
      list       = "AAA-EXEC"
      start_stop = true
      a1_tacacs  = true
      a2_radius  = true
      a3_group   = "AAA3"
      a4_none    = true
    }
  ]
  commands = [
    {
      list       = "AAA-COMMANDS"
      start_stop = true
      a1_tacacs  = true
      a2_group   = "AAA2"
      a3_local   = true
      a4_none    = true
    }
  ]
  system = [
    {
      list       = "AAA-SYSTEM"
      start_stop = true
      a1_tacacs  = true
      a2_radius  = true
      a3_group   = "AAA3"
      a4_none    = true
    }
  ]
  network = [
    {
      list       = "AAA-NETWORK"
      start_stop = true
      a1_tacacs  = true
      a2_radius  = true
      a3_group   = "AAA3"
      a4_none    = true
    }
  ]
}

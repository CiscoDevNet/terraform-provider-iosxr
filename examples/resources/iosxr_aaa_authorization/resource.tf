resource "iosxr_aaa_authorization" "example" {
  commands = [
    {
      a1_tacacs = true
      a2_group  = "AAA2"
      a3_local  = true
      a4_none   = true
      list      = "AAA-COMMANDS"
    }
  ]
  eventmanager = [
    {
      a1_tacacs = true
      list      = "AAA-EVENTMANAGER"
    }
  ]
  exec = [
    {
      a1_tacacs = true
      a2_radius = true
      a3_group  = "AAA3"
      a4_local  = true
      list      = "AAA-EXEC"
    }
  ]
  network = [
    {
      a1_tacacs = true
      a2_radius = true
      a3_group  = "AAA3"
      a4_local  = true
      list      = "AAA-NETWORK"
    }
  ]
}

resource "iosxr_aaa_authorization" "example" {
  exec = [
    {
      list = "AAA-EXEC"
      a1_tacacs = true
      a2_radius = true
      a3_group = "AAA3"
      a4_local = true
    }
  ]
  eventmanager = [
    {
      list = "AAA-EVENTMANAGER"
      a1_tacacs = true
    }
  ]
  commands = [
    {
      list = "AAA-COMMANDS"
      a1_tacacs = true
      a2_group = "AAA2"
      a3_local = true
      a4_none = true
    }
  ]
  network = [
    {
      list = "AAA-NETWORK"
      a1_tacacs = true
      a2_radius = true
      a3_group = "AAA3"
      a4_local = true
    }
  ]
}

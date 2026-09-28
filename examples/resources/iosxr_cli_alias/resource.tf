resource "iosxr_cli_alias" "example" {
  aliases = [
    {
      command = "show version"
      name    = "show-version"
    }
  ]
  config_aliases = [
    {
      command = "interface GigabitEthernet0/0/0/0"
      name    = "int-config"
    }
  ]
  exec_aliases = [
    {
      command = "show version"
      name    = "sv"
    }
  ]
}

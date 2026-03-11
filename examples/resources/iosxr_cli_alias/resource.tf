resource "iosxr_cli_alias" "example" {
  aliases = [
    {
      name = "show-version"
      command = "show version"
    }
  ]
  exec_aliases = [
    {
      name = "sv"
      command = "show version"
    }
  ]
  config_aliases = [
    {
      name = "int-config"
      command = "interface GigabitEthernet0/0/0/0"
    }
  ]
}

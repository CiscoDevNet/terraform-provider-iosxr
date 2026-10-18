resource "iosxr_srlg" "example" {
  groups = [
    {
      group_name = "GROUP1"
      indexes = [
        {
          index_number = 1
          priority     = "critical"
          value        = 200
        }
      ]
    }
  ]
  interfaces = [
    {
      groups = [
        {
          group_name   = "INTF-GROUP1"
          index_number = 1
        }
      ]
      include_optical          = true
      include_optical_priority = "critical"
      indexes = [
        {
          index_number = 1
          priority     = "critical"
          value        = 300
        }
      ]
      interface_name = "GigabitEthernet0/0/0/1"
      names = [
        {
          srlg_name = "INTF-SRLG1"
        }
      ]
    }
  ]
  names = [
    {
      description = "SRLG for critical links"
      srlg_name   = "SRLG1"
      value       = 100
    }
  ]
}

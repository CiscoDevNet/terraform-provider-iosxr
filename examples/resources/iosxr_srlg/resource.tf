resource "iosxr_srlg" "example" {
  names = [
    {
      srlg_name = "SRLG1"
      value = 100
      description = "SRLG for critical links"
    }
  ]
  interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
      include_optical = true
      include_optical_priority = "critical"
        indexes = [
          {
            index_number = 1
            value = 300
            priority = "critical"
          }
        ]
        names = [
          {
            srlg_name = "INTF-SRLG1"
          }
        ]
        groups = [
          {
            index_number = 1
            group_name = "INTF-GROUP1"
          }
        ]
    }
  ]
  groups = [
    {
      group_name = "GROUP1"
        indexes = [
          {
            index_number = 1
            value = 200
            priority = "critical"
          }
        ]
    }
  ]
}

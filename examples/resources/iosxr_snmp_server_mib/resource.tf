resource "iosxr_snmp_server_mib" "example" {
  cbqosmib_cache = true
  cbqosmib_cache_refresh_time = 30
  cbqosmib_cache_service_policy_count = 100
  cbqosmib_persist = true
  cbqosmib_member_stats = true
  ifindex_persist = true
  interfaces = [
    {
      interface_name = "GigabitEthernet0/0/0/1"
      notification_linkupdown_disable = true
      index_persistence = true
    }
  ]
  trap_link_ietf = true
  ifmib_ifalias_long = true
  ifmib_stats_cache = true
  ifmib_ipsubscriber = true
  ifmib_internal_cache_max_duration = 30
  rfmib_entphyindex = true
  sensormib_cache = true
  mplstemib_cache_timers_garbage_collect = 1800
  mplstemib_cache_timers_refresh = 300
  mplsp2mpmib_cache_timer = 300
  frrmib_cache_timer = 300
  cmplsteextmib_cache_timer = 300
  cmplsteextstdmib_cache_timer = 300
  mroutemib_send_all_vrf = true
  notification_log_mib_default = true
  notification_log_mib_global_age_out = 120
  notification_log_mib_global_size = 5000
  notification_log_mib_disable = true
  notification_log_mib_size = 1000
  entityindex_persist = true
}

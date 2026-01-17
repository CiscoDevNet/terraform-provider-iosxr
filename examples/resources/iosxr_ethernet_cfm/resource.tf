resource "iosxr_ethernet_cfm" "example" {
  traceroute_cache_hold_time = 60
  traceroute_cache_size      = 3000
  domains = [
    {
      domain_name            = "DOMAIN1"
      level                  = 5
      id_mac_address         = "00:11:22:33:44:55"
      id_mac_address_integer = 100
      services = [
        {
          service_name                             = "SERVICE1"
          xconnect_p2p_group_name                  = "XC-GROUP1"
          xconnect_p2p_xc_name                     = "XC-P2P1"
          id_icc_based_icc                         = "ICC1"
          id_icc_based_umc                         = "UMC1"
          tags                                     = "1"
          mip_auto_create_lower_mep_only           = true
          mip_auto_create_ccm_learning             = true
          continuity_check_interval                = "1s"
          continuity_check_interval_loss_threshold = 5
          continuity_check_archive_hold_time       = 60
          continuity_check_loss_auto_traceroute    = true
          maximum_meps                             = 100
          ais_transmission_interval                = "1s"
          ais_transmission_cos                     = 5
          log_continuity_check_mep_changes         = true
          log_continuity_check_errors              = true
          log_crosscheck_errors                    = true
          log_ais                                  = true
          log_csf                                  = true
          mep_crosschecks = [
            {
              mep_id      = 5
              mac_address = "00:11:22:33:44:55"
            }
          ]
          mep_crosscheck_auto = true
        }
      ]
    }
  ]
}

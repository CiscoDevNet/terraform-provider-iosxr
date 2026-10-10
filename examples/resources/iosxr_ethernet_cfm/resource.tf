resource "iosxr_ethernet_cfm" "example" {
  domains = [
    {
      domain_name            = "DOMAIN1"
      id_mac_address         = "00:11:22:33:44:55"
      id_mac_address_integer = 100
      level                  = 5
      services = [
        {
          ais_transmission_cos                     = 5
          ais_transmission_interval                = "1s"
          continuity_check_archive_hold_time       = 60
          continuity_check_interval                = "1s"
          continuity_check_interval_loss_threshold = 5
          continuity_check_loss_auto_traceroute    = true
          id_icc_based_icc                         = "ICC1"
          id_icc_based_umc                         = "UMC1"
          log_ais                                  = true
          log_continuity_check_errors              = true
          log_continuity_check_mep_changes         = true
          log_crosscheck_errors                    = true
          log_csf                                  = true
          maximum_meps                             = 100
          mep_crosscheck_auto                      = true
          mep_crosschecks = [
            {
              mac_address = "00:11:22:33:44:55"
              mep_id      = 5
            }
          ]
          mip_auto_create_ccm_learning   = true
          mip_auto_create_lower_mep_only = true
          service_name                   = "SERVICE1"
          tags                           = "1"
          xconnect_p2p_group_name        = "XC-GROUP1"
          xconnect_p2p_xc_name           = "XC-P2P1"
        }
      ]
    }
  ]
  traceroute_cache_hold_time = 60
  traceroute_cache_size      = 3000
}

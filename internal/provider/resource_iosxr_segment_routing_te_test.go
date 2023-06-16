// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrSegmentRoutingTE(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSegmentRoutingTEConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "logging_pcep_peer_status", "true"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "logging_policy_status", "true"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_report_all", "true"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_source_address", "88.88.88.8"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_delegation_timeout", "10"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_dead_timer", "60"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_initiated_state", "15"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_initiated_orphan", "10"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pce_peers.0.pce_address", "66.66.66.6"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pce_peers.0.precedence", "122"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.dynamic_anycast_sid_inclusion", "true"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.dynamic_metric_type", "te"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.color", "266"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.srv6_locator_name", "LOC11"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.srv6_locator_behavior", "ub6-insert-reduced"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.srv6_locator_binding_sid_type", "srv6-dynamic"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.source_address", "fccc:0:213::1"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.source_address_type", "end-point-type-ipv6"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.effective_metric_value", "4444"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.effective_metric_type", "igp"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.constraint_segments_protection_type", "protected-only"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.constraint_segments_sid_algorithm", "128"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.policy_name", "POLICY1"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.srv6_locator_name", "Locator11"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.srv6_locator_binding_sid_type", "srv6-dynamic"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.srv6_locator_behavior", "ub6-insert-reduced"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.source_address", "fccc:0:103::1"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.source_address_type", "end-point-type-ipv6"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.policy_color_endpoint_color", "65534"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.policy_color_endpoint_type", "end-point-type-ipv6"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.policy_color_endpoint_address", "fccc:0:215::1"),
				),
			},
			{
				ResourceName:  "iosxr_segment_routing_te.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-segment-routing-ms-cfg:/sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering",
			},
		},
	})
}

func testAccIosxrSegmentRoutingTEConfig_minimum() string {
	return `
	resource "iosxr_segment_routing_te" "test" {
	}
	`
}

func testAccIosxrSegmentRoutingTEConfig_all() string {
	return `
	resource "iosxr_segment_routing_te" "test" {
		logging_pcep_peer_status = true
		logging_policy_status = true
		pcc_report_all = true
		pcc_source_address = "88.88.88.8"
		pcc_delegation_timeout = 10
		pcc_dead_timer = 60
		pcc_initiated_state = 15
		pcc_initiated_orphan = 10
		pce_peers = [{
			pce_address = "66.66.66.6"
			precedence = 122
		}]
		on_demand_colors = [{
			dynamic_anycast_sid_inclusion = true
			dynamic_metric_type = "te"
			color = 266
			srv6_locator_name = "LOC11"
			srv6_locator_behavior = "ub6-insert-reduced"
			srv6_locator_binding_sid_type = "srv6-dynamic"
			source_address = "fccc:0:213::1"
			source_address_type = "end-point-type-ipv6"
			effective_metric_value = 4444
			effective_metric_type = "igp"
			constraint_segments_protection_type = "protected-only"
			constraint_segments_sid_algorithm = 128
		}]
		policies = [{
			policy_name = "POLICY1"
			srv6_locator_name = "Locator11"
			srv6_locator_binding_sid_type = "srv6-dynamic"
			srv6_locator_behavior = "ub6-insert-reduced"
			source_address = "fccc:0:103::1"
			source_address_type = "end-point-type-ipv6"
			policy_color_endpoint_color = 65534
			policy_color_endpoint_type = "end-point-type-ipv6"
			policy_color_endpoint_address = "fccc:0:215::1"
		}]
	}
	`
}

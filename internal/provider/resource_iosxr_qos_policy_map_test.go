// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrQOSPolicyMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrQOSPolicyMapConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "policy_map_name", "core-ingress-classifier"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_name", "class-default"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_type", "qos"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_set_mpls_experimental_topmost", "0"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_set_dscp", "0"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_queue_limits_queue_limit.0.value", "100"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_queue_limits_queue_limit.0.unit", "us"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_service_policy_name", "SERVICEPOLICY"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_police_rate_value", "5"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_police_rate_unit", "gbps"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_shape_average_rate_value", "100"),
					resource.TestCheckResourceAttr("iosxr_qos_policy_map.test", "class_shape_average_rate_unit", "gbps"),
				),
			},
			{
				ResourceName:  "iosxr_qos_policy_map.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name=core-ingress-classifier]",
			},
		},
	})
}

func testAccIosxrQOSPolicyMapConfig_minimum() string {
	return `
	resource "iosxr_qos_policy_map" "test" {
		policy_map_name = "core-ingress-classifier"
	}
	`
}

func testAccIosxrQOSPolicyMapConfig_all() string {
	return `
	resource "iosxr_qos_policy_map" "test" {
		policy_map_name = "core-ingress-classifier"
		class_name = "class-default"
		class_type = "qos"
		class_set_mpls_experimental_topmost = 0
		class_set_dscp = "0"
		class_queue_limits_queue_limit = [{
			value = "100"
			unit = "us"
		}]
		class_service_policy_name = "SERVICEPOLICY"
		class_police_rate_value = "5"
		class_police_rate_unit = "gbps"
		class_shape_average_rate_value = "100"
		class_shape_average_rate_unit = "gbps"
	}
	`
}

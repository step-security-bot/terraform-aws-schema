package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpnConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_attachment_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_gateway_configuration": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "customer_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_acceleration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_ipv4_network_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_ipv6_network_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "outside_ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_ipv4_network_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_ipv6_network_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "destination_cidr_block": "string",
              "source": "string",
              "state": "string"
            }
          ]
        ]
      },
      "static_routes_only": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "transit_gateway_attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transport_transit_gateway_attachment_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel1_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel1_bgp_asn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel1_bgp_holdtime": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tunnel1_cgw_inside_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel1_dpd_timeout_action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel1_dpd_timeout_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel1_enable_tunnel_lifecycle_control": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tunnel1_ike_versions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel1_inside_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel1_inside_ipv6_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel1_phase1_dh_group_numbers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "number"
        ]
      },
      "tunnel1_phase1_encryption_algorithms": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel1_phase1_integrity_algorithms": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel1_phase1_lifetime_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel1_phase2_dh_group_numbers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "number"
        ]
      },
      "tunnel1_phase2_encryption_algorithms": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel1_phase2_integrity_algorithms": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel1_phase2_lifetime_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel1_preshared_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "tunnel1_rekey_fuzz_percentage": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel1_rekey_margin_time_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel1_replay_window_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel1_startup_action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel1_vgw_inside_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel2_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel2_bgp_asn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel2_bgp_holdtime": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tunnel2_cgw_inside_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel2_dpd_timeout_action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel2_dpd_timeout_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel2_enable_tunnel_lifecycle_control": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tunnel2_ike_versions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel2_inside_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel2_inside_ipv6_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel2_phase1_dh_group_numbers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "number"
        ]
      },
      "tunnel2_phase1_encryption_algorithms": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel2_phase1_integrity_algorithms": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel2_phase1_lifetime_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel2_phase2_dh_group_numbers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "number"
        ]
      },
      "tunnel2_phase2_encryption_algorithms": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel2_phase2_integrity_algorithms": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tunnel2_phase2_lifetime_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel2_preshared_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "tunnel2_rekey_fuzz_percentage": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel2_rekey_margin_time_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel2_replay_window_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tunnel2_startup_action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel2_vgw_inside_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel_inside_ip_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vgw_telemetry": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "accepted_route_count": "number",
              "certificate_arn": "string",
              "last_status_change": "string",
              "outside_ip_address": "string",
              "status": "string",
              "status_message": "string"
            }
          ]
        ]
      },
      "vpn_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "tunnel1_log_options": {
        "block": {
          "block_types": {
            "cloudwatch_log_options": {
              "block": {
                "attributes": {
                  "log_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_output_format": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "tunnel2_log_options": {
        "block": {
          "block_types": {
            "cloudwatch_log_options": {
              "block": {
                "attributes": {
                  "log_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_output_format": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpnConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpnConnection), &result)
	return &result
}

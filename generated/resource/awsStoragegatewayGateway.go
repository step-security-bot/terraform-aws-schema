package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsStoragegatewayGateway = `{
  "block": {
    "attributes": {
      "activation_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "average_download_rate_limit_in_bits_per_sec": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "average_upload_rate_limit_in_bits_per_sec": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cloudwatch_log_group_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ec2_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gateway_network_interface": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ipv4_address": "string"
            }
          ]
        ]
      },
      "gateway_timezone": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gateway_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_vpc_endpoint": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_environment": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "medium_changer_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "smb_file_share_visibility": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "smb_guest_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "smb_security_strategy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "tape_drive_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "maintenance_start_time": {
        "block": {
          "attributes": {
            "day_of_month": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "day_of_week": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hour_of_day": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "minute_of_hour": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "smb_active_directory_settings": {
        "block": {
          "attributes": {
            "active_directory_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "domain_controllers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "domain_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "organizational_unit": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "username": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsStoragegatewayGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsStoragegatewayGateway), &result)
	return &result
}

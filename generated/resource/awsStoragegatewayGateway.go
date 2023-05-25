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
      "cloudwatch_log_group_arn": {
        "description_kind": "plain",
        "optional": true,
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
      "smb_guest_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "tape_drive_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "smb_active_directory_settings": {
        "block": {
          "attributes": {
            "domain_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
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

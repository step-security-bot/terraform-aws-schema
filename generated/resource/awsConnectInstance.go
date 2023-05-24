package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConnectInstance = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_resolve_best_voices_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "contact_flow_logs_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "contact_lens_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "early_media_enabled": {
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
      "identity_management_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "inbound_calls_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "instance_alias": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_party_conference_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "outbound_calls_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "service_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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

func AwsConnectInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConnectInstance), &result)
	return &result
}

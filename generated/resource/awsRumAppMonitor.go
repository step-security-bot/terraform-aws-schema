package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRumAppMonitor = `{
  "block": {
    "attributes": {
      "app_monitor_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cw_log_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cw_log_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
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
      }
    },
    "block_types": {
      "app_monitor_configuration": {
        "block": {
          "attributes": {
            "allow_cookies": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_xray": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "excluded_pages": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "favorite_pages": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "guest_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "identity_pool_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "included_pages": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "session_sample_rate": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "telemetries": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "custom_events": {
        "block": {
          "attributes": {
            "status": {
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
  "version": 0
}`

func AwsRumAppMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRumAppMonitor), &result)
	return &result
}

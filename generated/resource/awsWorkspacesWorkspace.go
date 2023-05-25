package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkspacesWorkspace = `{
  "block": {
    "attributes": {
      "bundle_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "computer_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
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
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_volume_encryption_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
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
      "user_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_volume_encryption_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "volume_encryption_key": {
        "description_kind": "plain",
        "optional": true,
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
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "workspace_properties": {
        "block": {
          "attributes": {
            "compute_type_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "root_volume_size_gib": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "running_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "running_mode_auto_stop_timeout_in_minutes": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "user_volume_size_gib": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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

func AwsWorkspacesWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkspacesWorkspace), &result)
	return &result
}

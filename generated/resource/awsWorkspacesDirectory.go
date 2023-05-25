package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkspacesDirectory = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_user_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "directory_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "iam_role_id": {
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
      "ip_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "registration_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "workspace_security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "self_service_permissions": {
        "block": {
          "attributes": {
            "change_compute_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "increase_volume_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "rebuild_workspace": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "restart_workspace": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "switch_running_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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

func AwsWorkspacesDirectorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkspacesDirectory), &result)
	return &result
}

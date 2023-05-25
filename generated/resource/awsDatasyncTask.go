package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDatasyncTask = `{
  "block": {
    "attributes": {
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
      "destination_location_arn": {
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
        "optional": true,
        "type": "string"
      },
      "source_location_arn": {
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
      }
    },
    "block_types": {
      "options": {
        "block": {
          "attributes": {
            "atime": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bytes_per_second": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "gid": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mtime": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "posix_permissions": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preserve_deleted_files": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preserve_devices": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uid": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "verify_mode": {
              "description_kind": "plain",
              "optional": true,
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

func AwsDatasyncTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDatasyncTask), &result)
	return &result
}

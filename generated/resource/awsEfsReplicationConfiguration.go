package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEfsReplicationConfiguration = `{
  "block": {
    "attributes": {
      "creation_time": {
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
      "original_source_file_system_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_file_system_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_file_system_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_file_system_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "destination": {
        "block": {
          "attributes": {
            "availability_zone_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "file_system_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
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

func AwsEfsReplicationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEfsReplicationConfiguration), &result)
	return &result
}

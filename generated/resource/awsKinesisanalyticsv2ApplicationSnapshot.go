package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKinesisanalyticsv2ApplicationSnapshot = `{
  "block": {
    "attributes": {
      "application_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "application_version_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_creation_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_name": {
        "description_kind": "plain",
        "required": true,
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

func AwsKinesisanalyticsv2ApplicationSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKinesisanalyticsv2ApplicationSnapshot), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmMaintenanceWindow = `{
  "block": {
    "attributes": {
      "allow_unassociated_targets": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cutoff": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "duration": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "end_date": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_offset": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "schedule_timezone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_date": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsmMaintenanceWindowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmMaintenanceWindow), &result)
	return &result
}

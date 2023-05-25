package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDmsReplicationTask = `{
  "block": {
    "attributes": {
      "cdc_start_position": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cdc_start_time": {
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
      "migration_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_instance_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_task_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_task_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_task_settings": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_mappings": {
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
      },
      "target_endpoint_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDmsReplicationTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDmsReplicationTask), &result)
	return &result
}

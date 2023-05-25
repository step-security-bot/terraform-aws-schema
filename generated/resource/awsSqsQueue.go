package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSqsQueue = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_based_deduplication": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delay_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "fifo_queue": {
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
      "kms_data_key_reuse_period_seconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "kms_master_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_message_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "message_retention_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "receive_wait_time_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "redrive_policy": {
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
      "visibility_timeout_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSqsQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSqsQueue), &result)
	return &result
}

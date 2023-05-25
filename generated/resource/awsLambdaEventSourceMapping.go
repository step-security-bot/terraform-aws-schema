package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaEventSourceMapping = `{
  "block": {
    "attributes": {
      "batch_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "bisect_batch_on_function_error": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "event_source_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "function_name": {
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
      "last_modified": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_processing_result": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "maximum_batching_window_in_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maximum_record_age_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maximum_retry_attempts": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "parallelization_factor": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "starting_position": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "starting_position_timestamp": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state_transition_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "uuid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "destination_config": {
        "block": {
          "block_types": {
            "on_failure": {
              "block": {
                "attributes": {
                  "destination_arn": {
                    "description_kind": "plain",
                    "required": true,
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
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdaEventSourceMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaEventSourceMapping), &result)
	return &result
}

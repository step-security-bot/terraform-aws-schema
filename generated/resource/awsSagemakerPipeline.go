package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerPipeline = `{
  "block": {
    "attributes": {
      "arn": {
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
      "pipeline_definition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pipeline_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pipeline_display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pipeline_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
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
    "block_types": {
      "parallelism_configuration": {
        "block": {
          "attributes": {
            "max_parallel_execution_steps": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "pipeline_definition_s3_location": {
        "block": {
          "attributes": {
            "bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "object_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version_id": {
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

func AwsSagemakerPipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerPipeline), &result)
	return &result
}

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDatapipelinePipelineDefinition = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameter_object": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "attribute": [
                "set",
                [
                  "object",
                  {
                    "key": "string",
                    "string_value": "string"
                  }
                ]
              ],
              "id": "string"
            }
          ]
        ]
      },
      "pipeline_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pipeline_object": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "field": [
                "set",
                [
                  "object",
                  {
                    "key": "string",
                    "ref_value": "string",
                    "string_value": "string"
                  }
                ]
              ],
              "id": "string",
              "name": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "parameter_value": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "string_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDatapipelinePipelineDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDatapipelinePipelineDefinition), &result)
	return &result
}

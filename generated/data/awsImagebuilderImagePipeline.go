package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsImagebuilderImagePipeline = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "container_recipe_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "date_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "date_last_run": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "date_next_run": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "date_updated": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "distribution_configuration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enhanced_image_metadata_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_recipe_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_scanning_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ecr_configuration": [
                "list",
                [
                  "object",
                  {
                    "container_tags": [
                      "set",
                      "string"
                    ],
                    "repository_name": "string"
                  }
                ]
              ],
              "image_scanning_enabled": "bool"
            }
          ]
        ]
      },
      "image_tests_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "image_tests_enabled": "bool",
              "timeout_minutes": "number"
            }
          ]
        ]
      },
      "infrastructure_configuration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "pipeline_execution_start_condition": "string",
              "schedule_expression": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
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

func AwsImagebuilderImagePipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsImagebuilderImagePipeline), &result)
	return &result
}

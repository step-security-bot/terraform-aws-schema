package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsImagebuilderImage = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_recipe_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "date_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "distribution_configuration_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enhanced_image_metadata_enabled": {
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
      "image_recipe_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "infrastructure_configuration_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "os_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "output_resources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "amis": [
                "set",
                [
                  "object",
                  {
                    "account_id": "string",
                    "description": "string",
                    "image": "string",
                    "name": "string",
                    "region": "string"
                  }
                ]
              ],
              "containers": [
                "set",
                [
                  "object",
                  {
                    "image_uris": [
                      "set",
                      "string"
                    ],
                    "region": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "platform": {
        "computed": true,
        "description_kind": "plain",
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
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "image_tests_configuration": {
        "block": {
          "attributes": {
            "image_tests_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "timeout_minutes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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

func AwsImagebuilderImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsImagebuilderImage), &result)
	return &result
}

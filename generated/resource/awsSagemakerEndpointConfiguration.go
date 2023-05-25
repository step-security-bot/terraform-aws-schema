package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerEndpointConfiguration = `{
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
      "kms_key_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
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
      }
    },
    "block_types": {
      "production_variants": {
        "block": {
          "attributes": {
            "accelerator_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "initial_instance_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "initial_variant_weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "model_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "variant_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSagemakerEndpointConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerEndpointConfiguration), &result)
	return &result
}

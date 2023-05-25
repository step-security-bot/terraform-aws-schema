package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConfigConformancePack = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_s3_bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delivery_s3_key_prefix": {
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
      "template_body": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_s3_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "input_parameter": {
        "block": {
          "attributes": {
            "parameter_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parameter_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 60,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConfigConformancePackSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConfigConformancePack), &result)
	return &result
}

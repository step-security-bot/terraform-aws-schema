package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaLayerVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compatible_runtimes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filename": {
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
      "layer_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "layer_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "license_info": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_object_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_code_hash": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_code_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdaLayerVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaLayerVersion), &result)
	return &result
}

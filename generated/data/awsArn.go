package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsArn = `{
  "block": {
    "attributes": {
      "account": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
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
      "partition": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsArnSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsArn), &result)
	return &result
}

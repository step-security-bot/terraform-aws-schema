package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamSessionContext = `{
  "block": {
    "attributes": {
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
      "issuer_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "issuer_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "issuer_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "session_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIamSessionContextSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamSessionContext), &result)
	return &result
}

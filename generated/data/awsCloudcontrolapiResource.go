package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudcontrolapiResource = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "properties": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type_version_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudcontrolapiResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudcontrolapiResource), &result)
	return &result
}

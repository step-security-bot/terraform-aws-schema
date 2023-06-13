package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsControltowerControls = `{
  "block": {
    "attributes": {
      "enabled_controls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsControltowerControlsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsControltowerControls), &result)
	return &result
}

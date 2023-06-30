package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSfnStateMachineVersions = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "statemachine_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "statemachine_versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSfnStateMachineVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSfnStateMachineVersions), &result)
	return &result
}

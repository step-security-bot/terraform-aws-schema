package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsInspectorRulesPackages = `{
  "block": {
    "attributes": {
      "arns": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsInspectorRulesPackagesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsInspectorRulesPackages), &result)
	return &result
}

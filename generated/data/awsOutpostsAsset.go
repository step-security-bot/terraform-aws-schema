package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOutpostsAsset = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "asset_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "asset_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "host_id": {
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
      "rack_elevation": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "rack_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOutpostsAssetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOutpostsAsset), &result)
	return &result
}

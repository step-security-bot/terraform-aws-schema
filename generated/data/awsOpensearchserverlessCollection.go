package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessCollection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collection_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpensearchserverlessCollectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessCollection), &result)
	return &result
}

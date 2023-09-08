package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3Objects = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "common_prefixes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "delimiter": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encoding_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fetch_owner": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "keys": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "max_keys": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "owners": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_charged": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "request_payer": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_after": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3ObjectsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3Objects), &result)
	return &result
}

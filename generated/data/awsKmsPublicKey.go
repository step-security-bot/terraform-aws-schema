package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKmsPublicKey = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_master_key_spec": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_algorithms": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "grant_tokens": {
        "description_kind": "plain",
        "optional": true,
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
      "key_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_usage": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_pem": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "signing_algorithms": {
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

func AwsKmsPublicKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKmsPublicKey), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamUserLoginProfile = `{
  "block": {
    "attributes": {
      "encrypted_password": {
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
      "key_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "password_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "password_reset_required": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "pgp_key": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIamUserLoginProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamUserLoginProfile), &result)
	return &result
}

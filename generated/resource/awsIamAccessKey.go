package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamAccessKey = `{
  "block": {
    "attributes": {
      "create_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encrypted_secret": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encrypted_ses_smtp_password_v4": {
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
      "pgp_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "ses_smtp_password_v4": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
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

func AwsIamAccessKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamAccessKey), &result)
	return &result
}

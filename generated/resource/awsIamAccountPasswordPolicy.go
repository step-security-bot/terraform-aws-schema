package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamAccountPasswordPolicy = `{
  "block": {
    "attributes": {
      "allow_users_to_change_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "expire_passwords": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "hard_expiry": {
        "computed": true,
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
      "max_password_age": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "minimum_password_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "password_reuse_prevention": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "require_lowercase_characters": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "require_numbers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "require_symbols": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "require_uppercase_characters": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIamAccountPasswordPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamAccountPasswordPolicy), &result)
	return &result
}

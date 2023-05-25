package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAccountAlternateContact = `{
  "block": {
    "attributes": {
      "account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alternate_contact_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "email_address": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "phone_number": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "title": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAccountAlternateContactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAccountAlternateContact), &result)
	return &result
}

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSesv2EmailIdentityMailFromAttributes = `{
  "block": {
    "attributes": {
      "behavior_on_mx_failure": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "email_identity": {
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
      "mail_from_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSesv2EmailIdentityMailFromAttributesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSesv2EmailIdentityMailFromAttributes), &result)
	return &result
}

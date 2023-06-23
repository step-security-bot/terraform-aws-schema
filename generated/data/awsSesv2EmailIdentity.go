package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSesv2EmailIdentity = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_set_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dkim_signing_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "current_signing_key_length": "string",
              "domain_signing_private_key": "string",
              "domain_signing_selector": "string",
              "last_key_generation_timestamp": "string",
              "next_signing_key_length": "string",
              "signing_attributes_origin": "string",
              "status": "string",
              "tokens": [
                "list",
                "string"
              ]
            }
          ]
        ]
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
      "identity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "verified_for_sending_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSesv2EmailIdentitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSesv2EmailIdentity), &result)
	return &result
}

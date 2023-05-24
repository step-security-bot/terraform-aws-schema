package resource

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
        "description_kind": "plain",
        "optional": true,
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
      "identity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
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
    "block_types": {
      "dkim_signing_attributes": {
        "block": {
          "attributes": {
            "current_signing_key_length": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "domain_signing_private_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "domain_signing_selector": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_key_generation_timestamp": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "next_signing_key_length": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "signing_attributes_origin": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tokens": {
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
        "max_items": 1,
        "nesting_mode": "list"
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

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIdentitystoreUser = `{
  "block": {
    "attributes": {
      "addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "country": "string",
              "formatted": "string",
              "locality": "string",
              "postal_code": "string",
              "primary": "bool",
              "region": "string",
              "street_address": "string",
              "type": "string"
            }
          ]
        ]
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "emails": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "primary": "bool",
              "type": "string",
              "value": "string"
            }
          ]
        ]
      },
      "external_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "issuer": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "locale": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "family_name": "string",
              "formatted": "string",
              "given_name": "string",
              "honorific_prefix": "string",
              "honorific_suffix": "string",
              "middle_name": "string"
            }
          ]
        ]
      },
      "nickname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "phone_numbers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "primary": "bool",
              "type": "string",
              "value": "string"
            }
          ]
        ]
      },
      "preferred_language": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "timezone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "title": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "alternate_identifier": {
        "block": {
          "block_types": {
            "external_id": {
              "block": {
                "attributes": {
                  "id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "issuer": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "unique_attribute": {
              "block": {
                "attributes": {
                  "attribute_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "attribute_value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "filter": {
        "block": {
          "attributes": {
            "attribute_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "attribute_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "deprecated": true,
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

func AwsIdentitystoreUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIdentitystoreUser), &result)
	return &result
}

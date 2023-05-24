package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIdentitystoreUser = `{
  "block": {
    "attributes": {
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nickname": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "profile_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timezone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "title": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "addresses": {
        "block": {
          "attributes": {
            "country": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "formatted": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "locality": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "postal_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "primary": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "street_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "emails": {
        "block": {
          "attributes": {
            "primary": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "name": {
        "block": {
          "attributes": {
            "family_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "formatted": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "given_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "honorific_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "honorific_suffix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "middle_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "phone_numbers": {
        "block": {
          "attributes": {
            "primary": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
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
  "version": 0
}`

func AwsIdentitystoreUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIdentitystoreUser), &result)
	return &result
}

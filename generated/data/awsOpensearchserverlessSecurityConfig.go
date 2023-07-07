package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchserverlessSecurityConfig = `{
  "block": {
    "attributes": {
      "config_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "saml_options": {
        "block": {
          "attributes": {
            "group_attribute": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "metadata": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "session_timeout": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "user_attribute": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpensearchserverlessSecurityConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchserverlessSecurityConfig), &result)
	return &result
}

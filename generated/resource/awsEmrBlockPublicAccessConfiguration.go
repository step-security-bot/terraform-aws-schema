package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEmrBlockPublicAccessConfiguration = `{
  "block": {
    "attributes": {
      "block_public_security_group_rules": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "permitted_public_security_group_rule_range": {
        "block": {
          "attributes": {
            "max_range": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_range": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEmrBlockPublicAccessConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEmrBlockPublicAccessConfiguration), &result)
	return &result
}

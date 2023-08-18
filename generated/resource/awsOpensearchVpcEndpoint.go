package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpensearchVpcEndpoint = `{
  "block": {
    "attributes": {
      "domain_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "vpc_options": {
        "block": {
          "attributes": {
            "availability_zones": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpensearchVpcEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpensearchVpcEndpoint), &result)
	return &result
}

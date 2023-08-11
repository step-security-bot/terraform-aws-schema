package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcIpamPool = `{
  "block": {
    "attributes": {
      "address_family": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "allocation_default_netmask_length": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "allocation_max_netmask_length": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "allocation_min_netmask_length": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "allocation_resource_tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_import": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "aws_service": {
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
        "optional": true,
        "type": "string"
      },
      "ipam_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_scope_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_scope_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "locale": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pool_depth": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "publicly_advertisable": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "source_ipam_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
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
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
              "description_kind": "plain",
              "optional": true,
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

func AwsVpcIpamPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcIpamPool), &result)
	return &result
}

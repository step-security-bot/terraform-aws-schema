package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcIpamPool = `{
  "block": {
    "attributes": {
      "address_family": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "allocation_default_netmask_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_max_netmask_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_min_netmask_length": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_resource_tags": {
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
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "aws_service": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_scope_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_scope_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "locale": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pool_depth": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "public_ip_source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publicly_advertisable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source_ipam_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
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

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcIpamPools = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_pools": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "address_family": "string",
              "allocation_default_netmask_length": "number",
              "allocation_max_netmask_length": "number",
              "allocation_min_netmask_length": "number",
              "allocation_resource_tags": [
                "map",
                "string"
              ],
              "arn": "string",
              "auto_import": "bool",
              "aws_service": "string",
              "description": "string",
              "id": "string",
              "ipam_pool_id": "string",
              "ipam_scope_id": "string",
              "ipam_scope_type": "string",
              "locale": "string",
              "pool_depth": "number",
              "publicly_advertisable": "bool",
              "source_ipam_pool_id": "string",
              "state": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
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
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpcIpamPoolsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcIpamPools), &result)
	return &result
}

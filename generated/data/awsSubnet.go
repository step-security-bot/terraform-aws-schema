package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSubnet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assign_ipv6_address_on_creation": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "available_ip_address_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "customer_owned_ipv4_pool": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_for_az": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_dns64": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_resource_name_dns_a_record_on_launch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_resource_name_dns_aaaa_record_on_launch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_cidr_block": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_cidr_block_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ipv6_native": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "map_customer_owned_ip_on_launch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "map_public_ip_on_launch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "outpost_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_hostname_type_on_launch": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsSubnetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSubnet), &result)
	return &result
}

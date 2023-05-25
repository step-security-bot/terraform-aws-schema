package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAlb = `{
  "block": {
    "attributes": {
      "access_logs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bucket": "string",
              "enabled": "bool",
              "prefix": "string"
            }
          ]
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn_suffix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_owned_ipv4_pool": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "desync_mitigation_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "drop_invalid_header_fields": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_deletion_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_http2": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_waf_fail_open": {
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
      "idle_timeout": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "internal": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "subnet_mapping": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "allocation_id": "string",
              "ipv6_address": "string",
              "outpost_id": "string",
              "private_ipv4_address": "string",
              "subnet_id": "string"
            }
          ]
        ]
      },
      "subnets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
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
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAlbSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAlb), &result)
	return &result
}

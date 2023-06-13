package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElb = `{
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
              "bucket_prefix": "string",
              "enabled": "bool",
              "interval": "number"
            }
          ]
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "connection_draining": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "connection_draining_timeout": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "cross_zone_load_balancing": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "health_check": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "healthy_threshold": "number",
              "interval": "number",
              "target": "string",
              "timeout": "number",
              "unhealthy_threshold": "number"
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
      "idle_timeout": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "internal": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "listener": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "instance_port": "number",
              "instance_protocol": "string",
              "lb_port": "number",
              "lb_protocol": "string",
              "ssl_certificate_id": "string"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
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
      "source_security_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AwsElbSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElb), &result)
	return &result
}

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMqBroker = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "broker_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "broker_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "revision": "number"
            }
          ]
        ]
      },
      "deployment_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_id": "string",
              "use_aws_owned_key": "bool"
            }
          ]
        ]
      },
      "engine_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "host_instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "console_url": "string",
              "endpoints": [
                "list",
                "string"
              ],
              "ip_address": "string"
            }
          ]
        ]
      },
      "maintenance_window_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "day_of_week": "string",
              "time_of_day": "string",
              "time_zone": "string"
            }
          ]
        ]
      },
      "publicly_accessible": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "security_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "subnet_ids": {
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
      "user": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "console_access": "bool",
              "groups": [
                "set",
                "string"
              ],
              "username": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "logs": {
        "block": {
          "attributes": {
            "audit": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "general": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
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

func AwsMqBrokerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMqBroker), &result)
	return &result
}

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
      "authentication_strategy": {
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
      "ldap_server_metadata": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "hosts": [
                "list",
                "string"
              ],
              "role_base": "string",
              "role_name": "string",
              "role_search_matching": "string",
              "role_search_subtree": "bool",
              "service_account_password": "string",
              "service_account_username": "string",
              "user_base": "string",
              "user_role_name": "string",
              "user_search_matching": "string",
              "user_search_subtree": "bool"
            }
          ]
        ]
      },
      "logs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "audit": "string",
              "general": "bool"
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
      "storage_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
              "replication_user": "bool",
              "username": "string"
            }
          ]
        ]
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

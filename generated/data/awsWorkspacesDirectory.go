package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWorkspacesDirectory = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_user_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "directory_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "iam_role_id": {
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
      "ip_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "registration_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "self_service_permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "change_compute_type": "bool",
              "increase_volume_size": "bool",
              "rebuild_workspace": "bool",
              "restart_workspace": "bool",
              "switch_running_mode": "bool"
            }
          ]
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
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "workspace_access_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "device_type_android": "string",
              "device_type_chromeos": "string",
              "device_type_ios": "string",
              "device_type_linux": "string",
              "device_type_osx": "string",
              "device_type_web": "string",
              "device_type_windows": "string",
              "device_type_zeroclient": "string"
            }
          ]
        ]
      },
      "workspace_creation_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_security_group_id": "string",
              "default_ou": "string",
              "enable_internet_access": "bool",
              "enable_maintenance_mode": "bool",
              "user_enabled_as_local_administrator": "bool"
            }
          ]
        ]
      },
      "workspace_security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWorkspacesDirectorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWorkspacesDirectory), &result)
	return &result
}

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFsxWindowsFileSystem = `{
  "block": {
    "attributes": {
      "active_directory_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aliases": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "audit_log_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "audit_log_destination": "string",
              "file_access_audit_log_level": "string",
              "file_share_access_audit_log_level": "string"
            }
          ]
        ]
      },
      "automatic_backup_retention_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "backup_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "copy_tags_to_backups": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "daily_automatic_backup_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_iops_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "iops": "number",
              "mode": "string"
            }
          ]
        ]
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_file_server_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "skip_final_backup": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "storage_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "throughput_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "weekly_maintenance_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsFsxWindowsFileSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFsxWindowsFileSystem), &result)
	return &result
}

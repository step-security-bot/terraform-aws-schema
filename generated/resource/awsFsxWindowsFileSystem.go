package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFsxWindowsFileSystem = `{
  "block": {
    "attributes": {
      "active_directory_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "aliases": {
        "description_kind": "plain",
        "optional": true,
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
      "automatic_backup_retention_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "backup_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "copy_tags_to_backups": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "daily_automatic_backup_start_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_name": {
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
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "remote_administration_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "skip_final_backup": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "storage_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
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
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "throughput_capacity": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "audit_log_configuration": {
        "block": {
          "attributes": {
            "audit_log_destination": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "file_access_audit_log_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "file_share_access_audit_log_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "self_managed_active_directory": {
        "block": {
          "attributes": {
            "dns_ips": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "domain_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "file_system_administrators_group": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organizational_unit_distinguished_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "username": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func AwsFsxWindowsFileSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFsxWindowsFileSystem), &result)
	return &result
}

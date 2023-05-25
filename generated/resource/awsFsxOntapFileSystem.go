package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFsxOntapFileSystem = `{
  "block": {
    "attributes": {
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
      "daily_automatic_backup_start_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_ip_address_range": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "intercluster": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "management": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "fsx_admin_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
          "list",
          "string"
        ]
      },
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_subnet_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_table_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "security_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "storage_capacity": {
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
      "disk_iops_configuration": {
        "block": {
          "attributes": {
            "iops": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mode": {
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

func AwsFsxOntapFileSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFsxOntapFileSystem), &result)
	return &result
}

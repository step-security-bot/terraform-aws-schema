package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFsxFileCache = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "copy_tags_to_data_repository_associations": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "data_repository_association_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_cache_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_cache_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "file_cache_type_version": {
        "description_kind": "plain",
        "required": true,
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
        "required": true,
        "type": "number"
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
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "data_repository_association": {
        "block": {
          "attributes": {
            "association_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "data_repository_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "data_repository_subdirectories": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "file_cache_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "file_cache_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "file_system_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "file_system_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "imported_file_chunk_size": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "resource_arn": {
              "computed": true,
              "description_kind": "plain",
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
            }
          },
          "block_types": {
            "nfs": {
              "block": {
                "attributes": {
                  "dns_ips": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "version": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 8,
        "nesting_mode": "set"
      },
      "lustre_configuration": {
        "block": {
          "attributes": {
            "deployment_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_configuration": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                [
                  "object",
                  {
                    "destination": "string",
                    "level": "string"
                  }
                ]
              ]
            },
            "mount_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "per_unit_storage_throughput": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "weekly_maintenance_start_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "metadata_configuration": {
              "block": {
                "attributes": {
                  "storage_capacity": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 8,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func AwsFsxFileCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFsxFileCache), &result)
	return &result
}

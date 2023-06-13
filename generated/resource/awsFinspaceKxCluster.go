package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFinspaceKxCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "az_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "command_line_arguments": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "created_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "execution_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "initialization_script": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "release_label": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "auto_scaling_configuration": {
        "block": {
          "attributes": {
            "auto_scaling_metric": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "max_node_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "metric_target": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_node_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "scale_in_cooldown_seconds": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "scale_out_cooldown_seconds": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cache_storage_configurations": {
        "block": {
          "attributes": {
            "size": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "capacity_configuration": {
        "block": {
          "attributes": {
            "node_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "node_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "code": {
        "block": {
          "attributes": {
            "s3_bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_object_version": {
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
      "database": {
        "block": {
          "attributes": {
            "changeset_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "database_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "cache_configurations": {
              "block": {
                "attributes": {
                  "cache_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "db_paths": {
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "savedown_storage_configuration": {
        "block": {
          "attributes": {
            "size": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "type": {
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
      },
      "vpc_configuration": {
        "block": {
          "attributes": {
            "ip_address_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "security_group_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "vpc_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsFinspaceKxClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFinspaceKxCluster), &result)
	return &result
}

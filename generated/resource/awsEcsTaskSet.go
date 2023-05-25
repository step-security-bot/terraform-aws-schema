package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsTaskSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "external_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "launch_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "platform_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stability_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
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
      "task_definition": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_set_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "wait_until_stable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "wait_until_stable_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "capacity_provider_strategy": {
        "block": {
          "attributes": {
            "base": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "capacity_provider": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weight": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "load_balancer": {
        "block": {
          "attributes": {
            "container_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "container_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "load_balancer_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_group_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "network_configuration": {
        "block": {
          "attributes": {
            "assign_public_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "security_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnets": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "scale": {
        "block": {
          "attributes": {
            "unit": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "service_registries": {
        "block": {
          "attributes": {
            "container_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "registry_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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

func AwsEcsTaskSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsTaskSet), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2Fleet = `{
  "block": {
    "attributes": {
      "excess_capacity_termination_policy": {
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
      "replace_unhealthy_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "terminate_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "terminate_instances_with_expiration": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "launch_template_config": {
        "block": {
          "block_types": {
            "launch_template_specification": {
              "block": {
                "attributes": {
                  "launch_template_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "launch_template_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
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
            "override": {
              "block": {
                "attributes": {
                  "availability_zone": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_price": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "weighted_capacity": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 50,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "on_demand_options": {
        "block": {
          "attributes": {
            "allocation_strategy": {
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
      "spot_options": {
        "block": {
          "attributes": {
            "allocation_strategy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_interruption_behavior": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_pools_to_use_count": {
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
      "target_capacity_specification": {
        "block": {
          "attributes": {
            "default_target_capacity_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "on_demand_target_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "spot_target_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "total_target_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func AwsEc2FleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2Fleet), &result)
	return &result
}

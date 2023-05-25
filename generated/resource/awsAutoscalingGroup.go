package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAutoscalingGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zones": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "default_cooldown": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "desired_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enabled_metrics": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "force_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "health_check_grace_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_type": {
        "computed": true,
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
      "launch_configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "max_instance_lifetime": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "metrics_granularity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_elb_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "placement_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protect_from_scale_in": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "service_linked_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "suspended_processes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          [
            "map",
            "string"
          ]
        ]
      },
      "target_group_arns": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "termination_policies": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_zone_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "wait_for_capacity_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "wait_for_elb_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "initial_lifecycle_hook": {
        "block": {
          "attributes": {
            "default_result": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "heartbeat_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "lifecycle_transition": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "notification_metadata": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "notification_target_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "launch_template": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
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
      "mixed_instances_policy": {
        "block": {
          "block_types": {
            "instances_distribution": {
              "block": {
                "attributes": {
                  "on_demand_allocation_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "on_demand_base_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "on_demand_percentage_above_base_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "spot_allocation_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spot_instance_pools": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "spot_max_price": {
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
            "launch_template": {
              "block": {
                "block_types": {
                  "launch_template_specification": {
                    "block": {
                      "attributes": {
                        "launch_template_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "launch_template_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "version": {
                          "description_kind": "plain",
                          "optional": true,
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
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "weighted_capacity": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "tag": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "propagate_at_launch": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "delete": {
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

func AwsAutoscalingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAutoscalingGroup), &result)
	return &result
}

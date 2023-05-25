package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAutoscalingplansScalingPlan = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scaling_plan_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "application_source": {
        "block": {
          "attributes": {
            "cloudformation_stack_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "tag_filter": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 50,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "scaling_instruction": {
        "block": {
          "attributes": {
            "disable_dynamic_scaling": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "predictive_scaling_max_capacity_behavior": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "predictive_scaling_max_capacity_buffer": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "predictive_scaling_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scalable_dimension": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scaling_policy_update_behavior": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scheduled_action_buffer_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "service_namespace": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "customized_load_metric_specification": {
              "block": {
                "attributes": {
                  "dimensions": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "metric_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "statistic": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "unit": {
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
            "predefined_load_metric_specification": {
              "block": {
                "attributes": {
                  "predefined_load_metric_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "resource_label": {
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
            "target_tracking_configuration": {
              "block": {
                "attributes": {
                  "disable_scale_in": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "estimated_instance_warmup": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "scale_in_cooldown": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "scale_out_cooldown": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "target_value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "customized_scaling_metric_specification": {
                    "block": {
                      "attributes": {
                        "dimensions": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "metric_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "namespace": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "statistic": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "unit": {
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
                  "predefined_scaling_metric_specification": {
                    "block": {
                      "attributes": {
                        "predefined_scaling_metric_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "resource_label": {
                          "description_kind": "plain",
                          "optional": true,
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
              "max_items": 10,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAutoscalingplansScalingPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAutoscalingplansScalingPlan), &result)
	return &result
}

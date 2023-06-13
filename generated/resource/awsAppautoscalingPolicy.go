package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppautoscalingPolicy = `{
  "block": {
    "attributes": {
      "alarm_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_type": {
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
      "service_namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "step_scaling_policy_configuration": {
        "block": {
          "attributes": {
            "adjustment_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cooldown": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metric_aggregation_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_adjustment_magnitude": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "step_adjustment": {
              "block": {
                "attributes": {
                  "metric_interval_lower_bound": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metric_interval_upper_bound": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scaling_adjustment": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "target_tracking_scaling_policy_configuration": {
        "block": {
          "attributes": {
            "disable_scale_in": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
            "customized_metric_specification": {
              "block": {
                "attributes": {
                  "metric_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "statistic": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "unit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "dimensions": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
                  "metrics": {
                    "block": {
                      "attributes": {
                        "expression": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "label": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "return_data": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "metric_stat": {
                          "block": {
                            "attributes": {
                              "stat": {
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
                            "block_types": {
                              "metric": {
                                "block": {
                                  "attributes": {
                                    "metric_name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "namespace": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "dimensions": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
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
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "predefined_metric_specification": {
              "block": {
                "attributes": {
                  "predefined_metric_type": {
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
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAppautoscalingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppautoscalingPolicy), &result)
	return &result
}

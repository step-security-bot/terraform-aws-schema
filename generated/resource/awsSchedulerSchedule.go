package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSchedulerSchedule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_name": {
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
      "kms_key_arn": {
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
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule_expression": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_expression_timezone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "flexible_time_window": {
        "block": {
          "attributes": {
            "maximum_window_in_minutes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mode": {
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
      "target": {
        "block": {
          "attributes": {
            "arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "input": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "dead_letter_config": {
              "block": {
                "attributes": {
                  "arn": {
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
            "ecs_parameters": {
              "block": {
                "attributes": {
                  "enable_ecs_managed_tags": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_execute_command": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "group": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "launch_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "platform_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "propagate_tags": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reference_id": {
                    "description_kind": "plain",
                    "optional": true,
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
                  "task_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "task_definition_arn": {
                    "description_kind": "plain",
                    "required": true,
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
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 6,
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
                  "placement_constraints": {
                    "block": {
                      "attributes": {
                        "expression": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 10,
                    "nesting_mode": "set"
                  },
                  "placement_strategy": {
                    "block": {
                      "attributes": {
                        "field": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 5,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "eventbridge_parameters": {
              "block": {
                "attributes": {
                  "detail_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source": {
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
            "kinesis_parameters": {
              "block": {
                "attributes": {
                  "partition_key": {
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
            "retry_policy": {
              "block": {
                "attributes": {
                  "maximum_event_age_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_retry_attempts": {
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
            "sagemaker_pipeline_parameters": {
              "block": {
                "block_types": {
                  "pipeline_parameter": {
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
                    "max_items": 200,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sqs_parameters": {
              "block": {
                "attributes": {
                  "message_group_id": {
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
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSchedulerScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSchedulerSchedule), &result)
	return &result
}

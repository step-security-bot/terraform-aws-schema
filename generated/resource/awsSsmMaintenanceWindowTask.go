package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmMaintenanceWindowTask = `{
  "block": {
    "attributes": {
      "description": {
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
      "max_concurrency": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_errors": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "service_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "window_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "logging_info": {
        "block": {
          "attributes": {
            "s3_bucket_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_bucket_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_region": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "targets": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "task_invocation_parameters": {
        "block": {
          "block_types": {
            "automation_parameters": {
              "block": {
                "attributes": {
                  "document_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "parameter": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
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
            "lambda_parameters": {
              "block": {
                "attributes": {
                  "client_context": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "payload": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "qualifier": {
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
            "run_command_parameters": {
              "block": {
                "attributes": {
                  "comment": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "document_hash": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "document_hash_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_s3_bucket": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_s3_key_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timeout_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "notification_config": {
                    "block": {
                      "attributes": {
                        "notification_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "notification_events": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "notification_type": {
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
                  "parameter": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
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
            "step_functions_parameters": {
              "block": {
                "attributes": {
                  "input": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "name": {
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
      },
      "task_parameters": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsmMaintenanceWindowTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmMaintenanceWindowTask), &result)
	return &result
}

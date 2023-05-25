package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMwaaEnvironment = `{
  "block": {
    "attributes": {
      "airflow_configuration_options": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      },
      "airflow_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dag_s3_path": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment_class": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role_arn": {
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
      "kms_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_updated": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created_at": "string",
              "error": [
                "list",
                [
                  "object",
                  {
                    "error_code": "string",
                    "error_message": "string"
                  }
                ]
              ],
              "status": "string"
            }
          ]
        ]
      },
      "max_workers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_workers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "plugins_s3_object_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plugins_s3_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requirements_s3_object_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requirements_s3_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_bucket_arn": {
        "description_kind": "plain",
        "required": true,
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
      "webserver_access_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "webserver_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "weekly_maintenance_window_start": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "logging_configuration": {
        "block": {
          "block_types": {
            "dag_processing_logs": {
              "block": {
                "attributes": {
                  "cloud_watch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_level": {
                    "computed": true,
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
            "scheduler_logs": {
              "block": {
                "attributes": {
                  "cloud_watch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_level": {
                    "computed": true,
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
            "task_logs": {
              "block": {
                "attributes": {
                  "cloud_watch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_level": {
                    "computed": true,
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
            "webserver_logs": {
              "block": {
                "attributes": {
                  "cloud_watch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_level": {
                    "computed": true,
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
            "worker_logs": {
              "block": {
                "attributes": {
                  "cloud_watch_log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_level": {
                    "computed": true,
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
      "network_configuration": {
        "block": {
          "attributes": {
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

func AwsMwaaEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMwaaEnvironment), &result)
	return &result
}

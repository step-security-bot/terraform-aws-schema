package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEmrcontainersJobTemplate = `{
  "block": {
    "attributes": {
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
      "kms_key_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
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
      }
    },
    "block_types": {
      "job_template_data": {
        "block": {
          "attributes": {
            "execution_role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "job_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "release_label": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "configuration_overrides": {
              "block": {
                "block_types": {
                  "application_configuration": {
                    "block": {
                      "attributes": {
                        "classification": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "properties": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "configurations": {
                          "block": {
                            "attributes": {
                              "classification": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "properties": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 100,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 100,
                    "nesting_mode": "list"
                  },
                  "monitoring_configuration": {
                    "block": {
                      "attributes": {
                        "persistent_app_ui": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "cloud_watch_monitoring_configuration": {
                          "block": {
                            "attributes": {
                              "log_group_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "log_stream_name_prefix": {
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
                        "s3_monitoring_configuration": {
                          "block": {
                            "attributes": {
                              "log_uri": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "job_driver": {
              "block": {
                "block_types": {
                  "spark_sql_job_driver": {
                    "block": {
                      "attributes": {
                        "entry_point": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "spark_sql_parameters": {
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
                  "spark_submit_job_driver": {
                    "block": {
                      "attributes": {
                        "entry_point": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "entry_point_arguments": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "spark_submit_parameters": {
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
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
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

func AwsEmrcontainersJobTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEmrcontainersJobTemplate), &result)
	return &result
}

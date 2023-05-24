package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerEndpointConfiguration = `{
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
      "async_inference_config": {
        "block": {
          "block_types": {
            "client_config": {
              "block": {
                "attributes": {
                  "max_concurrent_invocations_per_instance": {
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
            "output_config": {
              "block": {
                "attributes": {
                  "kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_failure_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_output_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "notification_config": {
                    "block": {
                      "attributes": {
                        "error_topic": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "include_inference_response_in": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "success_topic": {
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
        "nesting_mode": "list"
      },
      "data_capture_config": {
        "block": {
          "attributes": {
            "destination_s3_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "enable_capture": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "initial_sampling_percentage": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "capture_content_type_header": {
              "block": {
                "attributes": {
                  "csv_content_types": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "json_content_types": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "capture_options": {
              "block": {
                "attributes": {
                  "capture_mode": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 2,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "production_variants": {
        "block": {
          "attributes": {
            "accelerator_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_startup_health_check_timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enable_ssm_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "initial_instance_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "initial_variant_weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_data_download_timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "model_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "variant_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_size_in_gb": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "core_dump_config": {
              "block": {
                "attributes": {
                  "destination_s3_uri": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "kms_key_id": {
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
            "serverless_config": {
              "block": {
                "attributes": {
                  "max_concurrency": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "memory_size_in_mb": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
        "nesting_mode": "list"
      },
      "shadow_production_variants": {
        "block": {
          "attributes": {
            "accelerator_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_startup_health_check_timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enable_ssm_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "initial_instance_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "initial_variant_weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_data_download_timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "model_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "variant_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_size_in_gb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "core_dump_config": {
              "block": {
                "attributes": {
                  "destination_s3_uri": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "kms_key_id": {
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
            "serverless_config": {
              "block": {
                "attributes": {
                  "max_concurrency": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "memory_size_in_mb": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSagemakerEndpointConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerEndpointConfiguration), &result)
	return &result
}

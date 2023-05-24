package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerDataQualityJobDefinition = `{
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
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
      "data_quality_app_specification": {
        "block": {
          "attributes": {
            "environment": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "image_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "post_analytics_processor_source_uri": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "record_preprocessor_source_uri": {
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
      "data_quality_baseline_config": {
        "block": {
          "block_types": {
            "constraints_resource": {
              "block": {
                "attributes": {
                  "s3_uri": {
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
            "statistics_resource": {
              "block": {
                "attributes": {
                  "s3_uri": {
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
      "data_quality_job_input": {
        "block": {
          "block_types": {
            "batch_transform_input": {
              "block": {
                "attributes": {
                  "data_captured_destination_s3_uri": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "local_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_data_distribution_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_input_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "dataset_format": {
                    "block": {
                      "block_types": {
                        "csv": {
                          "block": {
                            "attributes": {
                              "header": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "json": {
                          "block": {
                            "attributes": {
                              "line": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
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
            "endpoint_input": {
              "block": {
                "attributes": {
                  "endpoint_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "local_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_data_distribution_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_input_mode": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "data_quality_job_output_config": {
        "block": {
          "attributes": {
            "kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "monitoring_outputs": {
              "block": {
                "block_types": {
                  "s3_output": {
                    "block": {
                      "attributes": {
                        "local_path": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_upload_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
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
      "job_resources": {
        "block": {
          "block_types": {
            "cluster_config": {
              "block": {
                "attributes": {
                  "instance_count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "volume_kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_size_in_gb": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "attributes": {
            "enable_inter_container_traffic_encryption": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_network_isolation": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "vpc_config": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "stopping_condition": {
        "block": {
          "attributes": {
            "max_runtime_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
  "version": 0
}`

func AwsSagemakerDataQualityJobDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerDataQualityJobDefinition), &result)
	return &result
}

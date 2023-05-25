package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerDomain = `{
  "block": {
    "attributes": {
      "app_network_access_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auth_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "home_efs_file_system_id": {
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
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "single_sign_on_managed_application_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
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
      "url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "default_user_settings": {
        "block": {
          "attributes": {
            "execution_role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "security_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "jupyter_server_app_settings": {
              "block": {
                "attributes": {
                  "lifecycle_config_arns": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "default_resource_spec": {
                    "block": {
                      "attributes": {
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_arn": {
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
            "kernel_gateway_app_settings": {
              "block": {
                "attributes": {
                  "lifecycle_config_arns": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "custom_image": {
                    "block": {
                      "attributes": {
                        "app_image_config_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "image_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "image_version_number": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 30,
                    "nesting_mode": "list"
                  },
                  "default_resource_spec": {
                    "block": {
                      "attributes": {
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_arn": {
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
            "sharing_settings": {
              "block": {
                "attributes": {
                  "notebook_output_option": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_output_path": {
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
            "tensor_board_app_settings": {
              "block": {
                "block_types": {
                  "default_resource_spec": {
                    "block": {
                      "attributes": {
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "lifecycle_config_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sagemaker_image_version_arn": {
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
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "retention_policy": {
        "block": {
          "attributes": {
            "home_efs_file_system": {
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
  "version": 0
}`

func AwsSagemakerDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerDomain), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApprunnerService = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_scaling_configuration_arn": {
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
      "service_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_url": {
        "computed": true,
        "description_kind": "plain",
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
      }
    },
    "block_types": {
      "encryption_configuration": {
        "block": {
          "attributes": {
            "kms_key": {
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
      "health_check_configuration": {
        "block": {
          "attributes": {
            "healthy_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "unhealthy_threshold": {
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
      "instance_configuration": {
        "block": {
          "attributes": {
            "cpu": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "memory": {
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
      "source_configuration": {
        "block": {
          "attributes": {
            "auto_deployments_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "authentication_configuration": {
              "block": {
                "attributes": {
                  "access_role_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "connection_arn": {
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
            "code_repository": {
              "block": {
                "attributes": {
                  "repository_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "code_configuration": {
                    "block": {
                      "attributes": {
                        "configuration_source": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "code_configuration_values": {
                          "block": {
                            "attributes": {
                              "build_command": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "runtime": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "runtime_environment_variables": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "start_command": {
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
                  "source_code_version": {
                    "block": {
                      "attributes": {
                        "type": {
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
            "image_repository": {
              "block": {
                "attributes": {
                  "image_identifier": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "image_repository_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "image_configuration": {
                    "block": {
                      "attributes": {
                        "port": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "runtime_environment_variables": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "start_command": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApprunnerServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApprunnerService), &result)
	return &result
}

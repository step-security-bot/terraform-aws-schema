package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcsTaskDefinition = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_definitions": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cpu": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "family": {
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
      "ipc_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pid_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requires_compatibilities": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "revision": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "task_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "inference_accelerator": {
        "block": {
          "attributes": {
            "device_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "device_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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
      "proxy_configuration": {
        "block": {
          "attributes": {
            "container_name": {
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
            },
            "type": {
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
      "volume": {
        "block": {
          "attributes": {
            "host_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "docker_volume_configuration": {
              "block": {
                "attributes": {
                  "autoprovision": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "driver": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "driver_opts": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "labels": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "scope": {
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
            "efs_volume_configuration": {
              "block": {
                "attributes": {
                  "file_system_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "root_directory": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "transit_encryption": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "transit_encryption_port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "authorization_config": {
                    "block": {
                      "attributes": {
                        "access_point_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "iam": {
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
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsEcsTaskDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcsTaskDefinition), &result)
	return &result
}

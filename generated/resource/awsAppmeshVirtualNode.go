package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppmeshVirtualNode = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_date": {
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
      "last_updated_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mesh_name": {
        "description_kind": "plain",
        "required": true,
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
      }
    },
    "block_types": {
      "spec": {
        "block": {
          "attributes": {
            "backends": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "backend": {
              "block": {
                "block_types": {
                  "virtual_service": {
                    "block": {
                      "attributes": {
                        "virtual_service_name": {
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
              "max_items": 25,
              "nesting_mode": "set"
            },
            "listener": {
              "block": {
                "block_types": {
                  "health_check": {
                    "block": {
                      "attributes": {
                        "healthy_threshold": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "interval_millis": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "path": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "protocol": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timeout_millis": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "unhealthy_threshold": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "port_mapping": {
                    "block": {
                      "attributes": {
                        "port": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "protocol": {
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
            "logging": {
              "block": {
                "block_types": {
                  "access_log": {
                    "block": {
                      "block_types": {
                        "file": {
                          "block": {
                            "attributes": {
                              "path": {
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
            "service_discovery": {
              "block": {
                "block_types": {
                  "aws_cloud_map": {
                    "block": {
                      "attributes": {
                        "attributes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "namespace_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "service_name": {
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
                  "dns": {
                    "block": {
                      "attributes": {
                        "hostname": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "service_name": {
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
  "version": 1
}`

func AwsAppmeshVirtualNodeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppmeshVirtualNode), &result)
	return &result
}

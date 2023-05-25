package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppmeshRoute = `{
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
      },
      "virtual_router_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "spec": {
        "block": {
          "attributes": {
            "priority": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "http_route": {
              "block": {
                "block_types": {
                  "action": {
                    "block": {
                      "block_types": {
                        "weighted_target": {
                          "block": {
                            "attributes": {
                              "virtual_node": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "weight": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 10,
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "match": {
                    "block": {
                      "attributes": {
                        "method": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "prefix": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scheme": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "header": {
                          "block": {
                            "attributes": {
                              "invert": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "match": {
                                "block": {
                                  "attributes": {
                                    "exact": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "prefix": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "regex": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "suffix": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "range": {
                                      "block": {
                                        "attributes": {
                                          "end": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "start": {
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
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 10,
                          "nesting_mode": "set"
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
            "tcp_route": {
              "block": {
                "block_types": {
                  "action": {
                    "block": {
                      "block_types": {
                        "weighted_target": {
                          "block": {
                            "attributes": {
                              "virtual_node": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "weight": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 10,
                          "min_items": 1,
                          "nesting_mode": "set"
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

func AwsAppmeshRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppmeshRoute), &result)
	return &result
}

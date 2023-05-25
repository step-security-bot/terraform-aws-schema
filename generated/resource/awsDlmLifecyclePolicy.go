package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDlmLifecyclePolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "required": true,
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
      "state": {
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
      "policy_details": {
        "block": {
          "attributes": {
            "resource_types": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "target_tags": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "schedule": {
              "block": {
                "attributes": {
                  "copy_tags": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "tags_to_add": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "create_rule": {
                    "block": {
                      "attributes": {
                        "interval": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "interval_unit": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "times": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "cross_region_copy_rule": {
                    "block": {
                      "attributes": {
                        "cmk_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "copy_tags": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "encrypted": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "target": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "deprecate_rule": {
                          "block": {
                            "attributes": {
                              "interval": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "interval_unit": {
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
                        "retain_rule": {
                          "block": {
                            "attributes": {
                              "interval": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "interval_unit": {
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
                    "max_items": 3,
                    "nesting_mode": "set"
                  },
                  "retain_rule": {
                    "block": {
                      "attributes": {
                        "count": {
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
  "version": 0
}`

func AwsDlmLifecyclePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDlmLifecyclePolicy), &result)
	return &result
}

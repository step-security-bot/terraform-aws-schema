package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53TrafficPolicyDocument = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "json": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "record_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_endpoint": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_rule": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "endpoint": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "region": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "rule": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "geo_proximity_location": {
              "block": {
                "attributes": {
                  "bias": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "endpoint_reference": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "evaluate_target_health": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "health_check": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "latitude": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "longitude": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rule_reference": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "items": {
              "block": {
                "attributes": {
                  "endpoint_reference": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "health_check": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "location": {
              "block": {
                "attributes": {
                  "continent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "country": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "endpoint_reference": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "evaluate_target_health": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "health_check": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "is_default": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "rule_reference": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subdivision": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "primary": {
              "block": {
                "attributes": {
                  "endpoint_reference": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "evaluate_target_health": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "health_check": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rule_reference": {
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
            "region": {
              "block": {
                "attributes": {
                  "endpoint_reference": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "evaluate_target_health": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "health_check": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rule_reference": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "secondary": {
              "block": {
                "attributes": {
                  "endpoint_reference": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "evaluate_target_health": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "health_check": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rule_reference": {
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
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRoute53TrafficPolicyDocumentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53TrafficPolicyDocument), &result)
	return &result
}

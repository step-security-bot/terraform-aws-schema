package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApiGatewayUsagePlan = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_code": {
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
      "api_stages": {
        "block": {
          "attributes": {
            "api_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "stage": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "throttle": {
              "block": {
                "attributes": {
                  "burst_limit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rate_limit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "quota_settings": {
        "block": {
          "attributes": {
            "limit": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "offset": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "period": {
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
      "throttle_settings": {
        "block": {
          "attributes": {
            "burst_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rate_limit": {
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

func AwsApiGatewayUsagePlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApiGatewayUsagePlan), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCodedeployDeploymentConfig = `{
  "block": {
    "attributes": {
      "compute_platform": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_config_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "minimum_healthy_hosts": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
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
      "traffic_routing_config": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "time_based_canary": {
              "block": {
                "attributes": {
                  "interval": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "percentage": {
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
            "time_based_linear": {
              "block": {
                "attributes": {
                  "interval": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "percentage": {
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
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCodedeployDeploymentConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCodedeployDeploymentConfig), &result)
	return &result
}

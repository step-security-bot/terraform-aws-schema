package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEmrInstanceFleet = `{
  "block": {
    "attributes": {
      "cluster_id": {
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
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioned_on_demand_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "provisioned_spot_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "target_on_demand_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "target_spot_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "instance_type_configs": {
        "block": {
          "attributes": {
            "bid_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bid_price_as_percentage_of_on_demand_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weighted_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "configurations": {
              "block": {
                "attributes": {
                  "classification": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "properties": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "ebs_config": {
              "block": {
                "attributes": {
                  "iops": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "size": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "volumes_per_instance": {
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
      "launch_specifications": {
        "block": {
          "block_types": {
            "on_demand_specification": {
              "block": {
                "attributes": {
                  "allocation_strategy": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "spot_specification": {
              "block": {
                "attributes": {
                  "allocation_strategy": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "block_duration_minutes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout_action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout_duration_minutes": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
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

func AwsEmrInstanceFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEmrInstanceFleet), &result)
	return &result
}

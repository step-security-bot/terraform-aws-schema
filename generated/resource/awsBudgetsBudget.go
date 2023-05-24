package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBudgetsBudget = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "budget_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cost_filters": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "limit_amount": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "limit_unit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_period_end": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_period_start": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_unit": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "auto_adjust_data": {
        "block": {
          "attributes": {
            "auto_adjust_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "last_auto_adjust_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "historical_options": {
              "block": {
                "attributes": {
                  "budget_adjustment_period": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "lookback_available_periods": {
                    "computed": true,
                    "description_kind": "plain",
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
      },
      "cost_filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "cost_types": {
        "block": {
          "attributes": {
            "include_credit": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_discount": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_other_subscription": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_recurring": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_refund": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_subscription": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_support": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_tax": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_upfront": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_amortized": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_blended": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "notification": {
        "block": {
          "attributes": {
            "comparison_operator": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "notification_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "subscriber_email_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subscriber_sns_topic_arns": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "threshold": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "threshold_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "planned_limit": {
        "block": {
          "attributes": {
            "amount": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "start_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "unit": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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

func AwsBudgetsBudgetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBudgetsBudget), &result)
	return &result
}

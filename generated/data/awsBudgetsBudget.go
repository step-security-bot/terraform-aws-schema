package data

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
      "auto_adjust_data": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_adjust_type": "string",
              "historical_options": [
                "list",
                [
                  "object",
                  {
                    "budget_adjustment_period": "number",
                    "lookback_available_periods": "number"
                  }
                ]
              ],
              "last_auto_adjust_time": "string"
            }
          ]
        ]
      },
      "budget_exceeded": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "budget_limit": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "amount": "string",
              "unit": "string"
            }
          ]
        ]
      },
      "budget_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "calculated_spend": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "actual_spend": [
                "list",
                [
                  "object",
                  {
                    "amount": "string",
                    "unit": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "cost_filter": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "values": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "cost_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "include_credit": "bool",
              "include_discount": "bool",
              "include_other_subscription": "bool",
              "include_recurring": "bool",
              "include_refund": "bool",
              "include_subscription": "bool",
              "include_support": "bool",
              "include_tax": "bool",
              "include_upfront": "bool",
              "use_amortized": "bool",
              "use_blended": "bool"
            }
          ]
        ]
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
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notification": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "comparison_operator": "string",
              "notification_type": "string",
              "subscriber_email_addresses": [
                "set",
                "string"
              ],
              "subscriber_sns_topic_arns": [
                "set",
                "string"
              ],
              "threshold": "number",
              "threshold_type": "string"
            }
          ]
        ]
      },
      "planned_limit": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "amount": "string",
              "start_time": "string",
              "unit": "string"
            }
          ]
        ]
      },
      "time_period_end": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_period_start": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_unit": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

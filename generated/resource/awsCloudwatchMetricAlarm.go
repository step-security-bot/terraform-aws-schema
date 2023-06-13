package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchMetricAlarm = `{
  "block": {
    "attributes": {
      "actions_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "alarm_actions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "alarm_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alarm_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "comparison_operator": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "datapoints_to_alarm": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "dimensions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "evaluate_low_sample_count_percentiles": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "evaluation_periods": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "extended_statistic": {
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
      "insufficient_data_actions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "metric_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ok_actions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "statistic": {
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
      },
      "threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "threshold_metric_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "treat_missing_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "metric_query": {
        "block": {
          "attributes": {
            "account_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "label": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "period": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "return_data": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "metric": {
              "block": {
                "attributes": {
                  "dimensions": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "metric_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "period": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "stat": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "unit": {
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
  "version": 1
}`

func AwsCloudwatchMetricAlarmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchMetricAlarm), &result)
	return &result
}

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicequotasServiceQuota = `{
  "block": {
    "attributes": {
      "adjustable": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_value": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "global_quota": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_code": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quota_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "usage_metric": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "metric_dimensions": [
                "list",
                [
                  "object",
                  {
                    "class": "string",
                    "resource": "string",
                    "service": "string",
                    "type": "string"
                  }
                ]
              ],
              "metric_name": "string",
              "metric_namespace": "string",
              "metric_statistic_recommendation": "string"
            }
          ]
        ]
      },
      "value": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsServicequotasServiceQuotaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicequotasServiceQuota), &result)
	return &result
}

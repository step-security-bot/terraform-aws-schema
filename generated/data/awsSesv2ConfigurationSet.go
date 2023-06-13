package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSesv2ConfigurationSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_set_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "delivery_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "sending_pool_name": "string",
              "tls_policy": "string"
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
      "reputation_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "last_fresh_start": "string",
              "reputation_metrics_enabled": "bool"
            }
          ]
        ]
      },
      "sending_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "sending_enabled": "bool"
            }
          ]
        ]
      },
      "suppression_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "suppressed_reasons": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tracking_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_redirect_domain": "string"
            }
          ]
        ]
      },
      "vdm_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dashboard_options": [
                "list",
                [
                  "object",
                  {
                    "engagement_metrics": "string"
                  }
                ]
              ],
              "guardian_options": [
                "list",
                [
                  "object",
                  {
                    "optimized_shared_delivery": "string"
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSesv2ConfigurationSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSesv2ConfigurationSet), &result)
	return &result
}

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmincidentsResponsePlan = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ssm_automation": [
                "list",
                [
                  "object",
                  {
                    "document_name": "string",
                    "document_version": "string",
                    "dynamic_parameters": [
                      "map",
                      "string"
                    ],
                    "parameter": [
                      "set",
                      [
                        "object",
                        {
                          "name": "string",
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "role_arn": "string",
                    "target_account": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "chat_channel": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engagements": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "incident_template": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dedupe_string": "string",
              "impact": "number",
              "incident_tags": [
                "map",
                "string"
              ],
              "notification_target": [
                "set",
                [
                  "object",
                  {
                    "sns_topic_arn": "string"
                  }
                ]
              ],
              "summary": "string",
              "title": "string"
            }
          ]
        ]
      },
      "integration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "pagerduty": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "secret_id": "string",
                    "service_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
  "version": 0
}`

func AwsSsmincidentsResponsePlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmincidentsResponsePlan), &result)
	return &result
}

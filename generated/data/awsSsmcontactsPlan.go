package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmcontactsPlan = `{
  "block": {
    "attributes": {
      "contact_id": {
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
      "stage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "duration_in_minutes": "number",
              "target": [
                "list",
                [
                  "object",
                  {
                    "channel_target_info": [
                      "list",
                      [
                        "object",
                        {
                          "contact_channel_id": "string",
                          "retry_interval_in_minutes": "number"
                        }
                      ]
                    ],
                    "contact_target_info": [
                      "list",
                      [
                        "object",
                        {
                          "contact_id": "string",
                          "is_essential": "bool"
                        }
                      ]
                    ]
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

func AwsSsmcontactsPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmcontactsPlan), &result)
	return &result
}

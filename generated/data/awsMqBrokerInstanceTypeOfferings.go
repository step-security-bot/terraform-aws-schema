package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMqBrokerInstanceTypeOfferings = `{
  "block": {
    "attributes": {
      "broker_instance_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_zones": [
                "set",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "engine_type": "string",
              "host_instance_type": "string",
              "storage_type": "string",
              "supported_deployment_modes": [
                "set",
                "string"
              ],
              "supported_engine_versions": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "engine_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_instance_type": {
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
      "storage_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMqBrokerInstanceTypeOfferingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMqBrokerInstanceTypeOfferings), &result)
	return &result
}

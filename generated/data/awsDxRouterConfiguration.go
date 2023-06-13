package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDxRouterConfiguration = `{
  "block": {
    "attributes": {
      "customer_router_config": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "platform": "string",
              "router_type_identifier": "string",
              "software": "string",
              "vendor": "string",
              "xslt_template_name": "string",
              "xslt_template_name_for_mac_sec": "string"
            }
          ]
        ]
      },
      "router_type_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_interface_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_interface_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDxRouterConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDxRouterConfiguration), &result)
	return &result
}

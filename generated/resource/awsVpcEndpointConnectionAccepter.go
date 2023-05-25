package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcEndpointConnectionAccepter = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_endpoint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_endpoint_service_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_endpoint_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpcEndpointConnectionAccepterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcEndpointConnectionAccepter), &result)
	return &result
}

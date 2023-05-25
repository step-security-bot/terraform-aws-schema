package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDxHostedConnection = `{
  "block": {
    "attributes": {
      "aws_device": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bandwidth": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "has_logical_redundancy": {
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
      "jumbo_frame_capable": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "lag_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "loa_issue_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "partner_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provider_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vlan": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDxHostedConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDxHostedConnection), &result)
	return &result
}

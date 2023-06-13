package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConnectInstance = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_resolve_best_voices_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "contact_flow_logs_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "contact_lens_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "early_media_enabled": {
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
      "identity_management_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "inbound_calls_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "instance_alias": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_party_conference_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "outbound_calls_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "service_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConnectInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConnectInstance), &result)
	return &result
}

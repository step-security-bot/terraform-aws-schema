package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDevicefarmNetworkProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "downlink_bandwidth_bits": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "downlink_delay_ms": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "downlink_jitter_ms": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "downlink_loss_percent": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_arn": {
        "description_kind": "plain",
        "required": true,
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
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uplink_bandwidth_bits": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "uplink_delay_ms": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "uplink_jitter_ms": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "uplink_loss_percent": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDevicefarmNetworkProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDevicefarmNetworkProfile), &result)
	return &result
}

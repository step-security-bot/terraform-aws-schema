package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsStoragegatewayStoredIscsiVolume = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "chap_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "disk_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gateway_arn": {
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
      "kms_encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lun_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "network_interface_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_interface_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "preserve_existing_data": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "snapshot_id": {
        "description_kind": "plain",
        "optional": true,
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
      "target_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume_attachment_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_size_in_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "volume_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsStoragegatewayStoredIscsiVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsStoragegatewayStoredIscsiVolume), &result)
	return &result
}

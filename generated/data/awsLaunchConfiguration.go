package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLaunchConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associate_public_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ebs_block_device": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "delete_on_termination": "bool",
              "device_name": "string",
              "encrypted": "bool",
              "iops": "number",
              "no_device": "bool",
              "snapshot_id": "string",
              "throughput": "number",
              "volume_size": "number",
              "volume_type": "string"
            }
          ]
        ]
      },
      "ebs_optimized": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_monitoring": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ephemeral_block_device": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "device_name": "string",
              "virtual_name": "string"
            }
          ]
        ]
      },
      "iam_instance_profile": {
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
      "image_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metadata_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "http_endpoint": "string",
              "http_put_response_hop_limit": "number",
              "http_tokens": "string"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "placement_tenancy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_block_device": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "delete_on_termination": "bool",
              "encrypted": "bool",
              "iops": "number",
              "throughput": "number",
              "volume_size": "number",
              "volume_type": "string"
            }
          ]
        ]
      },
      "security_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "spot_price": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLaunchConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLaunchConfiguration), &result)
	return &result
}

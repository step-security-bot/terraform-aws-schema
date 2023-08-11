package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAmi = `{
  "block": {
    "attributes": {
      "architecture": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "block_device_mappings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "device_name": "string",
              "ebs": [
                "map",
                "string"
              ],
              "no_device": "string",
              "virtual_name": "string"
            }
          ]
        ]
      },
      "boot_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deprecation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ena_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "executable_users": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "hypervisor": {
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
      "image_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_owner_alias": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "imds_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "include_deprecated": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kernel_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "most_recent": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "owners": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform_details": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_codes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "product_code_id": "string",
              "product_code_type": "string"
            }
          ]
        ]
      },
      "public": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ramdisk_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_device_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_device_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_snapshot_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sriov_net_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tpm_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "usage_operation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "virtualization_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAmiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAmi), &result)
	return &result
}

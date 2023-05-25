package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAmiFromInstance = `{
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
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ena_support": {
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
      "image_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kernel_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "manage_ebs_snapshots": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "root_snapshot_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_without_reboot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sriov_net_support": {
        "computed": true,
        "description_kind": "plain",
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
      "virtualization_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "ebs_block_device": {
        "block": {
          "attributes": {
            "delete_on_termination": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "device_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "encrypted": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "iops": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "snapshot_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "volume_size": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "volume_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "ephemeral_block_device": {
        "block": {
          "attributes": {
            "device_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "virtual_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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

func AwsAmiFromInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAmiFromInstance), &result)
	return &result
}

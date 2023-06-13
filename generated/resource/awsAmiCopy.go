package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAmiCopy = `{
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
      "boot_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deprecation_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_outpost_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ena_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "kernel_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "root_snapshot_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_ami_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_ami_region": {
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
      "tags_all": {
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
            "outpost_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "snapshot_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "throughput": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
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

func AwsAmiCopySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAmiCopy), &result)
	return &result
}

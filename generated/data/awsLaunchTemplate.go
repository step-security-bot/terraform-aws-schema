package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLaunchTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "block_device_mappings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "device_name": "string",
              "ebs": [
                "list",
                [
                  "object",
                  {
                    "delete_on_termination": "string",
                    "encrypted": "string",
                    "iops": "number",
                    "kms_key_id": "string",
                    "snapshot_id": "string",
                    "volume_size": "number",
                    "volume_type": "string"
                  }
                ]
              ],
              "no_device": "string",
              "virtual_name": "string"
            }
          ]
        ]
      },
      "credit_specification": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cpu_credits": "string"
            }
          ]
        ]
      },
      "default_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disable_api_termination": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ebs_optimized": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "elastic_gpu_specifications": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "type": "string"
            }
          ]
        ]
      },
      "hibernation_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "configured": "bool"
            }
          ]
        ]
      },
      "iam_instance_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "name": "string"
            }
          ]
        ]
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
      "instance_initiated_shutdown_behavior": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_market_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "market_type": "string",
              "spot_options": [
                "list",
                [
                  "object",
                  {
                    "block_duration_minutes": "number",
                    "instance_interruption_behavior": "string",
                    "max_price": "string",
                    "spot_instance_type": "string",
                    "valid_until": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kernel_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "latest_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "monitoring": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_interfaces": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "associate_public_ip_address": "string",
              "delete_on_termination": "bool",
              "description": "string",
              "device_index": "number",
              "ipv4_address_count": "number",
              "ipv4_addresses": [
                "set",
                "string"
              ],
              "ipv6_address_count": "number",
              "ipv6_addresses": [
                "set",
                "string"
              ],
              "network_interface_id": "string",
              "private_ip_address": "string",
              "security_groups": [
                "set",
                "string"
              ],
              "subnet_id": "string"
            }
          ]
        ]
      },
      "placement": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "affinity": "string",
              "availability_zone": "string",
              "group_name": "string",
              "host_id": "string",
              "partition_number": "number",
              "spread_domain": "string",
              "tenancy": "string"
            }
          ]
        ]
      },
      "ram_disk_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tag_specifications": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resource_type": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
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
      "user_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
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
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLaunchTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLaunchTemplate), &result)
	return &result
}

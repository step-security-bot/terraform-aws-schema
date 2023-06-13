package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsInstance = `{
  "block": {
    "attributes": {
      "ami": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
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
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "disable_api_stop": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "disable_api_termination": {
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
              "kms_key_id": "string",
              "snapshot_id": "string",
              "tags": [
                "map",
                "string"
              ],
              "throughput": "number",
              "volume_id": "string",
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
      "enclave_options": {
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
      "ephemeral_block_device": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "device_name": "string",
              "no_device": "bool",
              "virtual_name": "string"
            }
          ]
        ]
      },
      "get_password_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "get_user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "host_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "host_resource_group_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ipv6_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "key_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_recovery": "string"
            }
          ]
        ]
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
              "http_tokens": "string",
              "instance_metadata_tags": "string"
            }
          ]
        ]
      },
      "monitoring": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "network_interface_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "password_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "placement_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "placement_partition_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "private_dns": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_name_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_resource_name_dns_a_record": "bool",
              "enable_resource_name_dns_aaaa_record": "bool",
              "hostname_type": "string"
            }
          ]
        ]
      },
      "private_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_dns": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_block_device": {
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
              "kms_key_id": "string",
              "tags": [
                "map",
                "string"
              ],
              "throughput": "number",
              "volume_id": "string",
              "volume_size": "number",
              "volume_type": "string"
            }
          ]
        ]
      },
      "secondary_private_ips": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
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
      "source_dest_check": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "tenancy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_data_base64": {
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
  "version": 1
}`

func AwsInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsInstance), &result)
	return &result
}

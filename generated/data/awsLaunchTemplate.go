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
                    "throughput": "number",
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
      "capacity_reservation_specification": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity_reservation_preference": "string",
              "capacity_reservation_target": [
                "list",
                [
                  "object",
                  {
                    "capacity_reservation_id": "string",
                    "capacity_reservation_resource_group_arn": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "cpu_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "amd_sev_snp": "string",
              "core_count": "number",
              "threads_per_core": "number"
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
      "elastic_inference_accelerator": {
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
      "instance_requirements": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accelerator_count": [
                "list",
                [
                  "object",
                  {
                    "max": "number",
                    "min": "number"
                  }
                ]
              ],
              "accelerator_manufacturers": [
                "set",
                "string"
              ],
              "accelerator_names": [
                "set",
                "string"
              ],
              "accelerator_total_memory_mib": [
                "list",
                [
                  "object",
                  {
                    "max": "number",
                    "min": "number"
                  }
                ]
              ],
              "accelerator_types": [
                "set",
                "string"
              ],
              "allowed_instance_types": [
                "set",
                "string"
              ],
              "bare_metal": "string",
              "baseline_ebs_bandwidth_mbps": [
                "list",
                [
                  "object",
                  {
                    "max": "number",
                    "min": "number"
                  }
                ]
              ],
              "burstable_performance": "string",
              "cpu_manufacturers": [
                "set",
                "string"
              ],
              "excluded_instance_types": [
                "set",
                "string"
              ],
              "instance_generations": [
                "set",
                "string"
              ],
              "local_storage": "string",
              "local_storage_types": [
                "set",
                "string"
              ],
              "memory_gib_per_vcpu": [
                "list",
                [
                  "object",
                  {
                    "max": "number",
                    "min": "number"
                  }
                ]
              ],
              "memory_mib": [
                "list",
                [
                  "object",
                  {
                    "max": "number",
                    "min": "number"
                  }
                ]
              ],
              "network_bandwidth_gbps": [
                "list",
                [
                  "object",
                  {
                    "max": "number",
                    "min": "number"
                  }
                ]
              ],
              "network_interface_count": [
                "list",
                [
                  "object",
                  {
                    "max": "number",
                    "min": "number"
                  }
                ]
              ],
              "on_demand_max_price_percentage_over_lowest_price": "number",
              "require_hibernate_support": "bool",
              "spot_max_price_percentage_over_lowest_price": "number",
              "total_local_storage_gb": [
                "list",
                [
                  "object",
                  {
                    "max": "number",
                    "min": "number"
                  }
                ]
              ],
              "vcpu_count": [
                "list",
                [
                  "object",
                  {
                    "max": "number",
                    "min": "number"
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
      "license_specification": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "license_configuration_arn": "string"
            }
          ]
        ]
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
              "http_protocol_ipv6": "string",
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
        "computed": true,
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
              "associate_carrier_ip_address": "string",
              "associate_public_ip_address": "string",
              "delete_on_termination": "string",
              "description": "string",
              "device_index": "number",
              "interface_type": "string",
              "ipv4_address_count": "number",
              "ipv4_addresses": [
                "set",
                "string"
              ],
              "ipv4_prefix_count": "number",
              "ipv4_prefixes": [
                "set",
                "string"
              ],
              "ipv6_address_count": "number",
              "ipv6_addresses": [
                "set",
                "string"
              ],
              "ipv6_prefix_count": "number",
              "ipv6_prefixes": [
                "set",
                "string"
              ],
              "network_card_index": "number",
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
              "host_resource_group_arn": "string",
              "partition_number": "number",
              "spread_domain": "string",
              "tenancy": "string"
            }
          ]
        ]
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

func AwsLaunchTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLaunchTemplate), &result)
	return &result
}

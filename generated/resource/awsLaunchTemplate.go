package resource

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
      "default_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_api_termination": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ebs_optimized": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_initiated_shutdown_behavior": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kernel_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "latest_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ram_disk_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "update_default_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_security_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "block_device_mappings": {
        "block": {
          "attributes": {
            "device_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "no_device": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "virtual_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ebs": {
              "block": {
                "attributes": {
                  "delete_on_termination": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "encrypted": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "iops": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "kms_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "snapshot_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "volume_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "volume_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "capacity_reservation_specification": {
        "block": {
          "attributes": {
            "capacity_reservation_preference": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "capacity_reservation_target": {
              "block": {
                "attributes": {
                  "capacity_reservation_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cpu_options": {
        "block": {
          "attributes": {
            "core_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "threads_per_core": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "credit_specification": {
        "block": {
          "attributes": {
            "cpu_credits": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "elastic_gpu_specifications": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "elastic_inference_accelerator": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "hibernation_options": {
        "block": {
          "attributes": {
            "configured": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "iam_instance_profile": {
        "block": {
          "attributes": {
            "arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "instance_market_options": {
        "block": {
          "attributes": {
            "market_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "spot_options": {
              "block": {
                "attributes": {
                  "block_duration_minutes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_interruption_behavior": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_price": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "spot_instance_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "valid_until": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "license_specification": {
        "block": {
          "attributes": {
            "license_configuration_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "metadata_options": {
        "block": {
          "attributes": {
            "http_endpoint": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_put_response_hop_limit": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "http_tokens": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "monitoring": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_interfaces": {
        "block": {
          "attributes": {
            "associate_public_ip_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete_on_termination": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_index": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipv4_address_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipv4_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "ipv6_address_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipv6_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "network_interface_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_ip_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "placement": {
        "block": {
          "attributes": {
            "affinity": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "availability_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "group_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "partition_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "spread_domain": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tenancy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "tag_specifications": {
        "block": {
          "attributes": {
            "resource_type": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
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

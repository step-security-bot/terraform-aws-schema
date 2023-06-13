package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpsworksPhpAppLayer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_assign_elastic_ips": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_assign_public_ips": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_healing": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "custom_configure_recipes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "custom_deploy_recipes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "custom_instance_profile_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_json": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_security_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "custom_setup_recipes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "custom_shutdown_recipes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "custom_undeploy_recipes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "drain_elb_on_shutdown": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "elastic_load_balancer": {
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
      "install_updates_on_boot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance_shutdown_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stack_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "system_packages": {
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
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "use_ebs_optimized_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "cloudwatch_configuration": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "log_streams": {
              "block": {
                "attributes": {
                  "batch_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "batch_size": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "buffer_duration": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "datetime_format": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "encoding": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "file_fingerprint_lines": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "initial_position": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_group_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "multiline_start_pattern": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "time_zone": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ebs_volume": {
        "block": {
          "attributes": {
            "encrypted": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "iops": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mount_point": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "number_of_disks": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "raid_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "size": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "load_based_auto_scaling": {
        "block": {
          "attributes": {
            "enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "downscaling": {
              "block": {
                "attributes": {
                  "alarms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cpu_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ignore_metrics_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "load_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "memory_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "thresholds_wait_time": {
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
            "upscaling": {
              "block": {
                "attributes": {
                  "alarms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cpu_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ignore_metrics_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "load_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "memory_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "thresholds_wait_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpsworksPhpAppLayerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpsworksPhpAppLayer), &result)
	return &result
}

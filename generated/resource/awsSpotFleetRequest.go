package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSpotFleetRequest = `{
  "block": {
    "attributes": {
      "allocation_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "context": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "excess_capacity_termination_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fleet_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_fleet_role": {
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
      "instance_interruption_behaviour": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_pools_to_use_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "load_balancers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "on_demand_allocation_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "on_demand_max_total_price": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "on_demand_target_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "replace_unhealthy_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "spot_price": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_request_state": {
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
      "target_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "target_capacity_unit_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_group_arns": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "terminate_instances_on_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "terminate_instances_with_expiration": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "valid_from": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "valid_until": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "wait_for_fulfillment": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "launch_specification": {
        "block": {
          "attributes": {
            "ami": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "associate_public_ip_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "availability_zone": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ebs_optimized": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "iam_instance_profile": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "iam_instance_profile_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "monitoring": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "placement_group": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "placement_tenancy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "spot_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_id": {
              "computed": true,
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
            "user_data": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "weighted_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ebs_block_device": {
              "block": {
                "attributes": {
                  "delete_on_termination": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "device_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "encrypted": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "iops": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "snapshot_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "throughput": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
              "nesting_mode": "set"
            },
            "ephemeral_block_device": {
              "block": {
                "attributes": {
                  "device_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "virtual_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "root_block_device": {
              "block": {
                "attributes": {
                  "delete_on_termination": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "encrypted": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "iops": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "throughput": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "launch_template_config": {
        "block": {
          "block_types": {
            "launch_template_specification": {
              "block": {
                "attributes": {
                  "id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "overrides": {
              "block": {
                "attributes": {
                  "availability_zone": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "spot_price": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnet_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "weighted_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "instance_requirements": {
                    "block": {
                      "attributes": {
                        "accelerator_manufacturers": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "accelerator_names": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "accelerator_types": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "allowed_instance_types": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "bare_metal": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "burstable_performance": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cpu_manufacturers": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "excluded_instance_types": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "instance_generations": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "local_storage": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "local_storage_types": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "on_demand_max_price_percentage_over_lowest_price": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "require_hibernate_support": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "spot_max_price_percentage_over_lowest_price": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "accelerator_count": {
                          "block": {
                            "attributes": {
                              "max": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
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
                        "accelerator_total_memory_mib": {
                          "block": {
                            "attributes": {
                              "max": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
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
                        "baseline_ebs_bandwidth_mbps": {
                          "block": {
                            "attributes": {
                              "max": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
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
                        "memory_gib_per_vcpu": {
                          "block": {
                            "attributes": {
                              "max": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
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
                        "memory_mib": {
                          "block": {
                            "attributes": {
                              "max": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
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
                        "network_bandwidth_gbps": {
                          "block": {
                            "attributes": {
                              "max": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
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
                        "network_interface_count": {
                          "block": {
                            "attributes": {
                              "max": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
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
                        "total_local_storage_gb": {
                          "block": {
                            "attributes": {
                              "max": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
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
                        "vcpu_count": {
                          "block": {
                            "attributes": {
                              "max": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
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
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "spot_maintenance_strategies": {
        "block": {
          "block_types": {
            "capacity_rebalance": {
              "block": {
                "attributes": {
                  "replacement_strategy": {
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
  "version": 1
}`

func AwsSpotFleetRequestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSpotFleetRequest), &result)
	return &result
}

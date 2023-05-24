package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2Fleet = `{
  "block": {
    "attributes": {
      "arn": {
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
      "fleet_state": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fulfilled_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "fulfilled_on_demand_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replace_unhealthy_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "terminate_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "terminate_instances_with_expiration": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "block_types": {
      "fleet_instance_set": {
        "block": {
          "attributes": {
            "instance_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "instance_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lifecycle": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "platform": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "launch_template_config": {
        "block": {
          "block_types": {
            "launch_template_specification": {
              "block": {
                "attributes": {
                  "launch_template_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "launch_template_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
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
            "override": {
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
                  "max_price": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "weighted_capacity": {
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
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
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
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
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
              "max_items": 300,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 50,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "on_demand_options": {
        "block": {
          "attributes": {
            "allocation_strategy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_total_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_target_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "single_availability_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "single_instance_type": {
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
      "spot_options": {
        "block": {
          "attributes": {
            "allocation_strategy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_interruption_behavior": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_pools_to_use_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "maintenance_strategies": {
              "block": {
                "block_types": {
                  "capacity_rebalance": {
                    "block": {
                      "attributes": {
                        "replacement_strategy": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "termination_delay": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "target_capacity_specification": {
        "block": {
          "attributes": {
            "default_target_capacity_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "on_demand_target_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "spot_target_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_capacity_unit_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "total_target_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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
  "version": 0
}`

func AwsEc2FleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2Fleet), &result)
	return &result
}

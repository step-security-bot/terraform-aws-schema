package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEmrCluster = `{
  "block": {
    "attributes": {
      "additional_info": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "applications": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "autoscaling_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configurations": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "configurations_json": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_ami_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ebs_root_volume_size": {
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
      "keep_job_flow_alive_when_no_steps": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "list_steps_states": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "log_encryption_kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_public_dns": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "placement_group_config": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "instance_role": "string",
              "placement_strategy": "string"
            }
          ]
        ]
      },
      "release_label": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scale_down_behavior": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "step": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "action_on_failure": "string",
              "hadoop_jar_step": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "jar": "string",
                    "main_class": "string",
                    "properties": [
                      "map",
                      "string"
                    ]
                  }
                ]
              ],
              "name": "string"
            }
          ]
        ]
      },
      "step_concurrency_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "termination_protection": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "visible_to_all_users": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "auto_termination_policy": {
        "block": {
          "attributes": {
            "idle_timeout": {
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
      "bootstrap_action": {
        "block": {
          "attributes": {
            "args": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "core_instance_fleet": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "provisioned_on_demand_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "provisioned_spot_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "target_on_demand_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_spot_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "instance_type_configs": {
              "block": {
                "attributes": {
                  "bid_price": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bid_price_as_percentage_of_on_demand_price": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "weighted_capacity": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "configurations": {
                    "block": {
                      "attributes": {
                        "classification": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "properties": {
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
                    "nesting_mode": "set"
                  },
                  "ebs_config": {
                    "block": {
                      "attributes": {
                        "iops": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "size": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "volumes_per_instance": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            "launch_specifications": {
              "block": {
                "block_types": {
                  "on_demand_specification": {
                    "block": {
                      "attributes": {
                        "allocation_strategy": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "spot_specification": {
                    "block": {
                      "attributes": {
                        "allocation_strategy": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "block_duration_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_action": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timeout_duration_minutes": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "core_instance_group": {
        "block": {
          "attributes": {
            "autoscaling_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bid_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ebs_config": {
              "block": {
                "attributes": {
                  "iops": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "size": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "throughput": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "volumes_per_instance": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ec2_attributes": {
        "block": {
          "attributes": {
            "additional_master_security_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "additional_slave_security_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "emr_managed_master_security_group": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "emr_managed_slave_security_group": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_profile": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_access_security_group": {
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
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "kerberos_attributes": {
        "block": {
          "attributes": {
            "ad_domain_join_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "ad_domain_join_user": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cross_realm_trust_principal_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "kdc_admin_password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "realm": {
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
      "master_instance_fleet": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "provisioned_on_demand_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "provisioned_spot_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "target_on_demand_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_spot_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "instance_type_configs": {
              "block": {
                "attributes": {
                  "bid_price": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bid_price_as_percentage_of_on_demand_price": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "weighted_capacity": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "configurations": {
                    "block": {
                      "attributes": {
                        "classification": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "properties": {
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
                    "nesting_mode": "set"
                  },
                  "ebs_config": {
                    "block": {
                      "attributes": {
                        "iops": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "size": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "volumes_per_instance": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            "launch_specifications": {
              "block": {
                "block_types": {
                  "on_demand_specification": {
                    "block": {
                      "attributes": {
                        "allocation_strategy": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "spot_specification": {
                    "block": {
                      "attributes": {
                        "allocation_strategy": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "block_duration_minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout_action": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timeout_duration_minutes": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "master_instance_group": {
        "block": {
          "attributes": {
            "bid_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ebs_config": {
              "block": {
                "attributes": {
                  "iops": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "size": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "throughput": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "volumes_per_instance": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
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

func AwsEmrClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEmrCluster), &result)
	return &result
}

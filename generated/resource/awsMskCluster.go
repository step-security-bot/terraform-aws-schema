package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMskCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "current_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enhanced_monitoring": {
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
      "kafka_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number_of_broker_nodes": {
        "description_kind": "plain",
        "required": true,
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
      "zookeeper_connect_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "broker_node_group_info": {
        "block": {
          "attributes": {
            "az_distribution": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_subnets": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "ebs_volume_size": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "instance_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "security_groups": {
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
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "client_authentication": {
        "block": {
          "block_types": {
            "tls": {
              "block": {
                "attributes": {
                  "certificate_authority_arns": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "configuration_info": {
        "block": {
          "attributes": {
            "arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "revision": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "encryption_info": {
        "block": {
          "attributes": {
            "encryption_at_rest_kms_key_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "encryption_in_transit": {
              "block": {
                "attributes": {
                  "client_broker": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "in_cluster": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      "logging_info": {
        "block": {
          "block_types": {
            "broker_logs": {
              "block": {
                "block_types": {
                  "cloudwatch_logs": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "log_group": {
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
                  "firehose": {
                    "block": {
                      "attributes": {
                        "delivery_stream": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enabled": {
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
                  "s3": {
                    "block": {
                      "attributes": {
                        "bucket": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enabled": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "prefix": {
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "open_monitoring": {
        "block": {
          "block_types": {
            "prometheus": {
              "block": {
                "block_types": {
                  "jmx_exporter": {
                    "block": {
                      "attributes": {
                        "enabled_in_broker": {
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
                  "node_exporter": {
                    "block": {
                      "attributes": {
                        "enabled_in_broker": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
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
  "version": 0
}`

func AwsMskClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMskCluster), &result)
	return &result
}

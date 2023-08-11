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
      "bootstrap_brokers_public_sasl_iam": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_public_sasl_scram": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_public_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_sasl_iam": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_sasl_scram": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_vpc_connectivity_sasl_iam": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_vpc_connectivity_sasl_scram": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_vpc_connectivity_tls": {
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
      "storage_mode": {
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
      "tags_all": {
        "computed": true,
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
      },
      "zookeeper_connect_string_tls": {
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
                "set",
                "string"
              ]
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
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "connectivity_info": {
              "block": {
                "block_types": {
                  "public_access": {
                    "block": {
                      "attributes": {
                        "type": {
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
                  "vpc_connectivity": {
                    "block": {
                      "block_types": {
                        "client_authentication": {
                          "block": {
                            "attributes": {
                              "tls": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "sasl": {
                                "block": {
                                  "attributes": {
                                    "iam": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "scram": {
                                      "computed": true,
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
            "storage_info": {
              "block": {
                "block_types": {
                  "ebs_storage_info": {
                    "block": {
                      "attributes": {
                        "volume_size": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "provisioned_throughput": {
                          "block": {
                            "attributes": {
                              "enabled": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "volume_throughput": {
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
          "attributes": {
            "unauthenticated": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "sasl": {
              "block": {
                "attributes": {
                  "iam": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "scram": {
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

func AwsMskClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMskCluster), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsPipesPipe = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desired_state": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enrichment": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "required": true,
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
      "target": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "enrichment_parameters": {
        "block": {
          "attributes": {
            "input_template": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "http_parameters": {
              "block": {
                "attributes": {
                  "header_parameters": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "path_parameter_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "query_string_parameters": {
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
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_parameters": {
        "block": {
          "block_types": {
            "activemq_broker_parameters": {
              "block": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "queue_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "credentials": {
                    "block": {
                      "attributes": {
                        "basic_auth": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
            "dynamodb_stream_parameters": {
              "block": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_record_age_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_retry_attempts": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "on_partial_batch_item_failure": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parallelization_factor": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "starting_position": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "dead_letter_config": {
                    "block": {
                      "attributes": {
                        "arn": {
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
            "filter_criteria": {
              "block": {
                "block_types": {
                  "filter": {
                    "block": {
                      "attributes": {
                        "pattern": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 5,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "kinesis_stream_parameters": {
              "block": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_record_age_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_retry_attempts": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "on_partial_batch_item_failure": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parallelization_factor": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "starting_position": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "starting_position_timestamp": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "dead_letter_config": {
                    "block": {
                      "attributes": {
                        "arn": {
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
            "managed_streaming_kafka_parameters": {
              "block": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "consumer_group_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "starting_position": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "topic_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "credentials": {
                    "block": {
                      "attributes": {
                        "client_certificate_tls_auth": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sasl_scram_512_auth": {
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
            "rabbitmq_broker_parameters": {
              "block": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "queue_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "virtual_host": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "credentials": {
                    "block": {
                      "attributes": {
                        "basic_auth": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
            "self_managed_kafka_parameters": {
              "block": {
                "attributes": {
                  "additional_bootstrap_servers": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "consumer_group_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "server_root_ca_certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "starting_position": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "topic_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "credentials": {
                    "block": {
                      "attributes": {
                        "basic_auth": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_certificate_tls_auth": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sasl_scram_256_auth": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sasl_scram_512_auth": {
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
                  "vpc": {
                    "block": {
                      "attributes": {
                        "security_groups": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "subnets": {
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
            "sqs_queue_parameters": {
              "block": {
                "attributes": {
                  "batch_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "maximum_batching_window_in_seconds": {
                    "computed": true,
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
      },
      "target_parameters": {
        "block": {
          "attributes": {
            "input_template": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "batch_job_parameters": {
              "block": {
                "attributes": {
                  "job_definition": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "job_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "parameters": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "array_properties": {
                    "block": {
                      "attributes": {
                        "size": {
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
                  "container_overrides": {
                    "block": {
                      "attributes": {
                        "command": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "instance_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "environment": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "resource_requirement": {
                          "block": {
                            "attributes": {
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
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
                  "depends_on": {
                    "block": {
                      "attributes": {
                        "job_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 20,
                    "nesting_mode": "list"
                  },
                  "retry_strategy": {
                    "block": {
                      "attributes": {
                        "attempts": {
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
            },
            "cloudwatch_logs_parameters": {
              "block": {
                "attributes": {
                  "log_stream_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timestamp": {
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
            "ecs_task_parameters": {
              "block": {
                "attributes": {
                  "enable_ecs_managed_tags": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_execute_command": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "group": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "launch_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "platform_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "propagate_tags": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reference_id": {
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
                  "task_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "task_definition_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "capacity_provider_strategy": {
                    "block": {
                      "attributes": {
                        "base": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "capacity_provider": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "weight": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 6,
                    "nesting_mode": "list"
                  },
                  "network_configuration": {
                    "block": {
                      "block_types": {
                        "aws_vpc_configuration": {
                          "block": {
                            "attributes": {
                              "assign_public_ip": {
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
                              "subnets": {
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
                  "overrides": {
                    "block": {
                      "attributes": {
                        "cpu": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "execution_role_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "memory": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "task_role_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "container_override": {
                          "block": {
                            "attributes": {
                              "command": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "cpu": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "memory": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "memory_reservation": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "environment": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "environment_file": {
                                "block": {
                                  "attributes": {
                                    "type": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              },
                              "resource_requirement": {
                                "block": {
                                  "attributes": {
                                    "type": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "required": true,
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
                          "nesting_mode": "list"
                        },
                        "ephemeral_storage": {
                          "block": {
                            "attributes": {
                              "size_in_gib": {
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
                        "inference_accelerator_override": {
                          "block": {
                            "attributes": {
                              "device_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "device_type": {
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
                  "placement_constraint": {
                    "block": {
                      "attributes": {
                        "expression": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 10,
                    "nesting_mode": "list"
                  },
                  "placement_strategy": {
                    "block": {
                      "attributes": {
                        "field": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 5,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "eventbridge_event_bus_parameters": {
              "block": {
                "attributes": {
                  "detail_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "endpoint_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resources": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "source": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "time": {
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
            "http_parameters": {
              "block": {
                "attributes": {
                  "header_parameters": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "path_parameter_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "query_string_parameters": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "kinesis_stream_parameters": {
              "block": {
                "attributes": {
                  "partition_key": {
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
            "lambda_function_parameters": {
              "block": {
                "attributes": {
                  "invocation_type": {
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
            "redshift_data_parameters": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "db_user": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secret_manager_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sqls": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "statement_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "with_event": {
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
            "sagemaker_pipeline_parameters": {
              "block": {
                "block_types": {
                  "pipeline_parameter": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 200,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sqs_queue_parameters": {
              "block": {
                "attributes": {
                  "message_deduplication_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_group_id": {
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
            "step_function_state_machine_parameters": {
              "block": {
                "attributes": {
                  "invocation_type": {
                    "description_kind": "plain",
                    "required": true,
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
  "version": 0
}`

func AwsPipesPipeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsPipesPipe), &result)
	return &result
}

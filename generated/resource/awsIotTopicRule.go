package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIotTopicRule = `{
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
      "enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sql": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sql_version": {
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
      }
    },
    "block_types": {
      "cloudwatch_alarm": {
        "block": {
          "attributes": {
            "alarm_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "state_reason": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "state_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "cloudwatch_metric": {
        "block": {
          "attributes": {
            "metric_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_namespace": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_timestamp": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_unit": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "dynamodb": {
        "block": {
          "attributes": {
            "hash_key_field": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hash_key_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hash_key_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "operation": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "payload_field": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "range_key_field": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "range_key_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "range_key_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "table_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "dynamodbv2": {
        "block": {
          "attributes": {
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "put_item": {
              "block": {
                "attributes": {
                  "table_name": {
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
        "nesting_mode": "set"
      },
      "elasticsearch": {
        "block": {
          "attributes": {
            "endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "index": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "error_action": {
        "block": {
          "block_types": {
            "cloudwatch_alarm": {
              "block": {
                "attributes": {
                  "alarm_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "state_reason": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "state_value": {
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
            "cloudwatch_metric": {
              "block": {
                "attributes": {
                  "metric_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "metric_namespace": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "metric_timestamp": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metric_unit": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "metric_value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "role_arn": {
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
            "dynamodb": {
              "block": {
                "attributes": {
                  "hash_key_field": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "hash_key_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "hash_key_value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "operation": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "payload_field": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "range_key_field": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "range_key_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "range_key_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "table_name": {
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
            "dynamodbv2": {
              "block": {
                "attributes": {
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "put_item": {
                    "block": {
                      "attributes": {
                        "table_name": {
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
            "elasticsearch": {
              "block": {
                "attributes": {
                  "endpoint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "index": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
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
            "firehose": {
              "block": {
                "attributes": {
                  "delivery_stream_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "separator": {
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
            "iot_analytics": {
              "block": {
                "attributes": {
                  "channel_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "role_arn": {
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
            "iot_events": {
              "block": {
                "attributes": {
                  "input_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "message_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
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
            "kinesis": {
              "block": {
                "attributes": {
                  "partition_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "stream_name": {
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
            "lambda": {
              "block": {
                "attributes": {
                  "function_arn": {
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
            "republish": {
              "block": {
                "attributes": {
                  "qos": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "topic": {
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
            "s3": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "role_arn": {
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
            "sns": {
              "block": {
                "attributes": {
                  "message_format": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "target_arn": {
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
            "sqs": {
              "block": {
                "attributes": {
                  "queue_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "use_base64": {
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
            "step_functions": {
              "block": {
                "attributes": {
                  "execution_name_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "state_machine_name": {
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
      "firehose": {
        "block": {
          "attributes": {
            "delivery_stream_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "separator": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "iot_analytics": {
        "block": {
          "attributes": {
            "channel_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "iot_events": {
        "block": {
          "attributes": {
            "input_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "message_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "kinesis": {
        "block": {
          "attributes": {
            "partition_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "stream_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "lambda": {
        "block": {
          "attributes": {
            "function_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "republish": {
        "block": {
          "attributes": {
            "qos": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "topic": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "s3": {
        "block": {
          "attributes": {
            "bucket_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "sns": {
        "block": {
          "attributes": {
            "message_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "sqs": {
        "block": {
          "attributes": {
            "queue_url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_base64": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "step_functions": {
        "block": {
          "attributes": {
            "execution_name_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "state_machine_name": {
              "description_kind": "plain",
              "required": true,
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
  "version": 0
}`

func AwsIotTopicRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIotTopicRule), &result)
	return &result
}

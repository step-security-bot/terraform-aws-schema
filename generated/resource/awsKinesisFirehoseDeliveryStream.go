package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKinesisFirehoseDeliveryStream = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_id": {
        "computed": true,
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
      "version_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "elasticsearch_configuration": {
        "block": {
          "attributes": {
            "buffering_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "buffering_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cluster_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "domain_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "index_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "index_rotation_period": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retry_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_backup_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "cloudwatch_logging_options": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_stream_name": {
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
            "processing_configuration": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "processors": {
                    "block": {
                      "attributes": {
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "parameters": {
                          "block": {
                            "attributes": {
                              "parameter_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "parameter_value": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "vpc_config": {
              "block": {
                "attributes": {
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "security_group_ids": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "vpc_id": {
                    "computed": true,
                    "description_kind": "plain",
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
      "extended_s3_configuration": {
        "block": {
          "attributes": {
            "bucket_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "buffer_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "buffer_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "compression_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "error_output_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_backup_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "cloudwatch_logging_options": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_stream_name": {
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
            "data_format_conversion_configuration": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "input_format_configuration": {
                    "block": {
                      "block_types": {
                        "deserializer": {
                          "block": {
                            "block_types": {
                              "hive_json_ser_de": {
                                "block": {
                                  "attributes": {
                                    "timestamp_formats": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "open_x_json_ser_de": {
                                "block": {
                                  "attributes": {
                                    "case_insensitive": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "column_to_json_key_mappings": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "convert_dots_in_json_keys_to_underscores": {
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
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "output_format_configuration": {
                    "block": {
                      "block_types": {
                        "serializer": {
                          "block": {
                            "block_types": {
                              "orc_ser_de": {
                                "block": {
                                  "attributes": {
                                    "block_size_bytes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "bloom_filter_columns": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "bloom_filter_false_positive_probability": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "compression": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "dictionary_key_threshold": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "enable_padding": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "format_version": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "padding_tolerance": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "row_index_stride": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "stripe_size_bytes": {
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
                              "parquet_ser_de": {
                                "block": {
                                  "attributes": {
                                    "block_size_bytes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "compression": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "enable_dictionary_compression": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "max_padding_bytes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "page_size_bytes": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "writer_version": {
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "schema_configuration": {
                    "block": {
                      "attributes": {
                        "catalog_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "region": {
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
                        "table_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "version_id": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "dynamic_partitioning_configuration": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "retry_duration": {
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
            "processing_configuration": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "processors": {
                    "block": {
                      "attributes": {
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "parameters": {
                          "block": {
                            "attributes": {
                              "parameter_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "parameter_value": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "s3_backup_configuration": {
              "block": {
                "attributes": {
                  "bucket_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "buffer_interval": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "buffer_size": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "compression_format": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "error_output_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prefix": {
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
                "block_types": {
                  "cloudwatch_logging_options": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "log_group_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "log_stream_name": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_endpoint_configuration": {
        "block": {
          "attributes": {
            "access_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "buffering_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "buffering_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retry_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_backup_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "cloudwatch_logging_options": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_stream_name": {
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
            "processing_configuration": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "processors": {
                    "block": {
                      "attributes": {
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "parameters": {
                          "block": {
                            "attributes": {
                              "parameter_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "parameter_value": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "request_configuration": {
              "block": {
                "attributes": {
                  "content_encoding": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "common_attributes": {
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
      "kinesis_source_configuration": {
        "block": {
          "attributes": {
            "kinesis_stream_arn": {
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
      "opensearch_configuration": {
        "block": {
          "attributes": {
            "buffering_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "buffering_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cluster_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "domain_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "index_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "index_rotation_period": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retry_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_backup_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "cloudwatch_logging_options": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_stream_name": {
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
            "processing_configuration": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "processors": {
                    "block": {
                      "attributes": {
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "parameters": {
                          "block": {
                            "attributes": {
                              "parameter_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "parameter_value": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "vpc_config": {
              "block": {
                "attributes": {
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "security_group_ids": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "vpc_id": {
                    "computed": true,
                    "description_kind": "plain",
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
      "redshift_configuration": {
        "block": {
          "attributes": {
            "cluster_jdbcurl": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "copy_options": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_table_columns": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_table_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "retry_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_backup_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "cloudwatch_logging_options": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_stream_name": {
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
            "processing_configuration": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "processors": {
                    "block": {
                      "attributes": {
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "parameters": {
                          "block": {
                            "attributes": {
                              "parameter_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "parameter_value": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "s3_backup_configuration": {
              "block": {
                "attributes": {
                  "bucket_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "buffer_interval": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "buffer_size": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "compression_format": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "error_output_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prefix": {
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
                "block_types": {
                  "cloudwatch_logging_options": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "log_group_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "log_stream_name": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "s3_configuration": {
        "block": {
          "attributes": {
            "bucket_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "buffer_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "buffer_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "compression_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "error_output_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prefix": {
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
          "block_types": {
            "cloudwatch_logging_options": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_stream_name": {
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
      "server_side_encryption": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_type": {
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
      "splunk_configuration": {
        "block": {
          "attributes": {
            "hec_acknowledgment_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "hec_endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hec_endpoint_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hec_token": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "retry_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "s3_backup_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "cloudwatch_logging_options": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_stream_name": {
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
            "processing_configuration": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "processors": {
                    "block": {
                      "attributes": {
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "parameters": {
                          "block": {
                            "attributes": {
                              "parameter_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "parameter_value": {
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

func AwsKinesisFirehoseDeliveryStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKinesisFirehoseDeliveryStream), &result)
	return &result
}

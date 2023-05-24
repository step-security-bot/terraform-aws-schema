package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppflowFlow = `{
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_arn": {
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
      }
    },
    "block_types": {
      "destination_flow_config": {
        "block": {
          "attributes": {
            "api_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connector_profile_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connector_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "destination_connector_properties": {
              "block": {
                "block_types": {
                  "custom_connector": {
                    "block": {
                      "attributes": {
                        "custom_properties": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "entity_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "id_field_names": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "write_operation_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fail_on_first_destination_error": {
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
                  "customer_profiles": {
                    "block": {
                      "attributes": {
                        "domain_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "object_type_name": {
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
                  "event_bridge": {
                    "block": {
                      "attributes": {
                        "object": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fail_on_first_destination_error": {
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
                  "honeycode": {
                    "block": {
                      "attributes": {
                        "object": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fail_on_first_destination_error": {
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
                  "lookout_metrics": {
                    "block": {
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "marketo": {
                    "block": {
                      "attributes": {
                        "object": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fail_on_first_destination_error": {
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
                  "redshift": {
                    "block": {
                      "attributes": {
                        "bucket_prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "intermediate_bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "object": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fail_on_first_destination_error": {
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
                  "s3": {
                    "block": {
                      "attributes": {
                        "bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "bucket_prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "s3_output_format_config": {
                          "block": {
                            "attributes": {
                              "file_type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "preserve_source_data_typing": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "aggregation_config": {
                                "block": {
                                  "attributes": {
                                    "aggregation_type": {
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
                              "prefix_config": {
                                "block": {
                                  "attributes": {
                                    "prefix_format": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "prefix_type": {
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
                  "salesforce": {
                    "block": {
                      "attributes": {
                        "id_field_names": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "object": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "write_operation_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fail_on_first_destination_error": {
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
                  "sapo_data": {
                    "block": {
                      "attributes": {
                        "id_field_names": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "object_path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "write_operation_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fail_on_first_destination_error": {
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
                        "success_response_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
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
                  "snowflake": {
                    "block": {
                      "attributes": {
                        "bucket_prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "intermediate_bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "object": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fail_on_first_destination_error": {
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
                  "upsolver": {
                    "block": {
                      "attributes": {
                        "bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "bucket_prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "s3_output_format_config": {
                          "block": {
                            "attributes": {
                              "file_type": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "aggregation_config": {
                                "block": {
                                  "attributes": {
                                    "aggregation_type": {
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
                              "prefix_config": {
                                "block": {
                                  "attributes": {
                                    "prefix_format": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "prefix_type": {
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
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "zendesk": {
                    "block": {
                      "attributes": {
                        "id_field_names": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "object": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "write_operation_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "error_handling_config": {
                          "block": {
                            "attributes": {
                              "bucket_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "fail_on_first_destination_error": {
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "source_flow_config": {
        "block": {
          "attributes": {
            "api_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connector_profile_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connector_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "incremental_pull_config": {
              "block": {
                "attributes": {
                  "datetime_type_field_name": {
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
            "source_connector_properties": {
              "block": {
                "block_types": {
                  "amplitude": {
                    "block": {
                      "attributes": {
                        "object": {
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
                  "custom_connector": {
                    "block": {
                      "attributes": {
                        "custom_properties": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "entity_name": {
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
                  "datadog": {
                    "block": {
                      "attributes": {
                        "object": {
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
                  "dynatrace": {
                    "block": {
                      "attributes": {
                        "object": {
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
                  "google_analytics": {
                    "block": {
                      "attributes": {
                        "object": {
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
                  "infor_nexus": {
                    "block": {
                      "attributes": {
                        "object": {
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
                  "marketo": {
                    "block": {
                      "attributes": {
                        "object": {
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
                        "bucket_prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "s3_input_format_config": {
                          "block": {
                            "attributes": {
                              "s3_input_file_type": {
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
                  "salesforce": {
                    "block": {
                      "attributes": {
                        "enable_dynamic_field_update": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "include_deleted_records": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "object": {
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
                  "sapo_data": {
                    "block": {
                      "attributes": {
                        "object_path": {
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
                  "service_now": {
                    "block": {
                      "attributes": {
                        "object": {
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
                  "singular": {
                    "block": {
                      "attributes": {
                        "object": {
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
                  "slack": {
                    "block": {
                      "attributes": {
                        "object": {
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
                  "trendmicro": {
                    "block": {
                      "attributes": {
                        "object": {
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
                  "veeva": {
                    "block": {
                      "attributes": {
                        "document_type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "include_all_versions": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "include_renditions": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "include_source_files": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "object": {
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
                  "zendesk": {
                    "block": {
                      "attributes": {
                        "object": {
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
      "task": {
        "block": {
          "attributes": {
            "destination_field": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_fields": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "task_properties": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "task_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "connector_operator": {
              "block": {
                "attributes": {
                  "amplitude": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom_connector": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "datadog": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dynatrace": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "google_analytics": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "infor_nexus": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "marketo": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "salesforce": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sapo_data": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_now": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "singular": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "slack": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "trendmicro": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "veeva": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "zendesk": {
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
        "min_items": 1,
        "nesting_mode": "set"
      },
      "trigger_config": {
        "block": {
          "attributes": {
            "trigger_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "trigger_properties": {
              "block": {
                "block_types": {
                  "scheduled": {
                    "block": {
                      "attributes": {
                        "data_pull_mode": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "first_execution_from": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "schedule_end_time": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "schedule_expression": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "schedule_offset": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "schedule_start_time": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "timezone": {
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
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAppflowFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppflowFlow), &result)
	return &result
}

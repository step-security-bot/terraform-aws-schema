package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsQuicksightDataSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_set_id": {
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
      "import_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_columns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "name": "string",
              "type": "string"
            }
          ]
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
      }
    },
    "block_types": {
      "column_groups": {
        "block": {
          "block_types": {
            "geo_spatial_column_group": {
              "block": {
                "attributes": {
                  "columns": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "country_code": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
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
        "max_items": 8,
        "nesting_mode": "list"
      },
      "column_level_permission_rules": {
        "block": {
          "attributes": {
            "column_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principals": {
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
        "nesting_mode": "list"
      },
      "data_set_usage_configuration": {
        "block": {
          "attributes": {
            "disable_use_as_direct_query_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "disable_use_as_imported_source": {
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
      },
      "field_folders": {
        "block": {
          "attributes": {
            "columns": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "field_folders_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1000,
        "nesting_mode": "set"
      },
      "logical_table_map": {
        "block": {
          "attributes": {
            "alias": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "logical_table_map_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_transforms": {
              "block": {
                "block_types": {
                  "cast_column_type_operation": {
                    "block": {
                      "attributes": {
                        "column_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "format": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "new_column_type": {
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
                  "create_columns_operation": {
                    "block": {
                      "block_types": {
                        "columns": {
                          "block": {
                            "attributes": {
                              "column_id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "column_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "expression": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 128,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "filter_operation": {
                    "block": {
                      "attributes": {
                        "condition_expression": {
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
                  "project_operation": {
                    "block": {
                      "attributes": {
                        "projected_columns": {
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
                    "nesting_mode": "list"
                  },
                  "rename_column_operation": {
                    "block": {
                      "attributes": {
                        "column_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "new_column_name": {
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
                  "tag_column_operation": {
                    "block": {
                      "attributes": {
                        "column_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "tags": {
                          "block": {
                            "attributes": {
                              "column_geographic_role": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "column_description": {
                                "block": {
                                  "attributes": {
                                    "text": {
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
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 16,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "untag_column_operation": {
                    "block": {
                      "attributes": {
                        "column_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "tag_names": {
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
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 2048,
              "nesting_mode": "list"
            },
            "source": {
              "block": {
                "attributes": {
                  "data_set_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "physical_table_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "join_instruction": {
                    "block": {
                      "attributes": {
                        "left_operand": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "on_clause": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "right_operand": {
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
                      "block_types": {
                        "left_join_key_properties": {
                          "block": {
                            "attributes": {
                              "unique_key": {
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
                        },
                        "right_join_key_properties": {
                          "block": {
                            "attributes": {
                              "unique_key": {
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 64,
        "nesting_mode": "set"
      },
      "permissions": {
        "block": {
          "attributes": {
            "actions": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "principal": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 64,
        "nesting_mode": "set"
      },
      "physical_table_map": {
        "block": {
          "attributes": {
            "physical_table_map_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "custom_sql": {
              "block": {
                "attributes": {
                  "data_source_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "sql_query": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "columns": {
                    "block": {
                      "attributes": {
                        "name": {
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
                    "max_items": 2048,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "relational_table": {
              "block": {
                "attributes": {
                  "catalog": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_source_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "schema": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "input_columns": {
                    "block": {
                      "attributes": {
                        "name": {
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
                    "max_items": 2048,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "s3_source": {
              "block": {
                "attributes": {
                  "data_source_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "input_columns": {
                    "block": {
                      "attributes": {
                        "name": {
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
                    "max_items": 2048,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "upload_settings": {
                    "block": {
                      "attributes": {
                        "contains_header": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "delimiter": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "format": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "start_from_row": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "text_qualifier": {
                          "computed": true,
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 32,
        "nesting_mode": "set"
      },
      "refresh_properties": {
        "block": {
          "block_types": {
            "refresh_configuration": {
              "block": {
                "block_types": {
                  "incremental_refresh": {
                    "block": {
                      "block_types": {
                        "lookback_window": {
                          "block": {
                            "attributes": {
                              "column_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "size": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "size_unit": {
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "row_level_permission_data_set": {
        "block": {
          "attributes": {
            "arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "format_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "permission_policy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "status": {
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
      "row_level_permission_tag_configuration": {
        "block": {
          "attributes": {
            "status": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "tag_rules": {
              "block": {
                "attributes": {
                  "column_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "match_all_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tag_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "tag_multi_value_delimiter": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 50,
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

func AwsQuicksightDataSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsQuicksightDataSet), &result)
	return &result
}

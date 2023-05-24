package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKendraDataSource = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "error_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "index_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "language_code": {
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
      "role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
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
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "configuration": {
        "block": {
          "block_types": {
            "s3_configuration": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "exclusion_patterns": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "inclusion_patterns": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "inclusion_prefixes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "access_control_list_configuration": {
                    "block": {
                      "attributes": {
                        "key_path": {
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
                  "documents_metadata_configuration": {
                    "block": {
                      "attributes": {
                        "s3_prefix": {
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
            "web_crawler_configuration": {
              "block": {
                "attributes": {
                  "crawl_depth": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_content_size_per_page_in_mega_bytes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_links_per_page": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_urls_per_minute_crawl_rate": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "url_exclusion_patterns": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "url_inclusion_patterns": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "authentication_configuration": {
                    "block": {
                      "block_types": {
                        "basic_authentication": {
                          "block": {
                            "attributes": {
                              "credentials": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 10,
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "proxy_configuration": {
                    "block": {
                      "attributes": {
                        "credentials": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "port": {
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
                  "urls": {
                    "block": {
                      "block_types": {
                        "seed_url_configuration": {
                          "block": {
                            "attributes": {
                              "seed_urls": {
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "web_crawler_mode": {
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
                        "site_maps_configuration": {
                          "block": {
                            "attributes": {
                              "site_maps": {
                                "description_kind": "plain",
                                "required": true,
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "custom_document_enrichment_configuration": {
        "block": {
          "attributes": {
            "role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "inline_configurations": {
              "block": {
                "attributes": {
                  "document_content_deletion": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "condition": {
                    "block": {
                      "attributes": {
                        "condition_document_attribute_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "condition_on_value": {
                          "block": {
                            "attributes": {
                              "date_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "long_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "string_list_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "string_value": {
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
                  "target": {
                    "block": {
                      "attributes": {
                        "target_document_attribute_key": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "target_document_attribute_value_deletion": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "target_document_attribute_value": {
                          "block": {
                            "attributes": {
                              "date_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "long_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "string_list_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "string_value": {
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
              "max_items": 100,
              "nesting_mode": "set"
            },
            "post_extraction_hook_configuration": {
              "block": {
                "attributes": {
                  "lambda_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "s3_bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "invocation_condition": {
                    "block": {
                      "attributes": {
                        "condition_document_attribute_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "condition_on_value": {
                          "block": {
                            "attributes": {
                              "date_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "long_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "string_list_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "string_value": {
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
            "pre_extraction_hook_configuration": {
              "block": {
                "attributes": {
                  "lambda_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "s3_bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "invocation_condition": {
                    "block": {
                      "attributes": {
                        "condition_document_attribute_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "condition_on_value": {
                          "block": {
                            "attributes": {
                              "date_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "long_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "string_list_value": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "string_value": {
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

func AwsKendraDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKendraDataSource), &result)
	return &result
}

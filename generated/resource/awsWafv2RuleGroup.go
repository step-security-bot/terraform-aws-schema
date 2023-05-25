package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWafv2RuleGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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
      "lock_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
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
      "custom_response_body": {
        "block": {
          "attributes": {
            "content": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "content_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "rule": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "action": {
              "block": {
                "block_types": {
                  "allow": {
                    "block": {
                      "block_types": {
                        "custom_request_handling": {
                          "block": {
                            "block_types": {
                              "insert_header": {
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
                                "min_items": 1,
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "block": {
                    "block": {
                      "block_types": {
                        "custom_response": {
                          "block": {
                            "attributes": {
                              "custom_response_body_key": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "response_code": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "response_header": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "count": {
                    "block": {
                      "block_types": {
                        "custom_request_handling": {
                          "block": {
                            "block_types": {
                              "insert_header": {
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
                                "min_items": 1,
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
            "rule_label": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "statement": {
              "block": {
                "block_types": {
                  "and_statement": {
                    "block": {
                      "block_types": {
                        "statement": {
                          "block": {
                            "block_types": {
                              "and_statement": {
                                "block": {
                                  "block_types": {
                                    "statement": {
                                      "block": {
                                        "block_types": {
                                          "byte_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "positional_constraint": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "search_string": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "geo_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "country_codes": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
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
                                          "ip_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "ip_set_forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "position": {
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
                                          "label_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "scope": {
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
                                          "regex_pattern_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "size_constraint_statement": {
                                            "block": {
                                              "attributes": {
                                                "comparison_operator": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "size": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sqli_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "xss_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
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
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "byte_match_statement": {
                                "block": {
                                  "attributes": {
                                    "positional_constraint": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "search_string": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "geo_match_statement": {
                                "block": {
                                  "attributes": {
                                    "country_codes": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "forwarded_ip_config": {
                                      "block": {
                                        "attributes": {
                                          "fallback_behavior": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_name": {
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
                              "ip_set_reference_statement": {
                                "block": {
                                  "attributes": {
                                    "arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "ip_set_forwarded_ip_config": {
                                      "block": {
                                        "attributes": {
                                          "fallback_behavior": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "position": {
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
                              "label_match_statement": {
                                "block": {
                                  "attributes": {
                                    "key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "scope": {
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
                              "not_statement": {
                                "block": {
                                  "block_types": {
                                    "statement": {
                                      "block": {
                                        "block_types": {
                                          "byte_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "positional_constraint": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "search_string": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "geo_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "country_codes": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
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
                                          "ip_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "ip_set_forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "position": {
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
                                          "label_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "scope": {
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
                                          "regex_pattern_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "size_constraint_statement": {
                                            "block": {
                                              "attributes": {
                                                "comparison_operator": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "size": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sqli_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "xss_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
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
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "or_statement": {
                                "block": {
                                  "block_types": {
                                    "statement": {
                                      "block": {
                                        "block_types": {
                                          "byte_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "positional_constraint": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "search_string": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "geo_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "country_codes": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
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
                                          "ip_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "ip_set_forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "position": {
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
                                          "label_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "scope": {
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
                                          "regex_pattern_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "size_constraint_statement": {
                                            "block": {
                                              "attributes": {
                                                "comparison_operator": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "size": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sqli_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "xss_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
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
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "regex_pattern_set_reference_statement": {
                                "block": {
                                  "attributes": {
                                    "arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "size_constraint_statement": {
                                "block": {
                                  "attributes": {
                                    "comparison_operator": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "size": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "sqli_match_statement": {
                                "block": {
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "xss_match_statement": {
                                "block": {
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
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
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "byte_match_statement": {
                    "block": {
                      "attributes": {
                        "positional_constraint": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "search_string": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "field_to_match": {
                          "block": {
                            "block_types": {
                              "all_query_arguments": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "body": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "method": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "query_string": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "single_header": {
                                "block": {
                                  "attributes": {
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
                              },
                              "single_query_argument": {
                                "block": {
                                  "attributes": {
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
                              },
                              "uri_path": {
                                "block": {
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
                        "text_transformation": {
                          "block": {
                            "attributes": {
                              "priority": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "geo_match_statement": {
                    "block": {
                      "attributes": {
                        "country_codes": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "forwarded_ip_config": {
                          "block": {
                            "attributes": {
                              "fallback_behavior": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "header_name": {
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
                  "ip_set_reference_statement": {
                    "block": {
                      "attributes": {
                        "arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "ip_set_forwarded_ip_config": {
                          "block": {
                            "attributes": {
                              "fallback_behavior": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "header_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "position": {
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
                  "label_match_statement": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scope": {
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
                  "not_statement": {
                    "block": {
                      "block_types": {
                        "statement": {
                          "block": {
                            "block_types": {
                              "and_statement": {
                                "block": {
                                  "block_types": {
                                    "statement": {
                                      "block": {
                                        "block_types": {
                                          "byte_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "positional_constraint": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "search_string": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "geo_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "country_codes": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
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
                                          "ip_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "ip_set_forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "position": {
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
                                          "label_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "scope": {
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
                                          "regex_pattern_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "size_constraint_statement": {
                                            "block": {
                                              "attributes": {
                                                "comparison_operator": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "size": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sqli_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "xss_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
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
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "byte_match_statement": {
                                "block": {
                                  "attributes": {
                                    "positional_constraint": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "search_string": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "geo_match_statement": {
                                "block": {
                                  "attributes": {
                                    "country_codes": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "forwarded_ip_config": {
                                      "block": {
                                        "attributes": {
                                          "fallback_behavior": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_name": {
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
                              "ip_set_reference_statement": {
                                "block": {
                                  "attributes": {
                                    "arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "ip_set_forwarded_ip_config": {
                                      "block": {
                                        "attributes": {
                                          "fallback_behavior": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "position": {
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
                              "label_match_statement": {
                                "block": {
                                  "attributes": {
                                    "key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "scope": {
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
                              "not_statement": {
                                "block": {
                                  "block_types": {
                                    "statement": {
                                      "block": {
                                        "block_types": {
                                          "byte_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "positional_constraint": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "search_string": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "geo_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "country_codes": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
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
                                          "ip_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "ip_set_forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "position": {
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
                                          "label_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "scope": {
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
                                          "regex_pattern_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "size_constraint_statement": {
                                            "block": {
                                              "attributes": {
                                                "comparison_operator": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "size": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sqli_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "xss_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
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
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "or_statement": {
                                "block": {
                                  "block_types": {
                                    "statement": {
                                      "block": {
                                        "block_types": {
                                          "byte_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "positional_constraint": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "search_string": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "geo_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "country_codes": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
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
                                          "ip_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "ip_set_forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "position": {
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
                                          "label_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "scope": {
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
                                          "regex_pattern_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "size_constraint_statement": {
                                            "block": {
                                              "attributes": {
                                                "comparison_operator": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "size": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sqli_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "xss_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
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
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "regex_pattern_set_reference_statement": {
                                "block": {
                                  "attributes": {
                                    "arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "size_constraint_statement": {
                                "block": {
                                  "attributes": {
                                    "comparison_operator": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "size": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "sqli_match_statement": {
                                "block": {
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "xss_match_statement": {
                                "block": {
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
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
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "or_statement": {
                    "block": {
                      "block_types": {
                        "statement": {
                          "block": {
                            "block_types": {
                              "and_statement": {
                                "block": {
                                  "block_types": {
                                    "statement": {
                                      "block": {
                                        "block_types": {
                                          "byte_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "positional_constraint": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "search_string": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "geo_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "country_codes": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
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
                                          "ip_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "ip_set_forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "position": {
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
                                          "label_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "scope": {
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
                                          "regex_pattern_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "size_constraint_statement": {
                                            "block": {
                                              "attributes": {
                                                "comparison_operator": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "size": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sqli_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "xss_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
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
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "byte_match_statement": {
                                "block": {
                                  "attributes": {
                                    "positional_constraint": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "search_string": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "geo_match_statement": {
                                "block": {
                                  "attributes": {
                                    "country_codes": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "forwarded_ip_config": {
                                      "block": {
                                        "attributes": {
                                          "fallback_behavior": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_name": {
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
                              "ip_set_reference_statement": {
                                "block": {
                                  "attributes": {
                                    "arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "ip_set_forwarded_ip_config": {
                                      "block": {
                                        "attributes": {
                                          "fallback_behavior": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "header_name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "position": {
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
                              "label_match_statement": {
                                "block": {
                                  "attributes": {
                                    "key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "scope": {
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
                              "not_statement": {
                                "block": {
                                  "block_types": {
                                    "statement": {
                                      "block": {
                                        "block_types": {
                                          "byte_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "positional_constraint": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "search_string": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "geo_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "country_codes": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
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
                                          "ip_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "ip_set_forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "position": {
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
                                          "label_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "scope": {
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
                                          "regex_pattern_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "size_constraint_statement": {
                                            "block": {
                                              "attributes": {
                                                "comparison_operator": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "size": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sqli_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "xss_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
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
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "or_statement": {
                                "block": {
                                  "block_types": {
                                    "statement": {
                                      "block": {
                                        "block_types": {
                                          "byte_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "positional_constraint": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "search_string": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "geo_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "country_codes": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "block_types": {
                                                "forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
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
                                          "ip_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "ip_set_forwarded_ip_config": {
                                                  "block": {
                                                    "attributes": {
                                                      "fallback_behavior": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "header_name": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      },
                                                      "position": {
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
                                          "label_match_statement": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "scope": {
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
                                          "regex_pattern_set_reference_statement": {
                                            "block": {
                                              "attributes": {
                                                "arn": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "size_constraint_statement": {
                                            "block": {
                                              "attributes": {
                                                "comparison_operator": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "size": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sqli_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
                                                  "nesting_mode": "set"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "xss_match_statement": {
                                            "block": {
                                              "block_types": {
                                                "field_to_match": {
                                                  "block": {
                                                    "block_types": {
                                                      "all_query_arguments": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "body": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "method": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "query_string": {
                                                        "block": {
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "single_header": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "single_query_argument": {
                                                        "block": {
                                                          "attributes": {
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
                                                      },
                                                      "uri_path": {
                                                        "block": {
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
                                                "text_transformation": {
                                                  "block": {
                                                    "attributes": {
                                                      "priority": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "number"
                                                      },
                                                      "type": {
                                                        "description_kind": "plain",
                                                        "required": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description_kind": "plain"
                                                  },
                                                  "min_items": 1,
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
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "regex_pattern_set_reference_statement": {
                                "block": {
                                  "attributes": {
                                    "arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "size_constraint_statement": {
                                "block": {
                                  "attributes": {
                                    "comparison_operator": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "size": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "sqli_match_statement": {
                                "block": {
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
                                      "nesting_mode": "set"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "xss_match_statement": {
                                "block": {
                                  "block_types": {
                                    "field_to_match": {
                                      "block": {
                                        "block_types": {
                                          "all_query_arguments": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "body": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "method": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "query_string": {
                                            "block": {
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "single_header": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "single_query_argument": {
                                            "block": {
                                              "attributes": {
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
                                          },
                                          "uri_path": {
                                            "block": {
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
                                    "text_transformation": {
                                      "block": {
                                        "attributes": {
                                          "priority": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "min_items": 1,
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
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "regex_pattern_set_reference_statement": {
                    "block": {
                      "attributes": {
                        "arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "field_to_match": {
                          "block": {
                            "block_types": {
                              "all_query_arguments": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "body": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "method": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "query_string": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "single_header": {
                                "block": {
                                  "attributes": {
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
                              },
                              "single_query_argument": {
                                "block": {
                                  "attributes": {
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
                              },
                              "uri_path": {
                                "block": {
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
                        "text_transformation": {
                          "block": {
                            "attributes": {
                              "priority": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "size_constraint_statement": {
                    "block": {
                      "attributes": {
                        "comparison_operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "size": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "field_to_match": {
                          "block": {
                            "block_types": {
                              "all_query_arguments": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "body": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "method": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "query_string": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "single_header": {
                                "block": {
                                  "attributes": {
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
                              },
                              "single_query_argument": {
                                "block": {
                                  "attributes": {
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
                              },
                              "uri_path": {
                                "block": {
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
                        "text_transformation": {
                          "block": {
                            "attributes": {
                              "priority": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "sqli_match_statement": {
                    "block": {
                      "block_types": {
                        "field_to_match": {
                          "block": {
                            "block_types": {
                              "all_query_arguments": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "body": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "method": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "query_string": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "single_header": {
                                "block": {
                                  "attributes": {
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
                              },
                              "single_query_argument": {
                                "block": {
                                  "attributes": {
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
                              },
                              "uri_path": {
                                "block": {
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
                        "text_transformation": {
                          "block": {
                            "attributes": {
                              "priority": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "xss_match_statement": {
                    "block": {
                      "block_types": {
                        "field_to_match": {
                          "block": {
                            "block_types": {
                              "all_query_arguments": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "body": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "method": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "query_string": {
                                "block": {
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "single_header": {
                                "block": {
                                  "attributes": {
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
                              },
                              "single_query_argument": {
                                "block": {
                                  "attributes": {
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
                              },
                              "uri_path": {
                                "block": {
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
                        "text_transformation": {
                          "block": {
                            "attributes": {
                              "priority": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "min_items": 1,
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
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "visibility_config": {
              "block": {
                "attributes": {
                  "cloudwatch_metrics_enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "metric_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "sampled_requests_enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
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
        "nesting_mode": "set"
      },
      "visibility_config": {
        "block": {
          "attributes": {
            "cloudwatch_metrics_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "metric_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sampled_requests_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
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

func AwsWafv2RuleGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWafv2RuleGroup), &result)
	return &result
}

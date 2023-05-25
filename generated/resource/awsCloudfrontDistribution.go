package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudfrontDistribution = `{
  "block": {
    "attributes": {
      "aliases": {
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
      "caller_reference": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "comment": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_root_object": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "http_version": {
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
      "in_progress_validation_batches": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "is_ipv6_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "price_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retain_on_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "trusted_key_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "items": [
                "list",
                [
                  "object",
                  {
                    "key_group_id": "string",
                    "key_pair_ids": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "trusted_signers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "items": [
                "list",
                [
                  "object",
                  {
                    "aws_account_number": "string",
                    "key_pair_ids": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "wait_for_deployment": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "web_acl_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "custom_error_response": {
        "block": {
          "attributes": {
            "error_caching_min_ttl": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "error_code": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "response_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "response_page_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "default_cache_behavior": {
        "block": {
          "attributes": {
            "allowed_methods": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "cache_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cached_methods": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "compress": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "default_ttl": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "field_level_encryption_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_ttl": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_ttl": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "origin_request_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "realtime_log_config_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response_headers_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "smooth_streaming": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "target_origin_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "trusted_key_groups": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "trusted_signers": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "viewer_protocol_policy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "forwarded_values": {
              "block": {
                "attributes": {
                  "headers": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "query_string": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "query_string_cache_keys": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "cookies": {
                    "block": {
                      "attributes": {
                        "forward": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "whitelisted_names": {
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "function_association": {
              "block": {
                "attributes": {
                  "event_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "function_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 2,
              "nesting_mode": "set"
            },
            "lambda_function_association": {
              "block": {
                "attributes": {
                  "event_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "include_body": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "lambda_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 4,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "logging_config": {
        "block": {
          "attributes": {
            "bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "include_cookies": {
              "description_kind": "plain",
              "optional": true,
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
      },
      "ordered_cache_behavior": {
        "block": {
          "attributes": {
            "allowed_methods": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "cache_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cached_methods": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "compress": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "default_ttl": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "field_level_encryption_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_ttl": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_ttl": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "origin_request_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path_pattern": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "realtime_log_config_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "response_headers_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "smooth_streaming": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "target_origin_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "trusted_key_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "trusted_signers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "viewer_protocol_policy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "forwarded_values": {
              "block": {
                "attributes": {
                  "headers": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "query_string": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "query_string_cache_keys": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "cookies": {
                    "block": {
                      "attributes": {
                        "forward": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "whitelisted_names": {
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "function_association": {
              "block": {
                "attributes": {
                  "event_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "function_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 2,
              "nesting_mode": "set"
            },
            "lambda_function_association": {
              "block": {
                "attributes": {
                  "event_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "include_body": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "lambda_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 4,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "origin": {
        "block": {
          "attributes": {
            "connection_attempts": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "connection_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "domain_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "origin_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "origin_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "custom_header": {
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
            },
            "custom_origin_config": {
              "block": {
                "attributes": {
                  "http_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "https_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "origin_keepalive_timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "origin_protocol_policy": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "origin_read_timeout": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "origin_ssl_protocols": {
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
            },
            "origin_shield": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "origin_shield_region": {
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
            "s3_origin_config": {
              "block": {
                "attributes": {
                  "origin_access_identity": {
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
        "min_items": 1,
        "nesting_mode": "set"
      },
      "origin_group": {
        "block": {
          "attributes": {
            "origin_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "failover_criteria": {
              "block": {
                "attributes": {
                  "status_codes": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "number"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "member": {
              "block": {
                "attributes": {
                  "origin_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 2,
              "min_items": 2,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "restrictions": {
        "block": {
          "block_types": {
            "geo_restriction": {
              "block": {
                "attributes": {
                  "locations": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "restriction_type": {
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
      },
      "viewer_certificate": {
        "block": {
          "attributes": {
            "acm_certificate_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cloudfront_default_certificate": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "iam_certificate_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "minimum_protocol_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_support_method": {
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
  "version": 1
}`

func AwsCloudfrontDistributionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudfrontDistribution), &result)
	return &result
}

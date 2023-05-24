package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudfrontResponseHeadersPolicy = `{
  "block": {
    "attributes": {
      "comment": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
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
      }
    },
    "block_types": {
      "cors_config": {
        "block": {
          "attributes": {
            "access_control_allow_credentials": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "access_control_max_age_sec": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "origin_override": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "access_control_allow_headers": {
              "block": {
                "attributes": {
                  "items": {
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
            },
            "access_control_allow_methods": {
              "block": {
                "attributes": {
                  "items": {
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
            },
            "access_control_allow_origins": {
              "block": {
                "attributes": {
                  "items": {
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
            },
            "access_control_expose_headers": {
              "block": {
                "attributes": {
                  "items": {
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
      "custom_headers_config": {
        "block": {
          "block_types": {
            "items": {
              "block": {
                "attributes": {
                  "header": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "override": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
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
      },
      "remove_headers_config": {
        "block": {
          "block_types": {
            "items": {
              "block": {
                "attributes": {
                  "header": {
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
      },
      "security_headers_config": {
        "block": {
          "block_types": {
            "content_security_policy": {
              "block": {
                "attributes": {
                  "content_security_policy": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "override": {
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
            "content_type_options": {
              "block": {
                "attributes": {
                  "override": {
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
            "frame_options": {
              "block": {
                "attributes": {
                  "frame_option": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "override": {
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
            "referrer_policy": {
              "block": {
                "attributes": {
                  "override": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "referrer_policy": {
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
            "strict_transport_security": {
              "block": {
                "attributes": {
                  "access_control_max_age_sec": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "include_subdomains": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "override": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "preload": {
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
            "xss_protection": {
              "block": {
                "attributes": {
                  "mode_block": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "override": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "protection": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "report_uri": {
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
      "server_timing_headers_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "sampling_rate": {
              "description_kind": "plain",
              "required": true,
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
  "version": 0
}`

func AwsCloudfrontResponseHeadersPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudfrontResponseHeadersPolicy), &result)
	return &result
}

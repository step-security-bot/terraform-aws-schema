package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLbListener = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_arn": {
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
      "load_balancer_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "default_action": {
        "block": {
          "attributes": {
            "order": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_group_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "authenticate_cognito": {
              "block": {
                "attributes": {
                  "authentication_request_extra_params": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "on_unauthenticated_request": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scope": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_cookie_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "user_pool_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user_pool_client_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user_pool_domain": {
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
            "authenticate_oidc": {
              "block": {
                "attributes": {
                  "authentication_request_extra_params": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "authorization_endpoint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "issuer": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "on_unauthenticated_request": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scope": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_cookie_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "token_endpoint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user_info_endpoint": {
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
            "fixed_response": {
              "block": {
                "attributes": {
                  "content_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "message_body": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_code": {
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
            "forward": {
              "block": {
                "block_types": {
                  "stickiness": {
                    "block": {
                      "attributes": {
                        "duration": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "enabled": {
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
                  "target_group": {
                    "block": {
                      "attributes": {
                        "arn": {
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
                    "max_items": 5,
                    "min_items": 2,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "redirect": {
              "block": {
                "attributes": {
                  "host": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "query": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_code": {
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
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
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

func AwsLbListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLbListener), &result)
	return &result
}

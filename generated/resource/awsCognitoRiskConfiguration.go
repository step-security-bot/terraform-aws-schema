package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoRiskConfiguration = `{
  "block": {
    "attributes": {
      "client_id": {
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
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "account_takeover_risk_configuration": {
        "block": {
          "block_types": {
            "actions": {
              "block": {
                "block_types": {
                  "high_action": {
                    "block": {
                      "attributes": {
                        "event_action": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "notify": {
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
                  "low_action": {
                    "block": {
                      "attributes": {
                        "event_action": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "notify": {
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
                  "medium_action": {
                    "block": {
                      "attributes": {
                        "event_action": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "notify": {
                          "description_kind": "plain",
                          "required": true,
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
            },
            "notify_configuration": {
              "block": {
                "attributes": {
                  "from": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reply_to": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "block_email": {
                    "block": {
                      "attributes": {
                        "html_body": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "subject": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "text_body": {
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
                  "mfa_email": {
                    "block": {
                      "attributes": {
                        "html_body": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "subject": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "text_body": {
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
                  "no_action_email": {
                    "block": {
                      "attributes": {
                        "html_body": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "subject": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "text_body": {
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
        "nesting_mode": "list"
      },
      "compromised_credentials_risk_configuration": {
        "block": {
          "attributes": {
            "event_filter": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "actions": {
              "block": {
                "attributes": {
                  "event_action": {
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
      "risk_exception_configuration": {
        "block": {
          "attributes": {
            "blocked_ip_range_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "skipped_ip_range_list": {
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
  "version": 0
}`

func AwsCognitoRiskConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoRiskConfiguration), &result)
	return &result
}

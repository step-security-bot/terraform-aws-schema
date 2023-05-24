package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLexBot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "checksum": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "child_directed": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "create_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "detect_sentiment": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_model_improvements": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "failure_reason": {
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
      "idle_session_ttl_in_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "last_updated_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "locale": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nlu_intent_confidence_threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "process_behavior": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "voice_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "abort_statement": {
        "block": {
          "attributes": {
            "response_card": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "message": {
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
                  "group_number": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 15,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "clarification_prompt": {
        "block": {
          "attributes": {
            "max_attempts": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "response_card": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "message": {
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
                  "group_number": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 15,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "intent": {
        "block": {
          "attributes": {
            "intent_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "intent_version": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 250,
        "min_items": 1,
        "nesting_mode": "set"
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

func AwsLexBotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLexBot), &result)
	return &result
}

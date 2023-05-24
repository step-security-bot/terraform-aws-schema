package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSesReceiptRule = `{
  "block": {
    "attributes": {
      "after": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "recipients": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "rule_set_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scan_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tls_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "add_header_action": {
        "block": {
          "attributes": {
            "header_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "header_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "position": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "bounce_action": {
        "block": {
          "attributes": {
            "message": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "position": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "sender": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "smtp_reply_code": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "status_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "lambda_action": {
        "block": {
          "attributes": {
            "function_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "invocation_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "position": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "topic_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "s3_action": {
        "block": {
          "attributes": {
            "bucket_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "kms_key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "object_key_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "position": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "topic_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "sns_action": {
        "block": {
          "attributes": {
            "encoding": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "position": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "topic_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "stop_action": {
        "block": {
          "attributes": {
            "position": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "scope": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "topic_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "workmail_action": {
        "block": {
          "attributes": {
            "organization_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "position": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "topic_arn": {
              "description_kind": "plain",
              "optional": true,
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
  "version": 0
}`

func AwsSesReceiptRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSesReceiptRule), &result)
	return &result
}

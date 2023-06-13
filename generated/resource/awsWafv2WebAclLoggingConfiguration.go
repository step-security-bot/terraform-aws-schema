package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWafv2WebAclLoggingConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_destination_configs": {
        "description": "AWS Kinesis Firehose Delivery Stream ARNs",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "resource_arn": {
        "description": "AWS WebACL ARN",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "logging_filter": {
        "block": {
          "attributes": {
            "default_behavior": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "filter": {
              "block": {
                "attributes": {
                  "behavior": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "requirement": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "condition": {
                    "block": {
                      "block_types": {
                        "action_condition": {
                          "block": {
                            "attributes": {
                              "action": {
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
                        "label_name_condition": {
                          "block": {
                            "attributes": {
                              "label_name": {
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
      "redacted_fields": {
        "block": {
          "block_types": {
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
            "uri_path": {
              "block": {
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Parts of the request to exclude from logs",
          "description_kind": "plain"
        },
        "max_items": 100,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWafv2WebAclLoggingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWafv2WebAclLoggingConfiguration), &result)
	return &result
}

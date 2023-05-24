package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSesv2ConfigurationSetEventDestination = `{
  "block": {
    "attributes": {
      "configuration_set_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_destination_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "event_destination": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "matching_event_types": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "cloud_watch_destination": {
              "block": {
                "block_types": {
                  "dimension_configuration": {
                    "block": {
                      "attributes": {
                        "default_dimension_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "dimension_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "dimension_value_source": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
            "kinesis_firehose_destination": {
              "block": {
                "attributes": {
                  "delivery_stream_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "iam_role_arn": {
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
            "pinpoint_destination": {
              "block": {
                "attributes": {
                  "application_arn": {
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
            "sns_destination": {
              "block": {
                "attributes": {
                  "topic_arn": {
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
  "version": 0
}`

func AwsSesv2ConfigurationSetEventDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSesv2ConfigurationSetEventDestination), &result)
	return &result
}

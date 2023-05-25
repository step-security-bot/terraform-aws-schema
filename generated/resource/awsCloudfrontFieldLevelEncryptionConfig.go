package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudfrontFieldLevelEncryptionConfig = `{
  "block": {
    "attributes": {
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
      "etag": {
        "computed": true,
        "description_kind": "plain",
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
      "content_type_profile_config": {
        "block": {
          "attributes": {
            "forward_when_content_type_is_unknown": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "content_type_profiles": {
              "block": {
                "block_types": {
                  "items": {
                    "block": {
                      "attributes": {
                        "content_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "format": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "profile_id": {
                          "description_kind": "plain",
                          "optional": true,
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
      "query_arg_profile_config": {
        "block": {
          "attributes": {
            "forward_when_query_arg_profile_is_unknown": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "query_arg_profiles": {
              "block": {
                "block_types": {
                  "items": {
                    "block": {
                      "attributes": {
                        "profile_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "query_arg": {
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
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudfrontFieldLevelEncryptionConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudfrontFieldLevelEncryptionConfig), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAthenaWorkgroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_destroy": {
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
      "state": {
        "description_kind": "plain",
        "optional": true,
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
      "configuration": {
        "block": {
          "attributes": {
            "bytes_scanned_cutoff_per_query": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enforce_workgroup_configuration": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "execution_role": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "publish_cloudwatch_metrics_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "requester_pays_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "engine_version": {
              "block": {
                "attributes": {
                  "effective_engine_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "selected_engine_version": {
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
            "result_configuration": {
              "block": {
                "attributes": {
                  "expected_bucket_owner": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "acl_configuration": {
                    "block": {
                      "attributes": {
                        "s3_acl_option": {
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
                  "encryption_configuration": {
                    "block": {
                      "attributes": {
                        "encryption_option": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "kms_key_arn": {
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

func AwsAthenaWorkgroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAthenaWorkgroup), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3BucketLifecycleConfiguration = `{
  "block": {
    "attributes": {
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expected_bucket_owner": {
        "description_kind": "plain",
        "optional": true,
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
      "rule": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "prefix": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "abort_incomplete_multipart_upload": {
              "block": {
                "attributes": {
                  "days_after_initiation": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "expiration": {
              "block": {
                "attributes": {
                  "date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "expired_object_delete_marker": {
                    "computed": true,
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
            "filter": {
              "block": {
                "attributes": {
                  "object_size_greater_than": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "object_size_less_than": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "and": {
                    "block": {
                      "attributes": {
                        "object_size_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "object_size_less_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "prefix": {
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
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "tag": {
                    "block": {
                      "attributes": {
                        "key": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "noncurrent_version_expiration": {
              "block": {
                "attributes": {
                  "newer_noncurrent_versions": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "noncurrent_days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "noncurrent_version_transition": {
              "block": {
                "attributes": {
                  "newer_noncurrent_versions": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "noncurrent_days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "storage_class": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "transition": {
              "block": {
                "attributes": {
                  "date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "storage_class": {
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
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3BucketLifecycleConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3BucketLifecycleConfiguration), &result)
	return &result
}

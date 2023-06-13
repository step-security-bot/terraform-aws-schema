package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53RecoveryreadinessResourceSet = `{
  "block": {
    "attributes": {
      "arn": {
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
      "resource_set_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_set_type": {
        "description_kind": "plain",
        "required": true,
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
      "resources": {
        "block": {
          "attributes": {
            "component_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "readiness_scopes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "dns_target_resource": {
              "block": {
                "attributes": {
                  "domain_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "hosted_zone_arn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "record_set_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "record_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "target_resource": {
                    "block": {
                      "block_types": {
                        "nlb_resource": {
                          "block": {
                            "attributes": {
                              "arn": {
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
                        "r53_resource": {
                          "block": {
                            "attributes": {
                              "domain_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "record_set_id": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "delete": {
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

func AwsRoute53RecoveryreadinessResourceSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53RecoveryreadinessResourceSet), &result)
	return &result
}

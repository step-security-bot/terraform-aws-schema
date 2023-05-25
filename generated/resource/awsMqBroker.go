package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMqBroker = `{
  "block": {
    "attributes": {
      "apply_immediately": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authentication_strategy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_minor_version_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "broker_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deployment_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host_instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "console_url": "string",
              "endpoints": [
                "list",
                "string"
              ],
              "ip_address": "string"
            }
          ]
        ]
      },
      "publicly_accessible": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "security_groups": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "storage_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
            "id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "revision": {
              "computed": true,
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
      "encryption_options": {
        "block": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_aws_owned_key": {
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
      "ldap_server_metadata": {
        "block": {
          "attributes": {
            "hosts": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "role_base": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_search_matching": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_search_subtree": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "service_account_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "service_account_username": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_base": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_role_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_search_matching": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_search_subtree": {
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
      "logs": {
        "block": {
          "attributes": {
            "audit": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "general": {
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
      "maintenance_window_start_time": {
        "block": {
          "attributes": {
            "day_of_week": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "time_of_day": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "time_zone": {
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
      "user": {
        "block": {
          "attributes": {
            "console_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "groups": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "username": {
              "description_kind": "plain",
              "required": true,
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
  "version": 0
}`

func AwsMqBrokerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMqBroker), &result)
	return &result
}

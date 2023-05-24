package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpsworksStack = `{
  "block": {
    "attributes": {
      "agent_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "berkshelf_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "color": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "configuration_manager_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "configuration_manager_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_json": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_instance_profile_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_os": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_root_device_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_ssh_key_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hostname_theme": {
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
      "manage_berkshelf": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stack_endpoint": {
        "computed": true,
        "description_kind": "plain",
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
      },
      "use_custom_cookbooks": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "use_opsworks_security_groups": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "custom_cookbooks_source": {
        "block": {
          "attributes": {
            "password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "revision": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssh_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "username": {
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
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
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

func AwsOpsworksStackSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpsworksStack), &result)
	return &result
}

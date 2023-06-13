package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsFsxOntapStorageVirtualMachine = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "iscsi": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "management": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "nfs": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "smb": [
                "list",
                [
                  "object",
                  {
                    "dns_name": "string",
                    "ip_addresses": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "file_system_id": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "root_volume_security_style": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subtype": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "svm_admin_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "uuid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "active_directory_configuration": {
        "block": {
          "attributes": {
            "netbios_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "self_managed_active_directory_configuration": {
              "block": {
                "attributes": {
                  "dns_ips": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "domain_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "file_system_administrators_group": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "organizational_unit_distinguished_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
              "max_items": 1,
              "nesting_mode": "list"
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
  "version": 1
}`

func AwsFsxOntapStorageVirtualMachineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsFsxOntapStorageVirtualMachine), &result)
	return &result
}

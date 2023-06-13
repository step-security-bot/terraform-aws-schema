package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53DomainsRegisteredDomain = `{
  "block": {
    "attributes": {
      "abuse_contact_email": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "abuse_contact_phone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "admin_privacy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expiration_date": {
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
      "registrant_privacy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "registrar_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "registrar_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reseller": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
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
      },
      "tech_privacy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "transfer_lock": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "updated_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "whois_server": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "admin_contact": {
        "block": {
          "attributes": {
            "address_line_1": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_line_2": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "extra_params": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "fax": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "first_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organization_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "phone_number": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zip_code": {
              "computed": true,
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
      "name_server": {
        "block": {
          "attributes": {
            "glue_ips": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 6,
        "nesting_mode": "list"
      },
      "registrant_contact": {
        "block": {
          "attributes": {
            "address_line_1": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_line_2": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "extra_params": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "fax": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "first_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organization_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "phone_number": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zip_code": {
              "computed": true,
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
      "tech_contact": {
        "block": {
          "attributes": {
            "address_line_1": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_line_2": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "city": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "contact_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "extra_params": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "fax": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "first_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organization_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "phone_number": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zip_code": {
              "computed": true,
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
  "version": 0
}`

func AwsRoute53DomainsRegisteredDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53DomainsRegisteredDomain), &result)
	return &result
}

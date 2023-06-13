package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAcmpcaCertificateAuthority = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_chain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_signing_request": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
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
      "key_storage_security_standard": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "not_after": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "not_before": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "permanent_deletion_time_in_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "serial": {
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
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "usage_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "certificate_authority_configuration": {
        "block": {
          "attributes": {
            "key_algorithm": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "signing_algorithm": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "subject": {
              "block": {
                "attributes": {
                  "common_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "country": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "distinguished_name_qualifier": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "generation_qualifier": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "given_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "initials": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "locality": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "organization": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "organizational_unit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pseudonym": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "state": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "surname": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "title": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
      "revocation_configuration": {
        "block": {
          "block_types": {
            "crl_configuration": {
              "block": {
                "attributes": {
                  "custom_cname": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "expiration_in_days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "s3_bucket_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_object_acl": {
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
            "ocsp_configuration": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "ocsp_custom_cname": {
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
  "version": 1
}`

func AwsAcmpcaCertificateAuthoritySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAcmpcaCertificateAuthority), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoUserPool = `{
  "block": {
    "attributes": {
      "alias_attributes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_verified_attributes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "email_verification_message": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "email_verification_subject": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "estimated_number_of_users": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mfa_configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sms_authentication_message": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sms_verification_message": {
        "computed": true,
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
      },
      "username_attributes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "account_recovery_setting": {
        "block": {
          "block_types": {
            "recovery_mechanism": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "priority": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
        "nesting_mode": "list"
      },
      "admin_create_user_config": {
        "block": {
          "attributes": {
            "allow_admin_create_user_only": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "invite_message_template": {
              "block": {
                "attributes": {
                  "email_message": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "email_subject": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sms_message": {
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
      "device_configuration": {
        "block": {
          "attributes": {
            "challenge_required_on_new_device": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "device_only_remembered_on_user_prompt": {
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
      "email_configuration": {
        "block": {
          "attributes": {
            "configuration_set": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email_sending_account": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "from_email_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reply_to_email_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_arn": {
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
      "lambda_config": {
        "block": {
          "attributes": {
            "create_auth_challenge": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "custom_message": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "define_auth_challenge": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "post_authentication": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "post_confirmation": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pre_authentication": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pre_sign_up": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pre_token_generation": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_migration": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "verify_auth_challenge_response": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "custom_email_sender": {
              "block": {
                "attributes": {
                  "lambda_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "lambda_version": {
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
            "custom_sms_sender": {
              "block": {
                "attributes": {
                  "lambda_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "lambda_version": {
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
      "password_policy": {
        "block": {
          "attributes": {
            "minimum_length": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "require_lowercase": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "require_numbers": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "require_symbols": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "require_uppercase": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "temporary_password_validity_days": {
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
      "schema": {
        "block": {
          "attributes": {
            "attribute_data_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "developer_only_attribute": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "mutable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "required": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "number_attribute_constraints": {
              "block": {
                "attributes": {
                  "max_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_value": {
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
            "string_attribute_constraints": {
              "block": {
                "attributes": {
                  "max_length": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_length": {
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
        "max_items": 50,
        "nesting_mode": "set"
      },
      "sms_configuration": {
        "block": {
          "attributes": {
            "external_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sns_caller_arn": {
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
      "software_token_mfa_configuration": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "user_pool_add_ons": {
        "block": {
          "attributes": {
            "advanced_security_mode": {
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
      "username_configuration": {
        "block": {
          "attributes": {
            "case_sensitive": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "verification_message_template": {
        "block": {
          "attributes": {
            "default_email_option": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email_message": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email_message_by_link": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email_subject": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email_subject_by_link": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sms_message": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCognitoUserPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoUserPool), &result)
	return &result
}

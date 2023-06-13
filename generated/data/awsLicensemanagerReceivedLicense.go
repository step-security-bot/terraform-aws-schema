package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLicensemanagerReceivedLicense = `{
  "block": {
    "attributes": {
      "beneficiary": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "consumption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "borrow_configuration": [
                "list",
                [
                  "object",
                  {
                    "allow_early_check_in": "bool",
                    "max_time_to_live_in_minutes": "number"
                  }
                ]
              ],
              "provisional_configuration": [
                "list",
                [
                  "object",
                  {
                    "max_time_to_live_in_minutes": "number"
                  }
                ]
              ],
              "renew_type": "string"
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "entitlements": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "allow_check_in": "bool",
              "max_count": "number",
              "name": "string",
              "unit": "string",
              "value": "string"
            }
          ]
        ]
      },
      "home_region": {
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
      "issuer": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "key_fingerprint": "string",
              "name": "string",
              "sign_key": "string"
            }
          ]
        ]
      },
      "license_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "license_metadata": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "value": "string"
            }
          ]
        ]
      },
      "license_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_sku": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "received_metadata": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allowed_operations": [
                "set",
                "string"
              ],
              "received_status": "string",
              "received_status_reason": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "validity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "begin": "string",
              "end": "string"
            }
          ]
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLicensemanagerReceivedLicenseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLicensemanagerReceivedLicense), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLicensemanagerGrant = `{
  "block": {
    "attributes": {
      "allowed_operations": {
        "description": "Allowed operations for the grant. This is a subset of the allowed operations on the license.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "home_region": {
        "computed": true,
        "description": "Home Region of the grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_arn": {
        "description": "License ARN.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the grant.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent_arn": {
        "computed": true,
        "description": "Parent ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "description": "The grantee principal ARN. The target account for the grant in the form of the ARN for an account principal of the root user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Grant status.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "Grant version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLicensemanagerGrantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLicensemanagerGrant), &result)
	return &result
}

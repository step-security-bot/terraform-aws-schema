package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsImagebuilderDistributionConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "date_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "date_updated": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "distribution": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "ami_distribution_configuration": [
                "set",
                [
                  "object",
                  {
                    "ami_tags": [
                      "map",
                      "string"
                    ],
                    "description": "string",
                    "kms_key_id": "string",
                    "launch_permission": [
                      "set",
                      [
                        "object",
                        {
                          "user_groups": [
                            "set",
                            "string"
                          ],
                          "user_ids": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "name": "string",
                    "target_account_ids": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "license_configuration_arns": [
                "set",
                "string"
              ],
              "region": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
  "version": 0
}`

func AwsImagebuilderDistributionConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsImagebuilderDistributionConfiguration), &result)
	return &result
}

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
                          "organization_arns": [
                            "set",
                            "string"
                          ],
                          "organizational_unit_arns": [
                            "set",
                            "string"
                          ],
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
              "container_distribution_configuration": [
                "set",
                [
                  "object",
                  {
                    "container_tags": [
                      "set",
                      "string"
                    ],
                    "description": "string",
                    "target_repository": [
                      "set",
                      [
                        "object",
                        {
                          "repository_name": "string",
                          "service": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "fast_launch_configuration": [
                "set",
                [
                  "object",
                  {
                    "account_id": "string",
                    "enabled": "bool",
                    "launch_template": [
                      "set",
                      [
                        "object",
                        {
                          "launch_template_id": "string",
                          "launch_template_name": "string",
                          "launch_template_version": "string"
                        }
                      ]
                    ],
                    "max_parallel_launches": "number",
                    "snapshot_configuration": [
                      "set",
                      [
                        "object",
                        {
                          "target_resource_count": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "launch_template_configuration": [
                "set",
                [
                  "object",
                  {
                    "account_id": "string",
                    "default": "bool",
                    "launch_template_id": "string"
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

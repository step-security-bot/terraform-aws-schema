package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConnectInstanceStorageConfig = `{
  "block": {
    "attributes": {
      "association_id": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kinesis_firehose_config": [
                "list",
                [
                  "object",
                  {
                    "firehose_arn": "string"
                  }
                ]
              ],
              "kinesis_stream_config": [
                "list",
                [
                  "object",
                  {
                    "stream_arn": "string"
                  }
                ]
              ],
              "kinesis_video_stream_config": [
                "list",
                [
                  "object",
                  {
                    "encryption_config": [
                      "list",
                      [
                        "object",
                        {
                          "encryption_type": "string",
                          "key_id": "string"
                        }
                      ]
                    ],
                    "prefix": "string",
                    "retention_period_hours": "number"
                  }
                ]
              ],
              "s3_config": [
                "list",
                [
                  "object",
                  {
                    "bucket_name": "string",
                    "bucket_prefix": "string",
                    "encryption_config": [
                      "list",
                      [
                        "object",
                        {
                          "encryption_type": "string",
                          "key_id": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "storage_type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConnectInstanceStorageConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConnectInstanceStorageConfig), &result)
	return &result
}

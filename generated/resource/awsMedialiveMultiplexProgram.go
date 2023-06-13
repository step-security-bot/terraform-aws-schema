package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMedialiveMultiplexProgram = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "multiplex_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "program_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "multiplex_program_settings": {
        "block": {
          "attributes": {
            "preferred_channel_pipeline": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "program_number": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "service_descriptor": {
              "block": {
                "attributes": {
                  "provider_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "service_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "video_settings": {
              "block": {
                "attributes": {
                  "constant_bitrate": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "statmux_settings": {
                    "block": {
                      "attributes": {
                        "maximum_bitrate": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "minimum_bitrate": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "priority": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMedialiveMultiplexProgramSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMedialiveMultiplexProgram), &result)
	return &result
}

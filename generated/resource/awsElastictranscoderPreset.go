package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElastictranscoderPreset = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "optional": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "video_codec_options": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "audio": {
        "block": {
          "attributes": {
            "audio_packing_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bit_rate": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "channels": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "codec": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sample_rate": {
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
      "audio_codec_options": {
        "block": {
          "attributes": {
            "bit_depth": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bit_order": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "profile": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "signed": {
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
      "thumbnails": {
        "block": {
          "attributes": {
            "aspect_ratio": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_height": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_width": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "padding_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resolution": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sizing_policy": {
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
      "video": {
        "block": {
          "attributes": {
            "aspect_ratio": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bit_rate": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "codec": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "display_aspect_ratio": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fixed_gop": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "frame_rate": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "keyframes_max_dist": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_frame_rate": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_height": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_width": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "padding_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resolution": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sizing_policy": {
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
      "video_watermarks": {
        "block": {
          "attributes": {
            "horizontal_align": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "horizontal_offset": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_height": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_width": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "opacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sizing_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vertical_align": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vertical_offset": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElastictranscoderPresetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElastictranscoderPreset), &result)
	return &result
}

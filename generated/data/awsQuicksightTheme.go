package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsQuicksightTheme = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "base_theme_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_color_palette": [
                "list",
                [
                  "object",
                  {
                    "colors": [
                      "list",
                      "string"
                    ],
                    "empty_fill_color": "string",
                    "min_max_gradient": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "sheet": [
                "list",
                [
                  "object",
                  {
                    "tile": [
                      "list",
                      [
                        "object",
                        {
                          "border": [
                            "list",
                            [
                              "object",
                              {
                                "show": "bool"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "tile_layout": [
                      "list",
                      [
                        "object",
                        {
                          "gutter": [
                            "list",
                            [
                              "object",
                              {
                                "show": "bool"
                              }
                            ]
                          ],
                          "margin": [
                            "list",
                            [
                              "object",
                              {
                                "show": "bool"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "typography": [
                "list",
                [
                  "object",
                  {
                    "font_families": [
                      "list",
                      [
                        "object",
                        {
                          "font_family": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "ui_color_palette": [
                "list",
                [
                  "object",
                  {
                    "accent": "string",
                    "accent_foreground": "string",
                    "danger": "string",
                    "danger_foreground": "string",
                    "dimension": "string",
                    "dimension_foreground": "string",
                    "measure": "string",
                    "measure_foreground": "string",
                    "primary_background": "string",
                    "primary_foreground": "string",
                    "secondary_background": "string",
                    "secondary_foreground": "string",
                    "success": "string",
                    "success_foreground": "string",
                    "warning": "string",
                    "warning_foreground": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "created_time": {
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
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "actions": [
                "set",
                "string"
              ],
              "principal": "string"
            }
          ]
        ]
      },
      "status": {
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
      },
      "theme_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsQuicksightThemeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsQuicksightTheme), &result)
	return &result
}

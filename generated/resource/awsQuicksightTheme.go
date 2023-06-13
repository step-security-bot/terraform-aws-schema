package resource

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
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
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
      "theme_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "configuration": {
        "block": {
          "block_types": {
            "data_color_palette": {
              "block": {
                "attributes": {
                  "colors": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "empty_fill_color": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_max_gradient": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sheet": {
              "block": {
                "block_types": {
                  "tile": {
                    "block": {
                      "block_types": {
                        "border": {
                          "block": {
                            "attributes": {
                              "show": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
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
                  "tile_layout": {
                    "block": {
                      "block_types": {
                        "gutter": {
                          "block": {
                            "attributes": {
                              "show": {
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
                        "margin": {
                          "block": {
                            "attributes": {
                              "show": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "typography": {
              "block": {
                "block_types": {
                  "font_families": {
                    "block": {
                      "attributes": {
                        "font_family": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 5,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ui_color_palette": {
              "block": {
                "attributes": {
                  "accent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "accent_foreground": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "danger": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "danger_foreground": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dimension": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dimension_foreground": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "measure": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "measure_foreground": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "primary_background": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "primary_foreground": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secondary_background": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secondary_foreground": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "success": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "success_foreground": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "warning": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "warning_foreground": {
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
      "permissions": {
        "block": {
          "attributes": {
            "actions": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "principal": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 64,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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
  "version": 0
}`

func AwsQuicksightThemeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsQuicksightTheme), &result)
	return &result
}

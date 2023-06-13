package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLexIntent = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "checksum": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "created_date": {
        "computed": true,
        "description_kind": "plain",
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
      "last_updated_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent_intent_signature": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sample_utterances": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "conclusion_statement": {
        "block": {
          "attributes": {
            "response_card": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "message": {
              "block": {
                "attributes": {
                  "content": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "content_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "group_number": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 15,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "confirmation_prompt": {
        "block": {
          "attributes": {
            "max_attempts": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "response_card": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "message": {
              "block": {
                "attributes": {
                  "content": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "content_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "group_number": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 15,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "dialog_code_hook": {
        "block": {
          "attributes": {
            "message_version": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "uri": {
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
      "follow_up_prompt": {
        "block": {
          "block_types": {
            "prompt": {
              "block": {
                "attributes": {
                  "max_attempts": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "response_card": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "message": {
                    "block": {
                      "attributes": {
                        "content": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "content_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "group_number": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 15,
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "rejection_statement": {
              "block": {
                "attributes": {
                  "response_card": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "message": {
                    "block": {
                      "attributes": {
                        "content": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "content_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "group_number": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 15,
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "fulfillment_activity": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "code_hook": {
              "block": {
                "attributes": {
                  "message_version": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "uri": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "rejection_statement": {
        "block": {
          "attributes": {
            "response_card": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "message": {
              "block": {
                "attributes": {
                  "content": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "content_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "group_number": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 15,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "slot": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "response_card": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sample_utterances": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "slot_constraint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "slot_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "slot_type_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "value_elicitation_prompt": {
              "block": {
                "attributes": {
                  "max_attempts": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "response_card": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "message": {
                    "block": {
                      "attributes": {
                        "content": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "content_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "group_number": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 15,
                    "min_items": 1,
                    "nesting_mode": "set"
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
        "max_items": 100,
        "nesting_mode": "set"
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

func AwsLexIntentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLexIntent), &result)
	return &result
}

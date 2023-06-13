package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsChimesdkmediapipelinesMediaInsightsPipelineConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_access_role_arn": {
        "description_kind": "plain",
        "required": true,
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
      }
    },
    "block_types": {
      "elements": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "amazon_transcribe_call_analytics_processor_configuration": {
              "block": {
                "attributes": {
                  "call_analytics_stream_categories": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "content_identification_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "content_redaction_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_partial_results_stabilization": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "filter_partial_results": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "language_code": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "language_model_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partial_results_stability": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pii_entity_types": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_filter_method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_filter_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "post_call_analytics_settings": {
                    "block": {
                      "attributes": {
                        "content_redaction_output": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "data_access_role_arn": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "output_encryption_kms_key_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "output_location": {
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
              "nesting_mode": "list"
            },
            "amazon_transcribe_processor_configuration": {
              "block": {
                "attributes": {
                  "content_identification_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "content_redaction_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_partial_results_stabilization": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "filter_partial_results": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "language_code": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "language_model_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partial_results_stability": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pii_entity_types": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "show_speaker_label": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "vocabulary_filter_method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_filter_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_name": {
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
            "kinesis_data_stream_sink_configuration": {
              "block": {
                "attributes": {
                  "insights_target": {
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
            "lambda_function_sink_configuration": {
              "block": {
                "attributes": {
                  "insights_target": {
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
            "s3_recording_sink_configuration": {
              "block": {
                "attributes": {
                  "destination": {
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
            "sns_topic_sink_configuration": {
              "block": {
                "attributes": {
                  "insights_target": {
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
            "sqs_queue_sink_configuration": {
              "block": {
                "attributes": {
                  "insights_target": {
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
            "voice_analytics_processor_configuration": {
              "block": {
                "attributes": {
                  "speaker_search_status": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "voice_tone_analysis_status": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "real_time_alert_configuration": {
        "block": {
          "attributes": {
            "disabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "rules": {
              "block": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "issue_detection_configuration": {
                    "block": {
                      "attributes": {
                        "rule_name": {
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
                  "keyword_match_configuration": {
                    "block": {
                      "attributes": {
                        "keywords": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "negate": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "rule_name": {
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
                  "sentiment_configuration": {
                    "block": {
                      "attributes": {
                        "rule_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "sentiment_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "time_period": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
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
              "max_items": 3,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AwsChimesdkmediapipelinesMediaInsightsPipelineConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsChimesdkmediapipelinesMediaInsightsPipelineConfiguration), &result)
	return &result
}

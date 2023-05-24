package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAutoscalingPolicy = `{
  "block": {
    "attributes": {
      "adjustment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "autoscaling_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cooldown": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "estimated_instance_warmup": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metric_aggregation_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_adjustment_magnitude": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_adjustment": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "predictive_scaling_configuration": {
        "block": {
          "attributes": {
            "max_capacity_breach_behavior": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_capacity_buffer": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scheduling_buffer_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "metric_specification": {
              "block": {
                "attributes": {
                  "target_value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "customized_capacity_metric_specification": {
                    "block": {
                      "block_types": {
                        "metric_data_queries": {
                          "block": {
                            "attributes": {
                              "expression": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "label": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "return_data": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "metric_stat": {
                                "block": {
                                  "attributes": {
                                    "stat": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "unit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "metric": {
                                      "block": {
                                        "attributes": {
                                          "metric_name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "namespace": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "dimensions": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "description_kind": "plain",
                                                  "required": true,
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
                                      "max_items": 1,
                                      "min_items": 1,
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
                          "max_items": 10,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "customized_load_metric_specification": {
                    "block": {
                      "block_types": {
                        "metric_data_queries": {
                          "block": {
                            "attributes": {
                              "expression": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "label": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "return_data": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "metric_stat": {
                                "block": {
                                  "attributes": {
                                    "stat": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "unit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "metric": {
                                      "block": {
                                        "attributes": {
                                          "metric_name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "namespace": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "dimensions": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "description_kind": "plain",
                                                  "required": true,
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
                                      "max_items": 1,
                                      "min_items": 1,
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
                          "max_items": 10,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "customized_scaling_metric_specification": {
                    "block": {
                      "block_types": {
                        "metric_data_queries": {
                          "block": {
                            "attributes": {
                              "expression": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "label": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "return_data": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "metric_stat": {
                                "block": {
                                  "attributes": {
                                    "stat": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "unit": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "metric": {
                                      "block": {
                                        "attributes": {
                                          "metric_name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "namespace": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "dimensions": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "description_kind": "plain",
                                                  "required": true,
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
                                      "max_items": 1,
                                      "min_items": 1,
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
                          "max_items": 10,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "predefined_load_metric_specification": {
                    "block": {
                      "attributes": {
                        "predefined_metric_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "resource_label": {
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
                  "predefined_metric_pair_specification": {
                    "block": {
                      "attributes": {
                        "predefined_metric_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "resource_label": {
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
                  "predefined_scaling_metric_specification": {
                    "block": {
                      "attributes": {
                        "predefined_metric_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "resource_label": {
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "step_adjustment": {
        "block": {
          "attributes": {
            "metric_interval_lower_bound": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_interval_upper_bound": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scaling_adjustment": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "target_tracking_configuration": {
        "block": {
          "attributes": {
            "disable_scale_in": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "target_value": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "customized_metric_specification": {
              "block": {
                "attributes": {
                  "metric_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "namespace": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "statistic": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "unit": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "metric_dimension": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "metrics": {
                    "block": {
                      "attributes": {
                        "expression": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "label": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "return_data": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "metric_stat": {
                          "block": {
                            "attributes": {
                              "stat": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "unit": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "metric": {
                                "block": {
                                  "attributes": {
                                    "metric_name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "namespace": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "dimensions": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "description_kind": "plain",
                                            "required": true,
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
                                "max_items": 1,
                                "min_items": 1,
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
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "predefined_metric_specification": {
              "block": {
                "attributes": {
                  "predefined_metric_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "resource_label": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAutoscalingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAutoscalingPolicy), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecurityhubInsight = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_by_attribute": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "filters": {
        "block": {
          "block_types": {
            "aws_account_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "company_name": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "compliance_status": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "confidence": {
              "block": {
                "attributes": {
                  "eq": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "created_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "criticality": {
              "block": {
                "attributes": {
                  "eq": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "description": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "finding_provider_fields_confidence": {
              "block": {
                "attributes": {
                  "eq": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "finding_provider_fields_criticality": {
              "block": {
                "attributes": {
                  "eq": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "finding_provider_fields_related_findings_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "finding_provider_fields_related_findings_product_arn": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "finding_provider_fields_severity_label": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "finding_provider_fields_severity_original": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "finding_provider_fields_types": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "first_observed_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "generator_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "keyword": {
              "block": {
                "attributes": {
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "last_observed_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "malware_name": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "malware_path": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "malware_state": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "malware_type": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_destination_domain": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_destination_ipv4": {
              "block": {
                "attributes": {
                  "cidr": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_destination_ipv6": {
              "block": {
                "attributes": {
                  "cidr": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_destination_port": {
              "block": {
                "attributes": {
                  "eq": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_direction": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_protocol": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_source_domain": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_source_ipv4": {
              "block": {
                "attributes": {
                  "cidr": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_source_ipv6": {
              "block": {
                "attributes": {
                  "cidr": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_source_mac": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "network_source_port": {
              "block": {
                "attributes": {
                  "eq": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "note_text": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "note_updated_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "note_updated_by": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "process_launched_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "process_name": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "process_parent_pid": {
              "block": {
                "attributes": {
                  "eq": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "process_path": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "process_pid": {
              "block": {
                "attributes": {
                  "eq": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lte": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "process_terminated_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "product_arn": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "product_fields": {
              "block": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "product_name": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "recommendation_text": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "record_state": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "related_findings_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "related_findings_product_arn": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_ec2_instance_iam_instance_profile_arn": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_ec2_instance_image_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_ec2_instance_ipv4_addresses": {
              "block": {
                "attributes": {
                  "cidr": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_ec2_instance_ipv6_addresses": {
              "block": {
                "attributes": {
                  "cidr": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_ec2_instance_key_name": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_ec2_instance_launched_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_ec2_instance_subnet_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_ec2_instance_type": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_ec2_instance_vpc_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_iam_access_key_created_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_iam_access_key_status": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_iam_access_key_user_name": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_s3_bucket_owner_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_aws_s3_bucket_owner_name": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_container_image_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_container_image_name": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_container_launched_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_container_name": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_details_other": {
              "block": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_id": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_partition": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_region": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_tags": {
              "block": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "resource_type": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "severity_label": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "source_url": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "threat_intel_indicator_category": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "threat_intel_indicator_last_observed_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "threat_intel_indicator_source": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "threat_intel_indicator_source_url": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "threat_intel_indicator_type": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "threat_intel_indicator_value": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "title": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "type": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "updated_at": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "date_range": {
                    "block": {
                      "attributes": {
                        "unit": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "user_defined_values": {
              "block": {
                "attributes": {
                  "comparison": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "verification_state": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
              "nesting_mode": "set"
            },
            "workflow_status": {
              "block": {
                "attributes": {
                  "comparison": {
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
              "max_items": 20,
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
  "version": 0
}`

func AwsSecurityhubInsightSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecurityhubInsight), &result)
	return &result
}

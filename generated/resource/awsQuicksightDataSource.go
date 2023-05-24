package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsQuicksightDataSource = `{
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
      "data_source_id": {
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
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "credentials": {
        "block": {
          "attributes": {
            "copy_source_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "credential_pair": {
              "block": {
                "attributes": {
                  "password": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "username": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
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
      "parameters": {
        "block": {
          "block_types": {
            "amazon_elasticsearch": {
              "block": {
                "attributes": {
                  "domain": {
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
            "athena": {
              "block": {
                "attributes": {
                  "work_group": {
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
            "aurora": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "aurora_postgresql": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "aws_iot_analytics": {
              "block": {
                "attributes": {
                  "data_set_name": {
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
            "jira": {
              "block": {
                "attributes": {
                  "site_base_url": {
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
            "maria_db": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "mysql": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oracle": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "postgresql": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "presto": {
              "block": {
                "attributes": {
                  "catalog": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "rds": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "instance_id": {
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
            "redshift": {
              "block": {
                "attributes": {
                  "cluster_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "s3": {
              "block": {
                "block_types": {
                  "manifest_file_location": {
                    "block": {
                      "attributes": {
                        "bucket": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
            "service_now": {
              "block": {
                "attributes": {
                  "site_base_url": {
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
            "snowflake": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "warehouse": {
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
            "spark": {
              "block": {
                "attributes": {
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sql_server": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "teradata": {
              "block": {
                "attributes": {
                  "database": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "twitter": {
              "block": {
                "attributes": {
                  "max_rows": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "query": {
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
      "permission": {
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
        "nesting_mode": "set"
      },
      "ssl_properties": {
        "block": {
          "attributes": {
            "disable_ssl": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "vpc_connection_properties": {
        "block": {
          "attributes": {
            "vpc_connection_arn": {
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
  "version": 0
}`

func AwsQuicksightDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsQuicksightDataSource), &result)
	return &result
}

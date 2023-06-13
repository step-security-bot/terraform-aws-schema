package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCodebuildProject = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "badge_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "badge_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "build_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "concurrent_build_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key": {
        "computed": true,
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_visibility": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_project_alias": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "queued_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_access_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_version": {
        "description_kind": "plain",
        "optional": true,
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
      "artifacts": {
        "block": {
          "attributes": {
            "artifact_identifier": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bucket_owner_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encryption_disabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "location": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "override_artifact_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "packaging": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
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
      },
      "build_batch_config": {
        "block": {
          "attributes": {
            "combine_artifacts": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "service_role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "timeout_in_mins": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "restrictions": {
              "block": {
                "attributes": {
                  "compute_types_allowed": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "maximum_builds_allowed": {
                    "description_kind": "plain",
                    "optional": true,
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cache": {
        "block": {
          "attributes": {
            "location": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "modes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "type": {
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
      "environment": {
        "block": {
          "attributes": {
            "certificate": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "compute_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "image": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "image_pull_credentials_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "privileged_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "environment_variable": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
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
            "registry_credential": {
              "block": {
                "attributes": {
                  "credential": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "credential_provider": {
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
      "file_system_locations": {
        "block": {
          "attributes": {
            "identifier": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "location": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_options": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_point": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "logs_config": {
        "block": {
          "block_types": {
            "cloudwatch_logs": {
              "block": {
                "attributes": {
                  "group_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "stream_name": {
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
            "s3_logs": {
              "block": {
                "attributes": {
                  "bucket_owner_access": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "encryption_disabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status": {
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
      "secondary_artifacts": {
        "block": {
          "attributes": {
            "artifact_identifier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "bucket_owner_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encryption_disabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "location": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "namespace_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "override_artifact_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "packaging": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 12,
        "nesting_mode": "set"
      },
      "secondary_source_version": {
        "block": {
          "attributes": {
            "source_identifier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_version": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 12,
        "nesting_mode": "set"
      },
      "secondary_sources": {
        "block": {
          "attributes": {
            "buildspec": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "git_clone_depth": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "insecure_ssl": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "location": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "report_build_status": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "source_identifier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "build_status_config": {
              "block": {
                "attributes": {
                  "context": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_url": {
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
            "git_submodules_config": {
              "block": {
                "attributes": {
                  "fetch_submodules": {
                    "description_kind": "plain",
                    "required": true,
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
        "max_items": 12,
        "nesting_mode": "set"
      },
      "source": {
        "block": {
          "attributes": {
            "buildspec": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "git_clone_depth": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "insecure_ssl": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "location": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "report_build_status": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "build_status_config": {
              "block": {
                "attributes": {
                  "context": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_url": {
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
            "git_submodules_config": {
              "block": {
                "attributes": {
                  "fetch_submodules": {
                    "description_kind": "plain",
                    "required": true,
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "vpc_config": {
        "block": {
          "attributes": {
            "security_group_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnets": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "vpc_id": {
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

func AwsCodebuildProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCodebuildProject), &result)
	return &result
}

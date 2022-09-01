// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_macie_allow_list", allowListDataSourceType)
}

// allowListDataSourceType returns the Terraform awscc_macie_allow_list data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Macie::AllowList resource type.
func allowListDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "AllowList ARN.",
			//   "type": "string"
			// }
			Description: "AllowList ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"criteria": {
			// Property: Criteria
			// CloudFormation resource type schema:
			// {
			//   "description": "AllowList criteria.",
			//   "properties": {
			//     "Regex": {
			//       "description": "The S3 object key for the AllowList.",
			//       "type": "string"
			//     },
			//     "S3WordsList": {
			//       "additionalProperties": false,
			//       "description": "The S3 location for the AllowList.",
			//       "properties": {
			//         "BucketName": {
			//           "type": "string"
			//         },
			//         "ObjectKey": {
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "BucketName",
			//         "ObjectKey"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "AllowList criteria.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"regex": {
						// Property: Regex
						Description: "The S3 object key for the AllowList.",
						Type:        types.StringType,
						Computed:    true,
					},
					"s3_words_list": {
						// Property: S3WordsList
						Description: "The S3 location for the AllowList.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket_name": {
									// Property: BucketName
									Type:     types.StringType,
									Computed: true,
								},
								"object_key": {
									// Property: ObjectKey
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of AllowList.",
			//   "type": "string"
			// }
			Description: "Description of AllowList.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "AllowList ID.",
			//   "type": "string"
			// }
			Description: "AllowList ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of AllowList.",
			//   "type": "string"
			// }
			Description: "Name of AllowList.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "AllowList status.",
			//   "enum": [
			//     "OK",
			//     "S3_OBJECT_NOT_FOUND",
			//     "S3_USER_ACCESS_DENIED",
			//     "S3_OBJECT_ACCESS_DENIED",
			//     "S3_THROTTLED",
			//     "S3_OBJECT_OVERSIZE",
			//     "S3_OBJECT_EMPTY",
			//     "UNKNOWN_ERROR"
			//   ],
			//   "type": "string"
			// }
			Description: "AllowList status.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The tag's key.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A collection of tags associated with a resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The tag's key.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The tag's value.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Macie::AllowList",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::AllowList").WithTerraformTypeName("awscc_macie_allow_list")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":           "Arn",
		"bucket_name":   "BucketName",
		"criteria":      "Criteria",
		"description":   "Description",
		"id":            "Id",
		"key":           "Key",
		"name":          "Name",
		"object_key":    "ObjectKey",
		"regex":         "Regex",
		"s3_words_list": "S3WordsList",
		"status":        "Status",
		"tags":          "Tags",
		"value":         "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}

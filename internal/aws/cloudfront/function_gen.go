// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_cloudfront_function", functionResourceType)
}

// functionResourceType returns the Terraform awscc_cloudfront_function resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFront::Function resource type.
func functionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_publish": {
			// Property: AutoPublish
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
			// AutoPublish is a write-only attribute.
		},
		"function_arn": {
			// Property: FunctionARN
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"function_code": {
			// Property: FunctionCode
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			// FunctionCode is a write-only attribute.
		},
		"function_config": {
			// Property: FunctionConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Comment": {
			//       "type": "string"
			//     },
			//     "Runtime": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Comment",
			//     "Runtime"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"comment": {
						// Property: Comment
						Type:     types.StringType,
						Required: true,
					},
					"runtime": {
						// Property: Runtime
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"function_metadata": {
			// Property: FunctionMetadata
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "FunctionARN": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"function_arn": {
						// Property: FunctionARN
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"stage": {
			// Property: Stage
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::CloudFront::Function",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::Function").WithTerraformTypeName("awscc_cloudfront_function")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_publish":      "AutoPublish",
		"comment":           "Comment",
		"function_arn":      "FunctionARN",
		"function_code":     "FunctionCode",
		"function_config":   "FunctionConfig",
		"function_metadata": "FunctionMetadata",
		"name":              "Name",
		"runtime":           "Runtime",
		"stage":             "Stage",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/AutoPublish",
		"/properties/FunctionCode",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_cloudfront_function", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

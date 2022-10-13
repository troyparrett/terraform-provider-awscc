// Code generated by generators/resource/main.go; DO NOT EDIT.

package transfer

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_transfer_profile", profileResource)
}

// profileResource returns the Terraform awscc_transfer_profile resource.
// This Terraform resource corresponds to the CloudFormation AWS::Transfer::Profile resource.
func profileResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the unique Amazon Resource Name (ARN) for the profile.",
			//   "maxLength": 1600,
			//   "minLength": 20,
			//   "pattern": "arn:.*",
			//   "type": "string"
			// }
			Description: "Specifies the unique Amazon Resource Name (ARN) for the profile.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"as_2_id": {
			// Property: As2Id
			// CloudFormation resource type schema:
			// {
			//   "description": "AS2 identifier agreed with a trading partner.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "AS2 identifier agreed with a trading partner.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
		},
		"certificate_ids": {
			// Property: CertificateIds
			// CloudFormation resource type schema:
			// {
			//   "description": "List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages.",
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "A unique identifier for the certificate.",
			//     "maxLength": 22,
			//     "minLength": 22,
			//     "pattern": "^cert-([0-9a-f]{17})$",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringLenBetween(22, 22)),
				validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^cert-([0-9a-f]{17})$"), "")),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"profile_id": {
			// Property: ProfileId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the profile",
			//   "maxLength": 19,
			//   "minLength": 19,
			//   "pattern": "^p-([0-9a-f]{17})$",
			//   "type": "string"
			// }
			Description: "A unique identifier for the profile",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"profile_type": {
			// Property: ProfileType
			// CloudFormation resource type schema:
			// {
			//   "description": "Enum specifying whether the profile is local or associated with a trading partner.",
			//   "enum": [
			//     "LOCAL",
			//     "PARTNER"
			//   ],
			//   "type": "string"
			// }
			Description: "Enum specifying whether the profile is local or associated with a trading partner.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"LOCAL",
					"PARTNER",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Creates a key-value pair for a specific resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The name assigned to the tag that you create.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Contains one or more values that you assigned to the key name you create.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The name assigned to the tag that you create.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "Contains one or more values that you assigned to the key name you create.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Transfer::Profile",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Transfer::Profile").WithTerraformTypeName("awscc_transfer_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":             "Arn",
		"as_2_id":         "As2Id",
		"certificate_ids": "CertificateIds",
		"key":             "Key",
		"profile_id":      "ProfileId",
		"profile_type":    "ProfileType",
		"tags":            "Tags",
		"value":           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
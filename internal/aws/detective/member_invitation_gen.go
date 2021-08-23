// Code generated by generators/resource/main.go; DO NOT EDIT.

package detective

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
	registry.AddResourceTypeFactory("awscc_detective_member_invitation", memberInvitationResourceType)
}

// memberInvitationResourceType returns the Terraform awscc_detective_member_invitation resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Detective::MemberInvitation resource type.
func memberInvitationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"disable_email_notification": {
			// Property: DisableEmailNotification
			// CloudFormation resource type schema:
			// {
			//   "description": "When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.",
			//   "type": "boolean"
			// }
			Description: "When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"graph_arn": {
			// Property: GraphArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the graph to which the member account will be invited",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ARN of the graph to which the member account will be invited",
			Type:        types.StringType,
			Required:    true,
			// GraphArn is a force-new attribute.
		},
		"member_email_address": {
			// Property: MemberEmailAddress
			// CloudFormation resource type schema:
			// {
			//   "description": "The root email address for the account to be invited, for validation. Updating this field has no effect.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The root email address for the account to be invited, for validation. Updating this field has no effect.",
			Type:        types.StringType,
			Required:    true,
		},
		"member_id": {
			// Property: MemberId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS account ID to be invited to join the graph as a member",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The AWS account ID to be invited to join the graph as a member",
			Type:        types.StringType,
			Required:    true,
			// MemberId is a force-new attribute.
		},
		"message": {
			// Property: Message
			// CloudFormation resource type schema:
			// {
			//   "description": "A message to be included in the email invitation sent to the invited account. Updating this field has no effect.",
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A message to be included in the email invitation sent to the invited account. Updating this field has no effect.",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Detective::MemberInvitation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Detective::MemberInvitation").WithTerraformTypeName("awscc_detective_member_invitation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"disable_email_notification": "DisableEmailNotification",
		"graph_arn":                  "GraphArn",
		"member_email_address":       "MemberEmailAddress",
		"member_id":                  "MemberId",
		"message":                    "Message",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_detective_member_invitation", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

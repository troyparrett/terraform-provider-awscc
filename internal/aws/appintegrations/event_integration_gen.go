// Code generated by generators/resource/main.go; DO NOT EDIT.

package appintegrations

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
	registry.AddResourceTypeFactory("awscc_appintegrations_event_integration", eventIntegrationResourceType)
}

// eventIntegrationResourceType returns the Terraform awscc_appintegrations_event_integration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AppIntegrations::EventIntegration resource type.
func eventIntegrationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"associations": {
			// Property: Associations
			// CloudFormation resource type schema:
			// {
			//   "description": "The associations with the event integration.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ClientAssociationMetadata": {
			//         "description": "The metadata associated with the client.",
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "description": "A key to identify the metadata.",
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "Value": {
			//               "description": "Corresponding metadata value for the key.",
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "pattern": "",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Key",
			//             "Value"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array"
			//       },
			//       "ClientId": {
			//         "description": "The identifier for the client that is associated with the event integration.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "EventBridgeRuleName": {
			//         "description": "The name of the Eventbridge rule.",
			//         "maxLength": 2048,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "EventIntegrationAssociationArn": {
			//         "description": "The Amazon Resource Name (ARN) for the event integration association.",
			//         "maxLength": 2048,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "EventIntegrationAssociationId": {
			//         "description": "The identifier for the event integration association.",
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The associations with the event integration.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"client_association_metadata": {
						// Property: ClientAssociationMetadata
						Description: "The metadata associated with the client.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Description: "A key to identify the metadata.",
									Type:        types.StringType,
									Required:    true,
								},
								"value": {
									// Property: Value
									Description: "Corresponding metadata value for the key.",
									Type:        types.StringType,
									Required:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
					"client_id": {
						// Property: ClientId
						Description: "The identifier for the client that is associated with the event integration.",
						Type:        types.StringType,
						Optional:    true,
					},
					"event_bridge_rule_name": {
						// Property: EventBridgeRuleName
						Description: "The name of the Eventbridge rule.",
						Type:        types.StringType,
						Optional:    true,
					},
					"event_integration_association_arn": {
						// Property: EventIntegrationAssociationArn
						Description: "The Amazon Resource Name (ARN) for the event integration association.",
						Type:        types.StringType,
						Optional:    true,
					},
					"event_integration_association_id": {
						// Property: EventIntegrationAssociationId
						Description: "The identifier for the event integration association.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
				},
			),
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The event integration description.",
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The event integration description.",
			Type:        types.StringType,
			Optional:    true,
		},
		"event_bridge_bus": {
			// Property: EventBridgeBus
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Eventbridge bus for the event integration.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Eventbridge bus for the event integration.",
			Type:        types.StringType,
			Required:    true,
			// EventBridgeBus is a force-new attribute.
		},
		"event_filter": {
			// Property: EventFilter
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Source": {
			//       "description": "The source of the events.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Source"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"source": {
						// Property: Source
						Description: "The source of the events.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Required: true,
			// EventFilter is a force-new attribute.
		},
		"event_integration_arn": {
			// Property: EventIntegrationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the event integration.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the event integration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the event integration.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the event integration.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags (keys and values) associated with the event integration.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "A key to identify the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Corresponding tag value for the key.",
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
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags (keys and values) associated with the event integration.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A key to identify the tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "Corresponding tag value for the key.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 200,
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::AppIntegrations::EventIntegration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppIntegrations::EventIntegration").WithTerraformTypeName("awscc_appintegrations_event_integration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"associations":                      "Associations",
		"client_association_metadata":       "ClientAssociationMetadata",
		"client_id":                         "ClientId",
		"description":                       "Description",
		"event_bridge_bus":                  "EventBridgeBus",
		"event_bridge_rule_name":            "EventBridgeRuleName",
		"event_filter":                      "EventFilter",
		"event_integration_arn":             "EventIntegrationArn",
		"event_integration_association_arn": "EventIntegrationAssociationArn",
		"event_integration_association_id":  "EventIntegrationAssociationId",
		"key":                               "Key",
		"name":                              "Name",
		"source":                            "Source",
		"tags":                              "Tags",
		"value":                             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_appintegrations_event_integration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

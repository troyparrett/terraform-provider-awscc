// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

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
	registry.AddResourceTypeFactory("awscc_mediaconnect_flow_source", flowSourceResourceType)
}

// flowSourceResourceType returns the Terraform awscc_mediaconnect_flow_source resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MediaConnect::FlowSource resource type.
func flowSourceResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"decryption": {
			// Property: Decryption
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about the encryption of the flow.",
			//   "properties": {
			//     "Algorithm": {
			//       "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
			//       "enum": [
			//         "aes128",
			//         "aes192",
			//         "aes256"
			//       ],
			//       "type": "string"
			//     },
			//     "ConstantInitializationVector": {
			//       "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
			//       "type": "string"
			//     },
			//     "DeviceId": {
			//       "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//       "type": "string"
			//     },
			//     "KeyType": {
			//       "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
			//       "enum": [
			//         "speke",
			//         "static-key"
			//       ],
			//       "type": "string"
			//     },
			//     "Region": {
			//       "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//       "type": "string"
			//     },
			//     "ResourceId": {
			//       "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//       "type": "string"
			//     },
			//     "RoleArn": {
			//       "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
			//       "type": "string"
			//     },
			//     "SecretArn": {
			//       "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
			//       "type": "string"
			//     },
			//     "Url": {
			//       "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Algorithm",
			//     "RoleArn"
			//   ],
			//   "type": "object"
			// }
			Description: "Information about the encryption of the flow.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"algorithm": {
						// Property: Algorithm
						Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
						Type:        types.StringType,
						Required:    true,
					},
					"constant_initialization_vector": {
						// Property: ConstantInitializationVector
						Description: "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"device_id": {
						// Property: DeviceId
						Description: "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"key_type": {
						// Property: KeyType
						Description: "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
						Type:        types.StringType,
						Optional:    true,
					},
					"region": {
						// Property: Region
						Description: "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"resource_id": {
						// Property: ResourceId
						Description: "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"role_arn": {
						// Property: RoleArn
						Description: "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
						Type:        types.StringType,
						Required:    true,
					},
					"secret_arn": {
						// Property: SecretArn
						Description: " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
					"url": {
						// Property: Url
						Description: "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
			//   "type": "string"
			// }
			Description: "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
			Type:        types.StringType,
			Required:    true,
		},
		"entitlement_arn": {
			// Property: EntitlementArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
			//   "type": "string"
			// }
			Description: "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
			Type:        types.StringType,
			Optional:    true,
		},
		"flow_arn": {
			// Property: FlowArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the flow.",
			//   "type": "string"
			// }
			Description: "The ARN of the flow.",
			Type:        types.StringType,
			Optional:    true,
		},
		"ingest_ip": {
			// Property: IngestIp
			// CloudFormation resource type schema:
			// {
			//   "description": "The IP address that the flow will be listening on for incoming content.",
			//   "type": "string"
			// }
			Description: "The IP address that the flow will be listening on for incoming content.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ingest_port": {
			// Property: IngestPort
			// CloudFormation resource type schema:
			// {
			//   "description": "The port that the flow will be listening on for incoming content.",
			//   "type": "integer"
			// }
			Description: "The port that the flow will be listening on for incoming content.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"max_bitrate": {
			// Property: MaxBitrate
			// CloudFormation resource type schema:
			// {
			//   "description": "The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.",
			//   "type": "integer"
			// }
			Description: "The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"max_latency": {
			// Property: MaxLatency
			// CloudFormation resource type schema:
			// {
			//   "description": "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			//   "type": "integer"
			// }
			Description: "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the source.",
			//   "type": "string"
			// }
			Description: "The name of the source.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"protocol": {
			// Property: Protocol
			// CloudFormation resource type schema:
			// {
			//   "description": "The protocol that is used by the source.",
			//   "enum": [
			//     "zixi-push",
			//     "rtp-fec",
			//     "rtp",
			//     "rist"
			//   ],
			//   "type": "string"
			// }
			Description: "The protocol that is used by the source.",
			Type:        types.StringType,
			Optional:    true,
		},
		"source_arn": {
			// Property: SourceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the source.",
			//   "type": "string"
			// }
			Description: "The ARN of the source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"stream_id": {
			// Property: StreamId
			// CloudFormation resource type schema:
			// {
			//   "description": "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			//   "type": "string"
			// }
			Description: "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			Type:        types.StringType,
			Optional:    true,
		},
		"vpc_interface_name": {
			// Property: VpcInterfaceName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the VPC Interface this Source is configured with.",
			//   "type": "string"
			// }
			Description: "The name of the VPC Interface this Source is configured with.",
			Type:        types.StringType,
			Optional:    true,
		},
		"whitelist_cidr": {
			// Property: WhitelistCidr
			// CloudFormation resource type schema:
			// {
			//   "description": "The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			//   "type": "string"
			// }
			Description: "The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
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
		Description: "Resource schema for AWS::MediaConnect::FlowSource",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowSource").WithTerraformTypeName("awscc_mediaconnect_flow_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"algorithm":                      "Algorithm",
		"constant_initialization_vector": "ConstantInitializationVector",
		"decryption":                     "Decryption",
		"description":                    "Description",
		"device_id":                      "DeviceId",
		"entitlement_arn":                "EntitlementArn",
		"flow_arn":                       "FlowArn",
		"ingest_ip":                      "IngestIp",
		"ingest_port":                    "IngestPort",
		"key_type":                       "KeyType",
		"max_bitrate":                    "MaxBitrate",
		"max_latency":                    "MaxLatency",
		"name":                           "Name",
		"protocol":                       "Protocol",
		"region":                         "Region",
		"resource_id":                    "ResourceId",
		"role_arn":                       "RoleArn",
		"secret_arn":                     "SecretArn",
		"source_arn":                     "SourceArn",
		"stream_id":                      "StreamId",
		"url":                            "Url",
		"vpc_interface_name":             "VpcInterfaceName",
		"whitelist_cidr":                 "WhitelistCidr",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_mediaconnect_flow_source", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datasync_location_object_storage", locationObjectStorageDataSource)
}

// locationObjectStorageDataSource returns the Terraform awscc_datasync_location_object_storage data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataSync::LocationObjectStorage resource.
func locationObjectStorageDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^.+$",
		//	  "type": "string"
		//	}
		"access_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AgentArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 4,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"agent_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the bucket on the self-managed object storage server.",
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
		//	  "type": "string"
		//	}
		"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the bucket on the self-managed object storage server.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the location that is created.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"location_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the location that is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocationUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the object storage location that was described.",
		//	  "maxLength": 4356,
		//	  "pattern": "^(efs|nfs|s3|smb|fsxw|object-storage)://[a-zA-Z0-9./\\-]+$",
		//	  "type": "string"
		//	}
		"location_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the object storage location that was described.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecretKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
		//	  "maxLength": 200,
		//	  "minLength": 8,
		//	  "pattern": "^.+$",
		//	  "type": "string"
		//	}
		"secret_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServerCertificate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "X.509 PEM content containing a certificate authority or chain to trust.",
		//	  "maxLength": 32768,
		//	  "type": "string"
		//	}
		"server_certificate": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "X.509 PEM content containing a certificate authority or chain to trust.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServerHostname
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
		//	  "maxLength": 255,
		//	  "pattern": "^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$",
		//	  "type": "string"
		//	}
		"server_hostname": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServerPort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The port that your self-managed server accepts inbound network traffic on.",
		//	  "maximum": 65536,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"server_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The port that your self-managed server accepts inbound network traffic on.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServerProtocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The protocol that the object storage server uses to communicate.",
		//	  "enum": [
		//	    "HTTPS",
		//	    "HTTP"
		//	  ],
		//	  "type": "string"
		//	}
		"server_protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The protocol that the object storage server uses to communicate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Subdirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subdirectory in the self-managed object storage server that is used to read data from.",
		//	  "maxLength": 4096,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\p{Zs}]*$",
		//	  "type": "string"
		//	}
		"subdirectory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The subdirectory in the self-managed object storage server that is used to read data from.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key for an AWS resource tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for an AWS resource tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DataSync::LocationObjectStorage",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationObjectStorage").WithTerraformTypeName("awscc_datasync_location_object_storage")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_key":         "AccessKey",
		"agent_arns":         "AgentArns",
		"bucket_name":        "BucketName",
		"key":                "Key",
		"location_arn":       "LocationArn",
		"location_uri":       "LocationUri",
		"secret_key":         "SecretKey",
		"server_certificate": "ServerCertificate",
		"server_hostname":    "ServerHostname",
		"server_port":        "ServerPort",
		"server_protocol":    "ServerProtocol",
		"subdirectory":       "Subdirectory",
		"tags":               "Tags",
		"value":              "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

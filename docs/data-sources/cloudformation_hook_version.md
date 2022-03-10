---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudformation_hook_version Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CloudFormation::HookVersion
---

# awscc_cloudformation_hook_version (Data Source)

Data Source schema for AWS::CloudFormation::HookVersion



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource
- **execution_role_arn** (String) The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
- **is_default_version** (Boolean) Indicates if this type version is the current default version
- **logging_config** (Attributes) Specifies logging configuration information for a type. (see [below for nested schema](#nestedatt--logging_config))
- **schema_handler_package** (String) A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.

For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
- **type_arn** (String) The Amazon Resource Name (ARN) of the type without the versionID.
- **type_name** (String) The name of the type being registered.

We recommend that type names adhere to the following pattern: company_or_organization::service::type.
- **version_id** (String) The ID of the version of the type represented by this hook instance.
- **visibility** (String) The scope at which the type is visible and usable in CloudFormation operations.

Valid values include:

PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.

PUBLIC: The type is publically visible and usable within any Amazon account.

<a id="nestedatt--logging_config"></a>
### Nested Schema for `logging_config`

Read-Only:

- **log_group_name** (String) The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.
- **log_role_arn** (String) The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.


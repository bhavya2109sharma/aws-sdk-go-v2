// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A nested application summary.
type ApplicationDependencySummary struct {

	// The Amazon Resource Name (ARN) of the nested application.
	//
	// This member is required.
	ApplicationId *string

	// The semantic version of the nested application.
	//
	// This member is required.
	SemanticVersion *string

	noSmithyDocumentSerde
}

// Policy statement applied to the application.
type ApplicationPolicyStatement struct {

	// For the list of actions supported for this operation, see [Application Permissions].
	//
	// [Application Permissions]: https://docs.aws.amazon.com/serverlessrepo/latest/devguide/access-control-resource-based.html#application-permissions
	//
	// This member is required.
	Actions []string

	// An array of AWS account IDs, or * to make the application public.
	//
	// This member is required.
	Principals []string

	// An array of PrinciplalOrgIDs, which corresponds to AWS IAM [aws:PrincipalOrgID] global condition
	// key.
	//
	// [aws:PrincipalOrgID]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#principal-org-id
	PrincipalOrgIDs []string

	// A unique ID for the statement.
	StatementId *string

	noSmithyDocumentSerde
}

// Summary of details about the application.
type ApplicationSummary struct {

	// The application Amazon Resource Name (ARN).
	//
	// This member is required.
	ApplicationId *string

	// The name of the author publishing the app.
	//
	// Minimum length=1. Maximum length=127.
	//
	// Pattern "^[a-z0-9](([a-z0-9]|-(?!-))*[a-z0-9])?$";
	//
	// This member is required.
	Author *string

	// The description of the application.
	//
	// Minimum length=1. Maximum length=256
	//
	// This member is required.
	Description *string

	// The name of the application.
	//
	// Minimum length=1. Maximum length=140
	//
	// Pattern: "[a-zA-Z0-9\\-]+";
	//
	// This member is required.
	Name *string

	// The date and time this resource was created.
	CreationTime *string

	// A URL with more information about the application, for example the location of
	// your GitHub repository for the application.
	HomePageUrl *string

	// Labels to improve discovery of apps in search results.
	//
	// Minimum length=1. Maximum length=127. Maximum number of labels: 10
	//
	// Pattern: "^[a-zA-Z0-9+\\-_:\\/@]+$";
	Labels []string

	// A valid identifier from [https://spdx.org/licenses/].
	//
	// [https://spdx.org/licenses/]: https://spdx.org/licenses/
	SpdxLicenseId *string

	noSmithyDocumentSerde
}

// Parameters supported by the application.
type ParameterDefinition struct {

	// The name of the parameter.
	//
	// This member is required.
	Name *string

	// A list of AWS SAM resources that use this parameter.
	//
	// This member is required.
	ReferencedByResources []string

	// A regular expression that represents the patterns to allow for String types.
	AllowedPattern *string

	// An array containing the list of values allowed for the parameter.
	AllowedValues []string

	// A string that explains a constraint when the constraint is violated. For
	// example, without a constraint description, a parameter that has an allowed
	// pattern of [A-Za-z0-9]+ displays the following error message when the user
	// specifies an invalid value:
	//
	// Malformed input-Parameter MyParameter must match pattern [A-Za-z0-9]+
	//
	// By adding a constraint description, such as "must contain only uppercase and
	// lowercase letters and numbers," you can display the following customized error
	// message:
	//
	// Malformed input-Parameter MyParameter must contain only uppercase and lowercase
	// letters and numbers.
	ConstraintDescription *string

	// A value of the appropriate type for the template to use if no value is
	// specified when a stack is created. If you define constraints for the parameter,
	// you must specify a value that adheres to those constraints.
	DefaultValue *string

	// A string of up to 4,000 characters that describes the parameter.
	Description *string

	// An integer value that determines the largest number of characters that you want
	// to allow for String types.
	MaxLength *int32

	// A numeric value that determines the largest numeric value that you want to
	// allow for Number types.
	MaxValue *int32

	// An integer value that determines the smallest number of characters that you
	// want to allow for String types.
	MinLength *int32

	// A numeric value that determines the smallest numeric value that you want to
	// allow for Number types.
	MinValue *int32

	// Whether to mask the parameter value whenever anyone makes a call that describes
	// the stack. If you set the value to true, the parameter value is masked with
	// asterisks (*****).
	NoEcho *bool

	// The type of the parameter.
	//
	// Valid values: String | Number | List<Number> | CommaDelimitedList
	//
	// String: A literal string.
	//
	// For example, users can specify "MyUserName".
	//
	// Number: An integer or float. AWS CloudFormation validates the parameter value
	// as a number. However, when you use the parameter elsewhere in your template (for
	// example, by using the Ref intrinsic function), the parameter value becomes a
	// string.
	//
	// For example, users might specify "8888".
	//
	// List<Number>: An array of integers or floats that are separated by commas. AWS
	// CloudFormation validates the parameter value as numbers. However, when you use
	// the parameter elsewhere in your template (for example, by using the Ref
	// intrinsic function), the parameter value becomes a list of strings.
	//
	// For example, users might specify "80,20", and then Ref results in ["80","20"].
	//
	// CommaDelimitedList: An array of literal strings that are separated by commas.
	// The total number of strings should be one more than the total number of commas.
	// Also, each member string is space-trimmed.
	//
	// For example, users might specify "test,dev,prod", and then Ref results in
	// ["test","dev","prod"].
	Type *string

	noSmithyDocumentSerde
}

// Parameter value of the application.
type ParameterValue struct {

	// The key associated with the parameter. If you don't specify a key and value for
	// a particular parameter, AWS CloudFormation uses the default value that is
	// specified in your template.
	//
	// This member is required.
	Name *string

	// The input value associated with the parameter.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// This property corresponds to the AWS CloudFormation [RollbackConfiguration] Data Type.
//
// [RollbackConfiguration]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/RollbackConfiguration
type RollbackConfiguration struct {

	// This property corresponds to the content of the same name for the AWS
	// CloudFormation [RollbackConfiguration]Data Type.
	//
	// [RollbackConfiguration]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/RollbackConfiguration
	MonitoringTimeInMinutes *int32

	// This property corresponds to the content of the same name for the AWS
	// CloudFormation [RollbackConfiguration]Data Type.
	//
	// [RollbackConfiguration]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/RollbackConfiguration
	RollbackTriggers []RollbackTrigger

	noSmithyDocumentSerde
}

// This property corresponds to the AWS CloudFormation [RollbackTrigger] Data Type.
//
// [RollbackTrigger]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/RollbackTrigger
type RollbackTrigger struct {

	// This property corresponds to the content of the same name for the AWS
	// CloudFormation [RollbackTrigger]Data Type.
	//
	// [RollbackTrigger]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/RollbackTrigger
	//
	// This member is required.
	Arn *string

	// This property corresponds to the content of the same name for the AWS
	// CloudFormation [RollbackTrigger]Data Type.
	//
	// [RollbackTrigger]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/RollbackTrigger
	//
	// This member is required.
	Type *string

	noSmithyDocumentSerde
}

// This property corresponds to the AWS CloudFormation [Tag] Data Type.
//
// [Tag]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/Tag
type Tag struct {

	// This property corresponds to the content of the same name for the AWS
	// CloudFormation [Tag]Data Type.
	//
	// [Tag]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/Tag
	//
	// This member is required.
	Key *string

	// This property corresponds to the content of the same name for the AWS
	// CloudFormation [Tag]Data Type.
	//
	// [Tag]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/Tag
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Application version details.
type Version struct {

	// The application Amazon Resource Name (ARN).
	//
	// This member is required.
	ApplicationId *string

	// The date and time this resource was created.
	//
	// This member is required.
	CreationTime *string

	// An array of parameter types supported by the application.
	//
	// This member is required.
	ParameterDefinitions []ParameterDefinition

	// A list of values that you must specify before you can deploy certain
	// applications. Some applications might include resources that can affect
	// permissions in your AWS account, for example, by creating new AWS Identity and
	// Access Management (IAM) users. For those applications, you must explicitly
	// acknowledge their capabilities by specifying this parameter.
	//
	// The only valid values are CAPABILITY_IAM, CAPABILITY_NAMED_IAM,
	// CAPABILITY_RESOURCE_POLICY, and CAPABILITY_AUTO_EXPAND.
	//
	// The following resources require you to specify CAPABILITY_IAM or
	// CAPABILITY_NAMED_IAM: [AWS::IAM::Group], [AWS::IAM::InstanceProfile], [AWS::IAM::Policy], and [AWS::IAM::Role]. If the application contains IAM resources,
	// you can specify either CAPABILITY_IAM or CAPABILITY_NAMED_IAM. If the
	// application contains IAM resources with custom names, you must specify
	// CAPABILITY_NAMED_IAM.
	//
	// The following resources require you to specify CAPABILITY_RESOURCE_POLICY: [AWS::Lambda::Permission], [AWS::IAM:Policy], [AWS::ApplicationAutoScaling::ScalingPolicy]
	// , [AWS::S3::BucketPolicy], [AWS::SQS::QueuePolicy], and [AWS::SNS::TopicPolicy].
	//
	// Applications that contain one or more nested applications require you to
	// specify CAPABILITY_AUTO_EXPAND.
	//
	// If your application template contains any of the above resources, we recommend
	// that you review all permissions associated with the application before
	// deploying. If you don't specify this parameter for an application that requires
	// capabilities, the call will fail.
	//
	// [AWS::SQS::QueuePolicy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html
	// [AWS::SNS::TopicPolicy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
	// [AWS::ApplicationAutoScaling::ScalingPolicy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html
	// [AWS::S3::BucketPolicy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html
	// [AWS::IAM::InstanceProfile]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
	// [AWS::IAM::Policy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html
	// [AWS::IAM::Group]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html
	// [AWS::IAM::Role]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
	// [AWS::IAM:Policy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html
	// [AWS::Lambda::Permission]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html
	//
	// This member is required.
	RequiredCapabilities []Capability

	// Whether all of the AWS resources contained in this application are supported in
	// the region in which it is being retrieved.
	//
	// This member is required.
	ResourcesSupported *bool

	// The semantic version of the application:
	//
	// [https://semver.org/]
	//
	// [https://semver.org/]: https://semver.org/
	//
	// This member is required.
	SemanticVersion *string

	// A link to the packaged AWS SAM template of your application.
	//
	// This member is required.
	TemplateUrl *string

	// A link to the S3 object that contains the ZIP archive of the source code for
	// this version of your application.
	//
	// Maximum size 50 MB
	SourceCodeArchiveUrl *string

	// A link to a public repository for the source code of your application, for
	// example the URL of a specific GitHub commit.
	SourceCodeUrl *string

	noSmithyDocumentSerde
}

// An application version summary.
type VersionSummary struct {

	// The application Amazon Resource Name (ARN).
	//
	// This member is required.
	ApplicationId *string

	// The date and time this resource was created.
	//
	// This member is required.
	CreationTime *string

	// The semantic version of the application:
	//
	// [https://semver.org/]
	//
	// [https://semver.org/]: https://semver.org/
	//
	// This member is required.
	SemanticVersion *string

	// A link to a public repository for the source code of your application, for
	// example the URL of a specific GitHub commit.
	SourceCodeUrl *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

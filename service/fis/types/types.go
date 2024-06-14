// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Describes an action. For more information, see [FIS actions] in the Fault Injection Service
// User Guide.
//
// [FIS actions]: https://docs.aws.amazon.com/fis/latest/userguide/fis-actions-reference.html
type Action struct {

	// The Amazon Resource Name (ARN) of the action.
	Arn *string

	// The description for the action.
	Description *string

	// The ID of the action.
	Id *string

	// The action parameters, if applicable.
	Parameters map[string]ActionParameter

	// The tags for the action.
	Tags map[string]string

	// The supported targets for the action.
	Targets map[string]ActionTarget

	noSmithyDocumentSerde
}

// Describes a parameter for an action.
type ActionParameter struct {

	// The parameter description.
	Description *string

	// Indicates whether the parameter is required.
	Required *bool

	noSmithyDocumentSerde
}

// Provides a summary of an action.
type ActionSummary struct {

	// The Amazon Resource Name (ARN) of the action.
	Arn *string

	// The description for the action.
	Description *string

	// The ID of the action.
	Id *string

	// The tags for the action.
	Tags map[string]string

	// The targets for the action.
	Targets map[string]ActionTarget

	noSmithyDocumentSerde
}

// Describes a target for an action.
type ActionTarget struct {

	// The resource type of the target.
	ResourceType *string

	noSmithyDocumentSerde
}

// Specifies an action for an experiment template.
//
// For more information, see [Actions] in the Fault Injection Service User Guide.
//
// [Actions]: https://docs.aws.amazon.com/fis/latest/userguide/actions.html
type CreateExperimentTemplateActionInput struct {

	// The ID of the action. The format of the action ID is:
	// aws:service-name:action-type.
	//
	// This member is required.
	ActionId *string

	// A description for the action.
	Description *string

	// The parameters for the action, if applicable.
	Parameters map[string]string

	// The name of the action that must be completed before the current action starts.
	// Omit this parameter to run the action at the start of the experiment.
	StartAfter []string

	// The targets for the action.
	Targets map[string]string

	noSmithyDocumentSerde
}

// Specifies experiment options for an experiment template.
type CreateExperimentTemplateExperimentOptionsInput struct {

	// Specifies the account targeting setting for experiment options.
	AccountTargeting AccountTargeting

	// Specifies the empty target resolution mode for experiment options.
	EmptyTargetResolutionMode EmptyTargetResolutionMode

	noSmithyDocumentSerde
}

// Specifies the configuration for experiment logging.
type CreateExperimentTemplateLogConfigurationInput struct {

	// The schema version.
	//
	// This member is required.
	LogSchemaVersion *int32

	// The configuration for experiment logging to Amazon CloudWatch Logs.
	CloudWatchLogsConfiguration *ExperimentTemplateCloudWatchLogsLogConfigurationInput

	// The configuration for experiment logging to Amazon S3.
	S3Configuration *ExperimentTemplateS3LogConfigurationInput

	noSmithyDocumentSerde
}

// Specifies a stop condition for an experiment template.
type CreateExperimentTemplateStopConditionInput struct {

	// The source for the stop condition. Specify aws:cloudwatch:alarm if the stop
	// condition is defined by a CloudWatch alarm. Specify none if there is no stop
	// condition.
	//
	// This member is required.
	Source *string

	// The Amazon Resource Name (ARN) of the CloudWatch alarm. This is required if the
	// source is a CloudWatch alarm.
	Value *string

	noSmithyDocumentSerde
}

// Specifies a target for an experiment. You must specify at least one Amazon
// Resource Name (ARN) or at least one resource tag. You cannot specify both ARNs
// and tags.
//
// For more information, see [Targets] in the Fault Injection Service User Guide.
//
// [Targets]: https://docs.aws.amazon.com/fis/latest/userguide/targets.html
type CreateExperimentTemplateTargetInput struct {

	// The resource type. The resource type must be supported for the specified action.
	//
	// This member is required.
	ResourceType *string

	// Scopes the identified resources to a specific count of the resources at random,
	// or a percentage of the resources. All identified resources are included in the
	// target.
	//
	//   - ALL - Run the action on all identified targets. This is the default.
	//
	//   - COUNT(n) - Run the action on the specified number of targets, chosen from
	//   the identified targets at random. For example, COUNT(1) selects one of the
	//   targets.
	//
	//   - PERCENT(n) - Run the action on the specified percentage of targets, chosen
	//   from the identified targets at random. For example, PERCENT(25) selects 25% of
	//   the targets.
	//
	// This member is required.
	SelectionMode *string

	// The filters to apply to identify target resources using specific attributes.
	Filters []ExperimentTemplateTargetInputFilter

	// The resource type parameters.
	Parameters map[string]string

	// The Amazon Resource Names (ARNs) of the resources.
	ResourceArns []string

	// The tags for the target resources.
	ResourceTags map[string]string

	noSmithyDocumentSerde
}

// Describes an experiment.
type Experiment struct {

	// The actions for the experiment.
	Actions map[string]ExperimentAction

	// The Amazon Resource Name (ARN) of the experiment.
	Arn *string

	// The time that the experiment was created.
	CreationTime *time.Time

	// The time that the experiment ended.
	EndTime *time.Time

	// The experiment options for the experiment.
	ExperimentOptions *ExperimentOptions

	// The ID of the experiment template.
	ExperimentTemplateId *string

	// The ID of the experiment.
	Id *string

	// The configuration for experiment logging.
	LogConfiguration *ExperimentLogConfiguration

	// The Amazon Resource Name (ARN) of an IAM role that grants the FIS service
	// permission to perform service actions on your behalf.
	RoleArn *string

	// The time that the experiment started.
	StartTime *time.Time

	// The state of the experiment.
	State *ExperimentState

	// The stop conditions for the experiment.
	StopConditions []ExperimentStopCondition

	// The tags for the experiment.
	Tags map[string]string

	// The count of target account configurations for the experiment.
	TargetAccountConfigurationsCount *int64

	// The targets for the experiment.
	Targets map[string]ExperimentTarget

	noSmithyDocumentSerde
}

// Describes the action for an experiment.
type ExperimentAction struct {

	// The ID of the action.
	ActionId *string

	// The description for the action.
	Description *string

	// The time that the action ended.
	EndTime *time.Time

	// The parameters for the action.
	Parameters map[string]string

	// The name of the action that must be completed before this action starts.
	StartAfter []string

	// The time that the action started.
	StartTime *time.Time

	// The state of the action.
	State *ExperimentActionState

	// The targets for the action.
	Targets map[string]string

	noSmithyDocumentSerde
}

// Describes the state of an action.
type ExperimentActionState struct {

	// The reason for the state.
	Reason *string

	// The state of the action.
	Status ExperimentActionStatus

	noSmithyDocumentSerde
}

// Describes the configuration for experiment logging to Amazon CloudWatch Logs.
type ExperimentCloudWatchLogsLogConfiguration struct {

	// The Amazon Resource Name (ARN) of the destination Amazon CloudWatch Logs log
	// group.
	LogGroupArn *string

	noSmithyDocumentSerde
}

// Describes the configuration for experiment logging.
type ExperimentLogConfiguration struct {

	// The configuration for experiment logging to Amazon CloudWatch Logs.
	CloudWatchLogsConfiguration *ExperimentCloudWatchLogsLogConfiguration

	// The schema version.
	LogSchemaVersion *int32

	// The configuration for experiment logging to Amazon S3.
	S3Configuration *ExperimentS3LogConfiguration

	noSmithyDocumentSerde
}

// Describes the options for an experiment.
type ExperimentOptions struct {

	// The account targeting setting for an experiment.
	AccountTargeting AccountTargeting

	// The actions mode of the experiment that is set from the StartExperiment API
	// command.
	ActionsMode ActionsMode

	// The empty target resolution mode for an experiment.
	EmptyTargetResolutionMode EmptyTargetResolutionMode

	noSmithyDocumentSerde
}

// Describes the configuration for experiment logging to Amazon S3.
type ExperimentS3LogConfiguration struct {

	// The name of the destination bucket.
	BucketName *string

	// The bucket prefix.
	Prefix *string

	noSmithyDocumentSerde
}

// Describes the state of an experiment.
type ExperimentState struct {

	// The reason for the state.
	Reason *string

	// The state of the experiment.
	Status ExperimentStatus

	noSmithyDocumentSerde
}

// Describes the stop condition for an experiment.
type ExperimentStopCondition struct {

	// The source for the stop condition.
	Source *string

	// The Amazon Resource Name (ARN) of the CloudWatch alarm, if applicable.
	Value *string

	noSmithyDocumentSerde
}

// Provides a summary of an experiment.
type ExperimentSummary struct {

	// The Amazon Resource Name (ARN) of the experiment.
	Arn *string

	// The time that the experiment was created.
	CreationTime *time.Time

	// The experiment options for the experiment.
	ExperimentOptions *ExperimentOptions

	// The ID of the experiment template.
	ExperimentTemplateId *string

	// The ID of the experiment.
	Id *string

	// The state of the experiment.
	State *ExperimentState

	// The tags for the experiment.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Describes a target for an experiment.
type ExperimentTarget struct {

	// The filters to apply to identify target resources using specific attributes.
	Filters []ExperimentTargetFilter

	// The resource type parameters.
	Parameters map[string]string

	// The Amazon Resource Names (ARNs) of the resources.
	ResourceArns []string

	// The tags for the target resources.
	ResourceTags map[string]string

	// The resource type.
	ResourceType *string

	// Scopes the identified resources to a specific count or percentage.
	SelectionMode *string

	noSmithyDocumentSerde
}

// Describes a target account configuration for an experiment.
type ExperimentTargetAccountConfiguration struct {

	// The Amazon Web Services account ID of the target account.
	AccountId *string

	// The description of the target account.
	Description *string

	// The Amazon Resource Name (ARN) of an IAM role for the target account.
	RoleArn *string

	noSmithyDocumentSerde
}

// Provides a summary of a target account configuration.
type ExperimentTargetAccountConfigurationSummary struct {

	// The Amazon Web Services account ID of the target account.
	AccountId *string

	// The description of the target account.
	Description *string

	// The Amazon Resource Name (ARN) of an IAM role for the target account.
	RoleArn *string

	noSmithyDocumentSerde
}

// Describes a filter used for the target resources in an experiment.
type ExperimentTargetFilter struct {

	// The attribute path for the filter.
	Path *string

	// The attribute values for the filter.
	Values []string

	noSmithyDocumentSerde
}

// Describes an experiment template.
type ExperimentTemplate struct {

	// The actions for the experiment.
	Actions map[string]ExperimentTemplateAction

	// The Amazon Resource Name (ARN) of the experiment template.
	Arn *string

	// The time the experiment template was created.
	CreationTime *time.Time

	// The description for the experiment template.
	Description *string

	// The experiment options for an experiment template.
	ExperimentOptions *ExperimentTemplateExperimentOptions

	// The ID of the experiment template.
	Id *string

	// The time the experiment template was last updated.
	LastUpdateTime *time.Time

	// The configuration for experiment logging.
	LogConfiguration *ExperimentTemplateLogConfiguration

	// The Amazon Resource Name (ARN) of an IAM role.
	RoleArn *string

	// The stop conditions for the experiment.
	StopConditions []ExperimentTemplateStopCondition

	// The tags for the experiment template.
	Tags map[string]string

	// The count of target account configurations for the experiment template.
	TargetAccountConfigurationsCount *int64

	// The targets for the experiment.
	Targets map[string]ExperimentTemplateTarget

	noSmithyDocumentSerde
}

// Describes an action for an experiment template.
type ExperimentTemplateAction struct {

	// The ID of the action.
	ActionId *string

	// A description for the action.
	Description *string

	// The parameters for the action.
	Parameters map[string]string

	// The name of the action that must be completed before the current action starts.
	StartAfter []string

	// The targets for the action.
	Targets map[string]string

	noSmithyDocumentSerde
}

// Describes the configuration for experiment logging to Amazon CloudWatch Logs.
type ExperimentTemplateCloudWatchLogsLogConfiguration struct {

	// The Amazon Resource Name (ARN) of the destination Amazon CloudWatch Logs log
	// group.
	LogGroupArn *string

	noSmithyDocumentSerde
}

// Specifies the configuration for experiment logging to Amazon CloudWatch Logs.
type ExperimentTemplateCloudWatchLogsLogConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the destination Amazon CloudWatch Logs log
	// group.
	//
	// This member is required.
	LogGroupArn *string

	noSmithyDocumentSerde
}

// Describes the experiment options for an experiment template.
type ExperimentTemplateExperimentOptions struct {

	// The account targeting setting for an experiment template.
	AccountTargeting AccountTargeting

	// The empty target resolution mode for an experiment template.
	EmptyTargetResolutionMode EmptyTargetResolutionMode

	noSmithyDocumentSerde
}

// Describes the configuration for experiment logging.
type ExperimentTemplateLogConfiguration struct {

	// The configuration for experiment logging to Amazon CloudWatch Logs.
	CloudWatchLogsConfiguration *ExperimentTemplateCloudWatchLogsLogConfiguration

	// The schema version.
	LogSchemaVersion *int32

	// The configuration for experiment logging to Amazon S3.
	S3Configuration *ExperimentTemplateS3LogConfiguration

	noSmithyDocumentSerde
}

// Describes the configuration for experiment logging to Amazon S3.
type ExperimentTemplateS3LogConfiguration struct {

	// The name of the destination bucket.
	BucketName *string

	// The bucket prefix.
	Prefix *string

	noSmithyDocumentSerde
}

// Specifies the configuration for experiment logging to Amazon S3.
type ExperimentTemplateS3LogConfigurationInput struct {

	// The name of the destination bucket.
	//
	// This member is required.
	BucketName *string

	// The bucket prefix.
	Prefix *string

	noSmithyDocumentSerde
}

// Describes a stop condition for an experiment template.
type ExperimentTemplateStopCondition struct {

	// The source for the stop condition.
	Source *string

	// The Amazon Resource Name (ARN) of the CloudWatch alarm, if applicable.
	Value *string

	noSmithyDocumentSerde
}

// Provides a summary of an experiment template.
type ExperimentTemplateSummary struct {

	// The Amazon Resource Name (ARN) of the experiment template.
	Arn *string

	// The time that the experiment template was created.
	CreationTime *time.Time

	// The description of the experiment template.
	Description *string

	// The ID of the experiment template.
	Id *string

	// The time that the experiment template was last updated.
	LastUpdateTime *time.Time

	// The tags for the experiment template.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Describes a target for an experiment template.
type ExperimentTemplateTarget struct {

	// The filters to apply to identify target resources using specific attributes.
	Filters []ExperimentTemplateTargetFilter

	// The resource type parameters.
	Parameters map[string]string

	// The Amazon Resource Names (ARNs) of the targets.
	ResourceArns []string

	// The tags for the target resources.
	ResourceTags map[string]string

	// The resource type.
	ResourceType *string

	// Scopes the identified resources to a specific count or percentage.
	SelectionMode *string

	noSmithyDocumentSerde
}

// Describes a filter used for the target resources in an experiment template.
type ExperimentTemplateTargetFilter struct {

	// The attribute path for the filter.
	Path *string

	// The attribute values for the filter.
	Values []string

	noSmithyDocumentSerde
}

// Specifies a filter used for the target resource input in an experiment template.
//
// For more information, see [Resource filters] in the Fault Injection Service User Guide.
//
// [Resource filters]: https://docs.aws.amazon.com/fis/latest/userguide/targets.html#target-filters
type ExperimentTemplateTargetInputFilter struct {

	// The attribute path for the filter.
	//
	// This member is required.
	Path *string

	// The attribute values for the filter.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// Describes a resolved target.
type ResolvedTarget struct {

	// The resource type of the target.
	ResourceType *string

	// Information about the target.
	TargetInformation map[string]string

	// The name of the target.
	TargetName *string

	noSmithyDocumentSerde
}

// Specifies experiment options for running an experiment.
type StartExperimentExperimentOptionsInput struct {

	// Specifies the actions mode for experiment options.
	ActionsMode ActionsMode

	noSmithyDocumentSerde
}

// Describes a target account configuration.
type TargetAccountConfiguration struct {

	// The Amazon Web Services account ID of the target account.
	AccountId *string

	// The description of the target account.
	Description *string

	// The Amazon Resource Name (ARN) of an IAM role for the target account.
	RoleArn *string

	noSmithyDocumentSerde
}

// Provides a summary of a target account configuration.
type TargetAccountConfigurationSummary struct {

	// The Amazon Web Services account ID of the target account.
	AccountId *string

	// The description of the target account.
	Description *string

	// The Amazon Resource Name (ARN) of an IAM role for the target account.
	RoleArn *string

	noSmithyDocumentSerde
}

// Describes a resource type.
type TargetResourceType struct {

	// A description of the resource type.
	Description *string

	// The parameters for the resource type.
	Parameters map[string]TargetResourceTypeParameter

	// The resource type.
	ResourceType *string

	noSmithyDocumentSerde
}

// Describes the parameters for a resource type. Use parameters to determine which
// tasks are identified during target resolution.
type TargetResourceTypeParameter struct {

	// A description of the parameter.
	Description *string

	// Indicates whether the parameter is required.
	Required *bool

	noSmithyDocumentSerde
}

// Describes a resource type.
type TargetResourceTypeSummary struct {

	// A description of the resource type.
	Description *string

	// The resource type.
	ResourceType *string

	noSmithyDocumentSerde
}

// Specifies an action for an experiment template.
type UpdateExperimentTemplateActionInputItem struct {

	// The ID of the action.
	ActionId *string

	// A description for the action.
	Description *string

	// The parameters for the action, if applicable.
	Parameters map[string]string

	// The name of the action that must be completed before the current action starts.
	// Omit this parameter to run the action at the start of the experiment.
	StartAfter []string

	// The targets for the action.
	Targets map[string]string

	noSmithyDocumentSerde
}

// Specifies an experiment option for an experiment template.
type UpdateExperimentTemplateExperimentOptionsInput struct {

	// The empty target resolution mode of the experiment template.
	EmptyTargetResolutionMode EmptyTargetResolutionMode

	noSmithyDocumentSerde
}

// Specifies the configuration for experiment logging.
type UpdateExperimentTemplateLogConfigurationInput struct {

	// The configuration for experiment logging to Amazon CloudWatch Logs.
	CloudWatchLogsConfiguration *ExperimentTemplateCloudWatchLogsLogConfigurationInput

	// The schema version.
	LogSchemaVersion *int32

	// The configuration for experiment logging to Amazon S3.
	S3Configuration *ExperimentTemplateS3LogConfigurationInput

	noSmithyDocumentSerde
}

// Specifies a stop condition for an experiment. You can define a stop condition
// as a CloudWatch alarm.
type UpdateExperimentTemplateStopConditionInput struct {

	// The source for the stop condition. Specify aws:cloudwatch:alarm if the stop
	// condition is defined by a CloudWatch alarm. Specify none if there is no stop
	// condition.
	//
	// This member is required.
	Source *string

	// The Amazon Resource Name (ARN) of the CloudWatch alarm.
	Value *string

	noSmithyDocumentSerde
}

// Specifies a target for an experiment. You must specify at least one Amazon
// Resource Name (ARN) or at least one resource tag. You cannot specify both.
type UpdateExperimentTemplateTargetInput struct {

	// The resource type. The resource type must be supported for the specified action.
	//
	// This member is required.
	ResourceType *string

	// Scopes the identified resources to a specific count or percentage.
	//
	// This member is required.
	SelectionMode *string

	// The filters to apply to identify target resources using specific attributes.
	Filters []ExperimentTemplateTargetInputFilter

	// The resource type parameters.
	Parameters map[string]string

	// The Amazon Resource Names (ARNs) of the targets.
	ResourceArns []string

	// The tags for the target resources.
	ResourceTags map[string]string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

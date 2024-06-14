// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A structure containing the CloudWatch Logs log group where the project stores
// evaluation events.
type CloudWatchLogsDestination struct {

	// The name of the log group where the project stores evaluation events.
	LogGroup *string

	noSmithyDocumentSerde
}

// A structure containing the CloudWatch Logs log group where the project stores
// evaluation events.
type CloudWatchLogsDestinationConfig struct {

	// The name of the log group where the project stores evaluation events.
	LogGroup *string

	noSmithyDocumentSerde
}

// This structure assigns a feature variation to one user session.
type EvaluationRequest struct {

	// An internal ID that represents a unique user session of the application. This
	// entityID is checked against any override rules assigned for this feature.
	//
	// This member is required.
	EntityId *string

	// The name of the feature being evaluated.
	//
	// This member is required.
	Feature *string

	// A JSON block of attributes that you can optionally pass in. This JSON block is
	// included in the evaluation events sent to Evidently from the user session.
	//
	// This value conforms to the media type: application/json
	EvaluationContext *string

	noSmithyDocumentSerde
}

// This structure displays the results of one feature evaluation assignment to one
// user session.
type EvaluationResult struct {

	// An internal ID that represents a unique user session of the application.
	//
	// This member is required.
	EntityId *string

	// The name of the feature being evaluated.
	//
	// This member is required.
	Feature *string

	// If this user was assigned to a launch or experiment, this field lists the
	// launch or experiment name.
	//
	// This value conforms to the media type: application/json
	Details *string

	// The name or ARN of the project that contains the feature being evaluated.
	Project *string

	// Specifies the reason that the user session was assigned this variation.
	// Possible values include DEFAULT , meaning the user was served the default
	// variation; LAUNCH_RULE_MATCH , if the user session was enrolled in a launch; or
	// EXPERIMENT_RULE_MATCH , if the user session was enrolled in an experiment.
	Reason *string

	// The value assigned to this variation to differentiate it from the other
	// variations of this feature.
	Value VariableValue

	// The name of the variation that was served to the user session.
	Variation *string

	noSmithyDocumentSerde
}

// A structure that contains the information about an evaluation rule for this
// feature, if it is used in a launch or experiment.
type EvaluationRule struct {

	// This value is aws.evidently.splits if this is an evaluation rule for a launch,
	// and it is aws.evidently.onlineab if this is an evaluation rule for an
	// experiment.
	//
	// This member is required.
	Type *string

	// The name of the experiment or launch.
	Name *string

	noSmithyDocumentSerde
}

// A structure that contains the information about one evaluation event or custom
// event sent to Evidently. This is a JSON payload. If this event specifies a
// pre-defined event type, the payload must follow the defined event schema.
type Event struct {

	// The event data.
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Data *string

	// The timestamp of the event.
	//
	// This member is required.
	Timestamp *time.Time

	// aws.evidently.evaluation specifies an evaluation event, which determines which
	// feature variation that a user sees. aws.evidently.custom specifies a custom
	// event, which generates metrics from user actions such as clicks and checkouts.
	//
	// This member is required.
	Type EventType

	noSmithyDocumentSerde
}

// A structure containing the configuration details of an experiment.
type Experiment struct {

	// The ARN of the experiment.
	//
	// This member is required.
	Arn *string

	// The date and time that the experiment is first created.
	//
	// This member is required.
	CreatedTime *time.Time

	// The date and time that the experiment was most recently updated.
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The name of the experiment.
	//
	// This member is required.
	Name *string

	// The current state of the experiment.
	//
	// This member is required.
	Status ExperimentStatus

	// The type of this experiment. Currently, this value must be
	// aws.experiment.onlineab .
	//
	// This member is required.
	Type ExperimentType

	// A description of the experiment.
	Description *string

	// A structure that contains the date and time that the experiment started and
	// ended.
	Execution *ExperimentExecution

	// An array of structures that defines the metrics used for the experiment, and
	// whether a higher or lower value for each metric is the goal.
	MetricGoals []MetricGoal

	// A structure that contains the configuration of which variation to use as the
	// "control" version. The "control" version is used for comparison with other
	// variations. This structure also specifies how much experiment traffic is
	// allocated to each variation.
	OnlineAbDefinition *OnlineAbDefinition

	// The name or ARN of the project that contains this experiment.
	Project *string

	// This value is used when Evidently assigns a particular user session to the
	// experiment. It helps create a randomization ID to determine which variation the
	// user session is served. This randomization ID is a combination of the entity ID
	// and randomizationSalt .
	RandomizationSalt *string

	// In thousandths of a percent, the amount of the available audience that is
	// allocated to this experiment. The available audience is the total audience minus
	// the audience that you have allocated to overrides or current launches of this
	// feature.
	//
	// This is represented in thousandths of a percent, so a value of 10,000 is 10% of
	// the available audience.
	SamplingRate int64

	// A structure that contains the time and date that Evidently completed the
	// analysis of the experiment.
	Schedule *ExperimentSchedule

	// The audience segment being used for the experiment, if a segment is being used.
	Segment *string

	// If the experiment was stopped, this is the string that was entered by the
	// person who stopped the experiment, to explain why it was stopped.
	StatusReason *string

	// The list of tag keys and values associated with this experiment.
	Tags map[string]string

	// An array of structures that describe the configuration of each feature
	// variation used in the experiment.
	Treatments []Treatment

	noSmithyDocumentSerde
}

// This structure contains the date and time that the experiment started and ended.
type ExperimentExecution struct {

	// The date and time that the experiment ended.
	EndedTime *time.Time

	// The date and time that the experiment started.
	StartedTime *time.Time

	noSmithyDocumentSerde
}

// A structure that contains results of an experiment.
type ExperimentReport struct {

	// The content of the report.
	//
	// This value conforms to the media type: application/json
	Content *string

	// The name of the metric that is analyzed in this experiment report.
	MetricName *string

	// The type of analysis used for this report.
	ReportName ExperimentReportName

	// The name of the variation that this report pertains to.
	TreatmentName *string

	noSmithyDocumentSerde
}

// A structure that contains experiment results for one metric that is monitored
// in the experiment.
type ExperimentResultsData struct {

	// The name of the metric.
	MetricName *string

	// The experiment statistic that these results pertain to.
	ResultStat ExperimentResultResponseType

	// The treatment, or variation, that returned the values in this structure.
	TreatmentName *string

	// The values for the metricName that were recorded in the experiment.
	Values []float64

	noSmithyDocumentSerde
}

// This structure contains the time and date that Evidently completed the analysis
// of the experiment.
type ExperimentSchedule struct {

	// The time and date that Evidently completed the analysis of the experiment.
	AnalysisCompleteTime *time.Time

	noSmithyDocumentSerde
}

// This structure contains information about one Evidently feature in your account.
type Feature struct {

	// The ARN of the feature.
	//
	// This member is required.
	Arn *string

	// The date and time that the feature is created.
	//
	// This member is required.
	CreatedTime *time.Time

	// If this value is ALL_RULES , the traffic allocation specified by any ongoing
	// launches or experiments is being used. If this is DEFAULT_VARIATION , the
	// default variation is being served to all users.
	//
	// This member is required.
	EvaluationStrategy FeatureEvaluationStrategy

	// The date and time that the feature was most recently updated.
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The name of the feature.
	//
	// This member is required.
	Name *string

	// The current state of the feature.
	//
	// This member is required.
	Status FeatureStatus

	// Defines the type of value used to define the different feature variations. For
	// more information, see [Variation types]
	//
	// [Variation types]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-variationtypes.html
	//
	// This member is required.
	ValueType VariationValueType

	// An array of structures that contain the configuration of the feature's
	// different variations.
	//
	// This member is required.
	Variations []Variation

	// The name of the variation that is used as the default variation. The default
	// variation is served to users who are not allocated to any ongoing launches or
	// experiments of this feature.
	//
	// This variation must also be listed in the variations structure.
	//
	// If you omit defaultVariation , the first variation listed in the variations
	// structure is used as the default variation.
	DefaultVariation *string

	// The description of the feature.
	Description *string

	// A set of key-value pairs that specify users who should always be served a
	// specific variation of a feature. Each key specifies a user using their user ID,
	// account ID, or some other identifier. The value specifies the name of the
	// variation that the user is to be served.
	//
	// For the override to be successful, the value of the key must match the entityId
	// used in the [EvaluateFeature]operation.
	//
	// [EvaluateFeature]: https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_EvaluateFeature.html
	EntityOverrides map[string]string

	// An array of structures that define the evaluation rules for the feature.
	EvaluationRules []EvaluationRule

	// The name or ARN of the project that contains the feature.
	Project *string

	// The list of tag keys and values associated with this feature.
	Tags map[string]string

	noSmithyDocumentSerde
}

// This structure contains information about one Evidently feature in your account.
type FeatureSummary struct {

	// The ARN of the feature.
	//
	// This member is required.
	Arn *string

	// The date and time that the feature is created.
	//
	// This member is required.
	CreatedTime *time.Time

	// If this value is ALL_RULES , the traffic allocation specified by any ongoing
	// launches or experiments is being used. If this is DEFAULT_VARIATION , the
	// default variation is being served to all users.
	//
	// This member is required.
	EvaluationStrategy FeatureEvaluationStrategy

	// The date and time that the feature was most recently updated.
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The name of the feature.
	//
	// This member is required.
	Name *string

	// The current state of the feature.
	//
	// This member is required.
	Status FeatureStatus

	// The name of the variation that is used as the default variation. The default
	// variation is served to users who are not allocated to any ongoing launches or
	// experiments of this feature.
	DefaultVariation *string

	// An array of structures that define
	EvaluationRules []EvaluationRule

	// The name or ARN of the project that contains the feature.
	Project *string

	// The list of tag keys and values associated with this feature.
	Tags map[string]string

	noSmithyDocumentSerde
}

// This structure contains the configuration details of one Evidently launch.
type Launch struct {

	// The ARN of the launch.
	//
	// This member is required.
	Arn *string

	// The date and time that the launch is created.
	//
	// This member is required.
	CreatedTime *time.Time

	// The date and time that the launch was most recently updated.
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The name of the launch.
	//
	// This member is required.
	Name *string

	// The current state of the launch.
	//
	// This member is required.
	Status LaunchStatus

	// The type of launch.
	//
	// This member is required.
	Type LaunchType

	// The description of the launch.
	Description *string

	// A structure that contains information about the start and end times of the
	// launch.
	Execution *LaunchExecution

	// An array of structures that define the feature variations that are being used
	// in the launch.
	Groups []LaunchGroup

	// An array of structures that define the metrics that are being used to monitor
	// the launch performance.
	MetricMonitors []MetricMonitor

	// The name or ARN of the project that contains the launch.
	Project *string

	// This value is used when Evidently assigns a particular user session to the
	// launch, to help create a randomization ID to determine which variation the user
	// session is served. This randomization ID is a combination of the entity ID and
	// randomizationSalt .
	RandomizationSalt *string

	// An array of structures that define the traffic allocation percentages among the
	// feature variations during each step of the launch.
	ScheduledSplitsDefinition *ScheduledSplitsLaunchDefinition

	// If the launch was stopped, this is the string that was entered by the person
	// who stopped the launch, to explain why it was stopped.
	StatusReason *string

	// The list of tag keys and values associated with this launch.
	Tags map[string]string

	noSmithyDocumentSerde
}

// This structure contains information about the start and end times of the launch.
type LaunchExecution struct {

	// The date and time that the launch ended.
	EndedTime *time.Time

	// The date and time that the launch started.
	StartedTime *time.Time

	noSmithyDocumentSerde
}

// A structure that defines one launch group in a launch. A launch group is a
// variation of the feature that you are including in the launch.
type LaunchGroup struct {

	// The feature variation for this launch group. This is a key-value pair.
	//
	// This member is required.
	FeatureVariations map[string]string

	// The name of the launch group.
	//
	// This member is required.
	Name *string

	// A description of the launch group.
	Description *string

	noSmithyDocumentSerde
}

// A structure that defines one launch group in a launch. A launch group is a
// variation of the feature that you are including in the launch.
type LaunchGroupConfig struct {

	// The feature that this launch is using.
	//
	// This member is required.
	Feature *string

	// A name for this launch group.
	//
	// This member is required.
	Name *string

	// The feature variation to use for this launch group.
	//
	// This member is required.
	Variation *string

	// A description of the launch group.
	Description *string

	noSmithyDocumentSerde
}

// This structure defines a metric that is being used to evaluate the variations
// during a launch or experiment.
type MetricDefinition struct {

	// The entity, such as a user or session, that does an action that causes a metric
	// value to be recorded.
	EntityIdKey *string

	// The EventBridge event pattern that defines how the metric is recorded.
	//
	// For more information about EventBridge event patterns, see [Amazon EventBridge event patterns].
	//
	// [Amazon EventBridge event patterns]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html
	//
	// This value conforms to the media type: application/json
	EventPattern *string

	// The name of the metric.
	Name *string

	// The label for the units that the metric is measuring.
	UnitLabel *string

	// The value that is tracked to produce the metric.
	ValueKey *string

	noSmithyDocumentSerde
}

// This structure defines a metric that you want to use to evaluate the variations
// during a launch or experiment.
type MetricDefinitionConfig struct {

	// The entity, such as a user or session, that does an action that causes a metric
	// value to be recorded. An example is userDetails.userID .
	//
	// This member is required.
	EntityIdKey *string

	// A name for the metric.
	//
	// This member is required.
	Name *string

	// The value that is tracked to produce the metric.
	//
	// This member is required.
	ValueKey *string

	// The EventBridge event pattern that defines how the metric is recorded.
	//
	// For more information about EventBridge event patterns, see [Amazon EventBridge event patterns].
	//
	// [Amazon EventBridge event patterns]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html
	//
	// This value conforms to the media type: application/json
	EventPattern *string

	// A label for the units that the metric is measuring.
	UnitLabel *string

	noSmithyDocumentSerde
}

// A structure that tells Evidently whether higher or lower values are desired for
// a metric that is used in an experiment.
type MetricGoal struct {

	// A structure that contains details about the metric.
	//
	// This member is required.
	MetricDefinition *MetricDefinition

	// INCREASE means that a variation with a higher number for this metric is
	// performing better.
	//
	// DECREASE means that a variation with a lower number for this metric is
	// performing better.
	DesiredChange ChangeDirectionEnum

	noSmithyDocumentSerde
}

// Use this structure to tell Evidently whether higher or lower values are desired
// for a metric that is used in an experiment.
type MetricGoalConfig struct {

	// A structure that contains details about the metric.
	//
	// This member is required.
	MetricDefinition *MetricDefinitionConfig

	// INCREASE means that a variation with a higher number for this metric is
	// performing better.
	//
	// DECREASE means that a variation with a lower number for this metric is
	// performing better.
	DesiredChange ChangeDirectionEnum

	noSmithyDocumentSerde
}

// A structure that defines a metric to be used to monitor performance of the
// variations during a launch.
type MetricMonitor struct {

	// A structure that defines the metric.
	//
	// This member is required.
	MetricDefinition *MetricDefinition

	noSmithyDocumentSerde
}

// A structure that defines a metric to be used to monitor performance of the
// variations during a launch.
type MetricMonitorConfig struct {

	// A structure that defines the metric.
	//
	// This member is required.
	MetricDefinition *MetricDefinitionConfig

	noSmithyDocumentSerde
}

// A structure that contains the configuration of which variation to use as the
// "control" version. The "control" version is used for comparison with other
// variations. This structure also specifies how much experiment traffic is
// allocated to each variation.
type OnlineAbConfig struct {

	// The name of the variation that is to be the default variation that the other
	// variations are compared to.
	ControlTreatmentName *string

	// A set of key-value pairs. The keys are variation names, and the values are the
	// portion of experiment traffic to be assigned to that variation. Specify the
	// traffic portion in thousandths of a percent, so 20,000 for a variation would
	// allocate 20% of the experiment traffic to that variation.
	TreatmentWeights map[string]int64

	noSmithyDocumentSerde
}

// A structure that contains the configuration of which variation to use as the
// "control" version. The "control" version is used for comparison with other
// variations. This structure also specifies how much experiment traffic is
// allocated to each variation.
type OnlineAbDefinition struct {

	// The name of the variation that is the default variation that the other
	// variations are compared to.
	ControlTreatmentName *string

	// A set of key-value pairs. The keys are variation names, and the values are the
	// portion of experiment traffic to be assigned to that variation. The traffic
	// portion is specified in thousandths of a percent, so 20,000 for a variation
	// would allocate 20% of the experiment traffic to that variation.
	TreatmentWeights map[string]int64

	noSmithyDocumentSerde
}

// This structure defines a project, which is the logical object in Evidently that
// can contain features, launches, and experiments. Use projects to group similar
// features together.
type Project struct {

	// The name or ARN of the project.
	//
	// This member is required.
	Arn *string

	// The date and time that the project is created.
	//
	// This member is required.
	CreatedTime *time.Time

	// The date and time that the project was most recently updated.
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The name of the project.
	//
	// This member is required.
	Name *string

	// The current state of the project.
	//
	// This member is required.
	Status ProjectStatus

	// The number of ongoing experiments currently in the project.
	ActiveExperimentCount *int64

	// The number of ongoing launches currently in the project.
	ActiveLaunchCount *int64

	// This structure defines the configuration of how your application integrates
	// with AppConfig to run client-side evaluation.
	AppConfigResource *ProjectAppConfigResource

	// A structure that contains information about where Evidently is to store
	// evaluation events for longer term storage.
	DataDelivery *ProjectDataDelivery

	// The user-entered description of the project.
	Description *string

	// The number of experiments currently in the project. This includes all
	// experiments that have been created and not deleted, whether they are ongoing or
	// not.
	ExperimentCount *int64

	// The number of features currently in the project.
	FeatureCount *int64

	// The number of launches currently in the project. This includes all launches
	// that have been created and not deleted, whether they are ongoing or not.
	LaunchCount *int64

	// The list of tag keys and values associated with this project.
	Tags map[string]string

	noSmithyDocumentSerde
}

// This is a structure that defines the configuration of how your application
// integrates with AppConfig to run client-side evaluation.
type ProjectAppConfigResource struct {

	// The ID of the AppConfig application to use for client-side evaluation.
	//
	// This member is required.
	ApplicationId *string

	// The ID of the AppConfig profile to use for client-side evaluation.
	//
	// This member is required.
	ConfigurationProfileId *string

	// The ID of the AppConfig environment to use for client-side evaluation. This
	// must be an environment that is within the application that you specify for
	// applicationId .
	//
	// This member is required.
	EnvironmentId *string

	noSmithyDocumentSerde
}

// Use this parameter to configure client-side evaluation for your project.
// Client-side evaluation allows your application to assign variations to user
// sessions locally instead of by calling the [EvaluateFeature]operation to assign the variations.
// This mitigates the latency and availability risks that come with an API call.
//
// ProjectAppConfigResource is a structure that defines the configuration of how
// your application integrates with AppConfig to run client-side evaluation.
//
// [EvaluateFeature]: https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_EvaluateFeature.html
type ProjectAppConfigResourceConfig struct {

	// The ID of the AppConfig application to use for client-side evaluation.
	ApplicationId *string

	// The ID of the AppConfig environment to use for client-side evaluation. This
	// must be an environment that is within the application that you specify for
	// applicationId .
	EnvironmentId *string

	noSmithyDocumentSerde
}

// A structure that contains information about where Evidently is to store
// evaluation events for longer term storage.
type ProjectDataDelivery struct {

	// If the project stores evaluation events in CloudWatch Logs, this structure
	// stores the log group name.
	CloudWatchLogs *CloudWatchLogsDestination

	// If the project stores evaluation events in an Amazon S3 bucket, this structure
	// stores the bucket name and bucket prefix.
	S3Destination *S3Destination

	noSmithyDocumentSerde
}

// A structure that contains information about where Evidently is to store
// evaluation events for longer term storage.
type ProjectDataDeliveryConfig struct {

	// If the project stores evaluation events in CloudWatch Logs, this structure
	// stores the log group name.
	CloudWatchLogs *CloudWatchLogsDestinationConfig

	// If the project stores evaluation events in an Amazon S3 bucket, this structure
	// stores the bucket name and bucket prefix.
	S3Destination *S3DestinationConfig

	noSmithyDocumentSerde
}

// A structure that contains configuration information about an Evidently project.
type ProjectSummary struct {

	// The name or ARN of the project.
	//
	// This member is required.
	Arn *string

	// The date and time that the project is created.
	//
	// This member is required.
	CreatedTime *time.Time

	// The date and time that the project was most recently updated.
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The name of the project.
	//
	// This member is required.
	Name *string

	// The current state of the project.
	//
	// This member is required.
	Status ProjectStatus

	// The number of experiments currently in the project.
	ActiveExperimentCount *int64

	// The number of ongoing launches currently in the project.
	ActiveLaunchCount *int64

	// The description of the project.
	Description *string

	// The number of experiments currently in the project.
	ExperimentCount *int64

	// The number of features currently in the project.
	FeatureCount *int64

	// The number of launches currently in the project, including launches that are
	// ongoing, completed, and not started yet.
	LaunchCount *int64

	// The list of tag keys and values associated with this project.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A structure that contains Evidently's response to the sent events, including an
// event ID and error codes, if any.
type PutProjectEventsResultEntry struct {

	// If the PutProjectEvents operation has an error, the error code is returned here.
	ErrorCode *string

	// If the PutProjectEvents operation has an error, the error message is returned
	// here.
	ErrorMessage *string

	// A unique ID assigned to this PutProjectEvents operation.
	EventId *string

	noSmithyDocumentSerde
}

// A structure that contains information about one experiment or launch that uses
// the specified segment.
type RefResource struct {

	// The name of the experiment or launch.
	//
	// This member is required.
	Name *string

	// Specifies whether the resource that this structure contains information about
	// is an experiment or a launch.
	//
	// This member is required.
	Type *string

	// The ARN of the experiment or launch.
	Arn *string

	// The day and time that this experiment or launch ended.
	EndTime *string

	// The day and time that this experiment or launch was most recently updated.
	LastUpdatedOn *string

	// The day and time that this experiment or launch started.
	StartTime *string

	// The status of the experiment or launch.
	Status *string

	noSmithyDocumentSerde
}

// If the project stores evaluation events in an Amazon S3 bucket, this structure
// stores the bucket name and bucket prefix.
type S3Destination struct {

	// The name of the bucket in which Evidently stores evaluation events.
	Bucket *string

	// The bucket prefix in which Evidently stores evaluation events.
	Prefix *string

	noSmithyDocumentSerde
}

// If the project stores evaluation events in an Amazon S3 bucket, this structure
// stores the bucket name and bucket prefix.
type S3DestinationConfig struct {

	// The name of the bucket in which Evidently stores evaluation events.
	Bucket *string

	// The bucket prefix in which Evidently stores evaluation events.
	Prefix *string

	noSmithyDocumentSerde
}

// This structure defines the traffic allocation percentages among the feature
// variations during one step of a launch, and the start time of that step.
type ScheduledSplit struct {

	// The date and time that this step of the launch starts.
	//
	// This member is required.
	StartTime *time.Time

	// The traffic allocation percentages among the feature variations during one step
	// of a launch. This is a set of key-value pairs. The keys are variation names. The
	// values represent the percentage of traffic to allocate to that variation during
	// this step.
	//
	// The values is expressed in thousandths of a percent, so assigning a weight of
	// 50000 assigns 50% of traffic to that variation.
	//
	// If the sum of the weights for all the variations in a segment override does not
	// add up to 100,000, then the remaining traffic that matches this segment is not
	// assigned by this segment override, and instead moves on to the next segment
	// override or the default traffic split.
	GroupWeights map[string]int64

	// Use this parameter to specify different traffic splits for one or more audience
	// segments. A segment is a portion of your audience that share one or more
	// characteristics. Examples could be Chrome browser users, users in Europe, or
	// Firefox browser users in Europe who also fit other criteria that your
	// application collects, such as age.
	//
	// This parameter is an array of up to six segment override objects. Each of these
	// objects specifies a segment that you have already created, and defines the
	// traffic split for that segment.
	SegmentOverrides []SegmentOverride

	noSmithyDocumentSerde
}

// This structure defines the traffic allocation percentages among the feature
// variations during one step of a launch, and the start time of that step.
type ScheduledSplitConfig struct {

	// The traffic allocation percentages among the feature variations during one step
	// of a launch. This is a set of key-value pairs. The keys are variation names. The
	// values represent the percentage of traffic to allocate to that variation during
	// this step.
	//
	// The values is expressed in thousandths of a percent, so assigning a weight of
	// 50000 assigns 50% of traffic to that variation.
	//
	// If the sum of the weights for all the variations in a segment override does not
	// add up to 100,000, then the remaining traffic that matches this segment is not
	// assigned by this segment override, and instead moves on to the next segment
	// override or the default traffic split.
	//
	// This member is required.
	GroupWeights map[string]int64

	// The date and time that this step of the launch starts.
	//
	// This member is required.
	StartTime *time.Time

	// Use this parameter to specify different traffic splits for one or more audience
	// segments. A segment is a portion of your audience that share one or more
	// characteristics. Examples could be Chrome browser users, users in Europe, or
	// Firefox browser users in Europe who also fit other criteria that your
	// application collects, such as age.
	//
	// This parameter is an array of up to six segment override objects. Each of these
	// objects specifies a segment that you have already created, and defines the
	// traffic split for that segment.
	SegmentOverrides []SegmentOverride

	noSmithyDocumentSerde
}

// An array of structures that define the traffic allocation percentages among the
// feature variations during each step of a launch. This also defines the start
// time of each step.
type ScheduledSplitsLaunchConfig struct {

	// An array of structures that define the traffic allocation percentages among the
	// feature variations during each step of the launch. This also defines the start
	// time of each step.
	//
	// This member is required.
	Steps []ScheduledSplitConfig

	noSmithyDocumentSerde
}

// An array of structures that define the traffic allocation percentages among the
// feature variations during each step of a launch. This also defines the start
// time of each step.
type ScheduledSplitsLaunchDefinition struct {

	// An array of structures that define the traffic allocation percentages among the
	// feature variations during each step of the launch. This also defines the start
	// time of each step.
	Steps []ScheduledSplit

	noSmithyDocumentSerde
}

// This structure contains information about one audience segment. You can use
// segments in your experiments and launches to narrow the user sessions used for
// experiment or launch to only the user sessions that match one or more criteria.
type Segment struct {

	// The ARN of the segment.
	//
	// This member is required.
	Arn *string

	// The date and time that this segment was created.
	//
	// This member is required.
	CreatedTime *time.Time

	// The date and time that this segment was most recently updated.
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The name of the segment.
	//
	// This member is required.
	Name *string

	// The pattern that defines the attributes to use to evalute whether a user
	// session will be in the segment. For more information about the pattern syntax,
	// see [Segment rule pattern syntax].
	//
	// [Segment rule pattern syntax]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Pattern *string

	// The customer-created description for this segment.
	Description *string

	// The number of experiments that this segment is used in. This count includes all
	// current experiments, not just those that are currently running.
	ExperimentCount *int64

	// The number of launches that this segment is used in. This count includes all
	// current launches, not just those that are currently running.
	LaunchCount *int64

	// The list of tag keys and values associated with this launch.
	Tags map[string]string

	noSmithyDocumentSerde
}

// This structure specifies a segment that you have already created, and defines
// the traffic split for that segment to be used in a launch.
type SegmentOverride struct {

	// A number indicating the order to use to evaluate segment overrides, if there
	// are more than one. Segment overrides with lower numbers are evaluated first.
	//
	// This member is required.
	EvaluationOrder *int64

	// The ARN of the segment to use.
	//
	// This member is required.
	Segment *string

	// The traffic allocation percentages among the feature variations to assign to
	// this segment. This is a set of key-value pairs. The keys are variation names.
	// The values represent the amount of traffic to allocate to that variation for
	// this segment. This is expressed in thousandths of a percent, so a weight of
	// 50000 represents 50% of traffic.
	//
	// This member is required.
	Weights map[string]int64

	noSmithyDocumentSerde
}

// A structure that defines one treatment in an experiment. A treatment is a
// variation of the feature that you are including in the experiment.
type Treatment struct {

	// The name of this treatment.
	//
	// This member is required.
	Name *string

	// The description of the treatment.
	Description *string

	// The feature variation used for this treatment. This is a key-value pair. The
	// key is the feature name, and the value is the variation name.
	FeatureVariations map[string]string

	noSmithyDocumentSerde
}

// A structure that defines one treatment in an experiment. A treatment is a
// variation of the feature that you are including in the experiment.
type TreatmentConfig struct {

	// The feature that this experiment is testing.
	//
	// This member is required.
	Feature *string

	// A name for this treatment.
	//
	// This member is required.
	Name *string

	// The name of the variation to use as this treatment in the experiment.
	//
	// This member is required.
	Variation *string

	// A description for this treatment.
	Description *string

	noSmithyDocumentSerde
}

// A structure containing an error name and message.
type ValidationExceptionField struct {

	// The error message.
	//
	// This member is required.
	Message *string

	// The error name.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// The value assigned to a feature variation. This structure must contain exactly
// one field. It can be boolValue , doubleValue , longValue , or stringValue .
//
// The following types satisfy this interface:
//
//	VariableValueMemberBoolValue
//	VariableValueMemberDoubleValue
//	VariableValueMemberLongValue
//	VariableValueMemberStringValue
type VariableValue interface {
	isVariableValue()
}

// If this feature uses the Boolean variation type, this field contains the
// Boolean value of this variation.
type VariableValueMemberBoolValue struct {
	Value bool

	noSmithyDocumentSerde
}

func (*VariableValueMemberBoolValue) isVariableValue() {}

// If this feature uses the double integer variation type, this field contains the
// double integer value of this variation.
type VariableValueMemberDoubleValue struct {
	Value float64

	noSmithyDocumentSerde
}

func (*VariableValueMemberDoubleValue) isVariableValue() {}

// If this feature uses the long variation type, this field contains the long
// value of this variation.
type VariableValueMemberLongValue struct {
	Value int64

	noSmithyDocumentSerde
}

func (*VariableValueMemberLongValue) isVariableValue() {}

// If this feature uses the string variation type, this field contains the string
// value of this variation.
type VariableValueMemberStringValue struct {
	Value string

	noSmithyDocumentSerde
}

func (*VariableValueMemberStringValue) isVariableValue() {}

// This structure contains the name and variation value of one variation of a
// feature.
type Variation struct {

	// The name of the variation.
	Name *string

	// The value assigned to this variation.
	Value VariableValue

	noSmithyDocumentSerde
}

// This structure contains the name and variation value of one variation of a
// feature.
type VariationConfig struct {

	// The name of the variation.
	//
	// This member is required.
	Name *string

	// The value assigned to this variation.
	//
	// This member is required.
	Value VariableValue

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isVariableValue() {}

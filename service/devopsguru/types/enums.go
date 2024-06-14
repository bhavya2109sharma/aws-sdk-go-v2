// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AnomalySeverity string

// Enum values for AnomalySeverity
const (
	AnomalySeverityLow    AnomalySeverity = "LOW"
	AnomalySeverityMedium AnomalySeverity = "MEDIUM"
	AnomalySeverityHigh   AnomalySeverity = "HIGH"
)

// Values returns all known values for AnomalySeverity. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnomalySeverity) Values() []AnomalySeverity {
	return []AnomalySeverity{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

type AnomalyStatus string

// Enum values for AnomalyStatus
const (
	AnomalyStatusOngoing AnomalyStatus = "ONGOING"
	AnomalyStatusClosed  AnomalyStatus = "CLOSED"
)

// Values returns all known values for AnomalyStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnomalyStatus) Values() []AnomalyStatus {
	return []AnomalyStatus{
		"ONGOING",
		"CLOSED",
	}
}

type AnomalyType string

// Enum values for AnomalyType
const (
	AnomalyTypeCausal     AnomalyType = "CAUSAL"
	AnomalyTypeContextual AnomalyType = "CONTEXTUAL"
)

// Values returns all known values for AnomalyType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnomalyType) Values() []AnomalyType {
	return []AnomalyType{
		"CAUSAL",
		"CONTEXTUAL",
	}
}

type CloudWatchMetricDataStatusCode string

// Enum values for CloudWatchMetricDataStatusCode
const (
	CloudWatchMetricDataStatusCodeComplete      CloudWatchMetricDataStatusCode = "Complete"
	CloudWatchMetricDataStatusCodeInternalError CloudWatchMetricDataStatusCode = "InternalError"
	CloudWatchMetricDataStatusCodePartialData   CloudWatchMetricDataStatusCode = "PartialData"
)

// Values returns all known values for CloudWatchMetricDataStatusCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CloudWatchMetricDataStatusCode) Values() []CloudWatchMetricDataStatusCode {
	return []CloudWatchMetricDataStatusCode{
		"Complete",
		"InternalError",
		"PartialData",
	}
}

type CloudWatchMetricsStat string

// Enum values for CloudWatchMetricsStat
const (
	CloudWatchMetricsStatSum         CloudWatchMetricsStat = "Sum"
	CloudWatchMetricsStatAverage     CloudWatchMetricsStat = "Average"
	CloudWatchMetricsStatSampleCount CloudWatchMetricsStat = "SampleCount"
	CloudWatchMetricsStatMinimum     CloudWatchMetricsStat = "Minimum"
	CloudWatchMetricsStatMaximum     CloudWatchMetricsStat = "Maximum"
	CloudWatchMetricsStatP99         CloudWatchMetricsStat = "p99"
	CloudWatchMetricsStatP90         CloudWatchMetricsStat = "p90"
	CloudWatchMetricsStatP50         CloudWatchMetricsStat = "p50"
)

// Values returns all known values for CloudWatchMetricsStat. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CloudWatchMetricsStat) Values() []CloudWatchMetricsStat {
	return []CloudWatchMetricsStat{
		"Sum",
		"Average",
		"SampleCount",
		"Minimum",
		"Maximum",
		"p99",
		"p90",
		"p50",
	}
}

type CostEstimationServiceResourceState string

// Enum values for CostEstimationServiceResourceState
const (
	CostEstimationServiceResourceStateActive   CostEstimationServiceResourceState = "ACTIVE"
	CostEstimationServiceResourceStateInactive CostEstimationServiceResourceState = "INACTIVE"
)

// Values returns all known values for CostEstimationServiceResourceState. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostEstimationServiceResourceState) Values() []CostEstimationServiceResourceState {
	return []CostEstimationServiceResourceState{
		"ACTIVE",
		"INACTIVE",
	}
}

type CostEstimationStatus string

// Enum values for CostEstimationStatus
const (
	CostEstimationStatusOngoing   CostEstimationStatus = "ONGOING"
	CostEstimationStatusCompleted CostEstimationStatus = "COMPLETED"
)

// Values returns all known values for CostEstimationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostEstimationStatus) Values() []CostEstimationStatus {
	return []CostEstimationStatus{
		"ONGOING",
		"COMPLETED",
	}
}

type EventClass string

// Enum values for EventClass
const (
	EventClassInfrastructure EventClass = "INFRASTRUCTURE"
	EventClassDeployment     EventClass = "DEPLOYMENT"
	EventClassSecurityChange EventClass = "SECURITY_CHANGE"
	EventClassConfigChange   EventClass = "CONFIG_CHANGE"
	EventClassSchemaChange   EventClass = "SCHEMA_CHANGE"
)

// Values returns all known values for EventClass. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventClass) Values() []EventClass {
	return []EventClass{
		"INFRASTRUCTURE",
		"DEPLOYMENT",
		"SECURITY_CHANGE",
		"CONFIG_CHANGE",
		"SCHEMA_CHANGE",
	}
}

type EventDataSource string

// Enum values for EventDataSource
const (
	EventDataSourceAwsCloudTrail EventDataSource = "AWS_CLOUD_TRAIL"
	EventDataSourceAwsCodeDeploy EventDataSource = "AWS_CODE_DEPLOY"
)

// Values returns all known values for EventDataSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventDataSource) Values() []EventDataSource {
	return []EventDataSource{
		"AWS_CLOUD_TRAIL",
		"AWS_CODE_DEPLOY",
	}
}

type EventSourceOptInStatus string

// Enum values for EventSourceOptInStatus
const (
	EventSourceOptInStatusEnabled  EventSourceOptInStatus = "ENABLED"
	EventSourceOptInStatusDisabled EventSourceOptInStatus = "DISABLED"
)

// Values returns all known values for EventSourceOptInStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventSourceOptInStatus) Values() []EventSourceOptInStatus {
	return []EventSourceOptInStatus{
		"ENABLED",
		"DISABLED",
	}
}

type InsightFeedbackOption string

// Enum values for InsightFeedbackOption
const (
	InsightFeedbackOptionValidCollection      InsightFeedbackOption = "VALID_COLLECTION"
	InsightFeedbackOptionRecommendationUseful InsightFeedbackOption = "RECOMMENDATION_USEFUL"
	InsightFeedbackOptionAlertTooSensitive    InsightFeedbackOption = "ALERT_TOO_SENSITIVE"
	InsightFeedbackOptionDataNoisyAnomaly     InsightFeedbackOption = "DATA_NOISY_ANOMALY"
	InsightFeedbackOptionDataIncorrect        InsightFeedbackOption = "DATA_INCORRECT"
)

// Values returns all known values for InsightFeedbackOption. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InsightFeedbackOption) Values() []InsightFeedbackOption {
	return []InsightFeedbackOption{
		"VALID_COLLECTION",
		"RECOMMENDATION_USEFUL",
		"ALERT_TOO_SENSITIVE",
		"DATA_NOISY_ANOMALY",
		"DATA_INCORRECT",
	}
}

type InsightSeverity string

// Enum values for InsightSeverity
const (
	InsightSeverityLow    InsightSeverity = "LOW"
	InsightSeverityMedium InsightSeverity = "MEDIUM"
	InsightSeverityHigh   InsightSeverity = "HIGH"
)

// Values returns all known values for InsightSeverity. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InsightSeverity) Values() []InsightSeverity {
	return []InsightSeverity{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

type InsightStatus string

// Enum values for InsightStatus
const (
	InsightStatusOngoing InsightStatus = "ONGOING"
	InsightStatusClosed  InsightStatus = "CLOSED"
)

// Values returns all known values for InsightStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InsightStatus) Values() []InsightStatus {
	return []InsightStatus{
		"ONGOING",
		"CLOSED",
	}
}

type InsightType string

// Enum values for InsightType
const (
	InsightTypeReactive  InsightType = "REACTIVE"
	InsightTypeProactive InsightType = "PROACTIVE"
)

// Values returns all known values for InsightType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InsightType) Values() []InsightType {
	return []InsightType{
		"REACTIVE",
		"PROACTIVE",
	}
}

type Locale string

// Enum values for Locale
const (
	LocaleDeDe Locale = "DE_DE"
	LocaleEnUs Locale = "EN_US"
	LocaleEnGb Locale = "EN_GB"
	LocaleEsEs Locale = "ES_ES"
	LocaleFrFr Locale = "FR_FR"
	LocaleItIt Locale = "IT_IT"
	LocaleJaJp Locale = "JA_JP"
	LocaleKoKr Locale = "KO_KR"
	LocalePtBr Locale = "PT_BR"
	LocaleZhCn Locale = "ZH_CN"
	LocaleZhTw Locale = "ZH_TW"
)

// Values returns all known values for Locale. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Locale) Values() []Locale {
	return []Locale{
		"DE_DE",
		"EN_US",
		"EN_GB",
		"ES_ES",
		"FR_FR",
		"IT_IT",
		"JA_JP",
		"KO_KR",
		"PT_BR",
		"ZH_CN",
		"ZH_TW",
	}
}

type LogAnomalyType string

// Enum values for LogAnomalyType
const (
	LogAnomalyTypeKeyword        LogAnomalyType = "KEYWORD"
	LogAnomalyTypeKeywordToken   LogAnomalyType = "KEYWORD_TOKEN"
	LogAnomalyTypeFormat         LogAnomalyType = "FORMAT"
	LogAnomalyTypeHttpCode       LogAnomalyType = "HTTP_CODE"
	LogAnomalyTypeBlockFormat    LogAnomalyType = "BLOCK_FORMAT"
	LogAnomalyTypeNumericalPoint LogAnomalyType = "NUMERICAL_POINT"
	LogAnomalyTypeNumericalNan   LogAnomalyType = "NUMERICAL_NAN"
	LogAnomalyTypeNewFieldName   LogAnomalyType = "NEW_FIELD_NAME"
)

// Values returns all known values for LogAnomalyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LogAnomalyType) Values() []LogAnomalyType {
	return []LogAnomalyType{
		"KEYWORD",
		"KEYWORD_TOKEN",
		"FORMAT",
		"HTTP_CODE",
		"BLOCK_FORMAT",
		"NUMERICAL_POINT",
		"NUMERICAL_NAN",
		"NEW_FIELD_NAME",
	}
}

type NotificationMessageType string

// Enum values for NotificationMessageType
const (
	NotificationMessageTypeNewInsight        NotificationMessageType = "NEW_INSIGHT"
	NotificationMessageTypeClosedInsight     NotificationMessageType = "CLOSED_INSIGHT"
	NotificationMessageTypeNewAssociation    NotificationMessageType = "NEW_ASSOCIATION"
	NotificationMessageTypeSeverityUpgraded  NotificationMessageType = "SEVERITY_UPGRADED"
	NotificationMessageTypeNewRecommendation NotificationMessageType = "NEW_RECOMMENDATION"
)

// Values returns all known values for NotificationMessageType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NotificationMessageType) Values() []NotificationMessageType {
	return []NotificationMessageType{
		"NEW_INSIGHT",
		"CLOSED_INSIGHT",
		"NEW_ASSOCIATION",
		"SEVERITY_UPGRADED",
		"NEW_RECOMMENDATION",
	}
}

type OptInStatus string

// Enum values for OptInStatus
const (
	OptInStatusEnabled  OptInStatus = "ENABLED"
	OptInStatusDisabled OptInStatus = "DISABLED"
)

// Values returns all known values for OptInStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OptInStatus) Values() []OptInStatus {
	return []OptInStatus{
		"ENABLED",
		"DISABLED",
	}
}

type OrganizationResourceCollectionType string

// Enum values for OrganizationResourceCollectionType
const (
	OrganizationResourceCollectionTypeAwsCloudFormation OrganizationResourceCollectionType = "AWS_CLOUD_FORMATION"
	OrganizationResourceCollectionTypeAwsService        OrganizationResourceCollectionType = "AWS_SERVICE"
	OrganizationResourceCollectionTypeAwsAccount        OrganizationResourceCollectionType = "AWS_ACCOUNT"
	OrganizationResourceCollectionTypeAwsTags           OrganizationResourceCollectionType = "AWS_TAGS"
)

// Values returns all known values for OrganizationResourceCollectionType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrganizationResourceCollectionType) Values() []OrganizationResourceCollectionType {
	return []OrganizationResourceCollectionType{
		"AWS_CLOUD_FORMATION",
		"AWS_SERVICE",
		"AWS_ACCOUNT",
		"AWS_TAGS",
	}
}

type ResourceCollectionType string

// Enum values for ResourceCollectionType
const (
	ResourceCollectionTypeAwsCloudFormation ResourceCollectionType = "AWS_CLOUD_FORMATION"
	ResourceCollectionTypeAwsService        ResourceCollectionType = "AWS_SERVICE"
	ResourceCollectionTypeAwsTags           ResourceCollectionType = "AWS_TAGS"
)

// Values returns all known values for ResourceCollectionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceCollectionType) Values() []ResourceCollectionType {
	return []ResourceCollectionType{
		"AWS_CLOUD_FORMATION",
		"AWS_SERVICE",
		"AWS_TAGS",
	}
}

type ResourcePermission string

// Enum values for ResourcePermission
const (
	ResourcePermissionFullPermission    ResourcePermission = "FULL_PERMISSION"
	ResourcePermissionMissingPermission ResourcePermission = "MISSING_PERMISSION"
)

// Values returns all known values for ResourcePermission. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourcePermission) Values() []ResourcePermission {
	return []ResourcePermission{
		"FULL_PERMISSION",
		"MISSING_PERMISSION",
	}
}

type ResourceTypeFilter string

// Enum values for ResourceTypeFilter
const (
	ResourceTypeFilterLogGroups                          ResourceTypeFilter = "LOG_GROUPS"
	ResourceTypeFilterCloudfrontDistribution             ResourceTypeFilter = "CLOUDFRONT_DISTRIBUTION"
	ResourceTypeFilterDynamodbTable                      ResourceTypeFilter = "DYNAMODB_TABLE"
	ResourceTypeFilterEc2NatGateway                      ResourceTypeFilter = "EC2_NAT_GATEWAY"
	ResourceTypeFilterEcsCluster                         ResourceTypeFilter = "ECS_CLUSTER"
	ResourceTypeFilterEcsService                         ResourceTypeFilter = "ECS_SERVICE"
	ResourceTypeFilterEksCluster                         ResourceTypeFilter = "EKS_CLUSTER"
	ResourceTypeFilterElasticBeanstalkEnvironment        ResourceTypeFilter = "ELASTIC_BEANSTALK_ENVIRONMENT"
	ResourceTypeFilterElasticLoadBalancerLoadBalancer    ResourceTypeFilter = "ELASTIC_LOAD_BALANCER_LOAD_BALANCER"
	ResourceTypeFilterElasticLoadBalancingV2LoadBalancer ResourceTypeFilter = "ELASTIC_LOAD_BALANCING_V2_LOAD_BALANCER"
	ResourceTypeFilterElasticLoadBalancingV2TargetGroup  ResourceTypeFilter = "ELASTIC_LOAD_BALANCING_V2_TARGET_GROUP"
	ResourceTypeFilterElasticacheCacheCluster            ResourceTypeFilter = "ELASTICACHE_CACHE_CLUSTER"
	ResourceTypeFilterElasticsearchDomain                ResourceTypeFilter = "ELASTICSEARCH_DOMAIN"
	ResourceTypeFilterKinesisStream                      ResourceTypeFilter = "KINESIS_STREAM"
	ResourceTypeFilterLambdaFunction                     ResourceTypeFilter = "LAMBDA_FUNCTION"
	ResourceTypeFilterOpenSearchServiceDomain            ResourceTypeFilter = "OPEN_SEARCH_SERVICE_DOMAIN"
	ResourceTypeFilterRdsDbInstance                      ResourceTypeFilter = "RDS_DB_INSTANCE"
	ResourceTypeFilterRdsDbCluster                       ResourceTypeFilter = "RDS_DB_CLUSTER"
	ResourceTypeFilterRedshiftCluster                    ResourceTypeFilter = "REDSHIFT_CLUSTER"
	ResourceTypeFilterRoute53HostedZone                  ResourceTypeFilter = "ROUTE53_HOSTED_ZONE"
	ResourceTypeFilterRoute53HealthCheck                 ResourceTypeFilter = "ROUTE53_HEALTH_CHECK"
	ResourceTypeFilterS3Bucket                           ResourceTypeFilter = "S3_BUCKET"
	ResourceTypeFilterSagemakerEndpoint                  ResourceTypeFilter = "SAGEMAKER_ENDPOINT"
	ResourceTypeFilterSnsTopic                           ResourceTypeFilter = "SNS_TOPIC"
	ResourceTypeFilterSqsQueue                           ResourceTypeFilter = "SQS_QUEUE"
	ResourceTypeFilterStepFunctionsActivity              ResourceTypeFilter = "STEP_FUNCTIONS_ACTIVITY"
	ResourceTypeFilterStepFunctionsStateMachine          ResourceTypeFilter = "STEP_FUNCTIONS_STATE_MACHINE"
)

// Values returns all known values for ResourceTypeFilter. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceTypeFilter) Values() []ResourceTypeFilter {
	return []ResourceTypeFilter{
		"LOG_GROUPS",
		"CLOUDFRONT_DISTRIBUTION",
		"DYNAMODB_TABLE",
		"EC2_NAT_GATEWAY",
		"ECS_CLUSTER",
		"ECS_SERVICE",
		"EKS_CLUSTER",
		"ELASTIC_BEANSTALK_ENVIRONMENT",
		"ELASTIC_LOAD_BALANCER_LOAD_BALANCER",
		"ELASTIC_LOAD_BALANCING_V2_LOAD_BALANCER",
		"ELASTIC_LOAD_BALANCING_V2_TARGET_GROUP",
		"ELASTICACHE_CACHE_CLUSTER",
		"ELASTICSEARCH_DOMAIN",
		"KINESIS_STREAM",
		"LAMBDA_FUNCTION",
		"OPEN_SEARCH_SERVICE_DOMAIN",
		"RDS_DB_INSTANCE",
		"RDS_DB_CLUSTER",
		"REDSHIFT_CLUSTER",
		"ROUTE53_HOSTED_ZONE",
		"ROUTE53_HEALTH_CHECK",
		"S3_BUCKET",
		"SAGEMAKER_ENDPOINT",
		"SNS_TOPIC",
		"SQS_QUEUE",
		"STEP_FUNCTIONS_ACTIVITY",
		"STEP_FUNCTIONS_STATE_MACHINE",
	}
}

type ServerSideEncryptionType string

// Enum values for ServerSideEncryptionType
const (
	ServerSideEncryptionTypeCustomerManagedKey ServerSideEncryptionType = "CUSTOMER_MANAGED_KEY"
	ServerSideEncryptionTypeAwsOwnedKmsKey     ServerSideEncryptionType = "AWS_OWNED_KMS_KEY"
)

// Values returns all known values for ServerSideEncryptionType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServerSideEncryptionType) Values() []ServerSideEncryptionType {
	return []ServerSideEncryptionType{
		"CUSTOMER_MANAGED_KEY",
		"AWS_OWNED_KMS_KEY",
	}
}

type ServiceName string

// Enum values for ServiceName
const (
	ServiceNameApiGateway       ServiceName = "API_GATEWAY"
	ServiceNameApplicationElb   ServiceName = "APPLICATION_ELB"
	ServiceNameAutoScalingGroup ServiceName = "AUTO_SCALING_GROUP"
	ServiceNameCloudFront       ServiceName = "CLOUD_FRONT"
	ServiceNameDynamoDb         ServiceName = "DYNAMO_DB"
	ServiceNameEc2              ServiceName = "EC2"
	ServiceNameEcs              ServiceName = "ECS"
	ServiceNameEks              ServiceName = "EKS"
	ServiceNameElasticBeanstalk ServiceName = "ELASTIC_BEANSTALK"
	ServiceNameElastiCache      ServiceName = "ELASTI_CACHE"
	ServiceNameElb              ServiceName = "ELB"
	ServiceNameEs               ServiceName = "ES"
	ServiceNameKinesis          ServiceName = "KINESIS"
	ServiceNameLambda           ServiceName = "LAMBDA"
	ServiceNameNatGateway       ServiceName = "NAT_GATEWAY"
	ServiceNameNetworkElb       ServiceName = "NETWORK_ELB"
	ServiceNameRds              ServiceName = "RDS"
	ServiceNameRedshift         ServiceName = "REDSHIFT"
	ServiceNameRoute53          ServiceName = "ROUTE_53"
	ServiceNameS3               ServiceName = "S3"
	ServiceNameSageMaker        ServiceName = "SAGE_MAKER"
	ServiceNameSns              ServiceName = "SNS"
	ServiceNameSqs              ServiceName = "SQS"
	ServiceNameStepFunctions    ServiceName = "STEP_FUNCTIONS"
	ServiceNameSwf              ServiceName = "SWF"
)

// Values returns all known values for ServiceName. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServiceName) Values() []ServiceName {
	return []ServiceName{
		"API_GATEWAY",
		"APPLICATION_ELB",
		"AUTO_SCALING_GROUP",
		"CLOUD_FRONT",
		"DYNAMO_DB",
		"EC2",
		"ECS",
		"EKS",
		"ELASTIC_BEANSTALK",
		"ELASTI_CACHE",
		"ELB",
		"ES",
		"KINESIS",
		"LAMBDA",
		"NAT_GATEWAY",
		"NETWORK_ELB",
		"RDS",
		"REDSHIFT",
		"ROUTE_53",
		"S3",
		"SAGE_MAKER",
		"SNS",
		"SQS",
		"STEP_FUNCTIONS",
		"SWF",
	}
}

type UpdateResourceCollectionAction string

// Enum values for UpdateResourceCollectionAction
const (
	UpdateResourceCollectionActionAdd    UpdateResourceCollectionAction = "ADD"
	UpdateResourceCollectionActionRemove UpdateResourceCollectionAction = "REMOVE"
)

// Values returns all known values for UpdateResourceCollectionAction. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UpdateResourceCollectionAction) Values() []UpdateResourceCollectionAction {
	return []UpdateResourceCollectionAction{
		"ADD",
		"REMOVE",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation                      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonCannotParse                           ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonFieldValidationFailed                 ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther                                 ValidationExceptionReason = "OTHER"
	ValidationExceptionReasonInvalidParameterCombination           ValidationExceptionReason = "INVALID_PARAMETER_COMBINATION"
	ValidationExceptionReasonParameterInconsistentWithServiceState ValidationExceptionReason = "PARAMETER_INCONSISTENT_WITH_SERVICE_STATE"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UNKNOWN_OPERATION",
		"CANNOT_PARSE",
		"FIELD_VALIDATION_FAILED",
		"OTHER",
		"INVALID_PARAMETER_COMBINATION",
		"PARAMETER_INCONSISTENT_WITH_SERVICE_STATE",
	}
}

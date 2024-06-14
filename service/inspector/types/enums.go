// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessDeniedErrorCode string

// Enum values for AccessDeniedErrorCode
const (
	AccessDeniedErrorCodeAccessDeniedToAssessmentTarget   AccessDeniedErrorCode = "ACCESS_DENIED_TO_ASSESSMENT_TARGET"
	AccessDeniedErrorCodeAccessDeniedToAssessmentTemplate AccessDeniedErrorCode = "ACCESS_DENIED_TO_ASSESSMENT_TEMPLATE"
	AccessDeniedErrorCodeAccessDeniedToAssessmentRun      AccessDeniedErrorCode = "ACCESS_DENIED_TO_ASSESSMENT_RUN"
	AccessDeniedErrorCodeAccessDeniedToFinding            AccessDeniedErrorCode = "ACCESS_DENIED_TO_FINDING"
	AccessDeniedErrorCodeAccessDeniedToResourceGroup      AccessDeniedErrorCode = "ACCESS_DENIED_TO_RESOURCE_GROUP"
	AccessDeniedErrorCodeAccessDeniedToRulesPackage       AccessDeniedErrorCode = "ACCESS_DENIED_TO_RULES_PACKAGE"
	AccessDeniedErrorCodeAccessDeniedToSnsTopic           AccessDeniedErrorCode = "ACCESS_DENIED_TO_SNS_TOPIC"
	AccessDeniedErrorCodeAccessDeniedToIamRole            AccessDeniedErrorCode = "ACCESS_DENIED_TO_IAM_ROLE"
)

// Values returns all known values for AccessDeniedErrorCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccessDeniedErrorCode) Values() []AccessDeniedErrorCode {
	return []AccessDeniedErrorCode{
		"ACCESS_DENIED_TO_ASSESSMENT_TARGET",
		"ACCESS_DENIED_TO_ASSESSMENT_TEMPLATE",
		"ACCESS_DENIED_TO_ASSESSMENT_RUN",
		"ACCESS_DENIED_TO_FINDING",
		"ACCESS_DENIED_TO_RESOURCE_GROUP",
		"ACCESS_DENIED_TO_RULES_PACKAGE",
		"ACCESS_DENIED_TO_SNS_TOPIC",
		"ACCESS_DENIED_TO_IAM_ROLE",
	}
}

type AgentHealth string

// Enum values for AgentHealth
const (
	AgentHealthHealthy   AgentHealth = "HEALTHY"
	AgentHealthUnhealthy AgentHealth = "UNHEALTHY"
	AgentHealthUnknown   AgentHealth = "UNKNOWN"
)

// Values returns all known values for AgentHealth. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AgentHealth) Values() []AgentHealth {
	return []AgentHealth{
		"HEALTHY",
		"UNHEALTHY",
		"UNKNOWN",
	}
}

type AgentHealthCode string

// Enum values for AgentHealthCode
const (
	AgentHealthCodeIdle      AgentHealthCode = "IDLE"
	AgentHealthCodeRunning   AgentHealthCode = "RUNNING"
	AgentHealthCodeShutdown  AgentHealthCode = "SHUTDOWN"
	AgentHealthCodeUnhealthy AgentHealthCode = "UNHEALTHY"
	AgentHealthCodeThrottled AgentHealthCode = "THROTTLED"
	AgentHealthCodeUnknown   AgentHealthCode = "UNKNOWN"
)

// Values returns all known values for AgentHealthCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AgentHealthCode) Values() []AgentHealthCode {
	return []AgentHealthCode{
		"IDLE",
		"RUNNING",
		"SHUTDOWN",
		"UNHEALTHY",
		"THROTTLED",
		"UNKNOWN",
	}
}

type AssessmentRunNotificationSnsStatusCode string

// Enum values for AssessmentRunNotificationSnsStatusCode
const (
	AssessmentRunNotificationSnsStatusCodeSuccess           AssessmentRunNotificationSnsStatusCode = "SUCCESS"
	AssessmentRunNotificationSnsStatusCodeTopicDoesNotExist AssessmentRunNotificationSnsStatusCode = "TOPIC_DOES_NOT_EXIST"
	AssessmentRunNotificationSnsStatusCodeAccessDenied      AssessmentRunNotificationSnsStatusCode = "ACCESS_DENIED"
	AssessmentRunNotificationSnsStatusCodeInternalError     AssessmentRunNotificationSnsStatusCode = "INTERNAL_ERROR"
)

// Values returns all known values for AssessmentRunNotificationSnsStatusCode.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssessmentRunNotificationSnsStatusCode) Values() []AssessmentRunNotificationSnsStatusCode {
	return []AssessmentRunNotificationSnsStatusCode{
		"SUCCESS",
		"TOPIC_DOES_NOT_EXIST",
		"ACCESS_DENIED",
		"INTERNAL_ERROR",
	}
}

type AssessmentRunState string

// Enum values for AssessmentRunState
const (
	AssessmentRunStateCreated                       AssessmentRunState = "CREATED"
	AssessmentRunStateStartDataCollectionPending    AssessmentRunState = "START_DATA_COLLECTION_PENDING"
	AssessmentRunStateStartDataCollectionInProgress AssessmentRunState = "START_DATA_COLLECTION_IN_PROGRESS"
	AssessmentRunStateCollectingData                AssessmentRunState = "COLLECTING_DATA"
	AssessmentRunStateStopDataCollectionPending     AssessmentRunState = "STOP_DATA_COLLECTION_PENDING"
	AssessmentRunStateDataCollected                 AssessmentRunState = "DATA_COLLECTED"
	AssessmentRunStateStartEvaluatingRulesPending   AssessmentRunState = "START_EVALUATING_RULES_PENDING"
	AssessmentRunStateEvaluatingRules               AssessmentRunState = "EVALUATING_RULES"
	AssessmentRunStateFailed                        AssessmentRunState = "FAILED"
	AssessmentRunStateError                         AssessmentRunState = "ERROR"
	AssessmentRunStateCompleted                     AssessmentRunState = "COMPLETED"
	AssessmentRunStateCompletedWithErrors           AssessmentRunState = "COMPLETED_WITH_ERRORS"
	AssessmentRunStateCanceled                      AssessmentRunState = "CANCELED"
)

// Values returns all known values for AssessmentRunState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssessmentRunState) Values() []AssessmentRunState {
	return []AssessmentRunState{
		"CREATED",
		"START_DATA_COLLECTION_PENDING",
		"START_DATA_COLLECTION_IN_PROGRESS",
		"COLLECTING_DATA",
		"STOP_DATA_COLLECTION_PENDING",
		"DATA_COLLECTED",
		"START_EVALUATING_RULES_PENDING",
		"EVALUATING_RULES",
		"FAILED",
		"ERROR",
		"COMPLETED",
		"COMPLETED_WITH_ERRORS",
		"CANCELED",
	}
}

type AssetType string

// Enum values for AssetType
const (
	AssetTypeEc2Instance AssetType = "ec2-instance"
)

// Values returns all known values for AssetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssetType) Values() []AssetType {
	return []AssetType{
		"ec2-instance",
	}
}

type FailedItemErrorCode string

// Enum values for FailedItemErrorCode
const (
	FailedItemErrorCodeInvalidArn       FailedItemErrorCode = "INVALID_ARN"
	FailedItemErrorCodeDuplicateArn     FailedItemErrorCode = "DUPLICATE_ARN"
	FailedItemErrorCodeItemDoesNotExist FailedItemErrorCode = "ITEM_DOES_NOT_EXIST"
	FailedItemErrorCodeAccessDenied     FailedItemErrorCode = "ACCESS_DENIED"
	FailedItemErrorCodeLimitExceeded    FailedItemErrorCode = "LIMIT_EXCEEDED"
	FailedItemErrorCodeInternalError    FailedItemErrorCode = "INTERNAL_ERROR"
)

// Values returns all known values for FailedItemErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FailedItemErrorCode) Values() []FailedItemErrorCode {
	return []FailedItemErrorCode{
		"INVALID_ARN",
		"DUPLICATE_ARN",
		"ITEM_DOES_NOT_EXIST",
		"ACCESS_DENIED",
		"LIMIT_EXCEEDED",
		"INTERNAL_ERROR",
	}
}

type InspectorEvent string

// Enum values for InspectorEvent
const (
	InspectorEventAssessmentRunStarted      InspectorEvent = "ASSESSMENT_RUN_STARTED"
	InspectorEventAssessmentRunCompleted    InspectorEvent = "ASSESSMENT_RUN_COMPLETED"
	InspectorEventAssessmentRunStateChanged InspectorEvent = "ASSESSMENT_RUN_STATE_CHANGED"
	InspectorEventFindingReported           InspectorEvent = "FINDING_REPORTED"
	InspectorEventOther                     InspectorEvent = "OTHER"
)

// Values returns all known values for InspectorEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InspectorEvent) Values() []InspectorEvent {
	return []InspectorEvent{
		"ASSESSMENT_RUN_STARTED",
		"ASSESSMENT_RUN_COMPLETED",
		"ASSESSMENT_RUN_STATE_CHANGED",
		"FINDING_REPORTED",
		"OTHER",
	}
}

type InvalidCrossAccountRoleErrorCode string

// Enum values for InvalidCrossAccountRoleErrorCode
const (
	InvalidCrossAccountRoleErrorCodeRoleDoesNotExistOrInvalidTrustRelationship InvalidCrossAccountRoleErrorCode = "ROLE_DOES_NOT_EXIST_OR_INVALID_TRUST_RELATIONSHIP"
	InvalidCrossAccountRoleErrorCodeRoleDoesNotHaveCorrectPolicy               InvalidCrossAccountRoleErrorCode = "ROLE_DOES_NOT_HAVE_CORRECT_POLICY"
)

// Values returns all known values for InvalidCrossAccountRoleErrorCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InvalidCrossAccountRoleErrorCode) Values() []InvalidCrossAccountRoleErrorCode {
	return []InvalidCrossAccountRoleErrorCode{
		"ROLE_DOES_NOT_EXIST_OR_INVALID_TRUST_RELATIONSHIP",
		"ROLE_DOES_NOT_HAVE_CORRECT_POLICY",
	}
}

type InvalidInputErrorCode string

// Enum values for InvalidInputErrorCode
const (
	InvalidInputErrorCodeInvalidAssessmentTargetArn               InvalidInputErrorCode = "INVALID_ASSESSMENT_TARGET_ARN"
	InvalidInputErrorCodeInvalidAssessmentTemplateArn             InvalidInputErrorCode = "INVALID_ASSESSMENT_TEMPLATE_ARN"
	InvalidInputErrorCodeInvalidAssessmentRunArn                  InvalidInputErrorCode = "INVALID_ASSESSMENT_RUN_ARN"
	InvalidInputErrorCodeInvalidFindingArn                        InvalidInputErrorCode = "INVALID_FINDING_ARN"
	InvalidInputErrorCodeInvalidResourceGroupArn                  InvalidInputErrorCode = "INVALID_RESOURCE_GROUP_ARN"
	InvalidInputErrorCodeInvalidRulesPackageArn                   InvalidInputErrorCode = "INVALID_RULES_PACKAGE_ARN"
	InvalidInputErrorCodeInvalidResourceArn                       InvalidInputErrorCode = "INVALID_RESOURCE_ARN"
	InvalidInputErrorCodeInvalidSnsTopicArn                       InvalidInputErrorCode = "INVALID_SNS_TOPIC_ARN"
	InvalidInputErrorCodeInvalidIamRoleArn                        InvalidInputErrorCode = "INVALID_IAM_ROLE_ARN"
	InvalidInputErrorCodeInvalidAssessmentTargetName              InvalidInputErrorCode = "INVALID_ASSESSMENT_TARGET_NAME"
	InvalidInputErrorCodeInvalidAssessmentTargetNamePattern       InvalidInputErrorCode = "INVALID_ASSESSMENT_TARGET_NAME_PATTERN"
	InvalidInputErrorCodeInvalidAssessmentTemplateName            InvalidInputErrorCode = "INVALID_ASSESSMENT_TEMPLATE_NAME"
	InvalidInputErrorCodeInvalidAssessmentTemplateNamePattern     InvalidInputErrorCode = "INVALID_ASSESSMENT_TEMPLATE_NAME_PATTERN"
	InvalidInputErrorCodeInvalidAssessmentTemplateDuration        InvalidInputErrorCode = "INVALID_ASSESSMENT_TEMPLATE_DURATION"
	InvalidInputErrorCodeInvalidAssessmentTemplateDurationRange   InvalidInputErrorCode = "INVALID_ASSESSMENT_TEMPLATE_DURATION_RANGE"
	InvalidInputErrorCodeInvalidAssessmentRunDurationRange        InvalidInputErrorCode = "INVALID_ASSESSMENT_RUN_DURATION_RANGE"
	InvalidInputErrorCodeInvalidAssessmentRunStartTimeRange       InvalidInputErrorCode = "INVALID_ASSESSMENT_RUN_START_TIME_RANGE"
	InvalidInputErrorCodeInvalidAssessmentRunCompletionTimeRange  InvalidInputErrorCode = "INVALID_ASSESSMENT_RUN_COMPLETION_TIME_RANGE"
	InvalidInputErrorCodeInvalidAssessmentRunStateChangeTimeRange InvalidInputErrorCode = "INVALID_ASSESSMENT_RUN_STATE_CHANGE_TIME_RANGE"
	InvalidInputErrorCodeInvalidAssessmentRunState                InvalidInputErrorCode = "INVALID_ASSESSMENT_RUN_STATE"
	InvalidInputErrorCodeInvalidTag                               InvalidInputErrorCode = "INVALID_TAG"
	InvalidInputErrorCodeInvalidTagKey                            InvalidInputErrorCode = "INVALID_TAG_KEY"
	InvalidInputErrorCodeInvalidTagValue                          InvalidInputErrorCode = "INVALID_TAG_VALUE"
	InvalidInputErrorCodeInvalidResourceGroupTagKey               InvalidInputErrorCode = "INVALID_RESOURCE_GROUP_TAG_KEY"
	InvalidInputErrorCodeInvalidResourceGroupTagValue             InvalidInputErrorCode = "INVALID_RESOURCE_GROUP_TAG_VALUE"
	InvalidInputErrorCodeInvalidAttribute                         InvalidInputErrorCode = "INVALID_ATTRIBUTE"
	InvalidInputErrorCodeInvalidUserAttribute                     InvalidInputErrorCode = "INVALID_USER_ATTRIBUTE"
	InvalidInputErrorCodeInvalidUserAttributeKey                  InvalidInputErrorCode = "INVALID_USER_ATTRIBUTE_KEY"
	InvalidInputErrorCodeInvalidUserAttributeValue                InvalidInputErrorCode = "INVALID_USER_ATTRIBUTE_VALUE"
	InvalidInputErrorCodeInvalidPaginationToken                   InvalidInputErrorCode = "INVALID_PAGINATION_TOKEN"
	InvalidInputErrorCodeInvalidMaxResults                        InvalidInputErrorCode = "INVALID_MAX_RESULTS"
	InvalidInputErrorCodeInvalidAgentId                           InvalidInputErrorCode = "INVALID_AGENT_ID"
	InvalidInputErrorCodeInvalidAutoScalingGroup                  InvalidInputErrorCode = "INVALID_AUTO_SCALING_GROUP"
	InvalidInputErrorCodeInvalidRuleName                          InvalidInputErrorCode = "INVALID_RULE_NAME"
	InvalidInputErrorCodeInvalidSeverity                          InvalidInputErrorCode = "INVALID_SEVERITY"
	InvalidInputErrorCodeInvalidLocale                            InvalidInputErrorCode = "INVALID_LOCALE"
	InvalidInputErrorCodeInvalidEvent                             InvalidInputErrorCode = "INVALID_EVENT"
	InvalidInputErrorCodeAssessmentTargetNameAlreadyTaken         InvalidInputErrorCode = "ASSESSMENT_TARGET_NAME_ALREADY_TAKEN"
	InvalidInputErrorCodeAssessmentTemplateNameAlreadyTaken       InvalidInputErrorCode = "ASSESSMENT_TEMPLATE_NAME_ALREADY_TAKEN"
	InvalidInputErrorCodeInvalidNumberOfAssessmentTargetArns      InvalidInputErrorCode = "INVALID_NUMBER_OF_ASSESSMENT_TARGET_ARNS"
	InvalidInputErrorCodeInvalidNumberOfAssessmentTemplateArns    InvalidInputErrorCode = "INVALID_NUMBER_OF_ASSESSMENT_TEMPLATE_ARNS"
	InvalidInputErrorCodeInvalidNumberOfAssessmentRunArns         InvalidInputErrorCode = "INVALID_NUMBER_OF_ASSESSMENT_RUN_ARNS"
	InvalidInputErrorCodeInvalidNumberOfFindingArns               InvalidInputErrorCode = "INVALID_NUMBER_OF_FINDING_ARNS"
	InvalidInputErrorCodeInvalidNumberOfResourceGroupArns         InvalidInputErrorCode = "INVALID_NUMBER_OF_RESOURCE_GROUP_ARNS"
	InvalidInputErrorCodeInvalidNumberOfRulesPackageArns          InvalidInputErrorCode = "INVALID_NUMBER_OF_RULES_PACKAGE_ARNS"
	InvalidInputErrorCodeInvalidNumberOfAssessmentRunStates       InvalidInputErrorCode = "INVALID_NUMBER_OF_ASSESSMENT_RUN_STATES"
	InvalidInputErrorCodeInvalidNumberOfTags                      InvalidInputErrorCode = "INVALID_NUMBER_OF_TAGS"
	InvalidInputErrorCodeInvalidNumberOfResourceGroupTags         InvalidInputErrorCode = "INVALID_NUMBER_OF_RESOURCE_GROUP_TAGS"
	InvalidInputErrorCodeInvalidNumberOfAttributes                InvalidInputErrorCode = "INVALID_NUMBER_OF_ATTRIBUTES"
	InvalidInputErrorCodeInvalidNumberOfUserAttributes            InvalidInputErrorCode = "INVALID_NUMBER_OF_USER_ATTRIBUTES"
	InvalidInputErrorCodeInvalidNumberOfAgentIds                  InvalidInputErrorCode = "INVALID_NUMBER_OF_AGENT_IDS"
	InvalidInputErrorCodeInvalidNumberOfAutoScalingGroups         InvalidInputErrorCode = "INVALID_NUMBER_OF_AUTO_SCALING_GROUPS"
	InvalidInputErrorCodeInvalidNumberOfRuleNames                 InvalidInputErrorCode = "INVALID_NUMBER_OF_RULE_NAMES"
	InvalidInputErrorCodeInvalidNumberOfSeverities                InvalidInputErrorCode = "INVALID_NUMBER_OF_SEVERITIES"
)

// Values returns all known values for InvalidInputErrorCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InvalidInputErrorCode) Values() []InvalidInputErrorCode {
	return []InvalidInputErrorCode{
		"INVALID_ASSESSMENT_TARGET_ARN",
		"INVALID_ASSESSMENT_TEMPLATE_ARN",
		"INVALID_ASSESSMENT_RUN_ARN",
		"INVALID_FINDING_ARN",
		"INVALID_RESOURCE_GROUP_ARN",
		"INVALID_RULES_PACKAGE_ARN",
		"INVALID_RESOURCE_ARN",
		"INVALID_SNS_TOPIC_ARN",
		"INVALID_IAM_ROLE_ARN",
		"INVALID_ASSESSMENT_TARGET_NAME",
		"INVALID_ASSESSMENT_TARGET_NAME_PATTERN",
		"INVALID_ASSESSMENT_TEMPLATE_NAME",
		"INVALID_ASSESSMENT_TEMPLATE_NAME_PATTERN",
		"INVALID_ASSESSMENT_TEMPLATE_DURATION",
		"INVALID_ASSESSMENT_TEMPLATE_DURATION_RANGE",
		"INVALID_ASSESSMENT_RUN_DURATION_RANGE",
		"INVALID_ASSESSMENT_RUN_START_TIME_RANGE",
		"INVALID_ASSESSMENT_RUN_COMPLETION_TIME_RANGE",
		"INVALID_ASSESSMENT_RUN_STATE_CHANGE_TIME_RANGE",
		"INVALID_ASSESSMENT_RUN_STATE",
		"INVALID_TAG",
		"INVALID_TAG_KEY",
		"INVALID_TAG_VALUE",
		"INVALID_RESOURCE_GROUP_TAG_KEY",
		"INVALID_RESOURCE_GROUP_TAG_VALUE",
		"INVALID_ATTRIBUTE",
		"INVALID_USER_ATTRIBUTE",
		"INVALID_USER_ATTRIBUTE_KEY",
		"INVALID_USER_ATTRIBUTE_VALUE",
		"INVALID_PAGINATION_TOKEN",
		"INVALID_MAX_RESULTS",
		"INVALID_AGENT_ID",
		"INVALID_AUTO_SCALING_GROUP",
		"INVALID_RULE_NAME",
		"INVALID_SEVERITY",
		"INVALID_LOCALE",
		"INVALID_EVENT",
		"ASSESSMENT_TARGET_NAME_ALREADY_TAKEN",
		"ASSESSMENT_TEMPLATE_NAME_ALREADY_TAKEN",
		"INVALID_NUMBER_OF_ASSESSMENT_TARGET_ARNS",
		"INVALID_NUMBER_OF_ASSESSMENT_TEMPLATE_ARNS",
		"INVALID_NUMBER_OF_ASSESSMENT_RUN_ARNS",
		"INVALID_NUMBER_OF_FINDING_ARNS",
		"INVALID_NUMBER_OF_RESOURCE_GROUP_ARNS",
		"INVALID_NUMBER_OF_RULES_PACKAGE_ARNS",
		"INVALID_NUMBER_OF_ASSESSMENT_RUN_STATES",
		"INVALID_NUMBER_OF_TAGS",
		"INVALID_NUMBER_OF_RESOURCE_GROUP_TAGS",
		"INVALID_NUMBER_OF_ATTRIBUTES",
		"INVALID_NUMBER_OF_USER_ATTRIBUTES",
		"INVALID_NUMBER_OF_AGENT_IDS",
		"INVALID_NUMBER_OF_AUTO_SCALING_GROUPS",
		"INVALID_NUMBER_OF_RULE_NAMES",
		"INVALID_NUMBER_OF_SEVERITIES",
	}
}

type LimitExceededErrorCode string

// Enum values for LimitExceededErrorCode
const (
	LimitExceededErrorCodeAssessmentTargetLimitExceeded   LimitExceededErrorCode = "ASSESSMENT_TARGET_LIMIT_EXCEEDED"
	LimitExceededErrorCodeAssessmentTemplateLimitExceeded LimitExceededErrorCode = "ASSESSMENT_TEMPLATE_LIMIT_EXCEEDED"
	LimitExceededErrorCodeAssessmentRunLimitExceeded      LimitExceededErrorCode = "ASSESSMENT_RUN_LIMIT_EXCEEDED"
	LimitExceededErrorCodeResourceGroupLimitExceeded      LimitExceededErrorCode = "RESOURCE_GROUP_LIMIT_EXCEEDED"
	LimitExceededErrorCodeEventSubscriptionLimitExceeded  LimitExceededErrorCode = "EVENT_SUBSCRIPTION_LIMIT_EXCEEDED"
)

// Values returns all known values for LimitExceededErrorCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LimitExceededErrorCode) Values() []LimitExceededErrorCode {
	return []LimitExceededErrorCode{
		"ASSESSMENT_TARGET_LIMIT_EXCEEDED",
		"ASSESSMENT_TEMPLATE_LIMIT_EXCEEDED",
		"ASSESSMENT_RUN_LIMIT_EXCEEDED",
		"RESOURCE_GROUP_LIMIT_EXCEEDED",
		"EVENT_SUBSCRIPTION_LIMIT_EXCEEDED",
	}
}

type Locale string

// Enum values for Locale
const (
	LocaleEnUs Locale = "EN_US"
)

// Values returns all known values for Locale. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Locale) Values() []Locale {
	return []Locale{
		"EN_US",
	}
}

type NoSuchEntityErrorCode string

// Enum values for NoSuchEntityErrorCode
const (
	NoSuchEntityErrorCodeAssessmentTargetDoesNotExist   NoSuchEntityErrorCode = "ASSESSMENT_TARGET_DOES_NOT_EXIST"
	NoSuchEntityErrorCodeAssessmentTemplateDoesNotExist NoSuchEntityErrorCode = "ASSESSMENT_TEMPLATE_DOES_NOT_EXIST"
	NoSuchEntityErrorCodeAssessmentRunDoesNotExist      NoSuchEntityErrorCode = "ASSESSMENT_RUN_DOES_NOT_EXIST"
	NoSuchEntityErrorCodeFindingDoesNotExist            NoSuchEntityErrorCode = "FINDING_DOES_NOT_EXIST"
	NoSuchEntityErrorCodeResourceGroupDoesNotExist      NoSuchEntityErrorCode = "RESOURCE_GROUP_DOES_NOT_EXIST"
	NoSuchEntityErrorCodeRulesPackageDoesNotExist       NoSuchEntityErrorCode = "RULES_PACKAGE_DOES_NOT_EXIST"
	NoSuchEntityErrorCodeSnsTopicDoesNotExist           NoSuchEntityErrorCode = "SNS_TOPIC_DOES_NOT_EXIST"
	NoSuchEntityErrorCodeIamRoleDoesNotExist            NoSuchEntityErrorCode = "IAM_ROLE_DOES_NOT_EXIST"
)

// Values returns all known values for NoSuchEntityErrorCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NoSuchEntityErrorCode) Values() []NoSuchEntityErrorCode {
	return []NoSuchEntityErrorCode{
		"ASSESSMENT_TARGET_DOES_NOT_EXIST",
		"ASSESSMENT_TEMPLATE_DOES_NOT_EXIST",
		"ASSESSMENT_RUN_DOES_NOT_EXIST",
		"FINDING_DOES_NOT_EXIST",
		"RESOURCE_GROUP_DOES_NOT_EXIST",
		"RULES_PACKAGE_DOES_NOT_EXIST",
		"SNS_TOPIC_DOES_NOT_EXIST",
		"IAM_ROLE_DOES_NOT_EXIST",
	}
}

type PreviewStatus string

// Enum values for PreviewStatus
const (
	PreviewStatusWorkInProgress PreviewStatus = "WORK_IN_PROGRESS"
	PreviewStatusCompleted      PreviewStatus = "COMPLETED"
)

// Values returns all known values for PreviewStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PreviewStatus) Values() []PreviewStatus {
	return []PreviewStatus{
		"WORK_IN_PROGRESS",
		"COMPLETED",
	}
}

type ReportFileFormat string

// Enum values for ReportFileFormat
const (
	ReportFileFormatHtml ReportFileFormat = "HTML"
	ReportFileFormatPdf  ReportFileFormat = "PDF"
)

// Values returns all known values for ReportFileFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReportFileFormat) Values() []ReportFileFormat {
	return []ReportFileFormat{
		"HTML",
		"PDF",
	}
}

type ReportStatus string

// Enum values for ReportStatus
const (
	ReportStatusWorkInProgress ReportStatus = "WORK_IN_PROGRESS"
	ReportStatusFailed         ReportStatus = "FAILED"
	ReportStatusCompleted      ReportStatus = "COMPLETED"
)

// Values returns all known values for ReportStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReportStatus) Values() []ReportStatus {
	return []ReportStatus{
		"WORK_IN_PROGRESS",
		"FAILED",
		"COMPLETED",
	}
}

type ReportType string

// Enum values for ReportType
const (
	ReportTypeFinding ReportType = "FINDING"
	ReportTypeFull    ReportType = "FULL"
)

// Values returns all known values for ReportType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReportType) Values() []ReportType {
	return []ReportType{
		"FINDING",
		"FULL",
	}
}

type ScopeType string

// Enum values for ScopeType
const (
	ScopeTypeInstanceId      ScopeType = "INSTANCE_ID"
	ScopeTypeRulesPackageArn ScopeType = "RULES_PACKAGE_ARN"
)

// Values returns all known values for ScopeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScopeType) Values() []ScopeType {
	return []ScopeType{
		"INSTANCE_ID",
		"RULES_PACKAGE_ARN",
	}
}

type Severity string

// Enum values for Severity
const (
	SeverityLow           Severity = "Low"
	SeverityMedium        Severity = "Medium"
	SeverityHigh          Severity = "High"
	SeverityInformational Severity = "Informational"
	SeverityUndefined     Severity = "Undefined"
)

// Values returns all known values for Severity. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Severity) Values() []Severity {
	return []Severity{
		"Low",
		"Medium",
		"High",
		"Informational",
		"Undefined",
	}
}

type StopAction string

// Enum values for StopAction
const (
	StopActionStartEvaluation StopAction = "START_EVALUATION"
	StopActionSkipEvaluation  StopAction = "SKIP_EVALUATION"
)

// Values returns all known values for StopAction. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StopAction) Values() []StopAction {
	return []StopAction{
		"START_EVALUATION",
		"SKIP_EVALUATION",
	}
}

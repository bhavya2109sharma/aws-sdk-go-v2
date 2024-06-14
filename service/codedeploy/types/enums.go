// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApplicationRevisionSortBy string

// Enum values for ApplicationRevisionSortBy
const (
	ApplicationRevisionSortByRegisterTime  ApplicationRevisionSortBy = "registerTime"
	ApplicationRevisionSortByFirstUsedTime ApplicationRevisionSortBy = "firstUsedTime"
	ApplicationRevisionSortByLastUsedTime  ApplicationRevisionSortBy = "lastUsedTime"
)

// Values returns all known values for ApplicationRevisionSortBy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationRevisionSortBy) Values() []ApplicationRevisionSortBy {
	return []ApplicationRevisionSortBy{
		"registerTime",
		"firstUsedTime",
		"lastUsedTime",
	}
}

type AutoRollbackEvent string

// Enum values for AutoRollbackEvent
const (
	AutoRollbackEventDeploymentFailure       AutoRollbackEvent = "DEPLOYMENT_FAILURE"
	AutoRollbackEventDeploymentStopOnAlarm   AutoRollbackEvent = "DEPLOYMENT_STOP_ON_ALARM"
	AutoRollbackEventDeploymentStopOnRequest AutoRollbackEvent = "DEPLOYMENT_STOP_ON_REQUEST"
)

// Values returns all known values for AutoRollbackEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AutoRollbackEvent) Values() []AutoRollbackEvent {
	return []AutoRollbackEvent{
		"DEPLOYMENT_FAILURE",
		"DEPLOYMENT_STOP_ON_ALARM",
		"DEPLOYMENT_STOP_ON_REQUEST",
	}
}

type BundleType string

// Enum values for BundleType
const (
	BundleTypeTar     BundleType = "tar"
	BundleTypeTarGZip BundleType = "tgz"
	BundleTypeZip     BundleType = "zip"
	BundleTypeYaml    BundleType = "YAML"
	BundleTypeJson    BundleType = "JSON"
)

// Values returns all known values for BundleType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BundleType) Values() []BundleType {
	return []BundleType{
		"tar",
		"tgz",
		"zip",
		"YAML",
		"JSON",
	}
}

type ComputePlatform string

// Enum values for ComputePlatform
const (
	ComputePlatformServer ComputePlatform = "Server"
	ComputePlatformLambda ComputePlatform = "Lambda"
	ComputePlatformEcs    ComputePlatform = "ECS"
)

// Values returns all known values for ComputePlatform. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ComputePlatform) Values() []ComputePlatform {
	return []ComputePlatform{
		"Server",
		"Lambda",
		"ECS",
	}
}

type DeploymentCreator string

// Enum values for DeploymentCreator
const (
	DeploymentCreatorUser                   DeploymentCreator = "user"
	DeploymentCreatorAutoscaling            DeploymentCreator = "autoscaling"
	DeploymentCreatorCodeDeployRollback     DeploymentCreator = "codeDeployRollback"
	DeploymentCreatorCodeDeploy             DeploymentCreator = "CodeDeploy"
	DeploymentCreatorCodeDeployAutoUpdate   DeploymentCreator = "CodeDeployAutoUpdate"
	DeploymentCreatorCloudFormation         DeploymentCreator = "CloudFormation"
	DeploymentCreatorCloudFormationRollback DeploymentCreator = "CloudFormationRollback"
	DeploymentCreatorAutoscalingTermination DeploymentCreator = "autoscalingTermination"
)

// Values returns all known values for DeploymentCreator. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentCreator) Values() []DeploymentCreator {
	return []DeploymentCreator{
		"user",
		"autoscaling",
		"codeDeployRollback",
		"CodeDeploy",
		"CodeDeployAutoUpdate",
		"CloudFormation",
		"CloudFormationRollback",
		"autoscalingTermination",
	}
}

type DeploymentOption string

// Enum values for DeploymentOption
const (
	DeploymentOptionWithTrafficControl    DeploymentOption = "WITH_TRAFFIC_CONTROL"
	DeploymentOptionWithoutTrafficControl DeploymentOption = "WITHOUT_TRAFFIC_CONTROL"
)

// Values returns all known values for DeploymentOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentOption) Values() []DeploymentOption {
	return []DeploymentOption{
		"WITH_TRAFFIC_CONTROL",
		"WITHOUT_TRAFFIC_CONTROL",
	}
}

type DeploymentReadyAction string

// Enum values for DeploymentReadyAction
const (
	DeploymentReadyActionContinueDeployment DeploymentReadyAction = "CONTINUE_DEPLOYMENT"
	DeploymentReadyActionStopDeployment     DeploymentReadyAction = "STOP_DEPLOYMENT"
)

// Values returns all known values for DeploymentReadyAction. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentReadyAction) Values() []DeploymentReadyAction {
	return []DeploymentReadyAction{
		"CONTINUE_DEPLOYMENT",
		"STOP_DEPLOYMENT",
	}
}

type DeploymentStatus string

// Enum values for DeploymentStatus
const (
	DeploymentStatusCreated    DeploymentStatus = "Created"
	DeploymentStatusQueued     DeploymentStatus = "Queued"
	DeploymentStatusInProgress DeploymentStatus = "InProgress"
	DeploymentStatusBaking     DeploymentStatus = "Baking"
	DeploymentStatusSucceeded  DeploymentStatus = "Succeeded"
	DeploymentStatusFailed     DeploymentStatus = "Failed"
	DeploymentStatusStopped    DeploymentStatus = "Stopped"
	DeploymentStatusReady      DeploymentStatus = "Ready"
)

// Values returns all known values for DeploymentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		"Created",
		"Queued",
		"InProgress",
		"Baking",
		"Succeeded",
		"Failed",
		"Stopped",
		"Ready",
	}
}

type DeploymentTargetType string

// Enum values for DeploymentTargetType
const (
	DeploymentTargetTypeInstanceTarget       DeploymentTargetType = "InstanceTarget"
	DeploymentTargetTypeLambdaTarget         DeploymentTargetType = "LambdaTarget"
	DeploymentTargetTypeEcsTarget            DeploymentTargetType = "ECSTarget"
	DeploymentTargetTypeCloudformationTarget DeploymentTargetType = "CloudFormationTarget"
)

// Values returns all known values for DeploymentTargetType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentTargetType) Values() []DeploymentTargetType {
	return []DeploymentTargetType{
		"InstanceTarget",
		"LambdaTarget",
		"ECSTarget",
		"CloudFormationTarget",
	}
}

type DeploymentType string

// Enum values for DeploymentType
const (
	DeploymentTypeInPlace   DeploymentType = "IN_PLACE"
	DeploymentTypeBlueGreen DeploymentType = "BLUE_GREEN"
)

// Values returns all known values for DeploymentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentType) Values() []DeploymentType {
	return []DeploymentType{
		"IN_PLACE",
		"BLUE_GREEN",
	}
}

type DeploymentWaitType string

// Enum values for DeploymentWaitType
const (
	DeploymentWaitTypeReadyWait       DeploymentWaitType = "READY_WAIT"
	DeploymentWaitTypeTerminationWait DeploymentWaitType = "TERMINATION_WAIT"
)

// Values returns all known values for DeploymentWaitType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentWaitType) Values() []DeploymentWaitType {
	return []DeploymentWaitType{
		"READY_WAIT",
		"TERMINATION_WAIT",
	}
}

type EC2TagFilterType string

// Enum values for EC2TagFilterType
const (
	EC2TagFilterTypeKeyOnly     EC2TagFilterType = "KEY_ONLY"
	EC2TagFilterTypeValueOnly   EC2TagFilterType = "VALUE_ONLY"
	EC2TagFilterTypeKeyAndValue EC2TagFilterType = "KEY_AND_VALUE"
)

// Values returns all known values for EC2TagFilterType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EC2TagFilterType) Values() []EC2TagFilterType {
	return []EC2TagFilterType{
		"KEY_ONLY",
		"VALUE_ONLY",
		"KEY_AND_VALUE",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeAgentIssue                              ErrorCode = "AGENT_ISSUE"
	ErrorCodeAlarmActive                             ErrorCode = "ALARM_ACTIVE"
	ErrorCodeApplicationMissing                      ErrorCode = "APPLICATION_MISSING"
	ErrorCodeAutoscalingValidationError              ErrorCode = "AUTOSCALING_VALIDATION_ERROR"
	ErrorCodeAutoScalingConfiguration                ErrorCode = "AUTO_SCALING_CONFIGURATION"
	ErrorCodeAutoScalingIamRolePermissions           ErrorCode = "AUTO_SCALING_IAM_ROLE_PERMISSIONS"
	ErrorCodeCodedeployResourceCannotBeFound         ErrorCode = "CODEDEPLOY_RESOURCE_CANNOT_BE_FOUND"
	ErrorCodeCustomerApplicationUnhealthy            ErrorCode = "CUSTOMER_APPLICATION_UNHEALTHY"
	ErrorCodeDeploymentGroupMissing                  ErrorCode = "DEPLOYMENT_GROUP_MISSING"
	ErrorCodeEcsUpdateError                          ErrorCode = "ECS_UPDATE_ERROR"
	ErrorCodeElasticLoadBalancingInvalid             ErrorCode = "ELASTIC_LOAD_BALANCING_INVALID"
	ErrorCodeElbInvalidInstance                      ErrorCode = "ELB_INVALID_INSTANCE"
	ErrorCodeHealthConstraints                       ErrorCode = "HEALTH_CONSTRAINTS"
	ErrorCodeHealthConstraintsInvalid                ErrorCode = "HEALTH_CONSTRAINTS_INVALID"
	ErrorCodeHookExecutionFailure                    ErrorCode = "HOOK_EXECUTION_FAILURE"
	ErrorCodeIamRoleMissing                          ErrorCode = "IAM_ROLE_MISSING"
	ErrorCodeIamRolePermissions                      ErrorCode = "IAM_ROLE_PERMISSIONS"
	ErrorCodeInternalError                           ErrorCode = "INTERNAL_ERROR"
	ErrorCodeInvalidEcsService                       ErrorCode = "INVALID_ECS_SERVICE"
	ErrorCodeInvalidLambdaConfiguration              ErrorCode = "INVALID_LAMBDA_CONFIGURATION"
	ErrorCodeInvalidLambdaFunction                   ErrorCode = "INVALID_LAMBDA_FUNCTION"
	ErrorCodeInvalidRevision                         ErrorCode = "INVALID_REVISION"
	ErrorCodeManualStop                              ErrorCode = "MANUAL_STOP"
	ErrorCodeMissingBlueGreenDeploymentConfiguration ErrorCode = "MISSING_BLUE_GREEN_DEPLOYMENT_CONFIGURATION"
	ErrorCodeMissingElbInformation                   ErrorCode = "MISSING_ELB_INFORMATION"
	ErrorCodeMissingGithubToken                      ErrorCode = "MISSING_GITHUB_TOKEN"
	ErrorCodeNoEc2Subscription                       ErrorCode = "NO_EC2_SUBSCRIPTION"
	ErrorCodeNoInstances                             ErrorCode = "NO_INSTANCES"
	ErrorCodeOverMaxInstances                        ErrorCode = "OVER_MAX_INSTANCES"
	ErrorCodeResourceLimitExceeded                   ErrorCode = "RESOURCE_LIMIT_EXCEEDED"
	ErrorCodeRevisionMissing                         ErrorCode = "REVISION_MISSING"
	ErrorCodeThrottled                               ErrorCode = "THROTTLED"
	ErrorCodeTimeout                                 ErrorCode = "TIMEOUT"
	ErrorCodeCloudformationStackFailure              ErrorCode = "CLOUDFORMATION_STACK_FAILURE"
)

// Values returns all known values for ErrorCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"AGENT_ISSUE",
		"ALARM_ACTIVE",
		"APPLICATION_MISSING",
		"AUTOSCALING_VALIDATION_ERROR",
		"AUTO_SCALING_CONFIGURATION",
		"AUTO_SCALING_IAM_ROLE_PERMISSIONS",
		"CODEDEPLOY_RESOURCE_CANNOT_BE_FOUND",
		"CUSTOMER_APPLICATION_UNHEALTHY",
		"DEPLOYMENT_GROUP_MISSING",
		"ECS_UPDATE_ERROR",
		"ELASTIC_LOAD_BALANCING_INVALID",
		"ELB_INVALID_INSTANCE",
		"HEALTH_CONSTRAINTS",
		"HEALTH_CONSTRAINTS_INVALID",
		"HOOK_EXECUTION_FAILURE",
		"IAM_ROLE_MISSING",
		"IAM_ROLE_PERMISSIONS",
		"INTERNAL_ERROR",
		"INVALID_ECS_SERVICE",
		"INVALID_LAMBDA_CONFIGURATION",
		"INVALID_LAMBDA_FUNCTION",
		"INVALID_REVISION",
		"MANUAL_STOP",
		"MISSING_BLUE_GREEN_DEPLOYMENT_CONFIGURATION",
		"MISSING_ELB_INFORMATION",
		"MISSING_GITHUB_TOKEN",
		"NO_EC2_SUBSCRIPTION",
		"NO_INSTANCES",
		"OVER_MAX_INSTANCES",
		"RESOURCE_LIMIT_EXCEEDED",
		"REVISION_MISSING",
		"THROTTLED",
		"TIMEOUT",
		"CLOUDFORMATION_STACK_FAILURE",
	}
}

type FileExistsBehavior string

// Enum values for FileExistsBehavior
const (
	FileExistsBehaviorDisallow  FileExistsBehavior = "DISALLOW"
	FileExistsBehaviorOverwrite FileExistsBehavior = "OVERWRITE"
	FileExistsBehaviorRetain    FileExistsBehavior = "RETAIN"
)

// Values returns all known values for FileExistsBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FileExistsBehavior) Values() []FileExistsBehavior {
	return []FileExistsBehavior{
		"DISALLOW",
		"OVERWRITE",
		"RETAIN",
	}
}

type GreenFleetProvisioningAction string

// Enum values for GreenFleetProvisioningAction
const (
	GreenFleetProvisioningActionDiscoverExisting     GreenFleetProvisioningAction = "DISCOVER_EXISTING"
	GreenFleetProvisioningActionCopyAutoScalingGroup GreenFleetProvisioningAction = "COPY_AUTO_SCALING_GROUP"
)

// Values returns all known values for GreenFleetProvisioningAction. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GreenFleetProvisioningAction) Values() []GreenFleetProvisioningAction {
	return []GreenFleetProvisioningAction{
		"DISCOVER_EXISTING",
		"COPY_AUTO_SCALING_GROUP",
	}
}

type InstanceAction string

// Enum values for InstanceAction
const (
	InstanceActionTerminate InstanceAction = "TERMINATE"
	InstanceActionKeepAlive InstanceAction = "KEEP_ALIVE"
)

// Values returns all known values for InstanceAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InstanceAction) Values() []InstanceAction {
	return []InstanceAction{
		"TERMINATE",
		"KEEP_ALIVE",
	}
}

type InstanceStatus string

// Enum values for InstanceStatus
const (
	InstanceStatusPending    InstanceStatus = "Pending"
	InstanceStatusInProgress InstanceStatus = "InProgress"
	InstanceStatusSucceeded  InstanceStatus = "Succeeded"
	InstanceStatusFailed     InstanceStatus = "Failed"
	InstanceStatusSkipped    InstanceStatus = "Skipped"
	InstanceStatusUnknown    InstanceStatus = "Unknown"
	InstanceStatusReady      InstanceStatus = "Ready"
)

// Values returns all known values for InstanceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InstanceStatus) Values() []InstanceStatus {
	return []InstanceStatus{
		"Pending",
		"InProgress",
		"Succeeded",
		"Failed",
		"Skipped",
		"Unknown",
		"Ready",
	}
}

type InstanceType string

// Enum values for InstanceType
const (
	InstanceTypeBlue  InstanceType = "Blue"
	InstanceTypeGreen InstanceType = "Green"
)

// Values returns all known values for InstanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InstanceType) Values() []InstanceType {
	return []InstanceType{
		"Blue",
		"Green",
	}
}

type LifecycleErrorCode string

// Enum values for LifecycleErrorCode
const (
	LifecycleErrorCodeSuccess             LifecycleErrorCode = "Success"
	LifecycleErrorCodeScriptMissing       LifecycleErrorCode = "ScriptMissing"
	LifecycleErrorCodeScriptNotExecutable LifecycleErrorCode = "ScriptNotExecutable"
	LifecycleErrorCodeScriptTimedOut      LifecycleErrorCode = "ScriptTimedOut"
	LifecycleErrorCodeScriptFailed        LifecycleErrorCode = "ScriptFailed"
	LifecycleErrorCodeUnknownError        LifecycleErrorCode = "UnknownError"
)

// Values returns all known values for LifecycleErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LifecycleErrorCode) Values() []LifecycleErrorCode {
	return []LifecycleErrorCode{
		"Success",
		"ScriptMissing",
		"ScriptNotExecutable",
		"ScriptTimedOut",
		"ScriptFailed",
		"UnknownError",
	}
}

type LifecycleEventStatus string

// Enum values for LifecycleEventStatus
const (
	LifecycleEventStatusPending    LifecycleEventStatus = "Pending"
	LifecycleEventStatusInProgress LifecycleEventStatus = "InProgress"
	LifecycleEventStatusSucceeded  LifecycleEventStatus = "Succeeded"
	LifecycleEventStatusFailed     LifecycleEventStatus = "Failed"
	LifecycleEventStatusSkipped    LifecycleEventStatus = "Skipped"
	LifecycleEventStatusUnknown    LifecycleEventStatus = "Unknown"
)

// Values returns all known values for LifecycleEventStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LifecycleEventStatus) Values() []LifecycleEventStatus {
	return []LifecycleEventStatus{
		"Pending",
		"InProgress",
		"Succeeded",
		"Failed",
		"Skipped",
		"Unknown",
	}
}

type ListStateFilterAction string

// Enum values for ListStateFilterAction
const (
	ListStateFilterActionInclude ListStateFilterAction = "include"
	ListStateFilterActionExclude ListStateFilterAction = "exclude"
	ListStateFilterActionIgnore  ListStateFilterAction = "ignore"
)

// Values returns all known values for ListStateFilterAction. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ListStateFilterAction) Values() []ListStateFilterAction {
	return []ListStateFilterAction{
		"include",
		"exclude",
		"ignore",
	}
}

type MinimumHealthyHostsPerZoneType string

// Enum values for MinimumHealthyHostsPerZoneType
const (
	MinimumHealthyHostsPerZoneTypeHostCount    MinimumHealthyHostsPerZoneType = "HOST_COUNT"
	MinimumHealthyHostsPerZoneTypeFleetPercent MinimumHealthyHostsPerZoneType = "FLEET_PERCENT"
)

// Values returns all known values for MinimumHealthyHostsPerZoneType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MinimumHealthyHostsPerZoneType) Values() []MinimumHealthyHostsPerZoneType {
	return []MinimumHealthyHostsPerZoneType{
		"HOST_COUNT",
		"FLEET_PERCENT",
	}
}

type MinimumHealthyHostsType string

// Enum values for MinimumHealthyHostsType
const (
	MinimumHealthyHostsTypeHostCount    MinimumHealthyHostsType = "HOST_COUNT"
	MinimumHealthyHostsTypeFleetPercent MinimumHealthyHostsType = "FLEET_PERCENT"
)

// Values returns all known values for MinimumHealthyHostsType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MinimumHealthyHostsType) Values() []MinimumHealthyHostsType {
	return []MinimumHealthyHostsType{
		"HOST_COUNT",
		"FLEET_PERCENT",
	}
}

type OutdatedInstancesStrategy string

// Enum values for OutdatedInstancesStrategy
const (
	OutdatedInstancesStrategyUpdate OutdatedInstancesStrategy = "UPDATE"
	OutdatedInstancesStrategyIgnore OutdatedInstancesStrategy = "IGNORE"
)

// Values returns all known values for OutdatedInstancesStrategy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OutdatedInstancesStrategy) Values() []OutdatedInstancesStrategy {
	return []OutdatedInstancesStrategy{
		"UPDATE",
		"IGNORE",
	}
}

type RegistrationStatus string

// Enum values for RegistrationStatus
const (
	RegistrationStatusRegistered   RegistrationStatus = "Registered"
	RegistrationStatusDeregistered RegistrationStatus = "Deregistered"
)

// Values returns all known values for RegistrationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RegistrationStatus) Values() []RegistrationStatus {
	return []RegistrationStatus{
		"Registered",
		"Deregistered",
	}
}

type RevisionLocationType string

// Enum values for RevisionLocationType
const (
	RevisionLocationTypeS3             RevisionLocationType = "S3"
	RevisionLocationTypeGitHub         RevisionLocationType = "GitHub"
	RevisionLocationTypeString         RevisionLocationType = "String"
	RevisionLocationTypeAppSpecContent RevisionLocationType = "AppSpecContent"
)

// Values returns all known values for RevisionLocationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RevisionLocationType) Values() []RevisionLocationType {
	return []RevisionLocationType{
		"S3",
		"GitHub",
		"String",
		"AppSpecContent",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ascending"
	SortOrderDescending SortOrder = "descending"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ascending",
		"descending",
	}
}

type StopStatus string

// Enum values for StopStatus
const (
	StopStatusPending   StopStatus = "Pending"
	StopStatusSucceeded StopStatus = "Succeeded"
)

// Values returns all known values for StopStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StopStatus) Values() []StopStatus {
	return []StopStatus{
		"Pending",
		"Succeeded",
	}
}

type TagFilterType string

// Enum values for TagFilterType
const (
	TagFilterTypeKeyOnly     TagFilterType = "KEY_ONLY"
	TagFilterTypeValueOnly   TagFilterType = "VALUE_ONLY"
	TagFilterTypeKeyAndValue TagFilterType = "KEY_AND_VALUE"
)

// Values returns all known values for TagFilterType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TagFilterType) Values() []TagFilterType {
	return []TagFilterType{
		"KEY_ONLY",
		"VALUE_ONLY",
		"KEY_AND_VALUE",
	}
}

type TargetFilterName string

// Enum values for TargetFilterName
const (
	TargetFilterNameTargetStatus        TargetFilterName = "TargetStatus"
	TargetFilterNameServerInstanceLabel TargetFilterName = "ServerInstanceLabel"
)

// Values returns all known values for TargetFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetFilterName) Values() []TargetFilterName {
	return []TargetFilterName{
		"TargetStatus",
		"ServerInstanceLabel",
	}
}

type TargetLabel string

// Enum values for TargetLabel
const (
	TargetLabelBlue  TargetLabel = "Blue"
	TargetLabelGreen TargetLabel = "Green"
)

// Values returns all known values for TargetLabel. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetLabel) Values() []TargetLabel {
	return []TargetLabel{
		"Blue",
		"Green",
	}
}

type TargetStatus string

// Enum values for TargetStatus
const (
	TargetStatusPending    TargetStatus = "Pending"
	TargetStatusInProgress TargetStatus = "InProgress"
	TargetStatusSucceeded  TargetStatus = "Succeeded"
	TargetStatusFailed     TargetStatus = "Failed"
	TargetStatusSkipped    TargetStatus = "Skipped"
	TargetStatusUnknown    TargetStatus = "Unknown"
	TargetStatusReady      TargetStatus = "Ready"
)

// Values returns all known values for TargetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetStatus) Values() []TargetStatus {
	return []TargetStatus{
		"Pending",
		"InProgress",
		"Succeeded",
		"Failed",
		"Skipped",
		"Unknown",
		"Ready",
	}
}

type TrafficRoutingType string

// Enum values for TrafficRoutingType
const (
	TrafficRoutingTypeTimeBasedCanary TrafficRoutingType = "TimeBasedCanary"
	TrafficRoutingTypeTimeBasedLinear TrafficRoutingType = "TimeBasedLinear"
	TrafficRoutingTypeAllAtOnce       TrafficRoutingType = "AllAtOnce"
)

// Values returns all known values for TrafficRoutingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TrafficRoutingType) Values() []TrafficRoutingType {
	return []TrafficRoutingType{
		"TimeBasedCanary",
		"TimeBasedLinear",
		"AllAtOnce",
	}
}

type TriggerEventType string

// Enum values for TriggerEventType
const (
	TriggerEventTypeDeploymentStart    TriggerEventType = "DeploymentStart"
	TriggerEventTypeDeploymentSuccess  TriggerEventType = "DeploymentSuccess"
	TriggerEventTypeDeploymentFailure  TriggerEventType = "DeploymentFailure"
	TriggerEventTypeDeploymentStop     TriggerEventType = "DeploymentStop"
	TriggerEventTypeDeploymentRollback TriggerEventType = "DeploymentRollback"
	TriggerEventTypeDeploymentReady    TriggerEventType = "DeploymentReady"
	TriggerEventTypeInstanceStart      TriggerEventType = "InstanceStart"
	TriggerEventTypeInstanceSuccess    TriggerEventType = "InstanceSuccess"
	TriggerEventTypeInstanceFailure    TriggerEventType = "InstanceFailure"
	TriggerEventTypeInstanceReady      TriggerEventType = "InstanceReady"
)

// Values returns all known values for TriggerEventType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TriggerEventType) Values() []TriggerEventType {
	return []TriggerEventType{
		"DeploymentStart",
		"DeploymentSuccess",
		"DeploymentFailure",
		"DeploymentStop",
		"DeploymentRollback",
		"DeploymentReady",
		"InstanceStart",
		"InstanceSuccess",
		"InstanceFailure",
		"InstanceReady",
	}
}

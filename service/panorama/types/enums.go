// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApplicationInstanceHealthStatus string

// Enum values for ApplicationInstanceHealthStatus
const (
	ApplicationInstanceHealthStatusRunning      ApplicationInstanceHealthStatus = "RUNNING"
	ApplicationInstanceHealthStatusError        ApplicationInstanceHealthStatus = "ERROR"
	ApplicationInstanceHealthStatusNotAvailable ApplicationInstanceHealthStatus = "NOT_AVAILABLE"
)

// Values returns all known values for ApplicationInstanceHealthStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationInstanceHealthStatus) Values() []ApplicationInstanceHealthStatus {
	return []ApplicationInstanceHealthStatus{
		"RUNNING",
		"ERROR",
		"NOT_AVAILABLE",
	}
}

type ApplicationInstanceStatus string

// Enum values for ApplicationInstanceStatus
const (
	ApplicationInstanceStatusDeploymentPending    ApplicationInstanceStatus = "DEPLOYMENT_PENDING"
	ApplicationInstanceStatusDeploymentRequested  ApplicationInstanceStatus = "DEPLOYMENT_REQUESTED"
	ApplicationInstanceStatusDeploymentInProgress ApplicationInstanceStatus = "DEPLOYMENT_IN_PROGRESS"
	ApplicationInstanceStatusDeploymentError      ApplicationInstanceStatus = "DEPLOYMENT_ERROR"
	ApplicationInstanceStatusDeploymentSucceeded  ApplicationInstanceStatus = "DEPLOYMENT_SUCCEEDED"
	ApplicationInstanceStatusRemovalPending       ApplicationInstanceStatus = "REMOVAL_PENDING"
	ApplicationInstanceStatusRemovalRequested     ApplicationInstanceStatus = "REMOVAL_REQUESTED"
	ApplicationInstanceStatusRemovalInProgress    ApplicationInstanceStatus = "REMOVAL_IN_PROGRESS"
	ApplicationInstanceStatusRemovalFailed        ApplicationInstanceStatus = "REMOVAL_FAILED"
	ApplicationInstanceStatusRemovalSucceeded     ApplicationInstanceStatus = "REMOVAL_SUCCEEDED"
	ApplicationInstanceStatusDeploymentFailed     ApplicationInstanceStatus = "DEPLOYMENT_FAILED"
)

// Values returns all known values for ApplicationInstanceStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationInstanceStatus) Values() []ApplicationInstanceStatus {
	return []ApplicationInstanceStatus{
		"DEPLOYMENT_PENDING",
		"DEPLOYMENT_REQUESTED",
		"DEPLOYMENT_IN_PROGRESS",
		"DEPLOYMENT_ERROR",
		"DEPLOYMENT_SUCCEEDED",
		"REMOVAL_PENDING",
		"REMOVAL_REQUESTED",
		"REMOVAL_IN_PROGRESS",
		"REMOVAL_FAILED",
		"REMOVAL_SUCCEEDED",
		"DEPLOYMENT_FAILED",
	}
}

type ConnectionType string

// Enum values for ConnectionType
const (
	ConnectionTypeStaticIp ConnectionType = "STATIC_IP"
	ConnectionTypeDhcp     ConnectionType = "DHCP"
)

// Values returns all known values for ConnectionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionType) Values() []ConnectionType {
	return []ConnectionType{
		"STATIC_IP",
		"DHCP",
	}
}

type DesiredState string

// Enum values for DesiredState
const (
	DesiredStateRunning DesiredState = "RUNNING"
	DesiredStateStopped DesiredState = "STOPPED"
	DesiredStateRemoved DesiredState = "REMOVED"
)

// Values returns all known values for DesiredState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DesiredState) Values() []DesiredState {
	return []DesiredState{
		"RUNNING",
		"STOPPED",
		"REMOVED",
	}
}

type DeviceAggregatedStatus string

// Enum values for DeviceAggregatedStatus
const (
	DeviceAggregatedStatusError                DeviceAggregatedStatus = "ERROR"
	DeviceAggregatedStatusAwaitingProvisioning DeviceAggregatedStatus = "AWAITING_PROVISIONING"
	DeviceAggregatedStatusPending              DeviceAggregatedStatus = "PENDING"
	DeviceAggregatedStatusFailed               DeviceAggregatedStatus = "FAILED"
	DeviceAggregatedStatusDeleting             DeviceAggregatedStatus = "DELETING"
	DeviceAggregatedStatusOnline               DeviceAggregatedStatus = "ONLINE"
	DeviceAggregatedStatusOffline              DeviceAggregatedStatus = "OFFLINE"
	DeviceAggregatedStatusLeaseExpired         DeviceAggregatedStatus = "LEASE_EXPIRED"
	DeviceAggregatedStatusUpdateNeeded         DeviceAggregatedStatus = "UPDATE_NEEDED"
	DeviceAggregatedStatusRebooting            DeviceAggregatedStatus = "REBOOTING"
)

// Values returns all known values for DeviceAggregatedStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeviceAggregatedStatus) Values() []DeviceAggregatedStatus {
	return []DeviceAggregatedStatus{
		"ERROR",
		"AWAITING_PROVISIONING",
		"PENDING",
		"FAILED",
		"DELETING",
		"ONLINE",
		"OFFLINE",
		"LEASE_EXPIRED",
		"UPDATE_NEEDED",
		"REBOOTING",
	}
}

type DeviceBrand string

// Enum values for DeviceBrand
const (
	DeviceBrandAwsPanorama DeviceBrand = "AWS_PANORAMA"
	DeviceBrandLenovo      DeviceBrand = "LENOVO"
)

// Values returns all known values for DeviceBrand. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeviceBrand) Values() []DeviceBrand {
	return []DeviceBrand{
		"AWS_PANORAMA",
		"LENOVO",
	}
}

type DeviceConnectionStatus string

// Enum values for DeviceConnectionStatus
const (
	DeviceConnectionStatusOnline              DeviceConnectionStatus = "ONLINE"
	DeviceConnectionStatusOffline             DeviceConnectionStatus = "OFFLINE"
	DeviceConnectionStatusAwaitingCredentials DeviceConnectionStatus = "AWAITING_CREDENTIALS"
	DeviceConnectionStatusNotAvailable        DeviceConnectionStatus = "NOT_AVAILABLE"
	DeviceConnectionStatusError               DeviceConnectionStatus = "ERROR"
)

// Values returns all known values for DeviceConnectionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeviceConnectionStatus) Values() []DeviceConnectionStatus {
	return []DeviceConnectionStatus{
		"ONLINE",
		"OFFLINE",
		"AWAITING_CREDENTIALS",
		"NOT_AVAILABLE",
		"ERROR",
	}
}

type DeviceReportedStatus string

// Enum values for DeviceReportedStatus
const (
	DeviceReportedStatusStopping          DeviceReportedStatus = "STOPPING"
	DeviceReportedStatusStopped           DeviceReportedStatus = "STOPPED"
	DeviceReportedStatusStopError         DeviceReportedStatus = "STOP_ERROR"
	DeviceReportedStatusRemovalFailed     DeviceReportedStatus = "REMOVAL_FAILED"
	DeviceReportedStatusRemovalInProgress DeviceReportedStatus = "REMOVAL_IN_PROGRESS"
	DeviceReportedStatusStarting          DeviceReportedStatus = "STARTING"
	DeviceReportedStatusRunning           DeviceReportedStatus = "RUNNING"
	DeviceReportedStatusInstallError      DeviceReportedStatus = "INSTALL_ERROR"
	DeviceReportedStatusLaunched          DeviceReportedStatus = "LAUNCHED"
	DeviceReportedStatusLaunchError       DeviceReportedStatus = "LAUNCH_ERROR"
	DeviceReportedStatusInstallInProgress DeviceReportedStatus = "INSTALL_IN_PROGRESS"
)

// Values returns all known values for DeviceReportedStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeviceReportedStatus) Values() []DeviceReportedStatus {
	return []DeviceReportedStatus{
		"STOPPING",
		"STOPPED",
		"STOP_ERROR",
		"REMOVAL_FAILED",
		"REMOVAL_IN_PROGRESS",
		"STARTING",
		"RUNNING",
		"INSTALL_ERROR",
		"LAUNCHED",
		"LAUNCH_ERROR",
		"INSTALL_IN_PROGRESS",
	}
}

type DeviceStatus string

// Enum values for DeviceStatus
const (
	DeviceStatusAwaitingProvisioning DeviceStatus = "AWAITING_PROVISIONING"
	DeviceStatusPending              DeviceStatus = "PENDING"
	DeviceStatusSucceeded            DeviceStatus = "SUCCEEDED"
	DeviceStatusFailed               DeviceStatus = "FAILED"
	DeviceStatusError                DeviceStatus = "ERROR"
	DeviceStatusDeleting             DeviceStatus = "DELETING"
)

// Values returns all known values for DeviceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeviceStatus) Values() []DeviceStatus {
	return []DeviceStatus{
		"AWAITING_PROVISIONING",
		"PENDING",
		"SUCCEEDED",
		"FAILED",
		"ERROR",
		"DELETING",
	}
}

type DeviceType string

// Enum values for DeviceType
const (
	DeviceTypePanoramaApplianceDeveloperKit DeviceType = "PANORAMA_APPLIANCE_DEVELOPER_KIT"
	DeviceTypePanoramaAppliance             DeviceType = "PANORAMA_APPLIANCE"
)

// Values returns all known values for DeviceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeviceType) Values() []DeviceType {
	return []DeviceType{
		"PANORAMA_APPLIANCE_DEVELOPER_KIT",
		"PANORAMA_APPLIANCE",
	}
}

type JobResourceType string

// Enum values for JobResourceType
const (
	JobResourceTypePackage JobResourceType = "PACKAGE"
)

// Values returns all known values for JobResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobResourceType) Values() []JobResourceType {
	return []JobResourceType{
		"PACKAGE",
	}
}

type JobType string

// Enum values for JobType
const (
	JobTypeOta    JobType = "OTA"
	JobTypeReboot JobType = "REBOOT"
)

// Values returns all known values for JobType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobType) Values() []JobType {
	return []JobType{
		"OTA",
		"REBOOT",
	}
}

type ListDevicesSortBy string

// Enum values for ListDevicesSortBy
const (
	ListDevicesSortByDeviceId               ListDevicesSortBy = "DEVICE_ID"
	ListDevicesSortByCreatedTime            ListDevicesSortBy = "CREATED_TIME"
	ListDevicesSortByName                   ListDevicesSortBy = "NAME"
	ListDevicesSortByDeviceAggregatedStatus ListDevicesSortBy = "DEVICE_AGGREGATED_STATUS"
)

// Values returns all known values for ListDevicesSortBy. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ListDevicesSortBy) Values() []ListDevicesSortBy {
	return []ListDevicesSortBy{
		"DEVICE_ID",
		"CREATED_TIME",
		"NAME",
		"DEVICE_AGGREGATED_STATUS",
	}
}

type NetworkConnectionStatus string

// Enum values for NetworkConnectionStatus
const (
	NetworkConnectionStatusConnected    NetworkConnectionStatus = "CONNECTED"
	NetworkConnectionStatusNotConnected NetworkConnectionStatus = "NOT_CONNECTED"
	NetworkConnectionStatusConnecting   NetworkConnectionStatus = "CONNECTING"
)

// Values returns all known values for NetworkConnectionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NetworkConnectionStatus) Values() []NetworkConnectionStatus {
	return []NetworkConnectionStatus{
		"CONNECTED",
		"NOT_CONNECTED",
		"CONNECTING",
	}
}

type NodeCategory string

// Enum values for NodeCategory
const (
	NodeCategoryBusinessLogic NodeCategory = "BUSINESS_LOGIC"
	NodeCategoryMlModel       NodeCategory = "ML_MODEL"
	NodeCategoryMediaSource   NodeCategory = "MEDIA_SOURCE"
	NodeCategoryMediaSink     NodeCategory = "MEDIA_SINK"
)

// Values returns all known values for NodeCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NodeCategory) Values() []NodeCategory {
	return []NodeCategory{
		"BUSINESS_LOGIC",
		"ML_MODEL",
		"MEDIA_SOURCE",
		"MEDIA_SINK",
	}
}

type NodeFromTemplateJobStatus string

// Enum values for NodeFromTemplateJobStatus
const (
	NodeFromTemplateJobStatusPending   NodeFromTemplateJobStatus = "PENDING"
	NodeFromTemplateJobStatusSucceeded NodeFromTemplateJobStatus = "SUCCEEDED"
	NodeFromTemplateJobStatusFailed    NodeFromTemplateJobStatus = "FAILED"
)

// Values returns all known values for NodeFromTemplateJobStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NodeFromTemplateJobStatus) Values() []NodeFromTemplateJobStatus {
	return []NodeFromTemplateJobStatus{
		"PENDING",
		"SUCCEEDED",
		"FAILED",
	}
}

type NodeInstanceStatus string

// Enum values for NodeInstanceStatus
const (
	NodeInstanceStatusRunning      NodeInstanceStatus = "RUNNING"
	NodeInstanceStatusError        NodeInstanceStatus = "ERROR"
	NodeInstanceStatusNotAvailable NodeInstanceStatus = "NOT_AVAILABLE"
	NodeInstanceStatusPaused       NodeInstanceStatus = "PAUSED"
)

// Values returns all known values for NodeInstanceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NodeInstanceStatus) Values() []NodeInstanceStatus {
	return []NodeInstanceStatus{
		"RUNNING",
		"ERROR",
		"NOT_AVAILABLE",
		"PAUSED",
	}
}

type NodeSignalValue string

// Enum values for NodeSignalValue
const (
	NodeSignalValuePause  NodeSignalValue = "PAUSE"
	NodeSignalValueResume NodeSignalValue = "RESUME"
)

// Values returns all known values for NodeSignalValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NodeSignalValue) Values() []NodeSignalValue {
	return []NodeSignalValue{
		"PAUSE",
		"RESUME",
	}
}

type PackageImportJobStatus string

// Enum values for PackageImportJobStatus
const (
	PackageImportJobStatusPending   PackageImportJobStatus = "PENDING"
	PackageImportJobStatusSucceeded PackageImportJobStatus = "SUCCEEDED"
	PackageImportJobStatusFailed    PackageImportJobStatus = "FAILED"
)

// Values returns all known values for PackageImportJobStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PackageImportJobStatus) Values() []PackageImportJobStatus {
	return []PackageImportJobStatus{
		"PENDING",
		"SUCCEEDED",
		"FAILED",
	}
}

type PackageImportJobType string

// Enum values for PackageImportJobType
const (
	PackageImportJobTypeNodePackageVersion            PackageImportJobType = "NODE_PACKAGE_VERSION"
	PackageImportJobTypeMarketplaceNodePackageVersion PackageImportJobType = "MARKETPLACE_NODE_PACKAGE_VERSION"
)

// Values returns all known values for PackageImportJobType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PackageImportJobType) Values() []PackageImportJobType {
	return []PackageImportJobType{
		"NODE_PACKAGE_VERSION",
		"MARKETPLACE_NODE_PACKAGE_VERSION",
	}
}

type PackageVersionStatus string

// Enum values for PackageVersionStatus
const (
	PackageVersionStatusRegisterPending   PackageVersionStatus = "REGISTER_PENDING"
	PackageVersionStatusRegisterCompleted PackageVersionStatus = "REGISTER_COMPLETED"
	PackageVersionStatusFailed            PackageVersionStatus = "FAILED"
	PackageVersionStatusDeleting          PackageVersionStatus = "DELETING"
)

// Values returns all known values for PackageVersionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PackageVersionStatus) Values() []PackageVersionStatus {
	return []PackageVersionStatus{
		"REGISTER_PENDING",
		"REGISTER_COMPLETED",
		"FAILED",
		"DELETING",
	}
}

type PortType string

// Enum values for PortType
const (
	PortTypeBoolean PortType = "BOOLEAN"
	PortTypeString  PortType = "STRING"
	PortTypeInt32   PortType = "INT32"
	PortTypeFloat32 PortType = "FLOAT32"
	PortTypeMedia   PortType = "MEDIA"
)

// Values returns all known values for PortType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PortType) Values() []PortType {
	return []PortType{
		"BOOLEAN",
		"STRING",
		"INT32",
		"FLOAT32",
		"MEDIA",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}

type StatusFilter string

// Enum values for StatusFilter
const (
	StatusFilterDeploymentSucceeded  StatusFilter = "DEPLOYMENT_SUCCEEDED"
	StatusFilterDeploymentError      StatusFilter = "DEPLOYMENT_ERROR"
	StatusFilterRemovalSucceeded     StatusFilter = "REMOVAL_SUCCEEDED"
	StatusFilterRemovalFailed        StatusFilter = "REMOVAL_FAILED"
	StatusFilterProcessingDeployment StatusFilter = "PROCESSING_DEPLOYMENT"
	StatusFilterProcessingRemoval    StatusFilter = "PROCESSING_REMOVAL"
	StatusFilterDeploymentFailed     StatusFilter = "DEPLOYMENT_FAILED"
)

// Values returns all known values for StatusFilter. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StatusFilter) Values() []StatusFilter {
	return []StatusFilter{
		"DEPLOYMENT_SUCCEEDED",
		"DEPLOYMENT_ERROR",
		"REMOVAL_SUCCEEDED",
		"REMOVAL_FAILED",
		"PROCESSING_DEPLOYMENT",
		"PROCESSING_REMOVAL",
		"DEPLOYMENT_FAILED",
	}
}

type TemplateType string

// Enum values for TemplateType
const (
	TemplateTypeRtspCameraStream TemplateType = "RTSP_CAMERA_STREAM"
)

// Values returns all known values for TemplateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TemplateType) Values() []TemplateType {
	return []TemplateType{
		"RTSP_CAMERA_STREAM",
	}
}

type UpdateProgress string

// Enum values for UpdateProgress
const (
	UpdateProgressPending     UpdateProgress = "PENDING"
	UpdateProgressInProgress  UpdateProgress = "IN_PROGRESS"
	UpdateProgressVerifying   UpdateProgress = "VERIFYING"
	UpdateProgressRebooting   UpdateProgress = "REBOOTING"
	UpdateProgressDownloading UpdateProgress = "DOWNLOADING"
	UpdateProgressCompleted   UpdateProgress = "COMPLETED"
	UpdateProgressFailed      UpdateProgress = "FAILED"
)

// Values returns all known values for UpdateProgress. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UpdateProgress) Values() []UpdateProgress {
	return []UpdateProgress{
		"PENDING",
		"IN_PROGRESS",
		"VERIFYING",
		"REBOOTING",
		"DOWNLOADING",
		"COMPLETED",
		"FAILED",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
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
	}
}

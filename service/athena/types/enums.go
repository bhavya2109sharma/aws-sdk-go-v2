// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AuthenticationType string

// Enum values for AuthenticationType
const (
	AuthenticationTypeDirectoryIdentity AuthenticationType = "DIRECTORY_IDENTITY"
)

// Values returns all known values for AuthenticationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthenticationType) Values() []AuthenticationType {
	return []AuthenticationType{
		"DIRECTORY_IDENTITY",
	}
}

type CalculationExecutionState string

// Enum values for CalculationExecutionState
const (
	CalculationExecutionStateCreating  CalculationExecutionState = "CREATING"
	CalculationExecutionStateCreated   CalculationExecutionState = "CREATED"
	CalculationExecutionStateQueued    CalculationExecutionState = "QUEUED"
	CalculationExecutionStateRunning   CalculationExecutionState = "RUNNING"
	CalculationExecutionStateCanceling CalculationExecutionState = "CANCELING"
	CalculationExecutionStateCanceled  CalculationExecutionState = "CANCELED"
	CalculationExecutionStateCompleted CalculationExecutionState = "COMPLETED"
	CalculationExecutionStateFailed    CalculationExecutionState = "FAILED"
)

// Values returns all known values for CalculationExecutionState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CalculationExecutionState) Values() []CalculationExecutionState {
	return []CalculationExecutionState{
		"CREATING",
		"CREATED",
		"QUEUED",
		"RUNNING",
		"CANCELING",
		"CANCELED",
		"COMPLETED",
		"FAILED",
	}
}

type CapacityAllocationStatus string

// Enum values for CapacityAllocationStatus
const (
	CapacityAllocationStatusPending   CapacityAllocationStatus = "PENDING"
	CapacityAllocationStatusSucceeded CapacityAllocationStatus = "SUCCEEDED"
	CapacityAllocationStatusFailed    CapacityAllocationStatus = "FAILED"
)

// Values returns all known values for CapacityAllocationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CapacityAllocationStatus) Values() []CapacityAllocationStatus {
	return []CapacityAllocationStatus{
		"PENDING",
		"SUCCEEDED",
		"FAILED",
	}
}

type CapacityReservationStatus string

// Enum values for CapacityReservationStatus
const (
	CapacityReservationStatusPending       CapacityReservationStatus = "PENDING"
	CapacityReservationStatusActive        CapacityReservationStatus = "ACTIVE"
	CapacityReservationStatusCancelling    CapacityReservationStatus = "CANCELLING"
	CapacityReservationStatusCancelled     CapacityReservationStatus = "CANCELLED"
	CapacityReservationStatusFailed        CapacityReservationStatus = "FAILED"
	CapacityReservationStatusUpdatePending CapacityReservationStatus = "UPDATE_PENDING"
)

// Values returns all known values for CapacityReservationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CapacityReservationStatus) Values() []CapacityReservationStatus {
	return []CapacityReservationStatus{
		"PENDING",
		"ACTIVE",
		"CANCELLING",
		"CANCELLED",
		"FAILED",
		"UPDATE_PENDING",
	}
}

type ColumnNullable string

// Enum values for ColumnNullable
const (
	ColumnNullableNotNull  ColumnNullable = "NOT_NULL"
	ColumnNullableNullable ColumnNullable = "NULLABLE"
	ColumnNullableUnknown  ColumnNullable = "UNKNOWN"
)

// Values returns all known values for ColumnNullable. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ColumnNullable) Values() []ColumnNullable {
	return []ColumnNullable{
		"NOT_NULL",
		"NULLABLE",
		"UNKNOWN",
	}
}

type DataCatalogType string

// Enum values for DataCatalogType
const (
	DataCatalogTypeLambda DataCatalogType = "LAMBDA"
	DataCatalogTypeGlue   DataCatalogType = "GLUE"
	DataCatalogTypeHive   DataCatalogType = "HIVE"
)

// Values returns all known values for DataCatalogType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataCatalogType) Values() []DataCatalogType {
	return []DataCatalogType{
		"LAMBDA",
		"GLUE",
		"HIVE",
	}
}

type EncryptionOption string

// Enum values for EncryptionOption
const (
	EncryptionOptionSseS3  EncryptionOption = "SSE_S3"
	EncryptionOptionSseKms EncryptionOption = "SSE_KMS"
	EncryptionOptionCseKms EncryptionOption = "CSE_KMS"
)

// Values returns all known values for EncryptionOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionOption) Values() []EncryptionOption {
	return []EncryptionOption{
		"SSE_S3",
		"SSE_KMS",
		"CSE_KMS",
	}
}

type ExecutorState string

// Enum values for ExecutorState
const (
	ExecutorStateCreating    ExecutorState = "CREATING"
	ExecutorStateCreated     ExecutorState = "CREATED"
	ExecutorStateRegistered  ExecutorState = "REGISTERED"
	ExecutorStateTerminating ExecutorState = "TERMINATING"
	ExecutorStateTerminated  ExecutorState = "TERMINATED"
	ExecutorStateFailed      ExecutorState = "FAILED"
)

// Values returns all known values for ExecutorState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExecutorState) Values() []ExecutorState {
	return []ExecutorState{
		"CREATING",
		"CREATED",
		"REGISTERED",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

type ExecutorType string

// Enum values for ExecutorType
const (
	ExecutorTypeCoordinator ExecutorType = "COORDINATOR"
	ExecutorTypeGateway     ExecutorType = "GATEWAY"
	ExecutorTypeWorker      ExecutorType = "WORKER"
)

// Values returns all known values for ExecutorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExecutorType) Values() []ExecutorType {
	return []ExecutorType{
		"COORDINATOR",
		"GATEWAY",
		"WORKER",
	}
}

type NotebookType string

// Enum values for NotebookType
const (
	NotebookTypeIpynb NotebookType = "IPYNB"
)

// Values returns all known values for NotebookType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NotebookType) Values() []NotebookType {
	return []NotebookType{
		"IPYNB",
	}
}

type QueryExecutionState string

// Enum values for QueryExecutionState
const (
	QueryExecutionStateQueued    QueryExecutionState = "QUEUED"
	QueryExecutionStateRunning   QueryExecutionState = "RUNNING"
	QueryExecutionStateSucceeded QueryExecutionState = "SUCCEEDED"
	QueryExecutionStateFailed    QueryExecutionState = "FAILED"
	QueryExecutionStateCancelled QueryExecutionState = "CANCELLED"
)

// Values returns all known values for QueryExecutionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryExecutionState) Values() []QueryExecutionState {
	return []QueryExecutionState{
		"QUEUED",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
		"CANCELLED",
	}
}

type S3AclOption string

// Enum values for S3AclOption
const (
	S3AclOptionBucketOwnerFullControl S3AclOption = "BUCKET_OWNER_FULL_CONTROL"
)

// Values returns all known values for S3AclOption. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3AclOption) Values() []S3AclOption {
	return []S3AclOption{
		"BUCKET_OWNER_FULL_CONTROL",
	}
}

type SessionState string

// Enum values for SessionState
const (
	SessionStateCreating    SessionState = "CREATING"
	SessionStateCreated     SessionState = "CREATED"
	SessionStateIdle        SessionState = "IDLE"
	SessionStateBusy        SessionState = "BUSY"
	SessionStateTerminating SessionState = "TERMINATING"
	SessionStateTerminated  SessionState = "TERMINATED"
	SessionStateDegraded    SessionState = "DEGRADED"
	SessionStateFailed      SessionState = "FAILED"
)

// Values returns all known values for SessionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SessionState) Values() []SessionState {
	return []SessionState{
		"CREATING",
		"CREATED",
		"IDLE",
		"BUSY",
		"TERMINATING",
		"TERMINATED",
		"DEGRADED",
		"FAILED",
	}
}

type StatementType string

// Enum values for StatementType
const (
	StatementTypeDdl     StatementType = "DDL"
	StatementTypeDml     StatementType = "DML"
	StatementTypeUtility StatementType = "UTILITY"
)

// Values returns all known values for StatementType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StatementType) Values() []StatementType {
	return []StatementType{
		"DDL",
		"DML",
		"UTILITY",
	}
}

type ThrottleReason string

// Enum values for ThrottleReason
const (
	ThrottleReasonConcurrentQueryLimitExceeded ThrottleReason = "CONCURRENT_QUERY_LIMIT_EXCEEDED"
)

// Values returns all known values for ThrottleReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ThrottleReason) Values() []ThrottleReason {
	return []ThrottleReason{
		"CONCURRENT_QUERY_LIMIT_EXCEEDED",
	}
}

type WorkGroupState string

// Enum values for WorkGroupState
const (
	WorkGroupStateEnabled  WorkGroupState = "ENABLED"
	WorkGroupStateDisabled WorkGroupState = "DISABLED"
)

// Values returns all known values for WorkGroupState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkGroupState) Values() []WorkGroupState {
	return []WorkGroupState{
		"ENABLED",
		"DISABLED",
	}
}

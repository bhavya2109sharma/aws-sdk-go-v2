// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AuthPolicyState string

// Enum values for AuthPolicyState
const (
	AuthPolicyStateActive   AuthPolicyState = "Active"
	AuthPolicyStateInactive AuthPolicyState = "Inactive"
)

// Values returns all known values for AuthPolicyState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthPolicyState) Values() []AuthPolicyState {
	return []AuthPolicyState{
		"Active",
		"Inactive",
	}
}

type AuthType string

// Enum values for AuthType
const (
	AuthTypeNone   AuthType = "NONE"
	AuthTypeAwsIam AuthType = "AWS_IAM"
)

// Values returns all known values for AuthType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthType) Values() []AuthType {
	return []AuthType{
		"NONE",
		"AWS_IAM",
	}
}

type HealthCheckProtocolVersion string

// Enum values for HealthCheckProtocolVersion
const (
	// Indicates use of HTTP/1.1 to send requests to target
	HealthCheckProtocolVersionHttp1 HealthCheckProtocolVersion = "HTTP1"
	// Indicates use of HTTP/2 to send requests to target
	HealthCheckProtocolVersionHttp2 HealthCheckProtocolVersion = "HTTP2"
)

// Values returns all known values for HealthCheckProtocolVersion. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HealthCheckProtocolVersion) Values() []HealthCheckProtocolVersion {
	return []HealthCheckProtocolVersion{
		"HTTP1",
		"HTTP2",
	}
}

type IpAddressType string

// Enum values for IpAddressType
const (
	// Indicates IPv4 address type
	IpAddressTypeIpv4 IpAddressType = "IPV4"
	// Indicates IPv6 address type
	IpAddressTypeIpv6 IpAddressType = "IPV6"
)

// Values returns all known values for IpAddressType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IpAddressType) Values() []IpAddressType {
	return []IpAddressType{
		"IPV4",
		"IPV6",
	}
}

type LambdaEventStructureVersion string

// Enum values for LambdaEventStructureVersion
const (
	// This is the default lambda event structure version
	LambdaEventStructureVersionV1 LambdaEventStructureVersion = "V1"
	// Indicates use of lambda event structure version 2
	LambdaEventStructureVersionV2 LambdaEventStructureVersion = "V2"
)

// Values returns all known values for LambdaEventStructureVersion. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LambdaEventStructureVersion) Values() []LambdaEventStructureVersion {
	return []LambdaEventStructureVersion{
		"V1",
		"V2",
	}
}

type ListenerProtocol string

// Enum values for ListenerProtocol
const (
	// Indicates HTTP protocol
	ListenerProtocolHttp ListenerProtocol = "HTTP"
	// Indicates HTTPS protocol
	ListenerProtocolHttps ListenerProtocol = "HTTPS"
	// Indicates TLS_PASSTHROUGH protocol
	ListenerProtocolTlsPassthrough ListenerProtocol = "TLS_PASSTHROUGH"
)

// Values returns all known values for ListenerProtocol. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ListenerProtocol) Values() []ListenerProtocol {
	return []ListenerProtocol{
		"HTTP",
		"HTTPS",
		"TLS_PASSTHROUGH",
	}
}

type ServiceNetworkServiceAssociationStatus string

// Enum values for ServiceNetworkServiceAssociationStatus
const (
	// ServiceNetwork and Service association creation in progress
	ServiceNetworkServiceAssociationStatusCreateInProgress ServiceNetworkServiceAssociationStatus = "CREATE_IN_PROGRESS"
	// ServiceNetwork and Service association is active
	ServiceNetworkServiceAssociationStatusActive ServiceNetworkServiceAssociationStatus = "ACTIVE"
	// ServiceNetwork and Service association deletion in progress
	ServiceNetworkServiceAssociationStatusDeleteInProgress ServiceNetworkServiceAssociationStatus = "DELETE_IN_PROGRESS"
	// ServiceNetwork and Service association creation failed.
	ServiceNetworkServiceAssociationStatusCreateFailed ServiceNetworkServiceAssociationStatus = "CREATE_FAILED"
	// ServiceNetwork and Service association deletion failed
	ServiceNetworkServiceAssociationStatusDeleteFailed ServiceNetworkServiceAssociationStatus = "DELETE_FAILED"
)

// Values returns all known values for ServiceNetworkServiceAssociationStatus.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServiceNetworkServiceAssociationStatus) Values() []ServiceNetworkServiceAssociationStatus {
	return []ServiceNetworkServiceAssociationStatus{
		"CREATE_IN_PROGRESS",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"CREATE_FAILED",
		"DELETE_FAILED",
	}
}

type ServiceNetworkVpcAssociationStatus string

// Enum values for ServiceNetworkVpcAssociationStatus
const (
	// ServiceNetwork and Vpc association creation in progress
	ServiceNetworkVpcAssociationStatusCreateInProgress ServiceNetworkVpcAssociationStatus = "CREATE_IN_PROGRESS"
	// ServiceNetwork and Vpc association is active
	ServiceNetworkVpcAssociationStatusActive ServiceNetworkVpcAssociationStatus = "ACTIVE"
	// ServiceNetwork and Vpc association update in progress
	ServiceNetworkVpcAssociationStatusUpdateInProgress ServiceNetworkVpcAssociationStatus = "UPDATE_IN_PROGRESS"
	// ServiceNetwork and Vpc association deletion in progress
	ServiceNetworkVpcAssociationStatusDeleteInProgress ServiceNetworkVpcAssociationStatus = "DELETE_IN_PROGRESS"
	// ServiceNetwork and Vpc association creation failed.
	ServiceNetworkVpcAssociationStatusCreateFailed ServiceNetworkVpcAssociationStatus = "CREATE_FAILED"
	// ServiceNetwork and Vpc association deletion failed
	ServiceNetworkVpcAssociationStatusDeleteFailed ServiceNetworkVpcAssociationStatus = "DELETE_FAILED"
	// ServiceNetwork and Vpc association update failed
	ServiceNetworkVpcAssociationStatusUpdateFailed ServiceNetworkVpcAssociationStatus = "UPDATE_FAILED"
)

// Values returns all known values for ServiceNetworkVpcAssociationStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServiceNetworkVpcAssociationStatus) Values() []ServiceNetworkVpcAssociationStatus {
	return []ServiceNetworkVpcAssociationStatus{
		"CREATE_IN_PROGRESS",
		"ACTIVE",
		"UPDATE_IN_PROGRESS",
		"DELETE_IN_PROGRESS",
		"CREATE_FAILED",
		"DELETE_FAILED",
		"UPDATE_FAILED",
	}
}

type ServiceStatus string

// Enum values for ServiceStatus
const (
	// Service is active.
	ServiceStatusActive ServiceStatus = "ACTIVE"
	// Service creation in progress.
	ServiceStatusCreateInProgress ServiceStatus = "CREATE_IN_PROGRESS"
	// Service deletion in progress
	ServiceStatusDeleteInProgress ServiceStatus = "DELETE_IN_PROGRESS"
	// Service creation failed
	ServiceStatusCreateFailed ServiceStatus = "CREATE_FAILED"
	// Service deletion failed.
	ServiceStatusDeleteFailed ServiceStatus = "DELETE_FAILED"
)

// Values returns all known values for ServiceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServiceStatus) Values() []ServiceStatus {
	return []ServiceStatus{
		"ACTIVE",
		"CREATE_IN_PROGRESS",
		"DELETE_IN_PROGRESS",
		"CREATE_FAILED",
		"DELETE_FAILED",
	}
}

type TargetGroupProtocol string

// Enum values for TargetGroupProtocol
const (
	// Indicates HTTP protocol
	TargetGroupProtocolHttp TargetGroupProtocol = "HTTP"
	// Indicates HTTPS protocol
	TargetGroupProtocolHttps TargetGroupProtocol = "HTTPS"
	// Indicates TCP protocol
	TargetGroupProtocolTcp TargetGroupProtocol = "TCP"
)

// Values returns all known values for TargetGroupProtocol. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetGroupProtocol) Values() []TargetGroupProtocol {
	return []TargetGroupProtocol{
		"HTTP",
		"HTTPS",
		"TCP",
	}
}

type TargetGroupProtocolVersion string

// Enum values for TargetGroupProtocolVersion
const (
	// Indicates use of HTTP/1.1 to send requests to target
	TargetGroupProtocolVersionHttp1 TargetGroupProtocolVersion = "HTTP1"
	// Indicates use of HTTP/2 to send requests to target
	TargetGroupProtocolVersionHttp2 TargetGroupProtocolVersion = "HTTP2"
	// Indicates use of gRPC to send requests to target
	TargetGroupProtocolVersionGrpc TargetGroupProtocolVersion = "GRPC"
)

// Values returns all known values for TargetGroupProtocolVersion. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetGroupProtocolVersion) Values() []TargetGroupProtocolVersion {
	return []TargetGroupProtocolVersion{
		"HTTP1",
		"HTTP2",
		"GRPC",
	}
}

type TargetGroupStatus string

// Enum values for TargetGroupStatus
const (
	// TargetGroup creation in progress
	TargetGroupStatusCreateInProgress TargetGroupStatus = "CREATE_IN_PROGRESS"
	// TargetGroup is active
	TargetGroupStatusActive TargetGroupStatus = "ACTIVE"
	// TargetGroup deletion in progress
	TargetGroupStatusDeleteInProgress TargetGroupStatus = "DELETE_IN_PROGRESS"
	// TargetGroup creation failed.
	TargetGroupStatusCreateFailed TargetGroupStatus = "CREATE_FAILED"
	// TargetGroup deletion failed
	TargetGroupStatusDeleteFailed TargetGroupStatus = "DELETE_FAILED"
)

// Values returns all known values for TargetGroupStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetGroupStatus) Values() []TargetGroupStatus {
	return []TargetGroupStatus{
		"CREATE_IN_PROGRESS",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"CREATE_FAILED",
		"DELETE_FAILED",
	}
}

type TargetGroupType string

// Enum values for TargetGroupType
const (
	// Indicates targets in this target group are IP
	TargetGroupTypeIp TargetGroupType = "IP"
	// Indicates targets in this target group are Lambda
	TargetGroupTypeLambda TargetGroupType = "LAMBDA"
	// Indicates targets in this target group are EC2 instances
	TargetGroupTypeInstance TargetGroupType = "INSTANCE"
	// Indicates target in this target group is an ALB
	TargetGroupTypeAlb TargetGroupType = "ALB"
)

// Values returns all known values for TargetGroupType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetGroupType) Values() []TargetGroupType {
	return []TargetGroupType{
		"IP",
		"LAMBDA",
		"INSTANCE",
		"ALB",
	}
}

type TargetStatus string

// Enum values for TargetStatus
const (
	// The target is deregistering and connection draining is in process.
	TargetStatusDraining TargetStatus = "DRAINING"
	// Health checks are disabled.
	TargetStatusUnavailable TargetStatus = "UNAVAILABLE"
	// The target is healthy.
	TargetStatusHealthy TargetStatus = "HEALTHY"
	// The target failed the health check.
	TargetStatusUnhealthy TargetStatus = "UNHEALTHY"
	// The initial health check is in progress.
	TargetStatusInitial TargetStatus = "INITIAL"
	// The target group is not used in a listener rule.
	TargetStatusUnused TargetStatus = "UNUSED"
)

// Values returns all known values for TargetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetStatus) Values() []TargetStatus {
	return []TargetStatus{
		"DRAINING",
		"UNAVAILABLE",
		"HEALTHY",
		"UNHEALTHY",
		"INITIAL",
		"UNUSED",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"other",
	}
}

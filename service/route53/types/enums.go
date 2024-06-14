// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountLimitType string

// Enum values for AccountLimitType
const (
	AccountLimitTypeMaxHealthChecksByOwner           AccountLimitType = "MAX_HEALTH_CHECKS_BY_OWNER"
	AccountLimitTypeMaxHostedZonesByOwner            AccountLimitType = "MAX_HOSTED_ZONES_BY_OWNER"
	AccountLimitTypeMaxTrafficPolicyInstancesByOwner AccountLimitType = "MAX_TRAFFIC_POLICY_INSTANCES_BY_OWNER"
	AccountLimitTypeMaxReusableDelegationSetsByOwner AccountLimitType = "MAX_REUSABLE_DELEGATION_SETS_BY_OWNER"
	AccountLimitTypeMaxTrafficPoliciesByOwner        AccountLimitType = "MAX_TRAFFIC_POLICIES_BY_OWNER"
)

// Values returns all known values for AccountLimitType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccountLimitType) Values() []AccountLimitType {
	return []AccountLimitType{
		"MAX_HEALTH_CHECKS_BY_OWNER",
		"MAX_HOSTED_ZONES_BY_OWNER",
		"MAX_TRAFFIC_POLICY_INSTANCES_BY_OWNER",
		"MAX_REUSABLE_DELEGATION_SETS_BY_OWNER",
		"MAX_TRAFFIC_POLICIES_BY_OWNER",
	}
}

type ChangeAction string

// Enum values for ChangeAction
const (
	ChangeActionCreate ChangeAction = "CREATE"
	ChangeActionDelete ChangeAction = "DELETE"
	ChangeActionUpsert ChangeAction = "UPSERT"
)

// Values returns all known values for ChangeAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChangeAction) Values() []ChangeAction {
	return []ChangeAction{
		"CREATE",
		"DELETE",
		"UPSERT",
	}
}

type ChangeStatus string

// Enum values for ChangeStatus
const (
	ChangeStatusPending ChangeStatus = "PENDING"
	ChangeStatusInsync  ChangeStatus = "INSYNC"
)

// Values returns all known values for ChangeStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChangeStatus) Values() []ChangeStatus {
	return []ChangeStatus{
		"PENDING",
		"INSYNC",
	}
}

type CidrCollectionChangeAction string

// Enum values for CidrCollectionChangeAction
const (
	CidrCollectionChangeActionPut            CidrCollectionChangeAction = "PUT"
	CidrCollectionChangeActionDeleteIfExists CidrCollectionChangeAction = "DELETE_IF_EXISTS"
)

// Values returns all known values for CidrCollectionChangeAction. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CidrCollectionChangeAction) Values() []CidrCollectionChangeAction {
	return []CidrCollectionChangeAction{
		"PUT",
		"DELETE_IF_EXISTS",
	}
}

type CloudWatchRegion string

// Enum values for CloudWatchRegion
const (
	CloudWatchRegionUsEast1      CloudWatchRegion = "us-east-1"
	CloudWatchRegionUsEast2      CloudWatchRegion = "us-east-2"
	CloudWatchRegionUsWest1      CloudWatchRegion = "us-west-1"
	CloudWatchRegionUsWest2      CloudWatchRegion = "us-west-2"
	CloudWatchRegionCaCentral1   CloudWatchRegion = "ca-central-1"
	CloudWatchRegionEuCentral1   CloudWatchRegion = "eu-central-1"
	CloudWatchRegionEuCentral2   CloudWatchRegion = "eu-central-2"
	CloudWatchRegionEuWest1      CloudWatchRegion = "eu-west-1"
	CloudWatchRegionEuWest2      CloudWatchRegion = "eu-west-2"
	CloudWatchRegionEuWest3      CloudWatchRegion = "eu-west-3"
	CloudWatchRegionApEast1      CloudWatchRegion = "ap-east-1"
	CloudWatchRegionMeSouth1     CloudWatchRegion = "me-south-1"
	CloudWatchRegionMeCentral1   CloudWatchRegion = "me-central-1"
	CloudWatchRegionApSouth1     CloudWatchRegion = "ap-south-1"
	CloudWatchRegionApSouth2     CloudWatchRegion = "ap-south-2"
	CloudWatchRegionApSoutheast1 CloudWatchRegion = "ap-southeast-1"
	CloudWatchRegionApSoutheast2 CloudWatchRegion = "ap-southeast-2"
	CloudWatchRegionApSoutheast3 CloudWatchRegion = "ap-southeast-3"
	CloudWatchRegionApNortheast1 CloudWatchRegion = "ap-northeast-1"
	CloudWatchRegionApNortheast2 CloudWatchRegion = "ap-northeast-2"
	CloudWatchRegionApNortheast3 CloudWatchRegion = "ap-northeast-3"
	CloudWatchRegionEuNorth1     CloudWatchRegion = "eu-north-1"
	CloudWatchRegionSaEast1      CloudWatchRegion = "sa-east-1"
	CloudWatchRegionCnNorthwest1 CloudWatchRegion = "cn-northwest-1"
	CloudWatchRegionCnNorth1     CloudWatchRegion = "cn-north-1"
	CloudWatchRegionAfSouth1     CloudWatchRegion = "af-south-1"
	CloudWatchRegionEuSouth1     CloudWatchRegion = "eu-south-1"
	CloudWatchRegionEuSouth2     CloudWatchRegion = "eu-south-2"
	CloudWatchRegionUsGovWest1   CloudWatchRegion = "us-gov-west-1"
	CloudWatchRegionUsGovEast1   CloudWatchRegion = "us-gov-east-1"
	CloudWatchRegionUsIsoEast1   CloudWatchRegion = "us-iso-east-1"
	CloudWatchRegionUsIsoWest1   CloudWatchRegion = "us-iso-west-1"
	CloudWatchRegionUsIsobEast1  CloudWatchRegion = "us-isob-east-1"
	CloudWatchRegionApSoutheast4 CloudWatchRegion = "ap-southeast-4"
	CloudWatchRegionIlCentral1   CloudWatchRegion = "il-central-1"
	CloudWatchRegionCaWest1      CloudWatchRegion = "ca-west-1"
)

// Values returns all known values for CloudWatchRegion. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CloudWatchRegion) Values() []CloudWatchRegion {
	return []CloudWatchRegion{
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"ca-central-1",
		"eu-central-1",
		"eu-central-2",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"ap-east-1",
		"me-south-1",
		"me-central-1",
		"ap-south-1",
		"ap-south-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-southeast-3",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-northeast-3",
		"eu-north-1",
		"sa-east-1",
		"cn-northwest-1",
		"cn-north-1",
		"af-south-1",
		"eu-south-1",
		"eu-south-2",
		"us-gov-west-1",
		"us-gov-east-1",
		"us-iso-east-1",
		"us-iso-west-1",
		"us-isob-east-1",
		"ap-southeast-4",
		"il-central-1",
		"ca-west-1",
	}
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorGreaterThanOrEqualToThreshold ComparisonOperator = "GreaterThanOrEqualToThreshold"
	ComparisonOperatorGreaterThanThreshold          ComparisonOperator = "GreaterThanThreshold"
	ComparisonOperatorLessThanThreshold             ComparisonOperator = "LessThanThreshold"
	ComparisonOperatorLessThanOrEqualToThreshold    ComparisonOperator = "LessThanOrEqualToThreshold"
)

// Values returns all known values for ComparisonOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ComparisonOperator) Values() []ComparisonOperator {
	return []ComparisonOperator{
		"GreaterThanOrEqualToThreshold",
		"GreaterThanThreshold",
		"LessThanThreshold",
		"LessThanOrEqualToThreshold",
	}
}

type HealthCheckRegion string

// Enum values for HealthCheckRegion
const (
	HealthCheckRegionUsEast1      HealthCheckRegion = "us-east-1"
	HealthCheckRegionUsWest1      HealthCheckRegion = "us-west-1"
	HealthCheckRegionUsWest2      HealthCheckRegion = "us-west-2"
	HealthCheckRegionEuWest1      HealthCheckRegion = "eu-west-1"
	HealthCheckRegionApSoutheast1 HealthCheckRegion = "ap-southeast-1"
	HealthCheckRegionApSoutheast2 HealthCheckRegion = "ap-southeast-2"
	HealthCheckRegionApNortheast1 HealthCheckRegion = "ap-northeast-1"
	HealthCheckRegionSaEast1      HealthCheckRegion = "sa-east-1"
)

// Values returns all known values for HealthCheckRegion. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HealthCheckRegion) Values() []HealthCheckRegion {
	return []HealthCheckRegion{
		"us-east-1",
		"us-west-1",
		"us-west-2",
		"eu-west-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-northeast-1",
		"sa-east-1",
	}
}

type HealthCheckType string

// Enum values for HealthCheckType
const (
	HealthCheckTypeHttp             HealthCheckType = "HTTP"
	HealthCheckTypeHttps            HealthCheckType = "HTTPS"
	HealthCheckTypeHttpStrMatch     HealthCheckType = "HTTP_STR_MATCH"
	HealthCheckTypeHttpsStrMatch    HealthCheckType = "HTTPS_STR_MATCH"
	HealthCheckTypeTcp              HealthCheckType = "TCP"
	HealthCheckTypeCalculated       HealthCheckType = "CALCULATED"
	HealthCheckTypeCloudwatchMetric HealthCheckType = "CLOUDWATCH_METRIC"
	HealthCheckTypeRecoveryControl  HealthCheckType = "RECOVERY_CONTROL"
)

// Values returns all known values for HealthCheckType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HealthCheckType) Values() []HealthCheckType {
	return []HealthCheckType{
		"HTTP",
		"HTTPS",
		"HTTP_STR_MATCH",
		"HTTPS_STR_MATCH",
		"TCP",
		"CALCULATED",
		"CLOUDWATCH_METRIC",
		"RECOVERY_CONTROL",
	}
}

type HostedZoneLimitType string

// Enum values for HostedZoneLimitType
const (
	HostedZoneLimitTypeMaxRrsetsByZone         HostedZoneLimitType = "MAX_RRSETS_BY_ZONE"
	HostedZoneLimitTypeMaxVpcsAssociatedByZone HostedZoneLimitType = "MAX_VPCS_ASSOCIATED_BY_ZONE"
)

// Values returns all known values for HostedZoneLimitType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HostedZoneLimitType) Values() []HostedZoneLimitType {
	return []HostedZoneLimitType{
		"MAX_RRSETS_BY_ZONE",
		"MAX_VPCS_ASSOCIATED_BY_ZONE",
	}
}

type HostedZoneType string

// Enum values for HostedZoneType
const (
	HostedZoneTypePrivateHostedZone HostedZoneType = "PrivateHostedZone"
)

// Values returns all known values for HostedZoneType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HostedZoneType) Values() []HostedZoneType {
	return []HostedZoneType{
		"PrivateHostedZone",
	}
}

type InsufficientDataHealthStatus string

// Enum values for InsufficientDataHealthStatus
const (
	InsufficientDataHealthStatusHealthy         InsufficientDataHealthStatus = "Healthy"
	InsufficientDataHealthStatusUnhealthy       InsufficientDataHealthStatus = "Unhealthy"
	InsufficientDataHealthStatusLastKnownStatus InsufficientDataHealthStatus = "LastKnownStatus"
)

// Values returns all known values for InsufficientDataHealthStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InsufficientDataHealthStatus) Values() []InsufficientDataHealthStatus {
	return []InsufficientDataHealthStatus{
		"Healthy",
		"Unhealthy",
		"LastKnownStatus",
	}
}

type ResettableElementName string

// Enum values for ResettableElementName
const (
	ResettableElementNameFullyQualifiedDomainName ResettableElementName = "FullyQualifiedDomainName"
	ResettableElementNameRegions                  ResettableElementName = "Regions"
	ResettableElementNameResourcePath             ResettableElementName = "ResourcePath"
	ResettableElementNameChildHealthChecks        ResettableElementName = "ChildHealthChecks"
)

// Values returns all known values for ResettableElementName. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResettableElementName) Values() []ResettableElementName {
	return []ResettableElementName{
		"FullyQualifiedDomainName",
		"Regions",
		"ResourcePath",
		"ChildHealthChecks",
	}
}

type ResourceRecordSetFailover string

// Enum values for ResourceRecordSetFailover
const (
	ResourceRecordSetFailoverPrimary   ResourceRecordSetFailover = "PRIMARY"
	ResourceRecordSetFailoverSecondary ResourceRecordSetFailover = "SECONDARY"
)

// Values returns all known values for ResourceRecordSetFailover. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceRecordSetFailover) Values() []ResourceRecordSetFailover {
	return []ResourceRecordSetFailover{
		"PRIMARY",
		"SECONDARY",
	}
}

type ResourceRecordSetRegion string

// Enum values for ResourceRecordSetRegion
const (
	ResourceRecordSetRegionUsEast1      ResourceRecordSetRegion = "us-east-1"
	ResourceRecordSetRegionUsEast2      ResourceRecordSetRegion = "us-east-2"
	ResourceRecordSetRegionUsWest1      ResourceRecordSetRegion = "us-west-1"
	ResourceRecordSetRegionUsWest2      ResourceRecordSetRegion = "us-west-2"
	ResourceRecordSetRegionCaCentral1   ResourceRecordSetRegion = "ca-central-1"
	ResourceRecordSetRegionEuWest1      ResourceRecordSetRegion = "eu-west-1"
	ResourceRecordSetRegionEuWest2      ResourceRecordSetRegion = "eu-west-2"
	ResourceRecordSetRegionEuWest3      ResourceRecordSetRegion = "eu-west-3"
	ResourceRecordSetRegionEuCentral1   ResourceRecordSetRegion = "eu-central-1"
	ResourceRecordSetRegionEuCentral2   ResourceRecordSetRegion = "eu-central-2"
	ResourceRecordSetRegionApSoutheast1 ResourceRecordSetRegion = "ap-southeast-1"
	ResourceRecordSetRegionApSoutheast2 ResourceRecordSetRegion = "ap-southeast-2"
	ResourceRecordSetRegionApSoutheast3 ResourceRecordSetRegion = "ap-southeast-3"
	ResourceRecordSetRegionApNortheast1 ResourceRecordSetRegion = "ap-northeast-1"
	ResourceRecordSetRegionApNortheast2 ResourceRecordSetRegion = "ap-northeast-2"
	ResourceRecordSetRegionApNortheast3 ResourceRecordSetRegion = "ap-northeast-3"
	ResourceRecordSetRegionEuNorth1     ResourceRecordSetRegion = "eu-north-1"
	ResourceRecordSetRegionSaEast1      ResourceRecordSetRegion = "sa-east-1"
	ResourceRecordSetRegionCnNorth1     ResourceRecordSetRegion = "cn-north-1"
	ResourceRecordSetRegionCnNorthwest1 ResourceRecordSetRegion = "cn-northwest-1"
	ResourceRecordSetRegionApEast1      ResourceRecordSetRegion = "ap-east-1"
	ResourceRecordSetRegionMeSouth1     ResourceRecordSetRegion = "me-south-1"
	ResourceRecordSetRegionMeCentral1   ResourceRecordSetRegion = "me-central-1"
	ResourceRecordSetRegionApSouth1     ResourceRecordSetRegion = "ap-south-1"
	ResourceRecordSetRegionApSouth2     ResourceRecordSetRegion = "ap-south-2"
	ResourceRecordSetRegionAfSouth1     ResourceRecordSetRegion = "af-south-1"
	ResourceRecordSetRegionEuSouth1     ResourceRecordSetRegion = "eu-south-1"
	ResourceRecordSetRegionEuSouth2     ResourceRecordSetRegion = "eu-south-2"
	ResourceRecordSetRegionApSoutheast4 ResourceRecordSetRegion = "ap-southeast-4"
	ResourceRecordSetRegionIlCentral1   ResourceRecordSetRegion = "il-central-1"
	ResourceRecordSetRegionCaWest1      ResourceRecordSetRegion = "ca-west-1"
)

// Values returns all known values for ResourceRecordSetRegion. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceRecordSetRegion) Values() []ResourceRecordSetRegion {
	return []ResourceRecordSetRegion{
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"ca-central-1",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"eu-central-1",
		"eu-central-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-southeast-3",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-northeast-3",
		"eu-north-1",
		"sa-east-1",
		"cn-north-1",
		"cn-northwest-1",
		"ap-east-1",
		"me-south-1",
		"me-central-1",
		"ap-south-1",
		"ap-south-2",
		"af-south-1",
		"eu-south-1",
		"eu-south-2",
		"ap-southeast-4",
		"il-central-1",
		"ca-west-1",
	}
}

type ReusableDelegationSetLimitType string

// Enum values for ReusableDelegationSetLimitType
const (
	ReusableDelegationSetLimitTypeMaxZonesByReusableDelegationSet ReusableDelegationSetLimitType = "MAX_ZONES_BY_REUSABLE_DELEGATION_SET"
)

// Values returns all known values for ReusableDelegationSetLimitType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReusableDelegationSetLimitType) Values() []ReusableDelegationSetLimitType {
	return []ReusableDelegationSetLimitType{
		"MAX_ZONES_BY_REUSABLE_DELEGATION_SET",
	}
}

type RRType string

// Enum values for RRType
const (
	RRTypeSoa   RRType = "SOA"
	RRTypeA     RRType = "A"
	RRTypeTxt   RRType = "TXT"
	RRTypeNs    RRType = "NS"
	RRTypeCname RRType = "CNAME"
	RRTypeMx    RRType = "MX"
	RRTypeNaptr RRType = "NAPTR"
	RRTypePtr   RRType = "PTR"
	RRTypeSrv   RRType = "SRV"
	RRTypeSpf   RRType = "SPF"
	RRTypeAaaa  RRType = "AAAA"
	RRTypeCaa   RRType = "CAA"
	RRTypeDs    RRType = "DS"
)

// Values returns all known values for RRType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RRType) Values() []RRType {
	return []RRType{
		"SOA",
		"A",
		"TXT",
		"NS",
		"CNAME",
		"MX",
		"NAPTR",
		"PTR",
		"SRV",
		"SPF",
		"AAAA",
		"CAA",
		"DS",
	}
}

type Statistic string

// Enum values for Statistic
const (
	StatisticAverage     Statistic = "Average"
	StatisticSum         Statistic = "Sum"
	StatisticSampleCount Statistic = "SampleCount"
	StatisticMaximum     Statistic = "Maximum"
	StatisticMinimum     Statistic = "Minimum"
)

// Values returns all known values for Statistic. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Statistic) Values() []Statistic {
	return []Statistic{
		"Average",
		"Sum",
		"SampleCount",
		"Maximum",
		"Minimum",
	}
}

type TagResourceType string

// Enum values for TagResourceType
const (
	TagResourceTypeHealthcheck TagResourceType = "healthcheck"
	TagResourceTypeHostedzone  TagResourceType = "hostedzone"
)

// Values returns all known values for TagResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TagResourceType) Values() []TagResourceType {
	return []TagResourceType{
		"healthcheck",
		"hostedzone",
	}
}

type VPCRegion string

// Enum values for VPCRegion
const (
	VPCRegionUsEast1      VPCRegion = "us-east-1"
	VPCRegionUsEast2      VPCRegion = "us-east-2"
	VPCRegionUsWest1      VPCRegion = "us-west-1"
	VPCRegionUsWest2      VPCRegion = "us-west-2"
	VPCRegionEuWest1      VPCRegion = "eu-west-1"
	VPCRegionEuWest2      VPCRegion = "eu-west-2"
	VPCRegionEuWest3      VPCRegion = "eu-west-3"
	VPCRegionEuCentral1   VPCRegion = "eu-central-1"
	VPCRegionEuCentral2   VPCRegion = "eu-central-2"
	VPCRegionApEast1      VPCRegion = "ap-east-1"
	VPCRegionMeSouth1     VPCRegion = "me-south-1"
	VPCRegionUsGovWest1   VPCRegion = "us-gov-west-1"
	VPCRegionUsGovEast1   VPCRegion = "us-gov-east-1"
	VPCRegionUsIsoEast1   VPCRegion = "us-iso-east-1"
	VPCRegionUsIsoWest1   VPCRegion = "us-iso-west-1"
	VPCRegionUsIsobEast1  VPCRegion = "us-isob-east-1"
	VPCRegionMeCentral1   VPCRegion = "me-central-1"
	VPCRegionApSoutheast1 VPCRegion = "ap-southeast-1"
	VPCRegionApSoutheast2 VPCRegion = "ap-southeast-2"
	VPCRegionApSoutheast3 VPCRegion = "ap-southeast-3"
	VPCRegionApSouth1     VPCRegion = "ap-south-1"
	VPCRegionApSouth2     VPCRegion = "ap-south-2"
	VPCRegionApNortheast1 VPCRegion = "ap-northeast-1"
	VPCRegionApNortheast2 VPCRegion = "ap-northeast-2"
	VPCRegionApNortheast3 VPCRegion = "ap-northeast-3"
	VPCRegionEuNorth1     VPCRegion = "eu-north-1"
	VPCRegionSaEast1      VPCRegion = "sa-east-1"
	VPCRegionCaCentral1   VPCRegion = "ca-central-1"
	VPCRegionCnNorth1     VPCRegion = "cn-north-1"
	VPCRegionAfSouth1     VPCRegion = "af-south-1"
	VPCRegionEuSouth1     VPCRegion = "eu-south-1"
	VPCRegionEuSouth2     VPCRegion = "eu-south-2"
	VPCRegionApSoutheast4 VPCRegion = "ap-southeast-4"
	VPCRegionIlCentral1   VPCRegion = "il-central-1"
	VPCRegionCaWest1      VPCRegion = "ca-west-1"
)

// Values returns all known values for VPCRegion. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VPCRegion) Values() []VPCRegion {
	return []VPCRegion{
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"eu-central-1",
		"eu-central-2",
		"ap-east-1",
		"me-south-1",
		"us-gov-west-1",
		"us-gov-east-1",
		"us-iso-east-1",
		"us-iso-west-1",
		"us-isob-east-1",
		"me-central-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-southeast-3",
		"ap-south-1",
		"ap-south-2",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-northeast-3",
		"eu-north-1",
		"sa-east-1",
		"ca-central-1",
		"cn-north-1",
		"af-south-1",
		"eu-south-1",
		"eu-south-2",
		"ap-southeast-4",
		"il-central-1",
		"ca-west-1",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountScope string

// Enum values for AccountScope
const (
	AccountScopePayer  AccountScope = "PAYER"
	AccountScopeLinked AccountScope = "LINKED"
)

// Values returns all known values for AccountScope. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccountScope) Values() []AccountScope {
	return []AccountScope{
		"PAYER",
		"LINKED",
	}
}

type AnomalyFeedbackType string

// Enum values for AnomalyFeedbackType
const (
	AnomalyFeedbackTypeYes             AnomalyFeedbackType = "YES"
	AnomalyFeedbackTypeNo              AnomalyFeedbackType = "NO"
	AnomalyFeedbackTypePlannedActivity AnomalyFeedbackType = "PLANNED_ACTIVITY"
)

// Values returns all known values for AnomalyFeedbackType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnomalyFeedbackType) Values() []AnomalyFeedbackType {
	return []AnomalyFeedbackType{
		"YES",
		"NO",
		"PLANNED_ACTIVITY",
	}
}

type AnomalySubscriptionFrequency string

// Enum values for AnomalySubscriptionFrequency
const (
	AnomalySubscriptionFrequencyDaily     AnomalySubscriptionFrequency = "DAILY"
	AnomalySubscriptionFrequencyImmediate AnomalySubscriptionFrequency = "IMMEDIATE"
	AnomalySubscriptionFrequencyWeekly    AnomalySubscriptionFrequency = "WEEKLY"
)

// Values returns all known values for AnomalySubscriptionFrequency. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnomalySubscriptionFrequency) Values() []AnomalySubscriptionFrequency {
	return []AnomalySubscriptionFrequency{
		"DAILY",
		"IMMEDIATE",
		"WEEKLY",
	}
}

type ApproximationDimension string

// Enum values for ApproximationDimension
const (
	ApproximationDimensionService  ApproximationDimension = "SERVICE"
	ApproximationDimensionResource ApproximationDimension = "RESOURCE"
)

// Values returns all known values for ApproximationDimension. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApproximationDimension) Values() []ApproximationDimension {
	return []ApproximationDimension{
		"SERVICE",
		"RESOURCE",
	}
}

type Context string

// Enum values for Context
const (
	ContextCostAndUsage Context = "COST_AND_USAGE"
	ContextReservations Context = "RESERVATIONS"
	ContextSavingsPlans Context = "SAVINGS_PLANS"
)

// Values returns all known values for Context. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Context) Values() []Context {
	return []Context{
		"COST_AND_USAGE",
		"RESERVATIONS",
		"SAVINGS_PLANS",
	}
}

type CostAllocationTagBackfillStatus string

// Enum values for CostAllocationTagBackfillStatus
const (
	CostAllocationTagBackfillStatusSucceeded  CostAllocationTagBackfillStatus = "SUCCEEDED"
	CostAllocationTagBackfillStatusProcessing CostAllocationTagBackfillStatus = "PROCESSING"
	CostAllocationTagBackfillStatusFailed     CostAllocationTagBackfillStatus = "FAILED"
)

// Values returns all known values for CostAllocationTagBackfillStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostAllocationTagBackfillStatus) Values() []CostAllocationTagBackfillStatus {
	return []CostAllocationTagBackfillStatus{
		"SUCCEEDED",
		"PROCESSING",
		"FAILED",
	}
}

type CostAllocationTagStatus string

// Enum values for CostAllocationTagStatus
const (
	CostAllocationTagStatusActive   CostAllocationTagStatus = "Active"
	CostAllocationTagStatusInactive CostAllocationTagStatus = "Inactive"
)

// Values returns all known values for CostAllocationTagStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostAllocationTagStatus) Values() []CostAllocationTagStatus {
	return []CostAllocationTagStatus{
		"Active",
		"Inactive",
	}
}

type CostAllocationTagType string

// Enum values for CostAllocationTagType
const (
	CostAllocationTagTypeAwsGenerated CostAllocationTagType = "AWSGenerated"
	CostAllocationTagTypeUserDefined  CostAllocationTagType = "UserDefined"
)

// Values returns all known values for CostAllocationTagType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostAllocationTagType) Values() []CostAllocationTagType {
	return []CostAllocationTagType{
		"AWSGenerated",
		"UserDefined",
	}
}

type CostCategoryInheritedValueDimensionName string

// Enum values for CostCategoryInheritedValueDimensionName
const (
	CostCategoryInheritedValueDimensionNameLinkedAccountName CostCategoryInheritedValueDimensionName = "LINKED_ACCOUNT_NAME"
	CostCategoryInheritedValueDimensionNameTag               CostCategoryInheritedValueDimensionName = "TAG"
)

// Values returns all known values for CostCategoryInheritedValueDimensionName.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostCategoryInheritedValueDimensionName) Values() []CostCategoryInheritedValueDimensionName {
	return []CostCategoryInheritedValueDimensionName{
		"LINKED_ACCOUNT_NAME",
		"TAG",
	}
}

type CostCategoryRuleType string

// Enum values for CostCategoryRuleType
const (
	CostCategoryRuleTypeRegular        CostCategoryRuleType = "REGULAR"
	CostCategoryRuleTypeInheritedValue CostCategoryRuleType = "INHERITED_VALUE"
)

// Values returns all known values for CostCategoryRuleType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostCategoryRuleType) Values() []CostCategoryRuleType {
	return []CostCategoryRuleType{
		"REGULAR",
		"INHERITED_VALUE",
	}
}

type CostCategoryRuleVersion string

// Enum values for CostCategoryRuleVersion
const (
	CostCategoryRuleVersionCostCategoryExpressionV1 CostCategoryRuleVersion = "CostCategoryExpression.v1"
)

// Values returns all known values for CostCategoryRuleVersion. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostCategoryRuleVersion) Values() []CostCategoryRuleVersion {
	return []CostCategoryRuleVersion{
		"CostCategoryExpression.v1",
	}
}

type CostCategorySplitChargeMethod string

// Enum values for CostCategorySplitChargeMethod
const (
	CostCategorySplitChargeMethodFixed        CostCategorySplitChargeMethod = "FIXED"
	CostCategorySplitChargeMethodProportional CostCategorySplitChargeMethod = "PROPORTIONAL"
	CostCategorySplitChargeMethodEven         CostCategorySplitChargeMethod = "EVEN"
)

// Values returns all known values for CostCategorySplitChargeMethod. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostCategorySplitChargeMethod) Values() []CostCategorySplitChargeMethod {
	return []CostCategorySplitChargeMethod{
		"FIXED",
		"PROPORTIONAL",
		"EVEN",
	}
}

type CostCategorySplitChargeRuleParameterType string

// Enum values for CostCategorySplitChargeRuleParameterType
const (
	CostCategorySplitChargeRuleParameterTypeAllocationPercentages CostCategorySplitChargeRuleParameterType = "ALLOCATION_PERCENTAGES"
)

// Values returns all known values for CostCategorySplitChargeRuleParameterType.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostCategorySplitChargeRuleParameterType) Values() []CostCategorySplitChargeRuleParameterType {
	return []CostCategorySplitChargeRuleParameterType{
		"ALLOCATION_PERCENTAGES",
	}
}

type CostCategoryStatus string

// Enum values for CostCategoryStatus
const (
	CostCategoryStatusProcessing CostCategoryStatus = "PROCESSING"
	CostCategoryStatusApplied    CostCategoryStatus = "APPLIED"
)

// Values returns all known values for CostCategoryStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostCategoryStatus) Values() []CostCategoryStatus {
	return []CostCategoryStatus{
		"PROCESSING",
		"APPLIED",
	}
}

type CostCategoryStatusComponent string

// Enum values for CostCategoryStatusComponent
const (
	CostCategoryStatusComponentCostExplorer CostCategoryStatusComponent = "COST_EXPLORER"
)

// Values returns all known values for CostCategoryStatusComponent. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CostCategoryStatusComponent) Values() []CostCategoryStatusComponent {
	return []CostCategoryStatusComponent{
		"COST_EXPLORER",
	}
}

type Dimension string

// Enum values for Dimension
const (
	DimensionAz                           Dimension = "AZ"
	DimensionInstanceType                 Dimension = "INSTANCE_TYPE"
	DimensionLinkedAccount                Dimension = "LINKED_ACCOUNT"
	DimensionLinkedAccountName            Dimension = "LINKED_ACCOUNT_NAME"
	DimensionOperation                    Dimension = "OPERATION"
	DimensionPurchaseType                 Dimension = "PURCHASE_TYPE"
	DimensionRegion                       Dimension = "REGION"
	DimensionService                      Dimension = "SERVICE"
	DimensionServiceCode                  Dimension = "SERVICE_CODE"
	DimensionUsageType                    Dimension = "USAGE_TYPE"
	DimensionUsageTypeGroup               Dimension = "USAGE_TYPE_GROUP"
	DimensionRecordType                   Dimension = "RECORD_TYPE"
	DimensionOperatingSystem              Dimension = "OPERATING_SYSTEM"
	DimensionTenancy                      Dimension = "TENANCY"
	DimensionScope                        Dimension = "SCOPE"
	DimensionPlatform                     Dimension = "PLATFORM"
	DimensionSubscriptionId               Dimension = "SUBSCRIPTION_ID"
	DimensionLegalEntityName              Dimension = "LEGAL_ENTITY_NAME"
	DimensionDeploymentOption             Dimension = "DEPLOYMENT_OPTION"
	DimensionDatabaseEngine               Dimension = "DATABASE_ENGINE"
	DimensionCacheEngine                  Dimension = "CACHE_ENGINE"
	DimensionInstanceTypeFamily           Dimension = "INSTANCE_TYPE_FAMILY"
	DimensionBillingEntity                Dimension = "BILLING_ENTITY"
	DimensionReservationId                Dimension = "RESERVATION_ID"
	DimensionResourceId                   Dimension = "RESOURCE_ID"
	DimensionRightsizingType              Dimension = "RIGHTSIZING_TYPE"
	DimensionSavingsPlansType             Dimension = "SAVINGS_PLANS_TYPE"
	DimensionSavingsPlanArn               Dimension = "SAVINGS_PLAN_ARN"
	DimensionPaymentOption                Dimension = "PAYMENT_OPTION"
	DimensionAgreementEndDateTimeAfter    Dimension = "AGREEMENT_END_DATE_TIME_AFTER"
	DimensionAgreementEndDateTimeBefore   Dimension = "AGREEMENT_END_DATE_TIME_BEFORE"
	DimensionInvoicingEntity              Dimension = "INVOICING_ENTITY"
	DimensionAnomalyTotalImpactAbsolute   Dimension = "ANOMALY_TOTAL_IMPACT_ABSOLUTE"
	DimensionAnomalyTotalImpactPercentage Dimension = "ANOMALY_TOTAL_IMPACT_PERCENTAGE"
)

// Values returns all known values for Dimension. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Dimension) Values() []Dimension {
	return []Dimension{
		"AZ",
		"INSTANCE_TYPE",
		"LINKED_ACCOUNT",
		"LINKED_ACCOUNT_NAME",
		"OPERATION",
		"PURCHASE_TYPE",
		"REGION",
		"SERVICE",
		"SERVICE_CODE",
		"USAGE_TYPE",
		"USAGE_TYPE_GROUP",
		"RECORD_TYPE",
		"OPERATING_SYSTEM",
		"TENANCY",
		"SCOPE",
		"PLATFORM",
		"SUBSCRIPTION_ID",
		"LEGAL_ENTITY_NAME",
		"DEPLOYMENT_OPTION",
		"DATABASE_ENGINE",
		"CACHE_ENGINE",
		"INSTANCE_TYPE_FAMILY",
		"BILLING_ENTITY",
		"RESERVATION_ID",
		"RESOURCE_ID",
		"RIGHTSIZING_TYPE",
		"SAVINGS_PLANS_TYPE",
		"SAVINGS_PLAN_ARN",
		"PAYMENT_OPTION",
		"AGREEMENT_END_DATE_TIME_AFTER",
		"AGREEMENT_END_DATE_TIME_BEFORE",
		"INVOICING_ENTITY",
		"ANOMALY_TOTAL_IMPACT_ABSOLUTE",
		"ANOMALY_TOTAL_IMPACT_PERCENTAGE",
	}
}

type FindingReasonCode string

// Enum values for FindingReasonCode
const (
	FindingReasonCodeCpuOverProvisioned               FindingReasonCode = "CPU_OVER_PROVISIONED"
	FindingReasonCodeCpuUnderProvisioned              FindingReasonCode = "CPU_UNDER_PROVISIONED"
	FindingReasonCodeMemoryOverProvisioned            FindingReasonCode = "MEMORY_OVER_PROVISIONED"
	FindingReasonCodeMemoryUnderProvisioned           FindingReasonCode = "MEMORY_UNDER_PROVISIONED"
	FindingReasonCodeEbsThroughputOverProvisioned     FindingReasonCode = "EBS_THROUGHPUT_OVER_PROVISIONED"
	FindingReasonCodeEbsThroughputUnderProvisioned    FindingReasonCode = "EBS_THROUGHPUT_UNDER_PROVISIONED"
	FindingReasonCodeEbsIopsOverProvisioned           FindingReasonCode = "EBS_IOPS_OVER_PROVISIONED"
	FindingReasonCodeEbsIopsUnderProvisioned          FindingReasonCode = "EBS_IOPS_UNDER_PROVISIONED"
	FindingReasonCodeNetworkBandwidthOverProvisioned  FindingReasonCode = "NETWORK_BANDWIDTH_OVER_PROVISIONED"
	FindingReasonCodeNetworkBandwidthUnderProvisioned FindingReasonCode = "NETWORK_BANDWIDTH_UNDER_PROVISIONED"
	FindingReasonCodeNetworkPpsOverProvisioned        FindingReasonCode = "NETWORK_PPS_OVER_PROVISIONED"
	FindingReasonCodeNetworkPpsUnderProvisioned       FindingReasonCode = "NETWORK_PPS_UNDER_PROVISIONED"
	FindingReasonCodeDiskIopsOverProvisioned          FindingReasonCode = "DISK_IOPS_OVER_PROVISIONED"
	FindingReasonCodeDiskIopsUnderProvisioned         FindingReasonCode = "DISK_IOPS_UNDER_PROVISIONED"
	FindingReasonCodeDiskThroughputOverProvisioned    FindingReasonCode = "DISK_THROUGHPUT_OVER_PROVISIONED"
	FindingReasonCodeDiskThroughputUnderProvisioned   FindingReasonCode = "DISK_THROUGHPUT_UNDER_PROVISIONED"
)

// Values returns all known values for FindingReasonCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FindingReasonCode) Values() []FindingReasonCode {
	return []FindingReasonCode{
		"CPU_OVER_PROVISIONED",
		"CPU_UNDER_PROVISIONED",
		"MEMORY_OVER_PROVISIONED",
		"MEMORY_UNDER_PROVISIONED",
		"EBS_THROUGHPUT_OVER_PROVISIONED",
		"EBS_THROUGHPUT_UNDER_PROVISIONED",
		"EBS_IOPS_OVER_PROVISIONED",
		"EBS_IOPS_UNDER_PROVISIONED",
		"NETWORK_BANDWIDTH_OVER_PROVISIONED",
		"NETWORK_BANDWIDTH_UNDER_PROVISIONED",
		"NETWORK_PPS_OVER_PROVISIONED",
		"NETWORK_PPS_UNDER_PROVISIONED",
		"DISK_IOPS_OVER_PROVISIONED",
		"DISK_IOPS_UNDER_PROVISIONED",
		"DISK_THROUGHPUT_OVER_PROVISIONED",
		"DISK_THROUGHPUT_UNDER_PROVISIONED",
	}
}

type GenerationStatus string

// Enum values for GenerationStatus
const (
	GenerationStatusSucceeded  GenerationStatus = "SUCCEEDED"
	GenerationStatusProcessing GenerationStatus = "PROCESSING"
	GenerationStatusFailed     GenerationStatus = "FAILED"
)

// Values returns all known values for GenerationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GenerationStatus) Values() []GenerationStatus {
	return []GenerationStatus{
		"SUCCEEDED",
		"PROCESSING",
		"FAILED",
	}
}

type Granularity string

// Enum values for Granularity
const (
	GranularityDaily   Granularity = "DAILY"
	GranularityMonthly Granularity = "MONTHLY"
	GranularityHourly  Granularity = "HOURLY"
)

// Values returns all known values for Granularity. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Granularity) Values() []Granularity {
	return []Granularity{
		"DAILY",
		"MONTHLY",
		"HOURLY",
	}
}

type GroupDefinitionType string

// Enum values for GroupDefinitionType
const (
	GroupDefinitionTypeDimension    GroupDefinitionType = "DIMENSION"
	GroupDefinitionTypeTag          GroupDefinitionType = "TAG"
	GroupDefinitionTypeCostCategory GroupDefinitionType = "COST_CATEGORY"
)

// Values returns all known values for GroupDefinitionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GroupDefinitionType) Values() []GroupDefinitionType {
	return []GroupDefinitionType{
		"DIMENSION",
		"TAG",
		"COST_CATEGORY",
	}
}

type LookbackPeriodInDays string

// Enum values for LookbackPeriodInDays
const (
	LookbackPeriodInDaysSevenDays  LookbackPeriodInDays = "SEVEN_DAYS"
	LookbackPeriodInDaysThirtyDays LookbackPeriodInDays = "THIRTY_DAYS"
	LookbackPeriodInDaysSixtyDays  LookbackPeriodInDays = "SIXTY_DAYS"
)

// Values returns all known values for LookbackPeriodInDays. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LookbackPeriodInDays) Values() []LookbackPeriodInDays {
	return []LookbackPeriodInDays{
		"SEVEN_DAYS",
		"THIRTY_DAYS",
		"SIXTY_DAYS",
	}
}

type MatchOption string

// Enum values for MatchOption
const (
	MatchOptionEquals             MatchOption = "EQUALS"
	MatchOptionAbsent             MatchOption = "ABSENT"
	MatchOptionStartsWith         MatchOption = "STARTS_WITH"
	MatchOptionEndsWith           MatchOption = "ENDS_WITH"
	MatchOptionContains           MatchOption = "CONTAINS"
	MatchOptionCaseSensitive      MatchOption = "CASE_SENSITIVE"
	MatchOptionCaseInsensitive    MatchOption = "CASE_INSENSITIVE"
	MatchOptionGreaterThanOrEqual MatchOption = "GREATER_THAN_OR_EQUAL"
)

// Values returns all known values for MatchOption. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MatchOption) Values() []MatchOption {
	return []MatchOption{
		"EQUALS",
		"ABSENT",
		"STARTS_WITH",
		"ENDS_WITH",
		"CONTAINS",
		"CASE_SENSITIVE",
		"CASE_INSENSITIVE",
		"GREATER_THAN_OR_EQUAL",
	}
}

type Metric string

// Enum values for Metric
const (
	MetricBlendedCost           Metric = "BLENDED_COST"
	MetricUnblendedCost         Metric = "UNBLENDED_COST"
	MetricAmortizedCost         Metric = "AMORTIZED_COST"
	MetricNetUnblendedCost      Metric = "NET_UNBLENDED_COST"
	MetricNetAmortizedCost      Metric = "NET_AMORTIZED_COST"
	MetricUsageQuantity         Metric = "USAGE_QUANTITY"
	MetricNormalizedUsageAmount Metric = "NORMALIZED_USAGE_AMOUNT"
)

// Values returns all known values for Metric. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Metric) Values() []Metric {
	return []Metric{
		"BLENDED_COST",
		"UNBLENDED_COST",
		"AMORTIZED_COST",
		"NET_UNBLENDED_COST",
		"NET_AMORTIZED_COST",
		"USAGE_QUANTITY",
		"NORMALIZED_USAGE_AMOUNT",
	}
}

type MonitorDimension string

// Enum values for MonitorDimension
const (
	MonitorDimensionService MonitorDimension = "SERVICE"
)

// Values returns all known values for MonitorDimension. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MonitorDimension) Values() []MonitorDimension {
	return []MonitorDimension{
		"SERVICE",
	}
}

type MonitorType string

// Enum values for MonitorType
const (
	MonitorTypeDimensional MonitorType = "DIMENSIONAL"
	MonitorTypeCustom      MonitorType = "CUSTOM"
)

// Values returns all known values for MonitorType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MonitorType) Values() []MonitorType {
	return []MonitorType{
		"DIMENSIONAL",
		"CUSTOM",
	}
}

type NumericOperator string

// Enum values for NumericOperator
const (
	NumericOperatorEqual              NumericOperator = "EQUAL"
	NumericOperatorGreaterThanOrEqual NumericOperator = "GREATER_THAN_OR_EQUAL"
	NumericOperatorLessThanOrEqual    NumericOperator = "LESS_THAN_OR_EQUAL"
	NumericOperatorGreaterThan        NumericOperator = "GREATER_THAN"
	NumericOperatorLessThan           NumericOperator = "LESS_THAN"
	NumericOperatorBetween            NumericOperator = "BETWEEN"
)

// Values returns all known values for NumericOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NumericOperator) Values() []NumericOperator {
	return []NumericOperator{
		"EQUAL",
		"GREATER_THAN_OR_EQUAL",
		"LESS_THAN_OR_EQUAL",
		"GREATER_THAN",
		"LESS_THAN",
		"BETWEEN",
	}
}

type OfferingClass string

// Enum values for OfferingClass
const (
	OfferingClassStandard    OfferingClass = "STANDARD"
	OfferingClassConvertible OfferingClass = "CONVERTIBLE"
)

// Values returns all known values for OfferingClass. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OfferingClass) Values() []OfferingClass {
	return []OfferingClass{
		"STANDARD",
		"CONVERTIBLE",
	}
}

type PaymentOption string

// Enum values for PaymentOption
const (
	PaymentOptionNoUpfront         PaymentOption = "NO_UPFRONT"
	PaymentOptionPartialUpfront    PaymentOption = "PARTIAL_UPFRONT"
	PaymentOptionAllUpfront        PaymentOption = "ALL_UPFRONT"
	PaymentOptionLightUtilization  PaymentOption = "LIGHT_UTILIZATION"
	PaymentOptionMediumUtilization PaymentOption = "MEDIUM_UTILIZATION"
	PaymentOptionHeavyUtilization  PaymentOption = "HEAVY_UTILIZATION"
)

// Values returns all known values for PaymentOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PaymentOption) Values() []PaymentOption {
	return []PaymentOption{
		"NO_UPFRONT",
		"PARTIAL_UPFRONT",
		"ALL_UPFRONT",
		"LIGHT_UTILIZATION",
		"MEDIUM_UTILIZATION",
		"HEAVY_UTILIZATION",
	}
}

type PlatformDifference string

// Enum values for PlatformDifference
const (
	PlatformDifferenceHypervisor                PlatformDifference = "HYPERVISOR"
	PlatformDifferenceNetworkInterface          PlatformDifference = "NETWORK_INTERFACE"
	PlatformDifferenceStorageInterface          PlatformDifference = "STORAGE_INTERFACE"
	PlatformDifferenceInstanceStoreAvailability PlatformDifference = "INSTANCE_STORE_AVAILABILITY"
	PlatformDifferenceVirtualizationType        PlatformDifference = "VIRTUALIZATION_TYPE"
)

// Values returns all known values for PlatformDifference. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PlatformDifference) Values() []PlatformDifference {
	return []PlatformDifference{
		"HYPERVISOR",
		"NETWORK_INTERFACE",
		"STORAGE_INTERFACE",
		"INSTANCE_STORE_AVAILABILITY",
		"VIRTUALIZATION_TYPE",
	}
}

type RecommendationTarget string

// Enum values for RecommendationTarget
const (
	RecommendationTargetSameInstanceFamily  RecommendationTarget = "SAME_INSTANCE_FAMILY"
	RecommendationTargetCrossInstanceFamily RecommendationTarget = "CROSS_INSTANCE_FAMILY"
)

// Values returns all known values for RecommendationTarget. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationTarget) Values() []RecommendationTarget {
	return []RecommendationTarget{
		"SAME_INSTANCE_FAMILY",
		"CROSS_INSTANCE_FAMILY",
	}
}

type RightsizingType string

// Enum values for RightsizingType
const (
	RightsizingTypeTerminate RightsizingType = "TERMINATE"
	RightsizingTypeModify    RightsizingType = "MODIFY"
)

// Values returns all known values for RightsizingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RightsizingType) Values() []RightsizingType {
	return []RightsizingType{
		"TERMINATE",
		"MODIFY",
	}
}

type SavingsPlansDataType string

// Enum values for SavingsPlansDataType
const (
	SavingsPlansDataTypeAttributes          SavingsPlansDataType = "ATTRIBUTES"
	SavingsPlansDataTypeUtilization         SavingsPlansDataType = "UTILIZATION"
	SavingsPlansDataTypeAmortizedCommitment SavingsPlansDataType = "AMORTIZED_COMMITMENT"
	SavingsPlansDataTypeSavings             SavingsPlansDataType = "SAVINGS"
)

// Values returns all known values for SavingsPlansDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SavingsPlansDataType) Values() []SavingsPlansDataType {
	return []SavingsPlansDataType{
		"ATTRIBUTES",
		"UTILIZATION",
		"AMORTIZED_COMMITMENT",
		"SAVINGS",
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

type SubscriberStatus string

// Enum values for SubscriberStatus
const (
	SubscriberStatusConfirmed SubscriberStatus = "CONFIRMED"
	SubscriberStatusDeclined  SubscriberStatus = "DECLINED"
)

// Values returns all known values for SubscriberStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SubscriberStatus) Values() []SubscriberStatus {
	return []SubscriberStatus{
		"CONFIRMED",
		"DECLINED",
	}
}

type SubscriberType string

// Enum values for SubscriberType
const (
	SubscriberTypeEmail SubscriberType = "EMAIL"
	SubscriberTypeSns   SubscriberType = "SNS"
)

// Values returns all known values for SubscriberType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SubscriberType) Values() []SubscriberType {
	return []SubscriberType{
		"EMAIL",
		"SNS",
	}
}

type SupportedSavingsPlansType string

// Enum values for SupportedSavingsPlansType
const (
	SupportedSavingsPlansTypeComputeSp     SupportedSavingsPlansType = "COMPUTE_SP"
	SupportedSavingsPlansTypeEc2InstanceSp SupportedSavingsPlansType = "EC2_INSTANCE_SP"
	SupportedSavingsPlansTypeSagemakerSp   SupportedSavingsPlansType = "SAGEMAKER_SP"
)

// Values returns all known values for SupportedSavingsPlansType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SupportedSavingsPlansType) Values() []SupportedSavingsPlansType {
	return []SupportedSavingsPlansType{
		"COMPUTE_SP",
		"EC2_INSTANCE_SP",
		"SAGEMAKER_SP",
	}
}

type TermInYears string

// Enum values for TermInYears
const (
	TermInYearsOneYear    TermInYears = "ONE_YEAR"
	TermInYearsThreeYears TermInYears = "THREE_YEARS"
)

// Values returns all known values for TermInYears. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TermInYears) Values() []TermInYears {
	return []TermInYears{
		"ONE_YEAR",
		"THREE_YEARS",
	}
}

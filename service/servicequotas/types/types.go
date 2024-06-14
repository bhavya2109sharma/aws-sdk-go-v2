// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An error that explains why an action did not succeed.
type ErrorReason struct {

	// Service Quotas returns the following error values:
	//
	//   - DEPENDENCY_ACCESS_DENIED_ERROR - The caller does not have the required
	//   permissions to complete the action. To resolve the error, you must have
	//   permission to access the Amazon Web Service or quota.
	//
	//   - DEPENDENCY_THROTTLING_ERROR - The Amazon Web Service is throttling Service
	//   Quotas.
	//
	//   - DEPENDENCY_SERVICE_ERROR - The Amazon Web Service is not available.
	//
	//   - SERVICE_QUOTA_NOT_AVAILABLE_ERROR - There was an error in Service Quotas.
	ErrorCode ErrorCode

	// The error message.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// Information about the CloudWatch metric that reflects quota usage.
type MetricInfo struct {

	// The metric dimension. This is a name/value pair that is part of the identity of
	// a metric.
	MetricDimensions map[string]string

	// The name of the metric.
	MetricName *string

	// The namespace of the metric.
	MetricNamespace *string

	// The metric statistic that we recommend you use when determining quota usage.
	MetricStatisticRecommendation *string

	noSmithyDocumentSerde
}

// A structure that describes the context for a service quota. The context
// identifies what the quota applies to.
type QuotaContextInfo struct {

	// Specifies the Amazon Web Services account or resource to which the quota
	// applies. The value in this field depends on the context scope associated with
	// the specified service quota.
	ContextId *string

	// Specifies whether the quota applies to an Amazon Web Services account, or to a
	// resource.
	ContextScope QuotaContextScope

	// When the ContextScope is RESOURCE , then this specifies the resource type of the
	// specified resource.
	ContextScopeType *string

	noSmithyDocumentSerde
}

// Information about the quota period.
type QuotaPeriod struct {

	// The time unit.
	PeriodUnit PeriodUnit

	// The value associated with the reported PeriodUnit .
	PeriodValue *int32

	noSmithyDocumentSerde
}

// Information about a quota increase request.
type RequestedServiceQuotaChange struct {

	// The case ID.
	CaseId *string

	// The date and time when the quota increase request was received and the case ID
	// was created.
	Created *time.Time

	// The new, increased value for the quota.
	DesiredValue *float64

	// Indicates whether the quota is global.
	GlobalQuota bool

	// The unique identifier.
	Id *string

	// The date and time of the most recent change.
	LastUpdated *time.Time

	// The Amazon Resource Name (ARN) of the quota.
	QuotaArn *string

	// Specifies the quota identifier. To find the quota code for a specific quota,
	// use the ListServiceQuotasoperation, and look for the QuotaCode response in the output for the
	// quota you want.
	QuotaCode *string

	// The context for this service quota.
	QuotaContext *QuotaContextInfo

	// Specifies the quota name.
	QuotaName *string

	// Specifies at which level within the Amazon Web Services account the quota
	// request applies to.
	QuotaRequestedAtLevel AppliedLevelEnum

	// The IAM identity of the requester.
	Requester *string

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServicesoperation.
	ServiceCode *string

	// Specifies the service name.
	ServiceName *string

	// The state of the quota increase request.
	Status RequestStatus

	// The unit of measurement.
	Unit *string

	noSmithyDocumentSerde
}

// Information about an Amazon Web Service.
type ServiceInfo struct {

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServicesoperation.
	ServiceCode *string

	// Specifies the service name.
	ServiceName *string

	noSmithyDocumentSerde
}

// Information about a quota.
type ServiceQuota struct {

	// Indicates whether the quota value can be increased.
	Adjustable bool

	// The error code and error reason.
	ErrorReason *ErrorReason

	// Indicates whether the quota is global.
	GlobalQuota bool

	// The period of time.
	Period *QuotaPeriod

	// Specifies at which level of granularity that the quota value is applied.
	QuotaAppliedAtLevel AppliedLevelEnum

	// The Amazon Resource Name (ARN) of the quota.
	QuotaArn *string

	// Specifies the quota identifier. To find the quota code for a specific quota,
	// use the ListServiceQuotasoperation, and look for the QuotaCode response in the output for the
	// quota you want.
	QuotaCode *string

	// The context for this service quota.
	QuotaContext *QuotaContextInfo

	// Specifies the quota name.
	QuotaName *string

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServicesoperation.
	ServiceCode *string

	// Specifies the service name.
	ServiceName *string

	// The unit of measurement.
	Unit *string

	// Information about the measurement.
	UsageMetric *MetricInfo

	// The quota value.
	Value *float64

	noSmithyDocumentSerde
}

// Information about a quota increase request.
type ServiceQuotaIncreaseRequestInTemplate struct {

	// The Amazon Web Services Region.
	AwsRegion *string

	// The new, increased value of the quota.
	DesiredValue *float64

	// Indicates whether the quota is global.
	GlobalQuota bool

	// Specifies the quota identifier. To find the quota code for a specific quota,
	// use the ListServiceQuotasoperation, and look for the QuotaCode response in the output for the
	// quota you want.
	QuotaCode *string

	// Specifies the quota name.
	QuotaName *string

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServicesoperation.
	ServiceCode *string

	// Specifies the service name.
	ServiceName *string

	// The unit of measurement.
	Unit *string

	noSmithyDocumentSerde
}

// A complex data type that contains a tag key and tag value.
type Tag struct {

	// A string that contains a tag key. The string length should be between 1 and 128
	// characters. Valid characters include a-z, A-Z, 0-9, space, and the special
	// characters _ - . : / = + @.
	//
	// This member is required.
	Key *string

	// A string that contains an optional tag value. The string length should be
	// between 0 and 256 characters. Valid characters include a-z, A-Z, 0-9, space, and
	// the special characters _ - . : / = + @.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

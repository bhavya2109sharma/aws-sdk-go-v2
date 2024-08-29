// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The request is denied because of missing access permissions.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal server error occurred. Retry your request.
type InternalServerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request failed due to an error while processing the model.
type ModelErrorException struct {
	Message *string

	ErrorCodeOverride *string

	OriginalStatusCode *int32
	ResourceName       *string

	noSmithyDocumentSerde
}

func (e *ModelErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ModelErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ModelErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ModelErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *ModelErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The model specified in the request is not ready to serve inference requests.
// The AWS SDK will automatically retry the operation up to 5 times. For
// information about configuring automatic retries, see [Retry behavior]in the AWS SDKs and Tools
// reference guide.
//
// [Retry behavior]: https://docs.aws.amazon.com/sdkref/latest/guide/feature-retry-behavior.html
type ModelNotReadyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ModelNotReadyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ModelNotReadyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ModelNotReadyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ModelNotReadyException"
	}
	return *e.ErrorCodeOverride
}
func (e *ModelNotReadyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred while streaming the response. Retry your request.
type ModelStreamErrorException struct {
	Message *string

	ErrorCodeOverride *string

	OriginalStatusCode *int32
	OriginalMessage    *string

	noSmithyDocumentSerde
}

func (e *ModelStreamErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ModelStreamErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ModelStreamErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ModelStreamErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *ModelStreamErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request took too long to process. Processing time exceeded the model
// timeout length.
type ModelTimeoutException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ModelTimeoutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ModelTimeoutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ModelTimeoutException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ModelTimeoutException"
	}
	return *e.ErrorCodeOverride
}
func (e *ModelTimeoutException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource ARN was not found. Check the ARN and try your request
// again.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your request exceeds the service quota for your account. You can view your
// quotas at [Viewing service quotas]. You can resubmit your request later.
//
// [Viewing service quotas]: https://docs.aws.amazon.com/servicequotas/latest/userguide/gs-request-quota.html
type ServiceQuotaExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceQuotaExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service isn't currently available. Try again later.
type ServiceUnavailableException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Your request was throttled because of service-wide limitations. Resubmit your
// request later or in a different region. You can also purchase [Provisioned Throughput]to increase the
// rate or number of tokens you can process.
//
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
type ThrottlingException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ThrottlingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Input validation failed. Check your request parameters and retry the request.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

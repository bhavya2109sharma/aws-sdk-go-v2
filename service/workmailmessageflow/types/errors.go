// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// WorkMail could not access the updated email content. Possible reasons:
//
//   - You made the request in a region other than your S3 bucket region.
//
//   - The [S3 bucket owner]is not the same as the calling AWS account.
//
//   - You have an incomplete or missing S3 bucket policy. For more information
//     about policies, see [Updating message content with AWS Lambda]in the WorkMail Administrator Guide.
//
// [S3 bucket owner]: https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-owner-condition.html
// [Updating message content with AWS Lambda]: https://docs.aws.amazon.com/workmail/latest/adminguide/update-with-lambda.html
type InvalidContentLocation struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidContentLocation) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidContentLocation) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidContentLocation) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidContentLocation"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidContentLocation) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested email is not eligible for update. This is usually the case for a
// redirected email.
type MessageFrozen struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MessageFrozen) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MessageFrozen) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MessageFrozen) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MessageFrozen"
	}
	return *e.ErrorCodeOverride
}
func (e *MessageFrozen) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested email could not be updated due to an error in the MIME content.
// Check the error message for more information about what caused the error.
type MessageRejected struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MessageRejected) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MessageRejected) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MessageRejected) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MessageRejected"
	}
	return *e.ErrorCodeOverride
}
func (e *MessageRejected) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested email message is not found.
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

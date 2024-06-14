// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Permanently delete the specified registration attachment.
func (c *Client) DeleteRegistrationAttachment(ctx context.Context, params *DeleteRegistrationAttachmentInput, optFns ...func(*Options)) (*DeleteRegistrationAttachmentOutput, error) {
	if params == nil {
		params = &DeleteRegistrationAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRegistrationAttachment", params, optFns, c.addOperationDeleteRegistrationAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRegistrationAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRegistrationAttachmentInput struct {

	// The unique identifier for the registration attachment.
	//
	// This member is required.
	RegistrationAttachmentId *string

	noSmithyDocumentSerde
}

type DeleteRegistrationAttachmentOutput struct {

	// The status of the registration attachment.
	//
	//   - UPLOAD_IN_PROGRESS The attachment is being uploaded.
	//
	//   - UPLOAD_COMPLETE The attachment has been uploaded.
	//
	//   - UPLOAD_FAILED The attachment failed to uploaded.
	//
	//   - DELETED The attachment has been deleted..
	//
	// This member is required.
	AttachmentStatus types.AttachmentStatus

	// The time when the registration attachment was created, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	//
	// This member is required.
	CreatedTimestamp *time.Time

	// The Amazon Resource Name (ARN) for the registration attachment.
	//
	// This member is required.
	RegistrationAttachmentArn *string

	// The unique identifier for the registration attachment.
	//
	// This member is required.
	RegistrationAttachmentId *string

	// The error message if the upload failed.
	AttachmentUploadErrorReason types.AttachmentUploadErrorReason

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRegistrationAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteRegistrationAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteRegistrationAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteRegistrationAttachment"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpDeleteRegistrationAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRegistrationAttachment(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteRegistrationAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteRegistrationAttachment",
	}
}

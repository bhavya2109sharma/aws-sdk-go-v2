// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancel a specific image lifecycle policy runtime instance.
func (c *Client) CancelLifecycleExecution(ctx context.Context, params *CancelLifecycleExecutionInput, optFns ...func(*Options)) (*CancelLifecycleExecutionOutput, error) {
	if params == nil {
		params = &CancelLifecycleExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelLifecycleExecution", params, optFns, c.addOperationCancelLifecycleExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelLifecycleExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelLifecycleExecutionInput struct {

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see [Ensuring idempotency]in the Amazon EC2 API Reference.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	//
	// This member is required.
	ClientToken *string

	// Identifies the specific runtime instance of the image lifecycle to cancel.
	//
	// This member is required.
	LifecycleExecutionId *string

	noSmithyDocumentSerde
}

type CancelLifecycleExecutionOutput struct {

	// The unique identifier for the image lifecycle runtime instance that was
	// canceled.
	LifecycleExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelLifecycleExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelLifecycleExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelLifecycleExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelLifecycleExecution"); err != nil {
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
	if err = addIdempotencyToken_opCancelLifecycleExecutionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCancelLifecycleExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelLifecycleExecution(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCancelLifecycleExecution struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCancelLifecycleExecution) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCancelLifecycleExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CancelLifecycleExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CancelLifecycleExecutionInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCancelLifecycleExecutionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCancelLifecycleExecution{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCancelLifecycleExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelLifecycleExecution",
	}
}

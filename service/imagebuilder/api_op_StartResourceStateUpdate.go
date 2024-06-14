// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Begin asynchronous resource state update for lifecycle changes to the specified
// image resources.
func (c *Client) StartResourceStateUpdate(ctx context.Context, params *StartResourceStateUpdateInput, optFns ...func(*Options)) (*StartResourceStateUpdateOutput, error) {
	if params == nil {
		params = &StartResourceStateUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartResourceStateUpdate", params, optFns, c.addOperationStartResourceStateUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartResourceStateUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartResourceStateUpdateInput struct {

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see [Ensuring idempotency]in the Amazon EC2 API Reference.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	//
	// This member is required.
	ClientToken *string

	// The ARN of the Image Builder resource that is updated. The state update might
	// also impact associated resources.
	//
	// This member is required.
	ResourceArn *string

	// Indicates the lifecycle action to take for this request.
	//
	// This member is required.
	State *types.ResourceState

	// Skip action on the image resource and associated resources if specified
	// exclusion rules are met.
	ExclusionRules *types.ResourceStateUpdateExclusionRules

	// The name or Amazon Resource Name (ARN) of the IAM role that’s used to update
	// image state.
	ExecutionRole *string

	// A list of image resources to update state for.
	IncludeResources *types.ResourceStateUpdateIncludeResources

	// The timestamp that indicates when resources are updated by a lifecycle action.
	UpdateAt *time.Time

	noSmithyDocumentSerde
}

type StartResourceStateUpdateOutput struct {

	// Identifies the lifecycle runtime instance that started the resource state
	// update.
	LifecycleExecutionId *string

	// The requested ARN of the Image Builder resource for the asynchronous update.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartResourceStateUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartResourceStateUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartResourceStateUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartResourceStateUpdate"); err != nil {
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
	if err = addIdempotencyToken_opStartResourceStateUpdateMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartResourceStateUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartResourceStateUpdate(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartResourceStateUpdate struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartResourceStateUpdate) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartResourceStateUpdate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartResourceStateUpdateInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartResourceStateUpdateInput ")
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
func addIdempotencyToken_opStartResourceStateUpdateMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartResourceStateUpdate{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartResourceStateUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartResourceStateUpdate",
	}
}

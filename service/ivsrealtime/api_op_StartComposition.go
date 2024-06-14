// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a Composition from a stage based on the configuration provided in the
// request.
//
// A Composition is an ephemeral resource that exists after this endpoint returns
// successfully. Composition stops and the resource is deleted:
//
//   - When StopCompositionis called.
//
//   - After a 1-minute timeout, when all participants are disconnected from the
//     stage.
//
//   - After a 1-minute timeout, if there are no participants in the stage when
//     StartComposition is called.
//
//   - When broadcasting to the IVS channel fails and all retries are exhausted.
//
//   - When broadcasting is disconnected and all attempts to reconnect are
//     exhausted.
func (c *Client) StartComposition(ctx context.Context, params *StartCompositionInput, optFns ...func(*Options)) (*StartCompositionOutput, error) {
	if params == nil {
		params = &StartCompositionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartComposition", params, optFns, c.addOperationStartCompositionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartCompositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCompositionInput struct {

	// Array of destination configuration.
	//
	// This member is required.
	Destinations []types.DestinationConfiguration

	// ARN of the stage to be used for compositing.
	//
	// This member is required.
	StageArn *string

	// Idempotency token.
	IdempotencyToken *string

	// Layout object to configure composition parameters.
	Layout *types.LayoutConfiguration

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Tagging AWS Resources] for details, including restrictions that apply to tags and
	// "Tag naming limits and requirements"; Amazon IVS has no constraints on tags
	// beyond what is documented there.
	//
	// [Tagging AWS Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags map[string]string

	noSmithyDocumentSerde
}

type StartCompositionOutput struct {

	// The Composition that was created.
	Composition *types.Composition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartCompositionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartComposition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartComposition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartComposition"); err != nil {
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
	if err = addIdempotencyToken_opStartCompositionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartCompositionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartComposition(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartComposition struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartComposition) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartComposition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartCompositionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartCompositionInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartCompositionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartComposition{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartComposition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartComposition",
	}
}

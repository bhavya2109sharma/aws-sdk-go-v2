// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a version of the guardrail. Use this API to create a snapshot of the
// guardrail when you are satisfied with a configuration, or to compare the
// configuration with another version.
func (c *Client) CreateGuardrailVersion(ctx context.Context, params *CreateGuardrailVersionInput, optFns ...func(*Options)) (*CreateGuardrailVersionOutput, error) {
	if params == nil {
		params = &CreateGuardrailVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGuardrailVersion", params, optFns, c.addOperationCreateGuardrailVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGuardrailVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGuardrailVersionInput struct {

	// The unique identifier of the guardrail. This can be an ID or the ARN.
	//
	// This member is required.
	GuardrailIdentifier *string

	// A unique, case-sensitive identifier to ensure that the API request completes no
	// more than once. If this token matches a previous request, Amazon Bedrock ignores
	// the request, but does not return an error. For more information, see [Ensuring idempotency]in the
	// Amazon S3 User Guide.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientRequestToken *string

	// A description of the guardrail version.
	Description *string

	noSmithyDocumentSerde
}

type CreateGuardrailVersionOutput struct {

	// The unique identifier of the guardrail.
	//
	// This member is required.
	GuardrailId *string

	// The number of the version of the guardrail.
	//
	// This member is required.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGuardrailVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGuardrailVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGuardrailVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateGuardrailVersion"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateGuardrailVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateGuardrailVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGuardrailVersion(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateGuardrailVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateGuardrailVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateGuardrailVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateGuardrailVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateGuardrailVersionInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateGuardrailVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateGuardrailVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateGuardrailVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateGuardrailVersion",
	}
}

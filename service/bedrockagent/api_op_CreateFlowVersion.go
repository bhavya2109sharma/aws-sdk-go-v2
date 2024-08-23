// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a version of the flow that you can deploy. For more information, see [Deploy a flow in Amazon Bedrock]
// in the Amazon Bedrock User Guide.
//
// [Deploy a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-deploy.html
func (c *Client) CreateFlowVersion(ctx context.Context, params *CreateFlowVersionInput, optFns ...func(*Options)) (*CreateFlowVersionOutput, error) {
	if params == nil {
		params = &CreateFlowVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFlowVersion", params, optFns, c.addOperationCreateFlowVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFlowVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlowVersionInput struct {

	// The unique identifier of the flow that you want to create a version of.
	//
	// This member is required.
	FlowIdentifier *string

	// A unique, case-sensitive identifier to ensure that the API request completes no
	// more than one time. If this token matches a previous request, Amazon Bedrock
	// ignores the request, but does not return an error. For more information, see [Ensuring idempotency].
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientToken *string

	// A description of the version of the flow.
	Description *string

	noSmithyDocumentSerde
}

type CreateFlowVersionOutput struct {

	// The Amazon Resource Name (ARN) of the flow.
	//
	// This member is required.
	Arn *string

	// The time at which the flow was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the service role with permissions to create a
	// flow. For more information, see [Create a service role for flows in Amazon Bedrock]in the Amazon Bedrock User Guide.
	//
	// [Create a service role for flows in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-permissions.html
	//
	// This member is required.
	ExecutionRoleArn *string

	// The unique identifier of the flow.
	//
	// This member is required.
	Id *string

	// The name of the version.
	//
	// This member is required.
	Name *string

	// The status of the flow.
	//
	// This member is required.
	Status types.FlowStatus

	// The version of the flow that was created. Versions are numbered incrementally,
	// starting from 1.
	//
	// This member is required.
	Version *string

	// The KMS key that the flow is encrypted with.
	CustomerEncryptionKeyArn *string

	// A definition of the nodes and connections in the flow.
	Definition *types.FlowDefinition

	// The description of the version.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFlowVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFlowVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFlowVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFlowVersion"); err != nil {
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
	if err = addIdempotencyToken_opCreateFlowVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFlowVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlowVersion(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFlowVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFlowVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFlowVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFlowVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFlowVersionInput ")
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
func addIdempotencyToken_opCreateFlowVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFlowVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFlowVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFlowVersion",
	}
}

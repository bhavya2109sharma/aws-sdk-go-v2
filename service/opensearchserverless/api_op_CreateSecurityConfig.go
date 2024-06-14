// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies a security configuration for OpenSearch Serverless. For more
// information, see [SAML authentication for Amazon OpenSearch Serverless].
//
// [SAML authentication for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-saml.html
func (c *Client) CreateSecurityConfig(ctx context.Context, params *CreateSecurityConfigInput, optFns ...func(*Options)) (*CreateSecurityConfigOutput, error) {
	if params == nil {
		params = &CreateSecurityConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecurityConfig", params, optFns, c.addOperationCreateSecurityConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecurityConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityConfigInput struct {

	// The name of the security configuration.
	//
	// This member is required.
	Name *string

	// The type of security configuration.
	//
	// This member is required.
	Type types.SecurityConfigType

	// Unique, case-sensitive identifier to ensure idempotency of the request.
	ClientToken *string

	// A description of the security configuration.
	Description *string

	// Describes SAML options in in the form of a key-value map. This field is
	// required if you specify saml for the type parameter.
	SamlOptions *types.SamlConfigOptions

	noSmithyDocumentSerde
}

type CreateSecurityConfigOutput struct {

	// Details about the created security configuration.
	SecurityConfigDetail *types.SecurityConfigDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSecurityConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateSecurityConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateSecurityConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSecurityConfig"); err != nil {
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
	if err = addIdempotencyToken_opCreateSecurityConfigMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSecurityConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityConfig(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateSecurityConfig struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSecurityConfig) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSecurityConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateSecurityConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateSecurityConfigInput ")
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
func addIdempotencyToken_opCreateSecurityConfigMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateSecurityConfig{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSecurityConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSecurityConfig",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new container recipe. Container recipes define how images are
// configured, tested, and assessed.
func (c *Client) CreateContainerRecipe(ctx context.Context, params *CreateContainerRecipeInput, optFns ...func(*Options)) (*CreateContainerRecipeOutput, error) {
	if params == nil {
		params = &CreateContainerRecipeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContainerRecipe", params, optFns, c.addOperationCreateContainerRecipeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContainerRecipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContainerRecipeInput struct {

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see [Ensuring idempotency]in the Amazon EC2 API Reference.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	//
	// This member is required.
	ClientToken *string

	// Components for build and test that are included in the container recipe.
	// Recipes require a minimum of one build component, and can have a maximum of 20
	// build and test components in any combination.
	//
	// This member is required.
	Components []types.ComponentConfiguration

	// The type of container to create.
	//
	// This member is required.
	ContainerType types.ContainerType

	// The name of the container recipe.
	//
	// This member is required.
	Name *string

	// The base image for the container recipe.
	//
	// This member is required.
	ParentImage *string

	// The semantic version of the container recipe. This version follows the semantic
	// version syntax.
	//
	// The semantic version has four nodes: ../. You can assign values for the first
	// three, and can filter on all of them.
	//
	// Assignment: For the first three nodes you can assign any positive integer
	// value, including zero, with an upper limit of 2^30-1, or 1073741823 for each
	// node. Image Builder automatically assigns the build number to the fourth node.
	//
	// Patterns: You can use any numeric pattern that adheres to the assignment
	// requirements for the nodes that you can assign. For example, you might choose a
	// software version pattern, such as 1.0.0, or a date, such as 2021.01.01.
	//
	// This member is required.
	SemanticVersion *string

	// The destination repository for the container image.
	//
	// This member is required.
	TargetRepository *types.TargetContainerRepository

	// The description of the container recipe.
	Description *string

	// The Dockerfile template used to build your image as an inline data blob.
	DockerfileTemplateData *string

	// The Amazon S3 URI for the Dockerfile that will be used to build your container
	// image.
	DockerfileTemplateUri *string

	// Specifies the operating system version for the base image.
	ImageOsVersionOverride *string

	// A group of options that can be used to configure an instance for building and
	// testing container images.
	InstanceConfiguration *types.InstanceConfiguration

	// Identifies which KMS key is used to encrypt the container image.
	KmsKeyId *string

	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride types.Platform

	// Tags that are attached to the container recipe.
	Tags map[string]string

	// The working directory for use during build and test workflows.
	WorkingDirectory *string

	noSmithyDocumentSerde
}

type CreateContainerRecipeOutput struct {

	// The client token that uniquely identifies the request.
	ClientToken *string

	// Returns the Amazon Resource Name (ARN) of the container recipe that the request
	// created.
	ContainerRecipeArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContainerRecipeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateContainerRecipe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateContainerRecipe{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateContainerRecipe"); err != nil {
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
	if err = addIdempotencyToken_opCreateContainerRecipeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateContainerRecipeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContainerRecipe(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateContainerRecipe struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateContainerRecipe) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateContainerRecipe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateContainerRecipeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateContainerRecipeInput ")
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
func addIdempotencyToken_opCreateContainerRecipeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateContainerRecipe{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateContainerRecipe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateContainerRecipe",
	}
}

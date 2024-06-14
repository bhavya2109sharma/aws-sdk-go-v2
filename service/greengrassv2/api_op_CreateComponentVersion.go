// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a component. Components are software that run on Greengrass core
// devices. After you develop and test a component on your core device, you can use
// this operation to upload your component to IoT Greengrass. Then, you can deploy
// the component to other core devices.
//
// You can use this operation to do the following:
//
//   - Create components from recipes
//
// Create a component from a recipe, which is a file that defines the component's
//
//	metadata, parameters, dependencies, lifecycle, artifacts, and platform
//	capability. For more information, see [IoT Greengrass component recipe reference]in the IoT Greengrass V2 Developer
//	Guide.
//
// To create a component from a recipe, specify inlineRecipe when you call this
//
//	operation.
//
//	- Create components from Lambda functions
//
// Create a component from an Lambda function that runs on IoT Greengrass. This
//
//	creates a recipe and artifacts from the Lambda function's deployment package.
//	You can use this operation to migrate Lambda functions from IoT Greengrass V1 to
//	IoT Greengrass V2.
//
// This function accepts Lambda functions in all supported versions of Python,
//
//	Node.js, and Java runtimes. IoT Greengrass doesn't apply any additional
//	restrictions on deprecated Lambda runtime versions.
//
// To create a component from a Lambda function, specify lambdaFunction when you
//
//	call this operation.
//
// IoT Greengrass currently supports Lambda functions on only Linux core devices.
//
// [IoT Greengrass component recipe reference]: https://docs.aws.amazon.com/greengrass/v2/developerguide/component-recipe-reference.html
func (c *Client) CreateComponentVersion(ctx context.Context, params *CreateComponentVersionInput, optFns ...func(*Options)) (*CreateComponentVersionOutput, error) {
	if params == nil {
		params = &CreateComponentVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComponentVersion", params, optFns, c.addOperationCreateComponentVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateComponentVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateComponentVersionInput struct {

	// A unique, case-sensitive identifier that you can provide to ensure that the
	// request is idempotent. Idempotency means that the request is successfully
	// processed only once, even if you send the request multiple times. When a request
	// succeeds, and you specify the same client token for subsequent successful
	// requests, the IoT Greengrass V2 service returns the successful response that it
	// caches from the previous request. IoT Greengrass V2 caches successful responses
	// for idempotent requests for up to 8 hours.
	ClientToken *string

	// The recipe to use to create the component. The recipe defines the component's
	// metadata, parameters, dependencies, lifecycle, artifacts, and platform
	// compatibility.
	//
	// You must specify either inlineRecipe or lambdaFunction .
	InlineRecipe []byte

	// The parameters to create a component from a Lambda function.
	//
	// You must specify either inlineRecipe or lambdaFunction .
	LambdaFunction *types.LambdaFunctionRecipeSource

	// A list of key-value pairs that contain metadata for the resource. For more
	// information, see [Tag your resources]in the IoT Greengrass V2 Developer Guide.
	//
	// [Tag your resources]: https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateComponentVersionOutput struct {

	// The name of the component.
	//
	// This member is required.
	ComponentName *string

	// The version of the component.
	//
	// This member is required.
	ComponentVersion *string

	// The time at which the component was created, expressed in ISO 8601 format.
	//
	// This member is required.
	CreationTimestamp *time.Time

	// The status of the component version in IoT Greengrass V2. This status is
	// different from the status of the component on a core device.
	//
	// This member is required.
	Status *types.CloudComponentStatus

	// The [ARN] of the component version.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateComponentVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateComponentVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateComponentVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateComponentVersion"); err != nil {
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
	if err = addIdempotencyToken_opCreateComponentVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateComponentVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComponentVersion(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateComponentVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateComponentVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateComponentVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateComponentVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateComponentVersionInput ")
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
func addIdempotencyToken_opCreateComponentVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateComponentVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateComponentVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateComponentVersion",
	}
}

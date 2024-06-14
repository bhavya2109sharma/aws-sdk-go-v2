// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a virtual node within a service mesh.
//
// A virtual node acts as a logical pointer to a particular task group, such as an
// Amazon ECS service or a Kubernetes deployment. When you create a virtual node,
// you can specify the service discovery information for your task group, and
// whether the proxy running in a task group will communicate with other proxies
// using Transport Layer Security (TLS).
//
// You define a listener for any inbound traffic that your virtual node expects.
// Any virtual service that your virtual node expects to communicate to is
// specified as a backend .
//
// The response metadata for your new virtual node contains the arn that is
// associated with the virtual node. Set this value to the full ARN; for example,
// arn:aws:appmesh:us-west-2:123456789012:myMesh/default/virtualNode/myApp ) as the
// APPMESH_RESOURCE_ARN environment variable for your task group's Envoy proxy
// container in your task definition or pod spec. This is then mapped to the
// node.id and node.cluster Envoy parameters.
//
// By default, App Mesh uses the name of the resource you specified in
// APPMESH_RESOURCE_ARN when Envoy is referring to itself in metrics and traces.
// You can override this behavior by setting the APPMESH_RESOURCE_CLUSTER
// environment variable with your own name.
//
// For more information about virtual nodes, see [Virtual nodes]. You must be using 1.15.0 or
// later of the Envoy image when setting these variables. For more information
// aboutApp Mesh Envoy variables, see [Envoy image]in the App Mesh User Guide.
//
// [Virtual nodes]: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html
// [Envoy image]: https://docs.aws.amazon.com/app-mesh/latest/userguide/envoy.html
func (c *Client) CreateVirtualNode(ctx context.Context, params *CreateVirtualNodeInput, optFns ...func(*Options)) (*CreateVirtualNodeOutput, error) {
	if params == nil {
		params = &CreateVirtualNodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVirtualNode", params, optFns, c.addOperationCreateVirtualNodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVirtualNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVirtualNodeInput struct {

	// The name of the service mesh to create the virtual node in.
	//
	// This member is required.
	MeshName *string

	// The virtual node specification to apply.
	//
	// This member is required.
	Spec *types.VirtualNodeSpec

	// The name to use for the virtual node.
	//
	// This member is required.
	VirtualNodeName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The Amazon Web Services IAM account ID of the service mesh owner. If the
	// account ID is not your own, then the account that you specify must share the
	// mesh with your account before you can create the resource in the service mesh.
	// For more information about mesh sharing, see [Working with shared meshes].
	//
	// [Working with shared meshes]: https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html
	MeshOwner *string

	// Optional metadata that you can apply to the virtual node to assist with
	// categorization and organization. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []types.TagRef

	noSmithyDocumentSerde
}

type CreateVirtualNodeOutput struct {

	// The full description of your virtual node following the create call.
	//
	// This member is required.
	VirtualNode *types.VirtualNodeData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVirtualNodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVirtualNode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVirtualNode{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVirtualNode"); err != nil {
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
	if err = addIdempotencyToken_opCreateVirtualNodeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateVirtualNodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVirtualNode(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateVirtualNode struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateVirtualNode) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateVirtualNode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateVirtualNodeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateVirtualNodeInput ")
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
func addIdempotencyToken_opCreateVirtualNodeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateVirtualNode{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateVirtualNode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVirtualNode",
	}
}

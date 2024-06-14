// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Discovers the increasing revision associated with an instance.
func (c *Client) DiscoverInstancesRevision(ctx context.Context, params *DiscoverInstancesRevisionInput, optFns ...func(*Options)) (*DiscoverInstancesRevisionOutput, error) {
	if params == nil {
		params = &DiscoverInstancesRevisionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DiscoverInstancesRevision", params, optFns, c.addOperationDiscoverInstancesRevisionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DiscoverInstancesRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DiscoverInstancesRevisionInput struct {

	// The HttpName name of the namespace. It's found in the HttpProperties member of
	// the Properties member of the namespace.
	//
	// This member is required.
	NamespaceName *string

	// The name of the service that you specified when you registered the instance.
	//
	// This member is required.
	ServiceName *string

	noSmithyDocumentSerde
}

type DiscoverInstancesRevisionOutput struct {

	// The increasing revision associated to the response Instances list. If a new
	// instance is registered or deregistered, the InstancesRevision updates. The
	// health status updates don't update InstancesRevision .
	InstancesRevision *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDiscoverInstancesRevisionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDiscoverInstancesRevision{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDiscoverInstancesRevision{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DiscoverInstancesRevision"); err != nil {
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
	if err = addEndpointPrefix_opDiscoverInstancesRevisionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDiscoverInstancesRevisionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDiscoverInstancesRevision(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDiscoverInstancesRevisionMiddleware struct {
}

func (*endpointPrefix_opDiscoverInstancesRevisionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDiscoverInstancesRevisionMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDiscoverInstancesRevisionMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDiscoverInstancesRevisionMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDiscoverInstancesRevision(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DiscoverInstancesRevision",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specify the Amazon EC2 instance (destination) IP addresses and ports for a VPC
// subnet endpoint that can receive traffic for a custom routing accelerator. You
// can allow traffic to all destinations in the subnet endpoint, or allow traffic
// to a specified list of destination IP addresses and ports in the subnet. Note
// that you cannot specify IP addresses or ports outside of the range that you
// configured for the endpoint group.
//
// After you make changes, you can verify that the updates are complete by
// checking the status of your accelerator: the status changes from IN_PROGRESS to
// DEPLOYED.
func (c *Client) AllowCustomRoutingTraffic(ctx context.Context, params *AllowCustomRoutingTrafficInput, optFns ...func(*Options)) (*AllowCustomRoutingTrafficOutput, error) {
	if params == nil {
		params = &AllowCustomRoutingTrafficInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AllowCustomRoutingTraffic", params, optFns, c.addOperationAllowCustomRoutingTrafficMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AllowCustomRoutingTrafficOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllowCustomRoutingTrafficInput struct {

	// The Amazon Resource Name (ARN) of the endpoint group.
	//
	// This member is required.
	EndpointGroupArn *string

	// An ID for the endpoint. For custom routing accelerators, this is the virtual
	// private cloud (VPC) subnet ID.
	//
	// This member is required.
	EndpointId *string

	// Indicates whether all destination IP addresses and ports for a specified VPC
	// subnet endpoint can receive traffic from a custom routing accelerator. The value
	// is TRUE or FALSE.
	//
	// When set to TRUE, all destinations in the custom routing VPC subnet can receive
	// traffic. Note that you cannot specify destination IP addresses and ports when
	// the value is set to TRUE.
	//
	// When set to FALSE (or not specified), you must specify a list of destination IP
	// addresses that are allowed to receive traffic. A list of ports is optional. If
	// you don't specify a list of ports, the ports that can accept traffic is the same
	// as the ports configured for the endpoint group.
	//
	// The default value is FALSE.
	AllowAllTrafficToEndpoint *bool

	// A list of specific Amazon EC2 instance IP addresses (destination addresses) in
	// a subnet that you want to allow to receive traffic. The IP addresses must be a
	// subset of the IP addresses that you specified for the endpoint group.
	//
	// DestinationAddresses is required if AllowAllTrafficToEndpoint is FALSE or is
	// not specified.
	DestinationAddresses []string

	// A list of specific Amazon EC2 instance ports (destination ports) that you want
	// to allow to receive traffic.
	DestinationPorts []int32

	noSmithyDocumentSerde
}

type AllowCustomRoutingTrafficOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAllowCustomRoutingTrafficMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAllowCustomRoutingTraffic{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAllowCustomRoutingTraffic{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AllowCustomRoutingTraffic"); err != nil {
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
	if err = addOpAllowCustomRoutingTrafficValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAllowCustomRoutingTraffic(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAllowCustomRoutingTraffic(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AllowCustomRoutingTraffic",
	}
}

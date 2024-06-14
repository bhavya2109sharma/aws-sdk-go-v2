// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provisions an IP address range to use with your Amazon Web Services resources
// through bring your own IP addresses (BYOIP) and creates a corresponding address
// pool. After the address range is provisioned, it is ready to be advertised using
// [AdvertiseByoipCidr].
//
// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
//
// [AdvertiseByoipCidr]: https://docs.aws.amazon.com/global-accelerator/latest/api/AdvertiseByoipCidr.html
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
func (c *Client) ProvisionByoipCidr(ctx context.Context, params *ProvisionByoipCidrInput, optFns ...func(*Options)) (*ProvisionByoipCidrOutput, error) {
	if params == nil {
		params = &ProvisionByoipCidrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ProvisionByoipCidr", params, optFns, c.addOperationProvisionByoipCidrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ProvisionByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ProvisionByoipCidrInput struct {

	// The public IPv4 address range, in CIDR notation. The most specific IP prefix
	// that you can specify is /24. The address range cannot overlap with another
	// address range that you've brought to this Amazon Web Services Region or another
	// Region.
	//
	// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
	//
	// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
	//
	// This member is required.
	Cidr *string

	// A signed document that proves that you are authorized to bring the specified IP
	// address range to Amazon using BYOIP.
	//
	// This member is required.
	CidrAuthorizationContext *types.CidrAuthorizationContext

	noSmithyDocumentSerde
}

type ProvisionByoipCidrOutput struct {

	// Information about the address range.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationProvisionByoipCidrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpProvisionByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpProvisionByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ProvisionByoipCidr"); err != nil {
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
	if err = addOpProvisionByoipCidrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opProvisionByoipCidr(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opProvisionByoipCidr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ProvisionByoipCidr",
	}
}

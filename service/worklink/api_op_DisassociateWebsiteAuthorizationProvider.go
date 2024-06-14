// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a website authorization provider from a specified fleet. After
// the disassociation, users can't load any associated websites that require this
// authorization provider.
//
// Deprecated: Amazon WorkLink is no longer supported. This will be removed in a
// future version of the SDK.
func (c *Client) DisassociateWebsiteAuthorizationProvider(ctx context.Context, params *DisassociateWebsiteAuthorizationProviderInput, optFns ...func(*Options)) (*DisassociateWebsiteAuthorizationProviderOutput, error) {
	if params == nil {
		params = &DisassociateWebsiteAuthorizationProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateWebsiteAuthorizationProvider", params, optFns, c.addOperationDisassociateWebsiteAuthorizationProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateWebsiteAuthorizationProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateWebsiteAuthorizationProviderInput struct {

	// A unique identifier for the authorization provider.
	//
	// This member is required.
	AuthorizationProviderId *string

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	noSmithyDocumentSerde
}

type DisassociateWebsiteAuthorizationProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateWebsiteAuthorizationProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateWebsiteAuthorizationProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateWebsiteAuthorizationProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateWebsiteAuthorizationProvider"); err != nil {
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
	if err = addOpDisassociateWebsiteAuthorizationProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateWebsiteAuthorizationProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateWebsiteAuthorizationProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateWebsiteAuthorizationProvider",
	}
}

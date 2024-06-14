// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an OpenID Connect identity provider (IdP) resource object in IAM.
//
// Deleting an IAM OIDC provider resource does not update any roles that reference
// the provider as a principal in their trust policies. Any attempt to assume a
// role that references a deleted provider fails.
//
// This operation is idempotent; it does not fail or return an error if you call
// the operation for a provider that does not exist.
func (c *Client) DeleteOpenIDConnectProvider(ctx context.Context, params *DeleteOpenIDConnectProviderInput, optFns ...func(*Options)) (*DeleteOpenIDConnectProviderOutput, error) {
	if params == nil {
		params = &DeleteOpenIDConnectProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOpenIDConnectProvider", params, optFns, c.addOperationDeleteOpenIDConnectProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOpenIDConnectProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOpenIDConnectProviderInput struct {

	// The Amazon Resource Name (ARN) of the IAM OpenID Connect provider resource
	// object to delete. You can get a list of OpenID Connect provider resource ARNs by
	// using the ListOpenIDConnectProvidersoperation.
	//
	// This member is required.
	OpenIDConnectProviderArn *string

	noSmithyDocumentSerde
}

type DeleteOpenIDConnectProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOpenIDConnectProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteOpenIDConnectProvider"); err != nil {
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
	if err = addOpDeleteOpenIDConnectProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOpenIDConnectProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOpenIDConnectProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteOpenIDConnectProvider",
	}
}

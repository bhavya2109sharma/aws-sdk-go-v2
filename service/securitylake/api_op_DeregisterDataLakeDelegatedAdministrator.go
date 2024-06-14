// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the Amazon Security Lake delegated administrator account for the
// organization. This API can only be called by the organization management
// account. The organization management account cannot be the delegated
// administrator account.
func (c *Client) DeregisterDataLakeDelegatedAdministrator(ctx context.Context, params *DeregisterDataLakeDelegatedAdministratorInput, optFns ...func(*Options)) (*DeregisterDataLakeDelegatedAdministratorOutput, error) {
	if params == nil {
		params = &DeregisterDataLakeDelegatedAdministratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterDataLakeDelegatedAdministrator", params, optFns, c.addOperationDeregisterDataLakeDelegatedAdministratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterDataLakeDelegatedAdministratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterDataLakeDelegatedAdministratorInput struct {
	noSmithyDocumentSerde
}

type DeregisterDataLakeDelegatedAdministratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterDataLakeDelegatedAdministratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeregisterDataLakeDelegatedAdministrator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeregisterDataLakeDelegatedAdministrator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeregisterDataLakeDelegatedAdministrator"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterDataLakeDelegatedAdministrator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterDataLakeDelegatedAdministrator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeregisterDataLakeDelegatedAdministrator",
	}
}

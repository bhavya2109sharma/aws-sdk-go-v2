// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified tags from the contact resource. For more information
// about this API is used, see [Set up granular billing for a detailed view of your Amazon Connect usage].
//
// [Set up granular billing for a detailed view of your Amazon Connect usage]: https://docs.aws.amazon.com/connect/latest/adminguide/granular-billing.html
func (c *Client) UntagContact(ctx context.Context, params *UntagContactInput, optFns ...func(*Options)) (*UntagContactOutput, error) {
	if params == nil {
		params = &UntagContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UntagContact", params, optFns, c.addOperationUntagContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UntagContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagContactInput struct {

	// The identifier of the contact in this instance of Amazon Connect.
	//
	// This member is required.
	ContactId *string

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// A list of tag keys. Existing tags on the contact whose keys are members of this
	// list will be removed.
	//
	// This member is required.
	TagKeys []string

	noSmithyDocumentSerde
}

type UntagContactOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUntagContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUntagContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUntagContact{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UntagContact"); err != nil {
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
	if err = addOpUntagContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUntagContact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUntagContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UntagContact",
	}
}

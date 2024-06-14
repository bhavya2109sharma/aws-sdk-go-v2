// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified tags from the role. For more information about tagging,
// see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func (c *Client) UntagRole(ctx context.Context, params *UntagRoleInput, optFns ...func(*Options)) (*UntagRoleOutput, error) {
	if params == nil {
		params = &UntagRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UntagRole", params, optFns, c.addOperationUntagRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UntagRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagRoleInput struct {

	// The name of the IAM role from which you want to remove tags.
	//
	// This parameter accepts (through its [regex pattern]) a string of characters that consist of
	// upper and lowercase alphanumeric characters with no spaces. You can also include
	// any of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	RoleName *string

	// A list of key names as a simple array of strings. The tags with matching keys
	// are removed from the specified role.
	//
	// This member is required.
	TagKeys []string

	noSmithyDocumentSerde
}

type UntagRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUntagRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUntagRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUntagRole{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UntagRole"); err != nil {
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
	if err = addOpUntagRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUntagRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUntagRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UntagRole",
	}
}

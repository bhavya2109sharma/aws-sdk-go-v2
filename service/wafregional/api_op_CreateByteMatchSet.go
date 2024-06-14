// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a ByteMatchSet . You then use UpdateByteMatchSet to identify the part of a web request
// that you want AWS WAF to inspect, such as the values of the User-Agent header
// or the query string. For example, you can create a ByteMatchSet that matches
// any requests with User-Agent headers that contain the string BadBot . You can
// then configure AWS WAF to reject those requests.
//
// To create and configure a ByteMatchSet , perform the following steps:
//
//   - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
//     a CreateByteMatchSet request.
//
//   - Submit a CreateByteMatchSet request.
//
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of an UpdateByteMatchSet request.
//
//   - Submit an UpdateByteMatchSetrequest to specify the part of the request that you want AWS WAF
//     to inspect (for example, the header or the URI) and the value that you want AWS
//     WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
func (c *Client) CreateByteMatchSet(ctx context.Context, params *CreateByteMatchSetInput, optFns ...func(*Options)) (*CreateByteMatchSetOutput, error) {
	if params == nil {
		params = &CreateByteMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateByteMatchSet", params, optFns, c.addOperationCreateByteMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateByteMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description of the ByteMatchSet. You can't change Name after you create
	// a ByteMatchSet .
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateByteMatchSetOutput struct {

	// A ByteMatchSet that contains no ByteMatchTuple objects.
	ByteMatchSet *types.ByteMatchSet

	// The ChangeToken that you used to submit the CreateByteMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateByteMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateByteMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateByteMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateByteMatchSet"); err != nil {
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
	if err = addOpCreateByteMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateByteMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateByteMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateByteMatchSet",
	}
}

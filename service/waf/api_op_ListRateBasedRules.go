// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
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
// Returns an array of RuleSummary objects.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
func (c *Client) ListRateBasedRules(ctx context.Context, params *ListRateBasedRulesInput, optFns ...func(*Options)) (*ListRateBasedRulesOutput, error) {
	if params == nil {
		params = &ListRateBasedRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRateBasedRules", params, optFns, c.addOperationListRateBasedRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRateBasedRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRateBasedRulesInput struct {

	// Specifies the number of Rules that you want AWS WAF to return for this request.
	// If you have more Rules than the number that you specify for Limit , the response
	// includes a NextMarker value that you can use to get another batch of Rules .
	Limit int32

	// If you specify a value for Limit and you have more Rules than the value of Limit
	// , AWS WAF returns a NextMarker value in the response that allows you to list
	// another group of Rules . For the second and subsequent ListRateBasedRules
	// requests, specify the value of NextMarker from the previous response to get
	// information about another batch of Rules .
	NextMarker *string

	noSmithyDocumentSerde
}

type ListRateBasedRulesOutput struct {

	// If you have more Rules than the number that you specified for Limit in the
	// request, the response includes a NextMarker value. To list more Rules , submit
	// another ListRateBasedRules request, and specify the NextMarker value from the
	// response in the NextMarker value in the next request.
	NextMarker *string

	// An array of RuleSummary objects.
	Rules []types.RuleSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRateBasedRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRateBasedRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRateBasedRules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRateBasedRules"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRateBasedRules(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListRateBasedRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRateBasedRules",
	}
}

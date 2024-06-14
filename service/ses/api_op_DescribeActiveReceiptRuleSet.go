// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the metadata and receipt rules for the receipt rule set that is
// currently active.
//
// For information about setting up receipt rule sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html#receiving-email-concepts-rules
func (c *Client) DescribeActiveReceiptRuleSet(ctx context.Context, params *DescribeActiveReceiptRuleSetInput, optFns ...func(*Options)) (*DescribeActiveReceiptRuleSetOutput, error) {
	if params == nil {
		params = &DescribeActiveReceiptRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeActiveReceiptRuleSet", params, optFns, c.addOperationDescribeActiveReceiptRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeActiveReceiptRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the metadata and receipt rules for the receipt
// rule set that is currently active. You use receipt rule sets to receive email
// with Amazon SES. For more information, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html
type DescribeActiveReceiptRuleSetInput struct {
	noSmithyDocumentSerde
}

// Represents the metadata and receipt rules for the receipt rule set that is
// currently active.
type DescribeActiveReceiptRuleSetOutput struct {

	// The metadata for the currently active receipt rule set. The metadata consists
	// of the rule set name and a timestamp of when the rule set was created.
	Metadata *types.ReceiptRuleSetMetadata

	// The receipt rules that belong to the active rule set.
	Rules []types.ReceiptRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeActiveReceiptRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeActiveReceiptRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeActiveReceiptRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeActiveReceiptRuleSet"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActiveReceiptRuleSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeActiveReceiptRuleSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeActiveReceiptRuleSet",
	}
}

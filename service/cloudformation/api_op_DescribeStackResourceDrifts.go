// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns drift information for the resources that have been checked for drift in
// the specified stack. This includes actual and expected configuration values for
// resources where CloudFormation detects configuration drift.
//
// For a given stack, there will be one StackResourceDrift for each stack resource
// that has been checked for drift. Resources that haven't yet been checked for
// drift aren't included. Resources that don't currently support drift detection
// aren't checked, and so not included. For a list of resources that support drift
// detection, see [Resources that Support Drift Detection].
//
// Use DetectStackResourceDrift to detect drift on individual resources, or DetectStackDrift to detect drift on all
// supported resources for a given stack.
//
// [Resources that Support Drift Detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift-resource-list.html
func (c *Client) DescribeStackResourceDrifts(ctx context.Context, params *DescribeStackResourceDriftsInput, optFns ...func(*Options)) (*DescribeStackResourceDriftsOutput, error) {
	if params == nil {
		params = &DescribeStackResourceDriftsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStackResourceDrifts", params, optFns, c.addOperationDescribeStackResourceDriftsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStackResourceDriftsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStackResourceDriftsInput struct {

	// The name of the stack for which you want drift information.
	//
	// This member is required.
	StackName *string

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int32

	// A string that identifies the next page of stack resource drift results.
	NextToken *string

	// The resource drift status values to use as filters for the resource drift
	// results returned.
	//
	//   - DELETED : The resource differs from its expected template configuration in
	//   that the resource has been deleted.
	//
	//   - MODIFIED : One or more resource properties differ from their expected
	//   template values.
	//
	//   - IN_SYNC : The resource's actual configuration matches its expected template
	//   configuration.
	//
	//   - NOT_CHECKED : CloudFormation doesn't currently return this value.
	StackResourceDriftStatusFilters []types.StackResourceDriftStatus

	noSmithyDocumentSerde
}

type DescribeStackResourceDriftsOutput struct {

	// Drift information for the resources that have been checked for drift in the
	// specified stack. This includes actual and expected configuration values for
	// resources where CloudFormation detects drift.
	//
	// For a given stack, there will be one StackResourceDrift for each stack resource
	// that has been checked for drift. Resources that haven't yet been checked for
	// drift aren't included. Resources that do not currently support drift detection
	// aren't checked, and so not included. For a list of resources that support drift
	// detection, see [Resources that Support Drift Detection].
	//
	// [Resources that Support Drift Detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift-resource-list.html
	//
	// This member is required.
	StackResourceDrifts []types.StackResourceDrift

	// If the request doesn't return all the remaining results, NextToken is set to a
	// token. To retrieve the next set of results, call DescribeStackResourceDrifts
	// again and assign that token to the request object's NextToken parameter. If the
	// request returns all results, NextToken is set to null .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStackResourceDriftsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackResourceDrifts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackResourceDrifts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStackResourceDrifts"); err != nil {
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
	if err = addOpDescribeStackResourceDriftsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackResourceDrifts(options.Region), middleware.Before); err != nil {
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

// DescribeStackResourceDriftsAPIClient is a client that implements the
// DescribeStackResourceDrifts operation.
type DescribeStackResourceDriftsAPIClient interface {
	DescribeStackResourceDrifts(context.Context, *DescribeStackResourceDriftsInput, ...func(*Options)) (*DescribeStackResourceDriftsOutput, error)
}

var _ DescribeStackResourceDriftsAPIClient = (*Client)(nil)

// DescribeStackResourceDriftsPaginatorOptions is the paginator options for
// DescribeStackResourceDrifts
type DescribeStackResourceDriftsPaginatorOptions struct {
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeStackResourceDriftsPaginator is a paginator for
// DescribeStackResourceDrifts
type DescribeStackResourceDriftsPaginator struct {
	options   DescribeStackResourceDriftsPaginatorOptions
	client    DescribeStackResourceDriftsAPIClient
	params    *DescribeStackResourceDriftsInput
	nextToken *string
	firstPage bool
}

// NewDescribeStackResourceDriftsPaginator returns a new
// DescribeStackResourceDriftsPaginator
func NewDescribeStackResourceDriftsPaginator(client DescribeStackResourceDriftsAPIClient, params *DescribeStackResourceDriftsInput, optFns ...func(*DescribeStackResourceDriftsPaginatorOptions)) *DescribeStackResourceDriftsPaginator {
	if params == nil {
		params = &DescribeStackResourceDriftsInput{}
	}

	options := DescribeStackResourceDriftsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeStackResourceDriftsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeStackResourceDriftsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeStackResourceDrifts page.
func (p *DescribeStackResourceDriftsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeStackResourceDriftsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeStackResourceDrifts(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeStackResourceDrifts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStackResourceDrifts",
	}
}

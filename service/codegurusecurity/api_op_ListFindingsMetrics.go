// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurusecurity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codegurusecurity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metrics about all findings in an account within a specified time range.
func (c *Client) ListFindingsMetrics(ctx context.Context, params *ListFindingsMetricsInput, optFns ...func(*Options)) (*ListFindingsMetricsOutput, error) {
	if params == nil {
		params = &ListFindingsMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFindingsMetrics", params, optFns, c.addOperationListFindingsMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFindingsMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFindingsMetricsInput struct {

	// The end date of the interval which you want to retrieve metrics from. Round to
	// the nearest day.
	//
	// This member is required.
	EndDate *time.Time

	// The start date of the interval which you want to retrieve metrics from. Rounds
	// to the nearest day.
	//
	// This member is required.
	StartDate *time.Time

	// The maximum number of results to return in the response. Use this parameter
	// when paginating results. If additional results exist beyond the number you
	// specify, the nextToken element is returned in the response. Use nextToken in a
	// subsequent request to retrieve additional results. If not specified, returns
	// 1000 results.
	MaxResults *int32

	// A token to use for paginating results that are returned in the response. Set
	// the value of this parameter to null for the first request. For subsequent calls,
	// use the nextToken value returned from the previous request to continue listing
	// results after the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFindingsMetricsOutput struct {

	// A list of AccountFindingsMetric objects retrieved from the specified time
	// interval.
	FindingsMetrics []types.AccountFindingsMetric

	// A pagination token. You can use this in future calls to ListFindingMetrics to
	// continue listing results after the current page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFindingsMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFindingsMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFindingsMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFindingsMetrics"); err != nil {
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
	if err = addOpListFindingsMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFindingsMetrics(options.Region), middleware.Before); err != nil {
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

// ListFindingsMetricsAPIClient is a client that implements the
// ListFindingsMetrics operation.
type ListFindingsMetricsAPIClient interface {
	ListFindingsMetrics(context.Context, *ListFindingsMetricsInput, ...func(*Options)) (*ListFindingsMetricsOutput, error)
}

var _ ListFindingsMetricsAPIClient = (*Client)(nil)

// ListFindingsMetricsPaginatorOptions is the paginator options for
// ListFindingsMetrics
type ListFindingsMetricsPaginatorOptions struct {
	// The maximum number of results to return in the response. Use this parameter
	// when paginating results. If additional results exist beyond the number you
	// specify, the nextToken element is returned in the response. Use nextToken in a
	// subsequent request to retrieve additional results. If not specified, returns
	// 1000 results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFindingsMetricsPaginator is a paginator for ListFindingsMetrics
type ListFindingsMetricsPaginator struct {
	options   ListFindingsMetricsPaginatorOptions
	client    ListFindingsMetricsAPIClient
	params    *ListFindingsMetricsInput
	nextToken *string
	firstPage bool
}

// NewListFindingsMetricsPaginator returns a new ListFindingsMetricsPaginator
func NewListFindingsMetricsPaginator(client ListFindingsMetricsAPIClient, params *ListFindingsMetricsInput, optFns ...func(*ListFindingsMetricsPaginatorOptions)) *ListFindingsMetricsPaginator {
	if params == nil {
		params = &ListFindingsMetricsInput{}
	}

	options := ListFindingsMetricsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFindingsMetricsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFindingsMetricsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFindingsMetrics page.
func (p *ListFindingsMetricsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFindingsMetricsOutput, error) {
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

	result, err := p.client.ListFindingsMetrics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFindingsMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFindingsMetrics",
	}
}

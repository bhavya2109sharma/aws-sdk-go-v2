// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves search metrics data. The data provides a snapshot of how your users
// interact with your search application and how effective the application is.
func (c *Client) GetSnapshots(ctx context.Context, params *GetSnapshotsInput, optFns ...func(*Options)) (*GetSnapshotsOutput, error) {
	if params == nil {
		params = &GetSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSnapshots", params, optFns, c.addOperationGetSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSnapshotsInput struct {

	// The identifier of the index to get search metrics data.
	//
	// This member is required.
	IndexId *string

	// The time interval or time window to get search metrics data. The time interval
	// uses the time zone of your index. You can view data in the following time
	// windows:
	//
	//   - THIS_WEEK : The current week, starting on the Sunday and ending on the day
	//   before the current date.
	//
	//   - ONE_WEEK_AGO : The previous week, starting on the Sunday and ending on the
	//   following Saturday.
	//
	//   - TWO_WEEKS_AGO : The week before the previous week, starting on the Sunday
	//   and ending on the following Saturday.
	//
	//   - THIS_MONTH : The current month, starting on the first day of the month and
	//   ending on the day before the current date.
	//
	//   - ONE_MONTH_AGO : The previous month, starting on the first day of the month
	//   and ending on the last day of the month.
	//
	//   - TWO_MONTHS_AGO : The month before the previous month, starting on the first
	//   day of the month and ending on last day of the month.
	//
	// This member is required.
	Interval types.Interval

	// The metric you want to retrieve. You can specify only one metric per call.
	//
	// For more information about the metrics you can view, see [Gaining insights with search analytics].
	//
	// [Gaining insights with search analytics]: https://docs.aws.amazon.com/kendra/latest/dg/search-analytics.html
	//
	// This member is required.
	MetricType types.MetricType

	// The maximum number of returned data for the metric.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Kendra returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of search metrics data.
	NextToken *string

	noSmithyDocumentSerde
}

type GetSnapshotsOutput struct {

	// If the response is truncated, Amazon Kendra returns this token, which you can
	// use in a later request to retrieve the next set of search metrics data.
	NextToken *string

	// The Unix timestamp for the beginning and end of the time window for the search
	// metrics data.
	SnapShotTimeFilter *types.TimeRange

	// The search metrics data. The data returned depends on the metric type you
	// requested.
	SnapshotsData [][]string

	// The column headers for the search metrics data.
	SnapshotsDataHeader []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSnapshots"); err != nil {
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
	if err = addOpGetSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSnapshots(options.Region), middleware.Before); err != nil {
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

// GetSnapshotsAPIClient is a client that implements the GetSnapshots operation.
type GetSnapshotsAPIClient interface {
	GetSnapshots(context.Context, *GetSnapshotsInput, ...func(*Options)) (*GetSnapshotsOutput, error)
}

var _ GetSnapshotsAPIClient = (*Client)(nil)

// GetSnapshotsPaginatorOptions is the paginator options for GetSnapshots
type GetSnapshotsPaginatorOptions struct {
	// The maximum number of returned data for the metric.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSnapshotsPaginator is a paginator for GetSnapshots
type GetSnapshotsPaginator struct {
	options   GetSnapshotsPaginatorOptions
	client    GetSnapshotsAPIClient
	params    *GetSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewGetSnapshotsPaginator returns a new GetSnapshotsPaginator
func NewGetSnapshotsPaginator(client GetSnapshotsAPIClient, params *GetSnapshotsInput, optFns ...func(*GetSnapshotsPaginatorOptions)) *GetSnapshotsPaginator {
	if params == nil {
		params = &GetSnapshotsInput{}
	}

	options := GetSnapshotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetSnapshots page.
func (p *GetSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSnapshotsOutput, error) {
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

	result, err := p.client.GetSnapshots(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSnapshots",
	}
}

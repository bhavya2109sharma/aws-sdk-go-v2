// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the event types that meet the specified filter criteria. You can use
// this API operation to find information about the Health event, such as the
// category, Amazon Web Service, and event code. The metadata for each event
// appears in the [EventType]object.
//
// If you don't specify a filter criteria, the API operation returns all event
// types, in no particular order.
//
// This API operation uses pagination. Specify the nextToken parameter in the next
// request to return more results.
//
// [EventType]: https://docs.aws.amazon.com/health/latest/APIReference/API_EventType.html
func (c *Client) DescribeEventTypes(ctx context.Context, params *DescribeEventTypesInput, optFns ...func(*Options)) (*DescribeEventTypesOutput, error) {
	if params == nil {
		params = &DescribeEventTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventTypes", params, optFns, c.addOperationDescribeEventTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventTypesInput struct {

	// Values to narrow the results returned.
	Filter *types.EventTypeFilter

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string

	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	//
	// If you don't specify the maxResults parameter, this operation returns a maximum
	// of 30 items by default.
	MaxResults *int32

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeEventTypesOutput struct {

	// A list of event types that match the filter criteria. Event types have a
	// category ( issue , accountNotification , or scheduledChange ), a service (for
	// example, EC2 , RDS , DATAPIPELINE , BILLING ), and a code (in the format
	// AWS_SERVICE_DESCRIPTION ; for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT ).
	EventTypes []types.EventType

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEventTypes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventTypes(options.Region), middleware.Before); err != nil {
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

// DescribeEventTypesAPIClient is a client that implements the DescribeEventTypes
// operation.
type DescribeEventTypesAPIClient interface {
	DescribeEventTypes(context.Context, *DescribeEventTypesInput, ...func(*Options)) (*DescribeEventTypesOutput, error)
}

var _ DescribeEventTypesAPIClient = (*Client)(nil)

// DescribeEventTypesPaginatorOptions is the paginator options for
// DescribeEventTypes
type DescribeEventTypesPaginatorOptions struct {
	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	//
	// If you don't specify the maxResults parameter, this operation returns a maximum
	// of 30 items by default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEventTypesPaginator is a paginator for DescribeEventTypes
type DescribeEventTypesPaginator struct {
	options   DescribeEventTypesPaginatorOptions
	client    DescribeEventTypesAPIClient
	params    *DescribeEventTypesInput
	nextToken *string
	firstPage bool
}

// NewDescribeEventTypesPaginator returns a new DescribeEventTypesPaginator
func NewDescribeEventTypesPaginator(client DescribeEventTypesAPIClient, params *DescribeEventTypesInput, optFns ...func(*DescribeEventTypesPaginatorOptions)) *DescribeEventTypesPaginator {
	if params == nil {
		params = &DescribeEventTypesInput{}
	}

	options := DescribeEventTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEventTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEventTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEventTypes page.
func (p *DescribeEventTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEventTypesOutput, error) {
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

	result, err := p.client.DescribeEventTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEventTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEventTypes",
	}
}

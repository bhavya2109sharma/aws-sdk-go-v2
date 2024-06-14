// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the stored queries for a single Amazon Web Services account and a single
// Amazon Web Services Region. The default is 100.
func (c *Client) ListStoredQueries(ctx context.Context, params *ListStoredQueriesInput, optFns ...func(*Options)) (*ListStoredQueriesOutput, error) {
	if params == nil {
		params = &ListStoredQueriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStoredQueries", params, optFns, c.addOperationListStoredQueriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStoredQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStoredQueriesInput struct {

	// The maximum number of results to be returned with a single call.
	MaxResults *int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStoredQueriesOutput struct {

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call this action again and assign that token to the
	// request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null .
	NextToken *string

	// A list of StoredQueryMetadata objects.
	StoredQueryMetadata []types.StoredQueryMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStoredQueriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStoredQueries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStoredQueries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStoredQueries"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStoredQueries(options.Region), middleware.Before); err != nil {
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

// ListStoredQueriesAPIClient is a client that implements the ListStoredQueries
// operation.
type ListStoredQueriesAPIClient interface {
	ListStoredQueries(context.Context, *ListStoredQueriesInput, ...func(*Options)) (*ListStoredQueriesOutput, error)
}

var _ ListStoredQueriesAPIClient = (*Client)(nil)

// ListStoredQueriesPaginatorOptions is the paginator options for ListStoredQueries
type ListStoredQueriesPaginatorOptions struct {
	// The maximum number of results to be returned with a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStoredQueriesPaginator is a paginator for ListStoredQueries
type ListStoredQueriesPaginator struct {
	options   ListStoredQueriesPaginatorOptions
	client    ListStoredQueriesAPIClient
	params    *ListStoredQueriesInput
	nextToken *string
	firstPage bool
}

// NewListStoredQueriesPaginator returns a new ListStoredQueriesPaginator
func NewListStoredQueriesPaginator(client ListStoredQueriesAPIClient, params *ListStoredQueriesInput, optFns ...func(*ListStoredQueriesPaginatorOptions)) *ListStoredQueriesPaginator {
	if params == nil {
		params = &ListStoredQueriesInput{}
	}

	options := ListStoredQueriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStoredQueriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStoredQueriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStoredQueries page.
func (p *ListStoredQueriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStoredQueriesOutput, error) {
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

	result, err := p.client.ListStoredQueries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStoredQueries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStoredQueries",
	}
}

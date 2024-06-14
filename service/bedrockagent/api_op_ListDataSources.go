// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the data sources in a knowledge base and information about each one.
func (c *Client) ListDataSources(ctx context.Context, params *ListDataSourcesInput, optFns ...func(*Options)) (*ListDataSourcesOutput, error) {
	if params == nil {
		params = &ListDataSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSources", params, optFns, c.addOperationListDataSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSourcesInput struct {

	// The unique identifier of the knowledge base for which to return a list of
	// information.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	MaxResults *int32

	// If the total number of results is greater than the maxResults value provided in
	// the request, enter the token returned in the nextToken field in the response in
	// this field to return the next batch of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataSourcesOutput struct {

	// A list of objects, each of which contains information about a data source.
	//
	// This member is required.
	DataSourceSummaries []types.DataSourceSummary

	// If the total number of results is greater than the maxResults value provided in
	// the request, use this token when making another request in the nextToken field
	// to return the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataSources"); err != nil {
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
	if err = addOpListDataSourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSources(options.Region), middleware.Before); err != nil {
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

// ListDataSourcesAPIClient is a client that implements the ListDataSources
// operation.
type ListDataSourcesAPIClient interface {
	ListDataSources(context.Context, *ListDataSourcesInput, ...func(*Options)) (*ListDataSourcesOutput, error)
}

var _ ListDataSourcesAPIClient = (*Client)(nil)

// ListDataSourcesPaginatorOptions is the paginator options for ListDataSources
type ListDataSourcesPaginatorOptions struct {
	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataSourcesPaginator is a paginator for ListDataSources
type ListDataSourcesPaginator struct {
	options   ListDataSourcesPaginatorOptions
	client    ListDataSourcesAPIClient
	params    *ListDataSourcesInput
	nextToken *string
	firstPage bool
}

// NewListDataSourcesPaginator returns a new ListDataSourcesPaginator
func NewListDataSourcesPaginator(client ListDataSourcesAPIClient, params *ListDataSourcesInput, optFns ...func(*ListDataSourcesPaginatorOptions)) *ListDataSourcesPaginator {
	if params == nil {
		params = &ListDataSourcesInput{}
	}

	options := ListDataSourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataSourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataSourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataSources page.
func (p *ListDataSourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataSourcesOutput, error) {
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

	result, err := p.client.ListDataSources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataSources",
	}
}

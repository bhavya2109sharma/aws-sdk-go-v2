// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of media pipelines.
func (c *Client) ListMediaPipelines(ctx context.Context, params *ListMediaPipelinesInput, optFns ...func(*Options)) (*ListMediaPipelinesOutput, error) {
	if params == nil {
		params = &ListMediaPipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMediaPipelines", params, optFns, c.addOperationListMediaPipelinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMediaPipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMediaPipelinesInput struct {

	// The maximum number of results to return in a single call. Valid Range: 1 - 99.
	MaxResults *int32

	// The token used to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMediaPipelinesOutput struct {

	// The media pipeline objects in the list.
	MediaPipelines []types.MediaPipelineSummary

	// The token used to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMediaPipelinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMediaPipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMediaPipelines{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMediaPipelines"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMediaPipelines(options.Region), middleware.Before); err != nil {
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

// ListMediaPipelinesAPIClient is a client that implements the ListMediaPipelines
// operation.
type ListMediaPipelinesAPIClient interface {
	ListMediaPipelines(context.Context, *ListMediaPipelinesInput, ...func(*Options)) (*ListMediaPipelinesOutput, error)
}

var _ ListMediaPipelinesAPIClient = (*Client)(nil)

// ListMediaPipelinesPaginatorOptions is the paginator options for
// ListMediaPipelines
type ListMediaPipelinesPaginatorOptions struct {
	// The maximum number of results to return in a single call. Valid Range: 1 - 99.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMediaPipelinesPaginator is a paginator for ListMediaPipelines
type ListMediaPipelinesPaginator struct {
	options   ListMediaPipelinesPaginatorOptions
	client    ListMediaPipelinesAPIClient
	params    *ListMediaPipelinesInput
	nextToken *string
	firstPage bool
}

// NewListMediaPipelinesPaginator returns a new ListMediaPipelinesPaginator
func NewListMediaPipelinesPaginator(client ListMediaPipelinesAPIClient, params *ListMediaPipelinesInput, optFns ...func(*ListMediaPipelinesPaginatorOptions)) *ListMediaPipelinesPaginator {
	if params == nil {
		params = &ListMediaPipelinesInput{}
	}

	options := ListMediaPipelinesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMediaPipelinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMediaPipelinesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMediaPipelines page.
func (p *ListMediaPipelinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMediaPipelinesOutput, error) {
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

	result, err := p.client.ListMediaPipelines(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMediaPipelines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMediaPipelines",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation lists your jobs sorted by CreatedAt in descending order.
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if params == nil {
		params = &ListJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobs", params, optFns, c.addOperationListJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobsInput struct {

	// The unique identifier for a data set.
	DataSetId *string

	// The maximum number of results returned by a single call.
	MaxResults int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// The unique identifier for a revision.
	RevisionId *string

	noSmithyDocumentSerde
}

type ListJobsOutput struct {

	// The jobs listed by the request.
	Jobs []types.JobEntry

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobs(options.Region), middleware.Before); err != nil {
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

// ListJobsAPIClient is a client that implements the ListJobs operation.
type ListJobsAPIClient interface {
	ListJobs(context.Context, *ListJobsInput, ...func(*Options)) (*ListJobsOutput, error)
}

var _ ListJobsAPIClient = (*Client)(nil)

// ListJobsPaginatorOptions is the paginator options for ListJobs
type ListJobsPaginatorOptions struct {
	// The maximum number of results returned by a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobsPaginator is a paginator for ListJobs
type ListJobsPaginator struct {
	options   ListJobsPaginatorOptions
	client    ListJobsAPIClient
	params    *ListJobsInput
	nextToken *string
	firstPage bool
}

// NewListJobsPaginator returns a new ListJobsPaginator
func NewListJobsPaginator(client ListJobsAPIClient, params *ListJobsInput, optFns ...func(*ListJobsPaginatorOptions)) *ListJobsPaginator {
	if params == nil {
		params = &ListJobsInput{}
	}

	options := ListJobsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListJobs page.
func (p *ListJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListJobs",
	}
}

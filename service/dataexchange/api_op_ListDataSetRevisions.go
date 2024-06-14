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

// This operation lists a data set's revisions sorted by CreatedAt in descending
// order.
func (c *Client) ListDataSetRevisions(ctx context.Context, params *ListDataSetRevisionsInput, optFns ...func(*Options)) (*ListDataSetRevisionsOutput, error) {
	if params == nil {
		params = &ListDataSetRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSetRevisions", params, optFns, c.addOperationListDataSetRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSetRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSetRevisionsInput struct {

	// The unique identifier for a data set.
	//
	// This member is required.
	DataSetId *string

	// The maximum number of results returned by a single call.
	MaxResults int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataSetRevisionsOutput struct {

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// The asset objects listed by the request.
	Revisions []types.RevisionEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataSetRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataSetRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSetRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataSetRevisions"); err != nil {
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
	if err = addOpListDataSetRevisionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSetRevisions(options.Region), middleware.Before); err != nil {
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

// ListDataSetRevisionsAPIClient is a client that implements the
// ListDataSetRevisions operation.
type ListDataSetRevisionsAPIClient interface {
	ListDataSetRevisions(context.Context, *ListDataSetRevisionsInput, ...func(*Options)) (*ListDataSetRevisionsOutput, error)
}

var _ ListDataSetRevisionsAPIClient = (*Client)(nil)

// ListDataSetRevisionsPaginatorOptions is the paginator options for
// ListDataSetRevisions
type ListDataSetRevisionsPaginatorOptions struct {
	// The maximum number of results returned by a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataSetRevisionsPaginator is a paginator for ListDataSetRevisions
type ListDataSetRevisionsPaginator struct {
	options   ListDataSetRevisionsPaginatorOptions
	client    ListDataSetRevisionsAPIClient
	params    *ListDataSetRevisionsInput
	nextToken *string
	firstPage bool
}

// NewListDataSetRevisionsPaginator returns a new ListDataSetRevisionsPaginator
func NewListDataSetRevisionsPaginator(client ListDataSetRevisionsAPIClient, params *ListDataSetRevisionsInput, optFns ...func(*ListDataSetRevisionsPaginatorOptions)) *ListDataSetRevisionsPaginator {
	if params == nil {
		params = &ListDataSetRevisionsInput{}
	}

	options := ListDataSetRevisionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataSetRevisionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataSetRevisionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataSetRevisions page.
func (p *ListDataSetRevisionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataSetRevisionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDataSetRevisions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataSetRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataSetRevisions",
	}
}

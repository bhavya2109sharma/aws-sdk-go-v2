// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists LF-tags that the requester has permission to view.
func (c *Client) ListLFTags(ctx context.Context, params *ListLFTagsInput, optFns ...func(*Options)) (*ListLFTagsOutput, error) {
	if params == nil {
		params = &ListLFTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLFTags", params, optFns, c.addOperationListLFTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLFTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLFTagsInput struct {

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	// If resource share type is ALL , returns both in-account LF-tags and shared
	// LF-tags that the requester has permission to view. If resource share type is
	// FOREIGN , returns all share LF-tags that the requester can view. If no resource
	// share type is passed, lists LF-tags in the given catalog ID that the requester
	// has permission to view.
	ResourceShareType types.ResourceShareType

	noSmithyDocumentSerde
}

type ListLFTagsOutput struct {

	// A list of LF-tags that the requested has permission to view.
	LFTags []types.LFTagPair

	// A continuation token, present if the current list segment is not the last.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLFTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLFTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLFTags{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLFTags"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLFTags(options.Region), middleware.Before); err != nil {
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

// ListLFTagsAPIClient is a client that implements the ListLFTags operation.
type ListLFTagsAPIClient interface {
	ListLFTags(context.Context, *ListLFTagsInput, ...func(*Options)) (*ListLFTagsOutput, error)
}

var _ ListLFTagsAPIClient = (*Client)(nil)

// ListLFTagsPaginatorOptions is the paginator options for ListLFTags
type ListLFTagsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLFTagsPaginator is a paginator for ListLFTags
type ListLFTagsPaginator struct {
	options   ListLFTagsPaginatorOptions
	client    ListLFTagsAPIClient
	params    *ListLFTagsInput
	nextToken *string
	firstPage bool
}

// NewListLFTagsPaginator returns a new ListLFTagsPaginator
func NewListLFTagsPaginator(client ListLFTagsAPIClient, params *ListLFTagsInput, optFns ...func(*ListLFTagsPaginatorOptions)) *ListLFTagsPaginator {
	if params == nil {
		params = &ListLFTagsInput{}
	}

	options := ListLFTagsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLFTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLFTagsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLFTags page.
func (p *ListLFTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLFTagsOutput, error) {
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

	result, err := p.client.ListLFTags(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLFTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLFTags",
	}
}

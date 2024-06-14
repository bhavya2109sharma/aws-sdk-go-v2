// Code generated by smithy-go-codegen DO NOT EDIT.

package mailmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mailmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of email archive search jobs.
func (c *Client) ListArchiveSearches(ctx context.Context, params *ListArchiveSearchesInput, optFns ...func(*Options)) (*ListArchiveSearchesOutput, error) {
	if params == nil {
		params = &ListArchiveSearchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListArchiveSearches", params, optFns, c.addOperationListArchiveSearchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListArchiveSearchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list archive search jobs in your account.
type ListArchiveSearchesInput struct {

	// The identifier of the archive.
	//
	// This member is required.
	ArchiveId *string

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// The maximum number of archive search jobs that are returned per call. You can
	// use NextToken to obtain further pages of archives.
	PageSize *int32

	noSmithyDocumentSerde
}

// The response containing a list of archive search jobs and their statuses.
type ListArchiveSearchesOutput struct {

	// If present, use to retrieve the next page of results.
	NextToken *string

	// The list of search job identifiers and statuses.
	Searches []types.SearchSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListArchiveSearchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListArchiveSearches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListArchiveSearches{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListArchiveSearches"); err != nil {
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
	if err = addOpListArchiveSearchesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListArchiveSearches(options.Region), middleware.Before); err != nil {
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

// ListArchiveSearchesAPIClient is a client that implements the
// ListArchiveSearches operation.
type ListArchiveSearchesAPIClient interface {
	ListArchiveSearches(context.Context, *ListArchiveSearchesInput, ...func(*Options)) (*ListArchiveSearchesOutput, error)
}

var _ ListArchiveSearchesAPIClient = (*Client)(nil)

// ListArchiveSearchesPaginatorOptions is the paginator options for
// ListArchiveSearches
type ListArchiveSearchesPaginatorOptions struct {
	// The maximum number of archive search jobs that are returned per call. You can
	// use NextToken to obtain further pages of archives.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListArchiveSearchesPaginator is a paginator for ListArchiveSearches
type ListArchiveSearchesPaginator struct {
	options   ListArchiveSearchesPaginatorOptions
	client    ListArchiveSearchesAPIClient
	params    *ListArchiveSearchesInput
	nextToken *string
	firstPage bool
}

// NewListArchiveSearchesPaginator returns a new ListArchiveSearchesPaginator
func NewListArchiveSearchesPaginator(client ListArchiveSearchesAPIClient, params *ListArchiveSearchesInput, optFns ...func(*ListArchiveSearchesPaginatorOptions)) *ListArchiveSearchesPaginator {
	if params == nil {
		params = &ListArchiveSearchesInput{}
	}

	options := ListArchiveSearchesPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListArchiveSearchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListArchiveSearchesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListArchiveSearches page.
func (p *ListArchiveSearchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListArchiveSearchesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListArchiveSearches(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListArchiveSearches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListArchiveSearches",
	}
}

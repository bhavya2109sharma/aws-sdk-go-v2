// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the databases in the specified data catalog.
func (c *Client) ListDatabases(ctx context.Context, params *ListDatabasesInput, optFns ...func(*Options)) (*ListDatabasesOutput, error) {
	if params == nil {
		params = &ListDatabasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatabases", params, optFns, c.addOperationListDatabasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatabasesInput struct {

	// The name of the data catalog that contains the databases to return.
	//
	// This member is required.
	CatalogName *string

	// Specifies the maximum number of results to return.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// The name of the workgroup for which the metadata is being fetched. Required if
	// requesting an IAM Identity Center enabled Glue Data Catalog.
	WorkGroup *string

	noSmithyDocumentSerde
}

type ListDatabasesOutput struct {

	// A list of databases from a data catalog.
	DatabaseList []types.Database

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatabasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDatabases"); err != nil {
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
	if err = addOpListDatabasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatabases(options.Region), middleware.Before); err != nil {
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

// ListDatabasesAPIClient is a client that implements the ListDatabases operation.
type ListDatabasesAPIClient interface {
	ListDatabases(context.Context, *ListDatabasesInput, ...func(*Options)) (*ListDatabasesOutput, error)
}

var _ ListDatabasesAPIClient = (*Client)(nil)

// ListDatabasesPaginatorOptions is the paginator options for ListDatabases
type ListDatabasesPaginatorOptions struct {
	// Specifies the maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDatabasesPaginator is a paginator for ListDatabases
type ListDatabasesPaginator struct {
	options   ListDatabasesPaginatorOptions
	client    ListDatabasesAPIClient
	params    *ListDatabasesInput
	nextToken *string
	firstPage bool
}

// NewListDatabasesPaginator returns a new ListDatabasesPaginator
func NewListDatabasesPaginator(client ListDatabasesAPIClient, params *ListDatabasesInput, optFns ...func(*ListDatabasesPaginatorOptions)) *ListDatabasesPaginator {
	if params == nil {
		params = &ListDatabasesInput{}
	}

	options := ListDatabasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDatabasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDatabasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDatabases page.
func (p *ListDatabasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDatabasesOutput, error) {
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

	result, err := p.client.ListDatabases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDatabases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDatabases",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all groups in the identity store. Returns a paginated list of complete
// Group objects. Filtering for a Group by the DisplayName attribute is
// deprecated. Instead, use the GetGroupId API action.
//
// If you have administrator access to a member account, you can use this API from
// the member account. Read about [member accounts]in the Organizations User Guide.
//
// [member accounts]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html
func (c *Client) ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) {
	if params == nil {
		params = &ListGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroups", params, optFns, c.addOperationListGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupsInput struct {

	// The globally unique identifier for the identity store, such as d-1234567890 . In
	// this example, d- is a fixed prefix, and 1234567890 is a randomly generated
	// string that contains numbers and lower case letters. This value is generated at
	// the time that a new identity store is created.
	//
	// This member is required.
	IdentityStoreId *string

	// A list of Filter objects, which is used in the ListUsers and ListGroups
	// requests.
	//
	// Deprecated: Using filters with ListGroups API is deprecated, please use
	// GetGroupId API instead.
	Filters []types.Filter

	// The maximum number of results to be returned per request. This parameter is
	// used in the ListUsers and ListGroups requests to specify how many results to
	// return in one page. The length limit is 50 characters.
	MaxResults *int32

	// The pagination token used for the ListUsers and ListGroups API operations. This
	// value is generated by the identity store service. It is returned in the API
	// response if the total results are more than the size of one page. This token is
	// also returned when it is used in the API request to search for the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupsOutput struct {

	// A list of Group objects in the identity store.
	//
	// This member is required.
	Groups []types.Group

	// The pagination token used for the ListUsers and ListGroups API operations. This
	// value is generated by the identity store service. It is returned in the API
	// response if the total results are more than the size of one page. This token is
	// also returned when it1 is used in the API request to search for the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGroups"); err != nil {
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
	if err = addOpListGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroups(options.Region), middleware.Before); err != nil {
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

// ListGroupsAPIClient is a client that implements the ListGroups operation.
type ListGroupsAPIClient interface {
	ListGroups(context.Context, *ListGroupsInput, ...func(*Options)) (*ListGroupsOutput, error)
}

var _ ListGroupsAPIClient = (*Client)(nil)

// ListGroupsPaginatorOptions is the paginator options for ListGroups
type ListGroupsPaginatorOptions struct {
	// The maximum number of results to be returned per request. This parameter is
	// used in the ListUsers and ListGroups requests to specify how many results to
	// return in one page. The length limit is 50 characters.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupsPaginator is a paginator for ListGroups
type ListGroupsPaginator struct {
	options   ListGroupsPaginatorOptions
	client    ListGroupsAPIClient
	params    *ListGroupsInput
	nextToken *string
	firstPage bool
}

// NewListGroupsPaginator returns a new ListGroupsPaginator
func NewListGroupsPaginator(client ListGroupsAPIClient, params *ListGroupsInput, optFns ...func(*ListGroupsPaginatorOptions)) *ListGroupsPaginator {
	if params == nil {
		params = &ListGroupsInput{}
	}

	options := ListGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroups page.
func (p *ListGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupsOutput, error) {
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

	result, err := p.client.ListGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGroups",
	}
}

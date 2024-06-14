// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of ProtocolsListDataSummary objects.
func (c *Client) ListProtocolsLists(ctx context.Context, params *ListProtocolsListsInput, optFns ...func(*Options)) (*ListProtocolsListsOutput, error) {
	if params == nil {
		params = &ListProtocolsListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProtocolsLists", params, optFns, c.addOperationListProtocolsListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProtocolsListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProtocolsListsInput struct {

	// The maximum number of objects that you want Firewall Manager to return for this
	// request. If more objects are available, in the response, Firewall Manager
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	//
	// If you don't specify this, Firewall Manager returns all available objects.
	//
	// This member is required.
	MaxResults *int32

	// Specifies whether the lists to retrieve are default lists owned by Firewall
	// Manager.
	DefaultLists bool

	// If you specify a value for MaxResults in your list request, and you have more
	// objects than the maximum, Firewall Manager returns this token in the response.
	// For all but the first request, you provide the token returned by the prior
	// request in the request parameters, to retrieve the next batch of objects.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProtocolsListsOutput struct {

	// If you specify a value for MaxResults in your list request, and you have more
	// objects than the maximum, Firewall Manager returns this token in the response.
	// You can use this token in subsequent requests to retrieve the next batch of
	// objects.
	NextToken *string

	// An array of ProtocolsListDataSummary objects.
	ProtocolsLists []types.ProtocolsListDataSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProtocolsListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProtocolsLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProtocolsLists{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProtocolsLists"); err != nil {
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
	if err = addOpListProtocolsListsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProtocolsLists(options.Region), middleware.Before); err != nil {
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

// ListProtocolsListsAPIClient is a client that implements the ListProtocolsLists
// operation.
type ListProtocolsListsAPIClient interface {
	ListProtocolsLists(context.Context, *ListProtocolsListsInput, ...func(*Options)) (*ListProtocolsListsOutput, error)
}

var _ ListProtocolsListsAPIClient = (*Client)(nil)

// ListProtocolsListsPaginatorOptions is the paginator options for
// ListProtocolsLists
type ListProtocolsListsPaginatorOptions struct {
	// The maximum number of objects that you want Firewall Manager to return for this
	// request. If more objects are available, in the response, Firewall Manager
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	//
	// If you don't specify this, Firewall Manager returns all available objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProtocolsListsPaginator is a paginator for ListProtocolsLists
type ListProtocolsListsPaginator struct {
	options   ListProtocolsListsPaginatorOptions
	client    ListProtocolsListsAPIClient
	params    *ListProtocolsListsInput
	nextToken *string
	firstPage bool
}

// NewListProtocolsListsPaginator returns a new ListProtocolsListsPaginator
func NewListProtocolsListsPaginator(client ListProtocolsListsAPIClient, params *ListProtocolsListsInput, optFns ...func(*ListProtocolsListsPaginatorOptions)) *ListProtocolsListsPaginator {
	if params == nil {
		params = &ListProtocolsListsInput{}
	}

	options := ListProtocolsListsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProtocolsListsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProtocolsListsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProtocolsLists page.
func (p *ListProtocolsListsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProtocolsListsOutput, error) {
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

	result, err := p.client.ListProtocolsLists(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProtocolsLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProtocolsLists",
	}
}

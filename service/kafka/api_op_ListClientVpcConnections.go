// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the VPC connections in this Region.
func (c *Client) ListClientVpcConnections(ctx context.Context, params *ListClientVpcConnectionsInput, optFns ...func(*Options)) (*ListClientVpcConnectionsOutput, error) {
	if params == nil {
		params = &ListClientVpcConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClientVpcConnections", params, optFns, c.addOperationListClientVpcConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClientVpcConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClientVpcConnectionsInput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	//
	// This member is required.
	ClusterArn *string

	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	MaxResults *int32

	// The paginated results marker. When the result of the operation is truncated,
	// the call returns NextToken in the response. To get the next batch, provide this
	// token in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListClientVpcConnectionsOutput struct {

	// List of client VPC connections.
	ClientVpcConnections []types.ClientVpcConnection

	// The paginated results marker. When the result of a ListClientVpcConnections
	// operation is truncated, the call returns NextToken in the response. To get
	// another batch of configurations, provide this token in your next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClientVpcConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListClientVpcConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListClientVpcConnections{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListClientVpcConnections"); err != nil {
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
	if err = addOpListClientVpcConnectionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClientVpcConnections(options.Region), middleware.Before); err != nil {
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

// ListClientVpcConnectionsAPIClient is a client that implements the
// ListClientVpcConnections operation.
type ListClientVpcConnectionsAPIClient interface {
	ListClientVpcConnections(context.Context, *ListClientVpcConnectionsInput, ...func(*Options)) (*ListClientVpcConnectionsOutput, error)
}

var _ ListClientVpcConnectionsAPIClient = (*Client)(nil)

// ListClientVpcConnectionsPaginatorOptions is the paginator options for
// ListClientVpcConnections
type ListClientVpcConnectionsPaginatorOptions struct {
	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClientVpcConnectionsPaginator is a paginator for ListClientVpcConnections
type ListClientVpcConnectionsPaginator struct {
	options   ListClientVpcConnectionsPaginatorOptions
	client    ListClientVpcConnectionsAPIClient
	params    *ListClientVpcConnectionsInput
	nextToken *string
	firstPage bool
}

// NewListClientVpcConnectionsPaginator returns a new
// ListClientVpcConnectionsPaginator
func NewListClientVpcConnectionsPaginator(client ListClientVpcConnectionsAPIClient, params *ListClientVpcConnectionsInput, optFns ...func(*ListClientVpcConnectionsPaginatorOptions)) *ListClientVpcConnectionsPaginator {
	if params == nil {
		params = &ListClientVpcConnectionsInput{}
	}

	options := ListClientVpcConnectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListClientVpcConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClientVpcConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListClientVpcConnections page.
func (p *ListClientVpcConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClientVpcConnectionsOutput, error) {
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

	result, err := p.client.ListClientVpcConnections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListClientVpcConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListClientVpcConnections",
	}
}

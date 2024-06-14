// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more global networks. By default, all global networks are
// described. To describe the objects in your global network, you must use the
// appropriate Get* action. For example, to list the transit gateways in your
// global network, use GetTransitGatewayRegistrations.
func (c *Client) DescribeGlobalNetworks(ctx context.Context, params *DescribeGlobalNetworksInput, optFns ...func(*Options)) (*DescribeGlobalNetworksOutput, error) {
	if params == nil {
		params = &DescribeGlobalNetworksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGlobalNetworks", params, optFns, c.addOperationDescribeGlobalNetworksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGlobalNetworksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalNetworksInput struct {

	// The IDs of one or more global networks. The maximum is 10.
	GlobalNetworkIds []string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeGlobalNetworksOutput struct {

	// Information about the global networks.
	GlobalNetworks []types.GlobalNetwork

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGlobalNetworksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeGlobalNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeGlobalNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeGlobalNetworks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalNetworks(options.Region), middleware.Before); err != nil {
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

// DescribeGlobalNetworksAPIClient is a client that implements the
// DescribeGlobalNetworks operation.
type DescribeGlobalNetworksAPIClient interface {
	DescribeGlobalNetworks(context.Context, *DescribeGlobalNetworksInput, ...func(*Options)) (*DescribeGlobalNetworksOutput, error)
}

var _ DescribeGlobalNetworksAPIClient = (*Client)(nil)

// DescribeGlobalNetworksPaginatorOptions is the paginator options for
// DescribeGlobalNetworks
type DescribeGlobalNetworksPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeGlobalNetworksPaginator is a paginator for DescribeGlobalNetworks
type DescribeGlobalNetworksPaginator struct {
	options   DescribeGlobalNetworksPaginatorOptions
	client    DescribeGlobalNetworksAPIClient
	params    *DescribeGlobalNetworksInput
	nextToken *string
	firstPage bool
}

// NewDescribeGlobalNetworksPaginator returns a new DescribeGlobalNetworksPaginator
func NewDescribeGlobalNetworksPaginator(client DescribeGlobalNetworksAPIClient, params *DescribeGlobalNetworksInput, optFns ...func(*DescribeGlobalNetworksPaginatorOptions)) *DescribeGlobalNetworksPaginator {
	if params == nil {
		params = &DescribeGlobalNetworksInput{}
	}

	options := DescribeGlobalNetworksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeGlobalNetworksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeGlobalNetworksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeGlobalNetworks page.
func (p *DescribeGlobalNetworksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeGlobalNetworksOutput, error) {
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

	result, err := p.client.DescribeGlobalNetworks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeGlobalNetworks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeGlobalNetworks",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about all provisioned clusters if no cluster identifier is
// specified, or about a specific cluster if a cluster name is supplied.
func (c *Client) DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
	if params == nil {
		params = &DescribeClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusters", params, optFns, c.addOperationDescribeClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClustersInput struct {

	// The name of the cluster
	ClusterName *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// An optional flag that can be included in the request to retrieve information
	// about the individual shard(s).
	ShowShardDetails *bool

	noSmithyDocumentSerde
}

type DescribeClustersOutput struct {

	// A list of clusters
	Clusters []types.Cluster

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeClusters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeClusters"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusters(options.Region), middleware.Before); err != nil {
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

// DescribeClustersAPIClient is a client that implements the DescribeClusters
// operation.
type DescribeClustersAPIClient interface {
	DescribeClusters(context.Context, *DescribeClustersInput, ...func(*Options)) (*DescribeClustersOutput, error)
}

var _ DescribeClustersAPIClient = (*Client)(nil)

// DescribeClustersPaginatorOptions is the paginator options for DescribeClusters
type DescribeClustersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClustersPaginator is a paginator for DescribeClusters
type DescribeClustersPaginator struct {
	options   DescribeClustersPaginatorOptions
	client    DescribeClustersAPIClient
	params    *DescribeClustersInput
	nextToken *string
	firstPage bool
}

// NewDescribeClustersPaginator returns a new DescribeClustersPaginator
func NewDescribeClustersPaginator(client DescribeClustersAPIClient, params *DescribeClustersInput, optFns ...func(*DescribeClustersPaginatorOptions)) *DescribeClustersPaginator {
	if params == nil {
		params = &DescribeClustersInput{}
	}

	options := DescribeClustersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClustersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeClusters page.
func (p *DescribeClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
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

	result, err := p.client.DescribeClusters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeClusters",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about DB proxy endpoints.
func (c *Client) DescribeDBProxyEndpoints(ctx context.Context, params *DescribeDBProxyEndpointsInput, optFns ...func(*Options)) (*DescribeDBProxyEndpointsOutput, error) {
	if params == nil {
		params = &DescribeDBProxyEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBProxyEndpoints", params, optFns, c.addOperationDescribeDBProxyEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBProxyEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBProxyEndpointsInput struct {

	// The name of a DB proxy endpoint to describe. If you omit this parameter, the
	// output includes information about all DB proxy endpoints associated with the
	// specified proxy.
	DBProxyEndpointName *string

	// The name of the DB proxy whose endpoints you want to describe. If you omit this
	// parameter, the output includes information about all DB proxy endpoints
	// associated with all your DB proxies.
	DBProxyName *string

	// This parameter is not currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeDBProxyEndpointsOutput struct {

	// The list of ProxyEndpoint objects returned by the API operation.
	DBProxyEndpoints []types.DBProxyEndpoint

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBProxyEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBProxyEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBProxyEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBProxyEndpoints"); err != nil {
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
	if err = addOpDescribeDBProxyEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBProxyEndpoints(options.Region), middleware.Before); err != nil {
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

// DescribeDBProxyEndpointsAPIClient is a client that implements the
// DescribeDBProxyEndpoints operation.
type DescribeDBProxyEndpointsAPIClient interface {
	DescribeDBProxyEndpoints(context.Context, *DescribeDBProxyEndpointsInput, ...func(*Options)) (*DescribeDBProxyEndpointsOutput, error)
}

var _ DescribeDBProxyEndpointsAPIClient = (*Client)(nil)

// DescribeDBProxyEndpointsPaginatorOptions is the paginator options for
// DescribeDBProxyEndpoints
type DescribeDBProxyEndpointsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBProxyEndpointsPaginator is a paginator for DescribeDBProxyEndpoints
type DescribeDBProxyEndpointsPaginator struct {
	options   DescribeDBProxyEndpointsPaginatorOptions
	client    DescribeDBProxyEndpointsAPIClient
	params    *DescribeDBProxyEndpointsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBProxyEndpointsPaginator returns a new
// DescribeDBProxyEndpointsPaginator
func NewDescribeDBProxyEndpointsPaginator(client DescribeDBProxyEndpointsAPIClient, params *DescribeDBProxyEndpointsInput, optFns ...func(*DescribeDBProxyEndpointsPaginatorOptions)) *DescribeDBProxyEndpointsPaginator {
	if params == nil {
		params = &DescribeDBProxyEndpointsInput{}
	}

	options := DescribeDBProxyEndpointsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBProxyEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBProxyEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBProxyEndpoints page.
func (p *DescribeDBProxyEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBProxyEndpointsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBProxyEndpoints(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBProxyEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBProxyEndpoints",
	}
}

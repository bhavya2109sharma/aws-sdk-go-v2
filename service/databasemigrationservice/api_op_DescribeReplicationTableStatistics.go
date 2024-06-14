// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns table and schema statistics for one or more provisioned replications
// that use a given DMS Serverless replication configuration.
func (c *Client) DescribeReplicationTableStatistics(ctx context.Context, params *DescribeReplicationTableStatisticsInput, optFns ...func(*Options)) (*DescribeReplicationTableStatisticsOutput, error) {
	if params == nil {
		params = &DescribeReplicationTableStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationTableStatistics", params, optFns, c.addOperationDescribeReplicationTableStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationTableStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplicationTableStatisticsInput struct {

	// The replication config to describe.
	//
	// This member is required.
	ReplicationConfigArn *string

	// Filters applied to the replication table statistics.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeReplicationTableStatisticsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The Amazon Resource Name of the replication config.
	ReplicationConfigArn *string

	// Returns table statistics on the replication, including table name, rows
	// inserted, rows updated, and rows deleted.
	ReplicationTableStatistics []types.TableStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplicationTableStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationTableStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationTableStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReplicationTableStatistics"); err != nil {
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
	if err = addOpDescribeReplicationTableStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationTableStatistics(options.Region), middleware.Before); err != nil {
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

// DescribeReplicationTableStatisticsAPIClient is a client that implements the
// DescribeReplicationTableStatistics operation.
type DescribeReplicationTableStatisticsAPIClient interface {
	DescribeReplicationTableStatistics(context.Context, *DescribeReplicationTableStatisticsInput, ...func(*Options)) (*DescribeReplicationTableStatisticsOutput, error)
}

var _ DescribeReplicationTableStatisticsAPIClient = (*Client)(nil)

// DescribeReplicationTableStatisticsPaginatorOptions is the paginator options for
// DescribeReplicationTableStatistics
type DescribeReplicationTableStatisticsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReplicationTableStatisticsPaginator is a paginator for
// DescribeReplicationTableStatistics
type DescribeReplicationTableStatisticsPaginator struct {
	options   DescribeReplicationTableStatisticsPaginatorOptions
	client    DescribeReplicationTableStatisticsAPIClient
	params    *DescribeReplicationTableStatisticsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplicationTableStatisticsPaginator returns a new
// DescribeReplicationTableStatisticsPaginator
func NewDescribeReplicationTableStatisticsPaginator(client DescribeReplicationTableStatisticsAPIClient, params *DescribeReplicationTableStatisticsInput, optFns ...func(*DescribeReplicationTableStatisticsPaginatorOptions)) *DescribeReplicationTableStatisticsPaginator {
	if params == nil {
		params = &DescribeReplicationTableStatisticsInput{}
	}

	options := DescribeReplicationTableStatisticsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReplicationTableStatisticsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplicationTableStatisticsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReplicationTableStatistics page.
func (p *DescribeReplicationTableStatisticsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplicationTableStatisticsOutput, error) {
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

	result, err := p.client.DescribeReplicationTableStatistics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReplicationTableStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReplicationTableStatistics",
	}
}

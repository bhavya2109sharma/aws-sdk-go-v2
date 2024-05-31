// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of user groups.
func (c *Client) DescribeUserGroups(ctx context.Context, params *DescribeUserGroupsInput, optFns ...func(*Options)) (*DescribeUserGroupsOutput, error) {
	if params == nil {
		params = &DescribeUserGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUserGroups", params, optFns, c.addOperationDescribeUserGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserGroupsInput struct {

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords. >
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	MaxRecords *int32

	// The ID of the user group.
	UserGroupId *string

	noSmithyDocumentSerde
}

type DescribeUserGroupsOutput struct {

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.>
	Marker *string

	// Returns a list of user groups.
	UserGroups []types.UserGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeUserGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeUserGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeUserGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeUserGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUserGroups(options.Region), middleware.Before); err != nil {
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

// DescribeUserGroupsAPIClient is a client that implements the DescribeUserGroups
// operation.
type DescribeUserGroupsAPIClient interface {
	DescribeUserGroups(context.Context, *DescribeUserGroupsInput, ...func(*Options)) (*DescribeUserGroupsOutput, error)
}

var _ DescribeUserGroupsAPIClient = (*Client)(nil)

// DescribeUserGroupsPaginatorOptions is the paginator options for
// DescribeUserGroups
type DescribeUserGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeUserGroupsPaginator is a paginator for DescribeUserGroups
type DescribeUserGroupsPaginator struct {
	options   DescribeUserGroupsPaginatorOptions
	client    DescribeUserGroupsAPIClient
	params    *DescribeUserGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeUserGroupsPaginator returns a new DescribeUserGroupsPaginator
func NewDescribeUserGroupsPaginator(client DescribeUserGroupsAPIClient, params *DescribeUserGroupsInput, optFns ...func(*DescribeUserGroupsPaginatorOptions)) *DescribeUserGroupsPaginator {
	if params == nil {
		params = &DescribeUserGroupsInput{}
	}

	options := DescribeUserGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeUserGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeUserGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeUserGroups page.
func (p *DescribeUserGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeUserGroupsOutput, error) {
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

	result, err := p.client.DescribeUserGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeUserGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeUserGroups",
	}
}

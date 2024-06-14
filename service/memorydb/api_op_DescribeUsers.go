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

// Returns a list of users.
func (c *Client) DescribeUsers(ctx context.Context, params *DescribeUsersInput, optFns ...func(*Options)) (*DescribeUsersOutput, error) {
	if params == nil {
		params = &DescribeUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUsers", params, optFns, c.addOperationDescribeUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUsersInput struct {

	// Filter to determine the list of users to return.
	Filters []types.Filter

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

	// The name of the user
	UserName *string

	noSmithyDocumentSerde
}

type DescribeUsersOutput struct {

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// A list of users.
	Users []types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUsers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeUsers"); err != nil {
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
	if err = addOpDescribeUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUsers(options.Region), middleware.Before); err != nil {
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

// DescribeUsersAPIClient is a client that implements the DescribeUsers operation.
type DescribeUsersAPIClient interface {
	DescribeUsers(context.Context, *DescribeUsersInput, ...func(*Options)) (*DescribeUsersOutput, error)
}

var _ DescribeUsersAPIClient = (*Client)(nil)

// DescribeUsersPaginatorOptions is the paginator options for DescribeUsers
type DescribeUsersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeUsersPaginator is a paginator for DescribeUsers
type DescribeUsersPaginator struct {
	options   DescribeUsersPaginatorOptions
	client    DescribeUsersAPIClient
	params    *DescribeUsersInput
	nextToken *string
	firstPage bool
}

// NewDescribeUsersPaginator returns a new DescribeUsersPaginator
func NewDescribeUsersPaginator(client DescribeUsersAPIClient, params *DescribeUsersInput, optFns ...func(*DescribeUsersPaginatorOptions)) *DescribeUsersPaginator {
	if params == nil {
		params = &DescribeUsersInput{}
	}

	options := DescribeUsersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeUsersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeUsersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeUsers page.
func (p *DescribeUsersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeUsersOutput, error) {
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

	result, err := p.client.DescribeUsers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeUsers",
	}
}

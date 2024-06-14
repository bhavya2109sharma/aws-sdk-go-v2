// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Run queries to access information from your knowledge graph of entities within
// individual workspaces.
//
// The ExecuteQuery action only works with [Amazon Web Services Java SDK2]. ExecuteQuery will not work with any
// Amazon Web Services Java SDK version < 2.x.
//
// [Amazon Web Services Java SDK2]: https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide/home.html
func (c *Client) ExecuteQuery(ctx context.Context, params *ExecuteQueryInput, optFns ...func(*Options)) (*ExecuteQueryOutput, error) {
	if params == nil {
		params = &ExecuteQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteQuery", params, optFns, c.addOperationExecuteQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteQueryInput struct {

	// The query statement.
	//
	// This member is required.
	QueryStatement *string

	// The ID of the workspace.
	//
	// This member is required.
	WorkspaceId *string

	// The maximum number of results to return at one time. The default is 50.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ExecuteQueryOutput struct {

	// A list of ColumnDescription objects.
	ColumnDescriptions []types.ColumnDescription

	// The string that specifies the next page of results.
	NextToken *string

	// Represents a single row in the query results.
	Rows []types.Row

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExecuteQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExecuteQuery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExecuteQuery"); err != nil {
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
	if err = addEndpointPrefix_opExecuteQueryMiddleware(stack); err != nil {
		return err
	}
	if err = addOpExecuteQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteQuery(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opExecuteQueryMiddleware struct {
}

func (*endpointPrefix_opExecuteQueryMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opExecuteQueryMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opExecuteQueryMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opExecuteQueryMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ExecuteQueryAPIClient is a client that implements the ExecuteQuery operation.
type ExecuteQueryAPIClient interface {
	ExecuteQuery(context.Context, *ExecuteQueryInput, ...func(*Options)) (*ExecuteQueryOutput, error)
}

var _ ExecuteQueryAPIClient = (*Client)(nil)

// ExecuteQueryPaginatorOptions is the paginator options for ExecuteQuery
type ExecuteQueryPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ExecuteQueryPaginator is a paginator for ExecuteQuery
type ExecuteQueryPaginator struct {
	options   ExecuteQueryPaginatorOptions
	client    ExecuteQueryAPIClient
	params    *ExecuteQueryInput
	nextToken *string
	firstPage bool
}

// NewExecuteQueryPaginator returns a new ExecuteQueryPaginator
func NewExecuteQueryPaginator(client ExecuteQueryAPIClient, params *ExecuteQueryInput, optFns ...func(*ExecuteQueryPaginatorOptions)) *ExecuteQueryPaginator {
	if params == nil {
		params = &ExecuteQueryInput{}
	}

	options := ExecuteQueryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ExecuteQueryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ExecuteQueryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ExecuteQuery page.
func (p *ExecuteQueryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ExecuteQueryOutput, error) {
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

	result, err := p.client.ExecuteQuery(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opExecuteQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExecuteQuery",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of tasks for a run.
func (c *Client) ListRunTasks(ctx context.Context, params *ListRunTasksInput, optFns ...func(*Options)) (*ListRunTasksOutput, error) {
	if params == nil {
		params = &ListRunTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRunTasks", params, optFns, c.addOperationListRunTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRunTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRunTasksInput struct {

	// The run's ID.
	//
	// This member is required.
	Id *string

	// The maximum number of run tasks to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	StartingToken *string

	// Filter the list by status.
	Status types.TaskStatus

	noSmithyDocumentSerde
}

type ListRunTasksOutput struct {

	// A list of tasks.
	Items []types.TaskListItem

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRunTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRunTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRunTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRunTasks"); err != nil {
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
	if err = addEndpointPrefix_opListRunTasksMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListRunTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRunTasks(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListRunTasksMiddleware struct {
}

func (*endpointPrefix_opListRunTasksMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListRunTasksMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "workflows-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListRunTasksMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListRunTasksMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListRunTasksAPIClient is a client that implements the ListRunTasks operation.
type ListRunTasksAPIClient interface {
	ListRunTasks(context.Context, *ListRunTasksInput, ...func(*Options)) (*ListRunTasksOutput, error)
}

var _ ListRunTasksAPIClient = (*Client)(nil)

// ListRunTasksPaginatorOptions is the paginator options for ListRunTasks
type ListRunTasksPaginatorOptions struct {
	// The maximum number of run tasks to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRunTasksPaginator is a paginator for ListRunTasks
type ListRunTasksPaginator struct {
	options   ListRunTasksPaginatorOptions
	client    ListRunTasksAPIClient
	params    *ListRunTasksInput
	nextToken *string
	firstPage bool
}

// NewListRunTasksPaginator returns a new ListRunTasksPaginator
func NewListRunTasksPaginator(client ListRunTasksAPIClient, params *ListRunTasksInput, optFns ...func(*ListRunTasksPaginatorOptions)) *ListRunTasksPaginator {
	if params == nil {
		params = &ListRunTasksInput{}
	}

	options := ListRunTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRunTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.StartingToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRunTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRunTasks page.
func (p *ListRunTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRunTasksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.StartingToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListRunTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRunTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRunTasks",
	}
}

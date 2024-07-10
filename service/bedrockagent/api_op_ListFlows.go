// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of flows and information about each flow. For more information,
// see [Manage a flow in Amazon Bedrock]in the Amazon Bedrock User Guide.
//
// [Manage a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-manage.html
func (c *Client) ListFlows(ctx context.Context, params *ListFlowsInput, optFns ...func(*Options)) (*ListFlowsOutput, error) {
	if params == nil {
		params = &ListFlowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFlows", params, optFns, c.addOperationListFlowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFlowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlowsInput struct {

	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	MaxResults *int32

	// If the total number of results is greater than the maxResults value provided in
	// the request, enter the token returned in the nextToken field in the response in
	// this field to return the next batch of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFlowsOutput struct {

	// A list, each member of which contains information about a flow.
	//
	// This member is required.
	FlowSummaries []types.FlowSummary

	// If the total number of results is greater than the maxResults value provided in
	// the request, use this token when making another request in the nextToken field
	// to return the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFlowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFlows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFlows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFlows"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFlows(options.Region), middleware.Before); err != nil {
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

// ListFlowsPaginatorOptions is the paginator options for ListFlows
type ListFlowsPaginatorOptions struct {
	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFlowsPaginator is a paginator for ListFlows
type ListFlowsPaginator struct {
	options   ListFlowsPaginatorOptions
	client    ListFlowsAPIClient
	params    *ListFlowsInput
	nextToken *string
	firstPage bool
}

// NewListFlowsPaginator returns a new ListFlowsPaginator
func NewListFlowsPaginator(client ListFlowsAPIClient, params *ListFlowsInput, optFns ...func(*ListFlowsPaginatorOptions)) *ListFlowsPaginator {
	if params == nil {
		params = &ListFlowsInput{}
	}

	options := ListFlowsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFlowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFlowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFlows page.
func (p *ListFlowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFlowsOutput, error) {
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

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListFlows(ctx, &params, optFns...)
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

// ListFlowsAPIClient is a client that implements the ListFlows operation.
type ListFlowsAPIClient interface {
	ListFlows(context.Context, *ListFlowsInput, ...func(*Options)) (*ListFlowsOutput, error)
}

var _ ListFlowsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListFlows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFlows",
	}
}

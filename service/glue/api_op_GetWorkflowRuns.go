// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves metadata for all runs of a given workflow.
func (c *Client) GetWorkflowRuns(ctx context.Context, params *GetWorkflowRunsInput, optFns ...func(*Options)) (*GetWorkflowRunsOutput, error) {
	if params == nil {
		params = &GetWorkflowRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowRuns", params, optFns, c.addOperationGetWorkflowRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowRunsInput struct {

	// Name of the workflow whose metadata of runs should be returned.
	//
	// This member is required.
	Name *string

	// Specifies whether to include the workflow graph in response or not.
	IncludeGraph *bool

	// The maximum number of workflow runs to be included in the response.
	MaxResults *int32

	// The maximum size of the response.
	NextToken *string

	noSmithyDocumentSerde
}

type GetWorkflowRunsOutput struct {

	// A continuation token, if not all requested workflow runs have been returned.
	NextToken *string

	// A list of workflow run metadata objects.
	Runs []types.WorkflowRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkflowRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetWorkflowRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetWorkflowRuns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWorkflowRuns"); err != nil {
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
	if err = addOpGetWorkflowRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowRuns(options.Region), middleware.Before); err != nil {
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

// GetWorkflowRunsAPIClient is a client that implements the GetWorkflowRuns
// operation.
type GetWorkflowRunsAPIClient interface {
	GetWorkflowRuns(context.Context, *GetWorkflowRunsInput, ...func(*Options)) (*GetWorkflowRunsOutput, error)
}

var _ GetWorkflowRunsAPIClient = (*Client)(nil)

// GetWorkflowRunsPaginatorOptions is the paginator options for GetWorkflowRuns
type GetWorkflowRunsPaginatorOptions struct {
	// The maximum number of workflow runs to be included in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetWorkflowRunsPaginator is a paginator for GetWorkflowRuns
type GetWorkflowRunsPaginator struct {
	options   GetWorkflowRunsPaginatorOptions
	client    GetWorkflowRunsAPIClient
	params    *GetWorkflowRunsInput
	nextToken *string
	firstPage bool
}

// NewGetWorkflowRunsPaginator returns a new GetWorkflowRunsPaginator
func NewGetWorkflowRunsPaginator(client GetWorkflowRunsAPIClient, params *GetWorkflowRunsInput, optFns ...func(*GetWorkflowRunsPaginatorOptions)) *GetWorkflowRunsPaginator {
	if params == nil {
		params = &GetWorkflowRunsInput{}
	}

	options := GetWorkflowRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetWorkflowRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetWorkflowRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetWorkflowRuns page.
func (p *GetWorkflowRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetWorkflowRunsOutput, error) {
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

	result, err := p.client.GetWorkflowRuns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetWorkflowRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWorkflowRuns",
	}
}

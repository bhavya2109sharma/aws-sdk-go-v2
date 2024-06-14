// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

//	Lists mitigation actions executions for a Device Defender ML Detect Security
//
// Profile.
//
// Requires permission to access the [ListDetectMitigationActionsExecutions] action.
//
// [ListDetectMitigationActionsExecutions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListDetectMitigationActionsExecutions(ctx context.Context, params *ListDetectMitigationActionsExecutionsInput, optFns ...func(*Options)) (*ListDetectMitigationActionsExecutionsOutput, error) {
	if params == nil {
		params = &ListDetectMitigationActionsExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDetectMitigationActionsExecutions", params, optFns, c.addOperationListDetectMitigationActionsExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDetectMitigationActionsExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDetectMitigationActionsExecutionsInput struct {

	//  The end of the time period for which ML Detect mitigation actions executions
	// are returned.
	EndTime *time.Time

	//  The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	//  The token for the next set of results.
	NextToken *string

	//  A filter to limit results to those found after the specified time. You must
	// specify either the startTime and endTime or the taskId, but not both.
	StartTime *time.Time

	//  The unique identifier of the task.
	TaskId *string

	//  The name of the thing whose mitigation actions are listed.
	ThingName *string

	//  The unique identifier of the violation.
	ViolationId *string

	noSmithyDocumentSerde
}

type ListDetectMitigationActionsExecutionsOutput struct {

	//  List of actions executions.
	ActionsExecutions []types.DetectMitigationActionExecution

	//  A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDetectMitigationActionsExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDetectMitigationActionsExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDetectMitigationActionsExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDetectMitigationActionsExecutions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDetectMitigationActionsExecutions(options.Region), middleware.Before); err != nil {
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

// ListDetectMitigationActionsExecutionsAPIClient is a client that implements the
// ListDetectMitigationActionsExecutions operation.
type ListDetectMitigationActionsExecutionsAPIClient interface {
	ListDetectMitigationActionsExecutions(context.Context, *ListDetectMitigationActionsExecutionsInput, ...func(*Options)) (*ListDetectMitigationActionsExecutionsOutput, error)
}

var _ ListDetectMitigationActionsExecutionsAPIClient = (*Client)(nil)

// ListDetectMitigationActionsExecutionsPaginatorOptions is the paginator options
// for ListDetectMitigationActionsExecutions
type ListDetectMitigationActionsExecutionsPaginatorOptions struct {
	//  The maximum number of results to return at one time. The default is 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDetectMitigationActionsExecutionsPaginator is a paginator for
// ListDetectMitigationActionsExecutions
type ListDetectMitigationActionsExecutionsPaginator struct {
	options   ListDetectMitigationActionsExecutionsPaginatorOptions
	client    ListDetectMitigationActionsExecutionsAPIClient
	params    *ListDetectMitigationActionsExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListDetectMitigationActionsExecutionsPaginator returns a new
// ListDetectMitigationActionsExecutionsPaginator
func NewListDetectMitigationActionsExecutionsPaginator(client ListDetectMitigationActionsExecutionsAPIClient, params *ListDetectMitigationActionsExecutionsInput, optFns ...func(*ListDetectMitigationActionsExecutionsPaginatorOptions)) *ListDetectMitigationActionsExecutionsPaginator {
	if params == nil {
		params = &ListDetectMitigationActionsExecutionsInput{}
	}

	options := ListDetectMitigationActionsExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDetectMitigationActionsExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDetectMitigationActionsExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDetectMitigationActionsExecutions page.
func (p *ListDetectMitigationActionsExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDetectMitigationActionsExecutionsOutput, error) {
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

	result, err := p.client.ListDetectMitigationActionsExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDetectMitigationActionsExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDetectMitigationActionsExecutions",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of Dev Environments in a project.
func (c *Client) ListDevEnvironments(ctx context.Context, params *ListDevEnvironmentsInput, optFns ...func(*Options)) (*ListDevEnvironmentsOutput, error) {
	if params == nil {
		params = &ListDevEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevEnvironments", params, optFns, c.addOperationListDevEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDevEnvironmentsInput struct {

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// Information about filters to apply to narrow the results returned in the list.
	Filters []types.Filter

	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	MaxResults *int32

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	// The name of the project in the space.
	ProjectName *string

	noSmithyDocumentSerde
}

type ListDevEnvironmentsOutput struct {

	// Information about the Dev Environments in a project.
	//
	// This member is required.
	Items []types.DevEnvironmentSummary

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDevEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDevEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDevEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDevEnvironments"); err != nil {
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
	if err = addOpListDevEnvironmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDevEnvironments(options.Region), middleware.Before); err != nil {
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

// ListDevEnvironmentsAPIClient is a client that implements the
// ListDevEnvironments operation.
type ListDevEnvironmentsAPIClient interface {
	ListDevEnvironments(context.Context, *ListDevEnvironmentsInput, ...func(*Options)) (*ListDevEnvironmentsOutput, error)
}

var _ ListDevEnvironmentsAPIClient = (*Client)(nil)

// ListDevEnvironmentsPaginatorOptions is the paginator options for
// ListDevEnvironments
type ListDevEnvironmentsPaginatorOptions struct {
	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDevEnvironmentsPaginator is a paginator for ListDevEnvironments
type ListDevEnvironmentsPaginator struct {
	options   ListDevEnvironmentsPaginatorOptions
	client    ListDevEnvironmentsAPIClient
	params    *ListDevEnvironmentsInput
	nextToken *string
	firstPage bool
}

// NewListDevEnvironmentsPaginator returns a new ListDevEnvironmentsPaginator
func NewListDevEnvironmentsPaginator(client ListDevEnvironmentsAPIClient, params *ListDevEnvironmentsInput, optFns ...func(*ListDevEnvironmentsPaginatorOptions)) *ListDevEnvironmentsPaginator {
	if params == nil {
		params = &ListDevEnvironmentsInput{}
	}

	options := ListDevEnvironmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDevEnvironmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDevEnvironmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDevEnvironments page.
func (p *ListDevEnvironmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDevEnvironmentsOutput, error) {
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

	result, err := p.client.ListDevEnvironments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDevEnvironments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDevEnvironments",
	}
}

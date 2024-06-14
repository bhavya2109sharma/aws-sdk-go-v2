// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of the existing discovery jobs in the Amazon Web Services
// Region and Amazon Web Services account where you're using DataSync Discovery.
func (c *Client) ListDiscoveryJobs(ctx context.Context, params *ListDiscoveryJobsInput, optFns ...func(*Options)) (*ListDiscoveryJobsOutput, error) {
	if params == nil {
		params = &ListDiscoveryJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDiscoveryJobs", params, optFns, c.addOperationListDiscoveryJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDiscoveryJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDiscoveryJobsInput struct {

	// Specifies how many results you want in the response.
	MaxResults *int32

	// Specifies an opaque string that indicates the position to begin the next list
	// of results in the response.
	NextToken *string

	// Specifies the Amazon Resource Name (ARN) of an on-premises storage system. Use
	// this parameter if you only want to list the discovery jobs that are associated
	// with a specific storage system.
	StorageSystemArn *string

	noSmithyDocumentSerde
}

type ListDiscoveryJobsOutput struct {

	// The discovery jobs that you've run.
	DiscoveryJobs []types.DiscoveryJobListEntry

	// The opaque string that indicates the position to begin the next list of results
	// in the response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDiscoveryJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDiscoveryJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDiscoveryJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDiscoveryJobs"); err != nil {
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
	if err = addEndpointPrefix_opListDiscoveryJobsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDiscoveryJobs(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListDiscoveryJobsMiddleware struct {
}

func (*endpointPrefix_opListDiscoveryJobsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListDiscoveryJobsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "discovery-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListDiscoveryJobsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListDiscoveryJobsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListDiscoveryJobsAPIClient is a client that implements the ListDiscoveryJobs
// operation.
type ListDiscoveryJobsAPIClient interface {
	ListDiscoveryJobs(context.Context, *ListDiscoveryJobsInput, ...func(*Options)) (*ListDiscoveryJobsOutput, error)
}

var _ ListDiscoveryJobsAPIClient = (*Client)(nil)

// ListDiscoveryJobsPaginatorOptions is the paginator options for ListDiscoveryJobs
type ListDiscoveryJobsPaginatorOptions struct {
	// Specifies how many results you want in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDiscoveryJobsPaginator is a paginator for ListDiscoveryJobs
type ListDiscoveryJobsPaginator struct {
	options   ListDiscoveryJobsPaginatorOptions
	client    ListDiscoveryJobsAPIClient
	params    *ListDiscoveryJobsInput
	nextToken *string
	firstPage bool
}

// NewListDiscoveryJobsPaginator returns a new ListDiscoveryJobsPaginator
func NewListDiscoveryJobsPaginator(client ListDiscoveryJobsAPIClient, params *ListDiscoveryJobsInput, optFns ...func(*ListDiscoveryJobsPaginatorOptions)) *ListDiscoveryJobsPaginator {
	if params == nil {
		params = &ListDiscoveryJobsInput{}
	}

	options := ListDiscoveryJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDiscoveryJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDiscoveryJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDiscoveryJobs page.
func (p *ListDiscoveryJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDiscoveryJobsOutput, error) {
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

	result, err := p.client.ListDiscoveryJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDiscoveryJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDiscoveryJobs",
	}
}

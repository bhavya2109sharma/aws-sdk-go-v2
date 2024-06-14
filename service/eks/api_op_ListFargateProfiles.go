// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Fargate profiles associated with the specified cluster in your Amazon
// Web Services account in the specified Amazon Web Services Region.
func (c *Client) ListFargateProfiles(ctx context.Context, params *ListFargateProfilesInput, optFns ...func(*Options)) (*ListFargateProfilesOutput, error) {
	if params == nil {
		params = &ListFargateProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFargateProfiles", params, optFns, c.addOperationListFargateProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFargateProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFargateProfilesInput struct {

	// The name of your cluster.
	//
	// This member is required.
	ClusterName *string

	// The maximum number of results, returned in paginated output. You receive
	// maxResults in a single page, along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, 100 results and a nextToken value, if applicable, are
	// returned.
	MaxResults *int32

	// The nextToken value returned from a previous paginated request, where maxResults
	// was used and the results exceeded the value of that parameter. Pagination
	// continues from the end of the previous results that returned the nextToken
	// value. This value is null when there are no more results to return.
	//
	// This token should be treated as an opaque identifier that is used only to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFargateProfilesOutput struct {

	// A list of all of the Fargate profiles associated with the specified cluster.
	FargateProfileNames []string

	// The nextToken value returned from a previous paginated request, where maxResults
	// was used and the results exceeded the value of that parameter. Pagination
	// continues from the end of the previous results that returned the nextToken
	// value. This value is null when there are no more results to return.
	//
	// This token should be treated as an opaque identifier that is used only to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFargateProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFargateProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFargateProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFargateProfiles"); err != nil {
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
	if err = addOpListFargateProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFargateProfiles(options.Region), middleware.Before); err != nil {
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

// ListFargateProfilesAPIClient is a client that implements the
// ListFargateProfiles operation.
type ListFargateProfilesAPIClient interface {
	ListFargateProfiles(context.Context, *ListFargateProfilesInput, ...func(*Options)) (*ListFargateProfilesOutput, error)
}

var _ ListFargateProfilesAPIClient = (*Client)(nil)

// ListFargateProfilesPaginatorOptions is the paginator options for
// ListFargateProfiles
type ListFargateProfilesPaginatorOptions struct {
	// The maximum number of results, returned in paginated output. You receive
	// maxResults in a single page, along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, 100 results and a nextToken value, if applicable, are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFargateProfilesPaginator is a paginator for ListFargateProfiles
type ListFargateProfilesPaginator struct {
	options   ListFargateProfilesPaginatorOptions
	client    ListFargateProfilesAPIClient
	params    *ListFargateProfilesInput
	nextToken *string
	firstPage bool
}

// NewListFargateProfilesPaginator returns a new ListFargateProfilesPaginator
func NewListFargateProfilesPaginator(client ListFargateProfilesAPIClient, params *ListFargateProfilesInput, optFns ...func(*ListFargateProfilesPaginatorOptions)) *ListFargateProfilesPaginator {
	if params == nil {
		params = &ListFargateProfilesInput{}
	}

	options := ListFargateProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFargateProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFargateProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFargateProfiles page.
func (p *ListFargateProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFargateProfilesOutput, error) {
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

	result, err := p.client.ListFargateProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFargateProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFargateProfiles",
	}
}

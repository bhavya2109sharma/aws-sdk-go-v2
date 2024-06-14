// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the application versions for a specific application.
func (c *Client) ListApplicationVersions(ctx context.Context, params *ListApplicationVersionsInput, optFns ...func(*Options)) (*ListApplicationVersionsOutput, error) {
	if params == nil {
		params = &ListApplicationVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationVersions", params, optFns, c.addOperationListApplicationVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationVersionsInput struct {

	// The unique identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The maximum number of application versions to return.
	MaxResults *int32

	// A pagination token returned from a previous call to this operation. This
	// specifies the next item to return. To return to the beginning of the list,
	// exclude this parameter.
	NextToken *string

	noSmithyDocumentSerde
}

type ListApplicationVersionsOutput struct {

	// The list of application versions.
	//
	// This member is required.
	ApplicationVersions []types.ApplicationVersionSummary

	// If there are more items to return, this contains a token that is passed to a
	// subsequent call to this operation to retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListApplicationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListApplicationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListApplicationVersions"); err != nil {
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
	if err = addOpListApplicationVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationVersions(options.Region), middleware.Before); err != nil {
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

// ListApplicationVersionsAPIClient is a client that implements the
// ListApplicationVersions operation.
type ListApplicationVersionsAPIClient interface {
	ListApplicationVersions(context.Context, *ListApplicationVersionsInput, ...func(*Options)) (*ListApplicationVersionsOutput, error)
}

var _ ListApplicationVersionsAPIClient = (*Client)(nil)

// ListApplicationVersionsPaginatorOptions is the paginator options for
// ListApplicationVersions
type ListApplicationVersionsPaginatorOptions struct {
	// The maximum number of application versions to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationVersionsPaginator is a paginator for ListApplicationVersions
type ListApplicationVersionsPaginator struct {
	options   ListApplicationVersionsPaginatorOptions
	client    ListApplicationVersionsAPIClient
	params    *ListApplicationVersionsInput
	nextToken *string
	firstPage bool
}

// NewListApplicationVersionsPaginator returns a new
// ListApplicationVersionsPaginator
func NewListApplicationVersionsPaginator(client ListApplicationVersionsAPIClient, params *ListApplicationVersionsInput, optFns ...func(*ListApplicationVersionsPaginatorOptions)) *ListApplicationVersionsPaginator {
	if params == nil {
		params = &ListApplicationVersionsInput{}
	}

	options := ListApplicationVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationVersions page.
func (p *ListApplicationVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationVersionsOutput, error) {
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

	result, err := p.client.ListApplicationVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListApplicationVersions",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the most current information about the stream. The streamName or
// streamARN should be provided in the input.
func (c *Client) DescribeMappedResourceConfiguration(ctx context.Context, params *DescribeMappedResourceConfigurationInput, optFns ...func(*Options)) (*DescribeMappedResourceConfigurationOutput, error) {
	if params == nil {
		params = &DescribeMappedResourceConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMappedResourceConfiguration", params, optFns, c.addOperationDescribeMappedResourceConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMappedResourceConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMappedResourceConfigurationInput struct {

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The token to provide in your next request, to get another batch of results.
	NextToken *string

	// The Amazon Resource Name (ARN) of the stream.
	StreamARN *string

	// The name of the stream.
	StreamName *string

	noSmithyDocumentSerde
}

type DescribeMappedResourceConfigurationOutput struct {

	// A structure that encapsulates, or contains, the media storage configuration
	// properties.
	MappedResourceConfigurationList []types.MappedResourceConfigurationListItem

	// The token that was used in the NextToken request to fetch the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMappedResourceConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMappedResourceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMappedResourceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMappedResourceConfiguration"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMappedResourceConfiguration(options.Region), middleware.Before); err != nil {
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

// DescribeMappedResourceConfigurationAPIClient is a client that implements the
// DescribeMappedResourceConfiguration operation.
type DescribeMappedResourceConfigurationAPIClient interface {
	DescribeMappedResourceConfiguration(context.Context, *DescribeMappedResourceConfigurationInput, ...func(*Options)) (*DescribeMappedResourceConfigurationOutput, error)
}

var _ DescribeMappedResourceConfigurationAPIClient = (*Client)(nil)

// DescribeMappedResourceConfigurationPaginatorOptions is the paginator options
// for DescribeMappedResourceConfiguration
type DescribeMappedResourceConfigurationPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMappedResourceConfigurationPaginator is a paginator for
// DescribeMappedResourceConfiguration
type DescribeMappedResourceConfigurationPaginator struct {
	options   DescribeMappedResourceConfigurationPaginatorOptions
	client    DescribeMappedResourceConfigurationAPIClient
	params    *DescribeMappedResourceConfigurationInput
	nextToken *string
	firstPage bool
}

// NewDescribeMappedResourceConfigurationPaginator returns a new
// DescribeMappedResourceConfigurationPaginator
func NewDescribeMappedResourceConfigurationPaginator(client DescribeMappedResourceConfigurationAPIClient, params *DescribeMappedResourceConfigurationInput, optFns ...func(*DescribeMappedResourceConfigurationPaginatorOptions)) *DescribeMappedResourceConfigurationPaginator {
	if params == nil {
		params = &DescribeMappedResourceConfigurationInput{}
	}

	options := DescribeMappedResourceConfigurationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMappedResourceConfigurationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMappedResourceConfigurationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMappedResourceConfiguration page.
func (p *DescribeMappedResourceConfigurationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMappedResourceConfigurationOutput, error) {
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

	result, err := p.client.DescribeMappedResourceConfiguration(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMappedResourceConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMappedResourceConfiguration",
	}
}

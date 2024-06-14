// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the available Amazon OpenSearch Service Reserved Instance offerings
// for a given Region. For more information, see [Reserved Instances in Amazon OpenSearch Service].
//
// [Reserved Instances in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ri.html
func (c *Client) DescribeReservedInstanceOfferings(ctx context.Context, params *DescribeReservedInstanceOfferingsInput, optFns ...func(*Options)) (*DescribeReservedInstanceOfferingsOutput, error) {
	if params == nil {
		params = &DescribeReservedInstanceOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedInstanceOfferings", params, optFns, c.addOperationDescribeReservedInstanceOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedInstanceOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to a DescribeReservedInstanceOfferings
// operation.
type DescribeReservedInstanceOfferingsInput struct {

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults int32

	// If your initial DescribeReservedInstanceOfferings operation returns a nextToken
	// , you can include the returned nextToken in subsequent
	// DescribeReservedInstanceOfferings operations, which returns results in the next
	// page.
	NextToken *string

	// The Reserved Instance identifier filter value. Use this parameter to show only
	// the available instance types that match the specified reservation identifier.
	ReservedInstanceOfferingId *string

	noSmithyDocumentSerde
}

// Container for results of a DescribeReservedInstanceOfferings request.
type DescribeReservedInstanceOfferingsOutput struct {

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Send the request again
	// using the returned token to retrieve the next page.
	NextToken *string

	// List of Reserved Instance offerings.
	ReservedInstanceOfferings []types.ReservedInstanceOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedInstanceOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReservedInstanceOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReservedInstanceOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReservedInstanceOfferings"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedInstanceOfferings(options.Region), middleware.Before); err != nil {
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

// DescribeReservedInstanceOfferingsAPIClient is a client that implements the
// DescribeReservedInstanceOfferings operation.
type DescribeReservedInstanceOfferingsAPIClient interface {
	DescribeReservedInstanceOfferings(context.Context, *DescribeReservedInstanceOfferingsInput, ...func(*Options)) (*DescribeReservedInstanceOfferingsOutput, error)
}

var _ DescribeReservedInstanceOfferingsAPIClient = (*Client)(nil)

// DescribeReservedInstanceOfferingsPaginatorOptions is the paginator options for
// DescribeReservedInstanceOfferings
type DescribeReservedInstanceOfferingsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedInstanceOfferingsPaginator is a paginator for
// DescribeReservedInstanceOfferings
type DescribeReservedInstanceOfferingsPaginator struct {
	options   DescribeReservedInstanceOfferingsPaginatorOptions
	client    DescribeReservedInstanceOfferingsAPIClient
	params    *DescribeReservedInstanceOfferingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedInstanceOfferingsPaginator returns a new
// DescribeReservedInstanceOfferingsPaginator
func NewDescribeReservedInstanceOfferingsPaginator(client DescribeReservedInstanceOfferingsAPIClient, params *DescribeReservedInstanceOfferingsInput, optFns ...func(*DescribeReservedInstanceOfferingsPaginatorOptions)) *DescribeReservedInstanceOfferingsPaginator {
	if params == nil {
		params = &DescribeReservedInstanceOfferingsInput{}
	}

	options := DescribeReservedInstanceOfferingsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedInstanceOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedInstanceOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedInstanceOfferings page.
func (p *DescribeReservedInstanceOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedInstanceOfferingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeReservedInstanceOfferings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReservedInstanceOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReservedInstanceOfferings",
	}
}

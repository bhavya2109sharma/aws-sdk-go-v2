// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the Fleet Advisor collectors in your account.
func (c *Client) DescribeFleetAdvisorCollectors(ctx context.Context, params *DescribeFleetAdvisorCollectorsInput, optFns ...func(*Options)) (*DescribeFleetAdvisorCollectorsOutput, error) {
	if params == nil {
		params = &DescribeFleetAdvisorCollectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetAdvisorCollectors", params, optFns, c.addOperationDescribeFleetAdvisorCollectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetAdvisorCollectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetAdvisorCollectorsInput struct {

	//  If you specify any of the following filters, the output includes information
	// for only those collectors that meet the filter criteria:
	//
	//   - collector-referenced-id – The ID of the collector agent, for example
	//   d4610ac5-e323-4ad9-bc50-eaf7249dfe9d .
	//
	//   - collector-name – The name of the collector agent.
	//
	// An example is: describe-fleet-advisor-collectors --filter
	// Name="collector-referenced-id",Values="d4610ac5-e323-4ad9-bc50-eaf7249dfe9d"
	Filters []types.Filter

	// Sets the maximum number of records returned in the response.
	MaxRecords *int32

	// If NextToken is returned by a previous response, there are more results
	// available. The value of NextToken is a unique pagination token for each page.
	// Make the call again using the returned token to retrieve the next page. Keep all
	// other arguments unchanged.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetAdvisorCollectorsOutput struct {

	// Provides descriptions of the Fleet Advisor collectors, including the
	// collectors' name and ID, and the latest inventory data.
	Collectors []types.CollectorResponse

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetAdvisorCollectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetAdvisorCollectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetAdvisorCollectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFleetAdvisorCollectors"); err != nil {
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
	if err = addOpDescribeFleetAdvisorCollectorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetAdvisorCollectors(options.Region), middleware.Before); err != nil {
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

// DescribeFleetAdvisorCollectorsAPIClient is a client that implements the
// DescribeFleetAdvisorCollectors operation.
type DescribeFleetAdvisorCollectorsAPIClient interface {
	DescribeFleetAdvisorCollectors(context.Context, *DescribeFleetAdvisorCollectorsInput, ...func(*Options)) (*DescribeFleetAdvisorCollectorsOutput, error)
}

var _ DescribeFleetAdvisorCollectorsAPIClient = (*Client)(nil)

// DescribeFleetAdvisorCollectorsPaginatorOptions is the paginator options for
// DescribeFleetAdvisorCollectors
type DescribeFleetAdvisorCollectorsPaginatorOptions struct {
	// Sets the maximum number of records returned in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetAdvisorCollectorsPaginator is a paginator for
// DescribeFleetAdvisorCollectors
type DescribeFleetAdvisorCollectorsPaginator struct {
	options   DescribeFleetAdvisorCollectorsPaginatorOptions
	client    DescribeFleetAdvisorCollectorsAPIClient
	params    *DescribeFleetAdvisorCollectorsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetAdvisorCollectorsPaginator returns a new
// DescribeFleetAdvisorCollectorsPaginator
func NewDescribeFleetAdvisorCollectorsPaginator(client DescribeFleetAdvisorCollectorsAPIClient, params *DescribeFleetAdvisorCollectorsInput, optFns ...func(*DescribeFleetAdvisorCollectorsPaginatorOptions)) *DescribeFleetAdvisorCollectorsPaginator {
	if params == nil {
		params = &DescribeFleetAdvisorCollectorsInput{}
	}

	options := DescribeFleetAdvisorCollectorsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetAdvisorCollectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetAdvisorCollectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleetAdvisorCollectors page.
func (p *DescribeFleetAdvisorCollectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetAdvisorCollectorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeFleetAdvisorCollectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFleetAdvisorCollectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFleetAdvisorCollectors",
	}
}

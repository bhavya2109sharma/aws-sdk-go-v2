// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns compliance details for each rule in that conformance pack.
//
// You must provide exact rule names.
func (c *Client) DescribeConformancePackCompliance(ctx context.Context, params *DescribeConformancePackComplianceInput, optFns ...func(*Options)) (*DescribeConformancePackComplianceOutput, error) {
	if params == nil {
		params = &DescribeConformancePackComplianceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConformancePackCompliance", params, optFns, c.addOperationDescribeConformancePackComplianceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConformancePackComplianceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConformancePackComplianceInput struct {

	// Name of the conformance pack.
	//
	// This member is required.
	ConformancePackName *string

	// A ConformancePackComplianceFilters object.
	Filters *types.ConformancePackComplianceFilters

	// The maximum number of Config rules within a conformance pack are returned on
	// each page.
	Limit int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeConformancePackComplianceOutput struct {

	// Name of the conformance pack.
	//
	// This member is required.
	ConformancePackName *string

	// Returns a list of ConformancePackRuleCompliance objects.
	//
	// This member is required.
	ConformancePackRuleComplianceList []types.ConformancePackRuleCompliance

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConformancePackComplianceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConformancePackCompliance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConformancePackCompliance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeConformancePackCompliance"); err != nil {
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
	if err = addOpDescribeConformancePackComplianceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConformancePackCompliance(options.Region), middleware.Before); err != nil {
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

// DescribeConformancePackComplianceAPIClient is a client that implements the
// DescribeConformancePackCompliance operation.
type DescribeConformancePackComplianceAPIClient interface {
	DescribeConformancePackCompliance(context.Context, *DescribeConformancePackComplianceInput, ...func(*Options)) (*DescribeConformancePackComplianceOutput, error)
}

var _ DescribeConformancePackComplianceAPIClient = (*Client)(nil)

// DescribeConformancePackCompliancePaginatorOptions is the paginator options for
// DescribeConformancePackCompliance
type DescribeConformancePackCompliancePaginatorOptions struct {
	// The maximum number of Config rules within a conformance pack are returned on
	// each page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConformancePackCompliancePaginator is a paginator for
// DescribeConformancePackCompliance
type DescribeConformancePackCompliancePaginator struct {
	options   DescribeConformancePackCompliancePaginatorOptions
	client    DescribeConformancePackComplianceAPIClient
	params    *DescribeConformancePackComplianceInput
	nextToken *string
	firstPage bool
}

// NewDescribeConformancePackCompliancePaginator returns a new
// DescribeConformancePackCompliancePaginator
func NewDescribeConformancePackCompliancePaginator(client DescribeConformancePackComplianceAPIClient, params *DescribeConformancePackComplianceInput, optFns ...func(*DescribeConformancePackCompliancePaginatorOptions)) *DescribeConformancePackCompliancePaginator {
	if params == nil {
		params = &DescribeConformancePackComplianceInput{}
	}

	options := DescribeConformancePackCompliancePaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConformancePackCompliancePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConformancePackCompliancePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeConformancePackCompliance page.
func (p *DescribeConformancePackCompliancePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConformancePackComplianceOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeConformancePackCompliance(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeConformancePackCompliance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeConformancePackCompliance",
	}
}

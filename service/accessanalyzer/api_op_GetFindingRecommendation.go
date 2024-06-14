// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a finding recommendation for the specified analyzer.
func (c *Client) GetFindingRecommendation(ctx context.Context, params *GetFindingRecommendationInput, optFns ...func(*Options)) (*GetFindingRecommendationOutput, error) {
	if params == nil {
		params = &GetFindingRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFindingRecommendation", params, optFns, c.addOperationGetFindingRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFindingRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingRecommendationInput struct {

	// The [ARN of the analyzer] used to generate the finding recommendation.
	//
	// [ARN of the analyzer]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources
	//
	// This member is required.
	AnalyzerArn *string

	// The unique ID for the finding recommendation.
	//
	// This member is required.
	Id *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned.
	NextToken *string

	noSmithyDocumentSerde
}

type GetFindingRecommendationOutput struct {

	// The type of recommendation for the finding.
	//
	// This member is required.
	RecommendationType types.RecommendationType

	// The ARN of the resource of the finding.
	//
	// This member is required.
	ResourceArn *string

	// The time at which the retrieval of the finding recommendation was started.
	//
	// This member is required.
	StartedAt *time.Time

	// The status of the retrieval of the finding recommendation.
	//
	// This member is required.
	Status types.Status

	// The time at which the retrieval of the finding recommendation was completed.
	CompletedAt *time.Time

	// Detailed information about the reason that the retrieval of a recommendation
	// for the finding failed.
	Error *types.RecommendationError

	// A token used for pagination of results returned.
	NextToken *string

	// A group of recommended steps for the finding.
	RecommendedSteps []types.RecommendedStep

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFindingRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindingRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindingRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFindingRecommendation"); err != nil {
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
	if err = addOpGetFindingRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindingRecommendation(options.Region), middleware.Before); err != nil {
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

// GetFindingRecommendationAPIClient is a client that implements the
// GetFindingRecommendation operation.
type GetFindingRecommendationAPIClient interface {
	GetFindingRecommendation(context.Context, *GetFindingRecommendationInput, ...func(*Options)) (*GetFindingRecommendationOutput, error)
}

var _ GetFindingRecommendationAPIClient = (*Client)(nil)

// GetFindingRecommendationPaginatorOptions is the paginator options for
// GetFindingRecommendation
type GetFindingRecommendationPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetFindingRecommendationPaginator is a paginator for GetFindingRecommendation
type GetFindingRecommendationPaginator struct {
	options   GetFindingRecommendationPaginatorOptions
	client    GetFindingRecommendationAPIClient
	params    *GetFindingRecommendationInput
	nextToken *string
	firstPage bool
}

// NewGetFindingRecommendationPaginator returns a new
// GetFindingRecommendationPaginator
func NewGetFindingRecommendationPaginator(client GetFindingRecommendationAPIClient, params *GetFindingRecommendationInput, optFns ...func(*GetFindingRecommendationPaginatorOptions)) *GetFindingRecommendationPaginator {
	if params == nil {
		params = &GetFindingRecommendationInput{}
	}

	options := GetFindingRecommendationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetFindingRecommendationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetFindingRecommendationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetFindingRecommendation page.
func (p *GetFindingRecommendationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetFindingRecommendationOutput, error) {
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

	result, err := p.client.GetFindingRecommendation(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetFindingRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFindingRecommendation",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of measures that are potential causes or effects of an anomaly
// group.
func (c *Client) ListAnomalyGroupRelatedMetrics(ctx context.Context, params *ListAnomalyGroupRelatedMetricsInput, optFns ...func(*Options)) (*ListAnomalyGroupRelatedMetricsOutput, error) {
	if params == nil {
		params = &ListAnomalyGroupRelatedMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnomalyGroupRelatedMetrics", params, optFns, c.addOperationListAnomalyGroupRelatedMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnomalyGroupRelatedMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnomalyGroupRelatedMetricsInput struct {

	// The Amazon Resource Name (ARN) of the anomaly detector.
	//
	// This member is required.
	AnomalyDetectorArn *string

	// The ID of the anomaly group.
	//
	// This member is required.
	AnomalyGroupId *string

	// The maximum number of results to return.
	MaxResults *int32

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	NextToken *string

	// Filter for potential causes ( CAUSE_OF_INPUT_ANOMALY_GROUP ) or downstream
	// effects ( EFFECT_OF_INPUT_ANOMALY_GROUP ) of the anomaly group.
	RelationshipTypeFilter types.RelationshipType

	noSmithyDocumentSerde
}

type ListAnomalyGroupRelatedMetricsOutput struct {

	// Aggregated details about the measures contributing to the anomaly group, and
	// the measures potentially impacted by the anomaly group.
	InterMetricImpactList []types.InterMetricImpactDetails

	// The pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnomalyGroupRelatedMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnomalyGroupRelatedMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnomalyGroupRelatedMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAnomalyGroupRelatedMetrics"); err != nil {
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
	if err = addOpListAnomalyGroupRelatedMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnomalyGroupRelatedMetrics(options.Region), middleware.Before); err != nil {
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

// ListAnomalyGroupRelatedMetricsAPIClient is a client that implements the
// ListAnomalyGroupRelatedMetrics operation.
type ListAnomalyGroupRelatedMetricsAPIClient interface {
	ListAnomalyGroupRelatedMetrics(context.Context, *ListAnomalyGroupRelatedMetricsInput, ...func(*Options)) (*ListAnomalyGroupRelatedMetricsOutput, error)
}

var _ ListAnomalyGroupRelatedMetricsAPIClient = (*Client)(nil)

// ListAnomalyGroupRelatedMetricsPaginatorOptions is the paginator options for
// ListAnomalyGroupRelatedMetrics
type ListAnomalyGroupRelatedMetricsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnomalyGroupRelatedMetricsPaginator is a paginator for
// ListAnomalyGroupRelatedMetrics
type ListAnomalyGroupRelatedMetricsPaginator struct {
	options   ListAnomalyGroupRelatedMetricsPaginatorOptions
	client    ListAnomalyGroupRelatedMetricsAPIClient
	params    *ListAnomalyGroupRelatedMetricsInput
	nextToken *string
	firstPage bool
}

// NewListAnomalyGroupRelatedMetricsPaginator returns a new
// ListAnomalyGroupRelatedMetricsPaginator
func NewListAnomalyGroupRelatedMetricsPaginator(client ListAnomalyGroupRelatedMetricsAPIClient, params *ListAnomalyGroupRelatedMetricsInput, optFns ...func(*ListAnomalyGroupRelatedMetricsPaginatorOptions)) *ListAnomalyGroupRelatedMetricsPaginator {
	if params == nil {
		params = &ListAnomalyGroupRelatedMetricsInput{}
	}

	options := ListAnomalyGroupRelatedMetricsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnomalyGroupRelatedMetricsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnomalyGroupRelatedMetricsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnomalyGroupRelatedMetrics page.
func (p *ListAnomalyGroupRelatedMetricsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnomalyGroupRelatedMetricsOutput, error) {
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

	result, err := p.client.ListAnomalyGroupRelatedMetrics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAnomalyGroupRelatedMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAnomalyGroupRelatedMetrics",
	}
}

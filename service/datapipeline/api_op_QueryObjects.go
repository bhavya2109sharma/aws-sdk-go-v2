// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Queries the specified pipeline for the names of objects that match the
// specified set of conditions.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.QueryObjects Content-Length: 123 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE", "query": {"selectors": [ ] },
// "sphere": "INSTANCE", "marker": "", "limit": 10}
//
// x-amzn-RequestId: 14d704c1-0775-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 72 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"hasMoreResults": false, "ids": ["@SayHello_1_2012-09-25T17:00:00"] }
func (c *Client) QueryObjects(ctx context.Context, params *QueryObjectsInput, optFns ...func(*Options)) (*QueryObjectsOutput, error) {
	if params == nil {
		params = &QueryObjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QueryObjects", params, optFns, c.addOperationQueryObjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for QueryObjects.
type QueryObjectsInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// Indicates whether the query applies to components or instances. The possible
	// values are: COMPONENT , INSTANCE , and ATTEMPT .
	//
	// This member is required.
	Sphere *string

	// The maximum number of object names that QueryObjects will return in a single
	// call. The default value is 100.
	Limit *int32

	// The starting point for the results to be returned. For the first call, this
	// value should be empty. As long as there are more results, continue to call
	// QueryObjects with the marker value from the previous call to retrieve the next
	// set of results.
	Marker *string

	// The query that defines the objects to be returned. The Query object can contain
	// a maximum of ten selectors. The conditions in the query are limited to top-level
	// String fields in the object. These filters can be applied to components,
	// instances, and attempts.
	Query *types.Query

	noSmithyDocumentSerde
}

// Contains the output of QueryObjects.
type QueryObjectsOutput struct {

	// Indicates whether there are more results that can be obtained by a subsequent
	// call.
	HasMoreResults bool

	// The identifiers that match the query selectors.
	Ids []string

	// The starting point for the next page of results. To view the next page of
	// results, call QueryObjects again with this marker value. If the value is null,
	// there are no more results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryObjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpQueryObjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpQueryObjects{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "QueryObjects"); err != nil {
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
	if err = addOpQueryObjectsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQueryObjects(options.Region), middleware.Before); err != nil {
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

// QueryObjectsAPIClient is a client that implements the QueryObjects operation.
type QueryObjectsAPIClient interface {
	QueryObjects(context.Context, *QueryObjectsInput, ...func(*Options)) (*QueryObjectsOutput, error)
}

var _ QueryObjectsAPIClient = (*Client)(nil)

// QueryObjectsPaginatorOptions is the paginator options for QueryObjects
type QueryObjectsPaginatorOptions struct {
	// The maximum number of object names that QueryObjects will return in a single
	// call. The default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// QueryObjectsPaginator is a paginator for QueryObjects
type QueryObjectsPaginator struct {
	options   QueryObjectsPaginatorOptions
	client    QueryObjectsAPIClient
	params    *QueryObjectsInput
	nextToken *string
	firstPage bool
}

// NewQueryObjectsPaginator returns a new QueryObjectsPaginator
func NewQueryObjectsPaginator(client QueryObjectsAPIClient, params *QueryObjectsInput, optFns ...func(*QueryObjectsPaginatorOptions)) *QueryObjectsPaginator {
	if params == nil {
		params = &QueryObjectsInput{}
	}

	options := QueryObjectsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &QueryObjectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *QueryObjectsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next QueryObjects page.
func (p *QueryObjectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*QueryObjectsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.QueryObjects(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opQueryObjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "QueryObjects",
	}
}

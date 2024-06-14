// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists how the resources in an application version are mapped/sourced from.
// Mappings can be physical resource identifiers, CloudFormation stacks,
// resource-groups, or an application registry app.
func (c *Client) ListAppVersionResourceMappings(ctx context.Context, params *ListAppVersionResourceMappingsInput, optFns ...func(*Options)) (*ListAppVersionResourceMappingsOutput, error) {
	if params == nil {
		params = &ListAppVersionResourceMappingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppVersionResourceMappings", params, optFns, c.addOperationListAppVersionResourceMappingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppVersionResourceMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppVersionResourceMappingsInput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference
	// guide.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	AppArn *string

	// The version of the application.
	//
	// This member is required.
	AppVersion *string

	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppVersionResourceMappingsOutput struct {

	// Mappings used to map logical resources from the template to physical resources.
	// You can use the mapping type CFN_STACK if the application template uses a
	// logical stack name. Or you can map individual resources by using the mapping
	// type RESOURCE . We recommend using the mapping type CFN_STACK if the
	// application is backed by a CloudFormation stack.
	//
	// This member is required.
	ResourceMappings []types.ResourceMapping

	// Token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppVersionResourceMappingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppVersionResourceMappings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppVersionResourceMappings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAppVersionResourceMappings"); err != nil {
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
	if err = addOpListAppVersionResourceMappingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppVersionResourceMappings(options.Region), middleware.Before); err != nil {
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

// ListAppVersionResourceMappingsAPIClient is a client that implements the
// ListAppVersionResourceMappings operation.
type ListAppVersionResourceMappingsAPIClient interface {
	ListAppVersionResourceMappings(context.Context, *ListAppVersionResourceMappingsInput, ...func(*Options)) (*ListAppVersionResourceMappingsOutput, error)
}

var _ ListAppVersionResourceMappingsAPIClient = (*Client)(nil)

// ListAppVersionResourceMappingsPaginatorOptions is the paginator options for
// ListAppVersionResourceMappings
type ListAppVersionResourceMappingsPaginatorOptions struct {
	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppVersionResourceMappingsPaginator is a paginator for
// ListAppVersionResourceMappings
type ListAppVersionResourceMappingsPaginator struct {
	options   ListAppVersionResourceMappingsPaginatorOptions
	client    ListAppVersionResourceMappingsAPIClient
	params    *ListAppVersionResourceMappingsInput
	nextToken *string
	firstPage bool
}

// NewListAppVersionResourceMappingsPaginator returns a new
// ListAppVersionResourceMappingsPaginator
func NewListAppVersionResourceMappingsPaginator(client ListAppVersionResourceMappingsAPIClient, params *ListAppVersionResourceMappingsInput, optFns ...func(*ListAppVersionResourceMappingsPaginatorOptions)) *ListAppVersionResourceMappingsPaginator {
	if params == nil {
		params = &ListAppVersionResourceMappingsInput{}
	}

	options := ListAppVersionResourceMappingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppVersionResourceMappingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppVersionResourceMappingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppVersionResourceMappings page.
func (p *ListAppVersionResourceMappingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppVersionResourceMappingsOutput, error) {
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

	result, err := p.client.ListAppVersionResourceMappings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppVersionResourceMappings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAppVersionResourceMappings",
	}
}

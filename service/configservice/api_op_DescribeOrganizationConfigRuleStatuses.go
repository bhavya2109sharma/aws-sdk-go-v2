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

// Provides organization Config rule deployment status for an organization.
//
// The status is not considered successful until organization Config rule is
// successfully deployed in all the member accounts with an exception of excluded
// accounts.
//
// When you specify the limit and the next token, you receive a paginated
// response. Limit and next token are not applicable if you specify organization
// Config rule names. It is only applicable, when you request all the organization
// Config rules.
func (c *Client) DescribeOrganizationConfigRuleStatuses(ctx context.Context, params *DescribeOrganizationConfigRuleStatusesInput, optFns ...func(*Options)) (*DescribeOrganizationConfigRuleStatusesOutput, error) {
	if params == nil {
		params = &DescribeOrganizationConfigRuleStatusesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationConfigRuleStatuses", params, optFns, c.addOperationDescribeOrganizationConfigRuleStatusesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationConfigRuleStatusesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConfigRuleStatusesInput struct {

	// The maximum number of OrganizationConfigRuleStatuses returned on each page. If
	// you do no specify a number, Config uses the default. The default is 100.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The names of organization Config rules for which you want status details. If
	// you do not specify any names, Config returns details for all your organization
	// Config rules.
	OrganizationConfigRuleNames []string

	noSmithyDocumentSerde
}

type DescribeOrganizationConfigRuleStatusesOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of OrganizationConfigRuleStatus objects.
	OrganizationConfigRuleStatuses []types.OrganizationConfigRuleStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationConfigRuleStatusesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationConfigRuleStatuses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationConfigRuleStatuses{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeOrganizationConfigRuleStatuses"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationConfigRuleStatuses(options.Region), middleware.Before); err != nil {
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

// DescribeOrganizationConfigRuleStatusesAPIClient is a client that implements the
// DescribeOrganizationConfigRuleStatuses operation.
type DescribeOrganizationConfigRuleStatusesAPIClient interface {
	DescribeOrganizationConfigRuleStatuses(context.Context, *DescribeOrganizationConfigRuleStatusesInput, ...func(*Options)) (*DescribeOrganizationConfigRuleStatusesOutput, error)
}

var _ DescribeOrganizationConfigRuleStatusesAPIClient = (*Client)(nil)

// DescribeOrganizationConfigRuleStatusesPaginatorOptions is the paginator options
// for DescribeOrganizationConfigRuleStatuses
type DescribeOrganizationConfigRuleStatusesPaginatorOptions struct {
	// The maximum number of OrganizationConfigRuleStatuses returned on each page. If
	// you do no specify a number, Config uses the default. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOrganizationConfigRuleStatusesPaginator is a paginator for
// DescribeOrganizationConfigRuleStatuses
type DescribeOrganizationConfigRuleStatusesPaginator struct {
	options   DescribeOrganizationConfigRuleStatusesPaginatorOptions
	client    DescribeOrganizationConfigRuleStatusesAPIClient
	params    *DescribeOrganizationConfigRuleStatusesInput
	nextToken *string
	firstPage bool
}

// NewDescribeOrganizationConfigRuleStatusesPaginator returns a new
// DescribeOrganizationConfigRuleStatusesPaginator
func NewDescribeOrganizationConfigRuleStatusesPaginator(client DescribeOrganizationConfigRuleStatusesAPIClient, params *DescribeOrganizationConfigRuleStatusesInput, optFns ...func(*DescribeOrganizationConfigRuleStatusesPaginatorOptions)) *DescribeOrganizationConfigRuleStatusesPaginator {
	if params == nil {
		params = &DescribeOrganizationConfigRuleStatusesInput{}
	}

	options := DescribeOrganizationConfigRuleStatusesPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOrganizationConfigRuleStatusesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrganizationConfigRuleStatusesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOrganizationConfigRuleStatuses page.
func (p *DescribeOrganizationConfigRuleStatusesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrganizationConfigRuleStatusesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeOrganizationConfigRuleStatuses(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOrganizationConfigRuleStatuses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeOrganizationConfigRuleStatuses",
	}
}

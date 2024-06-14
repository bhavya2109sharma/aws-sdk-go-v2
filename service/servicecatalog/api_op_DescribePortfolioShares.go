// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a summary of each of the portfolio shares that were created for the
// specified portfolio.
//
// You can use this API to determine which accounts or organizational nodes this
// portfolio have been shared, whether the recipient entity has imported the share,
// and whether TagOptions are included with the share.
//
// The PortfolioId and Type parameters are both required.
func (c *Client) DescribePortfolioShares(ctx context.Context, params *DescribePortfolioSharesInput, optFns ...func(*Options)) (*DescribePortfolioSharesOutput, error) {
	if params == nil {
		params = &DescribePortfolioSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePortfolioShares", params, optFns, c.addOperationDescribePortfolioSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePortfolioSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePortfolioSharesInput struct {

	// The unique identifier of the portfolio for which shares will be retrieved.
	//
	// This member is required.
	PortfolioId *string

	// The type of portfolio share to summarize. This field acts as a filter on the
	// type of portfolio share, which can be one of the following:
	//
	// 1. ACCOUNT - Represents an external account to account share.
	//
	// 2. ORGANIZATION - Represents a share to an organization. This share is
	// available to every account in the organization.
	//
	// 3. ORGANIZATIONAL_UNIT - Represents a share to an organizational unit.
	//
	// 4. ORGANIZATION_MEMBER_ACCOUNT - Represents a share to an account in the
	// organization.
	//
	// This member is required.
	Type types.DescribePortfolioShareType

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	noSmithyDocumentSerde
}

type DescribePortfolioSharesOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Summaries about each of the portfolio shares.
	PortfolioShareDetails []types.PortfolioShareDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePortfolioSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePortfolioShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePortfolioShares{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePortfolioShares"); err != nil {
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
	if err = addOpDescribePortfolioSharesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePortfolioShares(options.Region), middleware.Before); err != nil {
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

// DescribePortfolioSharesAPIClient is a client that implements the
// DescribePortfolioShares operation.
type DescribePortfolioSharesAPIClient interface {
	DescribePortfolioShares(context.Context, *DescribePortfolioSharesInput, ...func(*Options)) (*DescribePortfolioSharesOutput, error)
}

var _ DescribePortfolioSharesAPIClient = (*Client)(nil)

// DescribePortfolioSharesPaginatorOptions is the paginator options for
// DescribePortfolioShares
type DescribePortfolioSharesPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePortfolioSharesPaginator is a paginator for DescribePortfolioShares
type DescribePortfolioSharesPaginator struct {
	options   DescribePortfolioSharesPaginatorOptions
	client    DescribePortfolioSharesAPIClient
	params    *DescribePortfolioSharesInput
	nextToken *string
	firstPage bool
}

// NewDescribePortfolioSharesPaginator returns a new
// DescribePortfolioSharesPaginator
func NewDescribePortfolioSharesPaginator(client DescribePortfolioSharesAPIClient, params *DescribePortfolioSharesInput, optFns ...func(*DescribePortfolioSharesPaginatorOptions)) *DescribePortfolioSharesPaginator {
	if params == nil {
		params = &DescribePortfolioSharesInput{}
	}

	options := DescribePortfolioSharesPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePortfolioSharesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePortfolioSharesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePortfolioShares page.
func (p *DescribePortfolioSharesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePortfolioSharesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.DescribePortfolioShares(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribePortfolioShares(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePortfolioShares",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get all users in a given launch profile membership.
func (c *Client) ListLaunchProfileMembers(ctx context.Context, params *ListLaunchProfileMembersInput, optFns ...func(*Options)) (*ListLaunchProfileMembersOutput, error) {
	if params == nil {
		params = &ListLaunchProfileMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLaunchProfileMembers", params, optFns, c.addOperationListLaunchProfileMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLaunchProfileMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLaunchProfileMembersInput struct {

	// The ID of the launch profile used to control access from the streaming session.
	//
	// This member is required.
	LaunchProfileId *string

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// The max number of results to return in the response.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLaunchProfileMembersOutput struct {

	// A list of members.
	Members []types.LaunchProfileMembership

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLaunchProfileMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLaunchProfileMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLaunchProfileMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLaunchProfileMembers"); err != nil {
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
	if err = addOpListLaunchProfileMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLaunchProfileMembers(options.Region), middleware.Before); err != nil {
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

// ListLaunchProfileMembersAPIClient is a client that implements the
// ListLaunchProfileMembers operation.
type ListLaunchProfileMembersAPIClient interface {
	ListLaunchProfileMembers(context.Context, *ListLaunchProfileMembersInput, ...func(*Options)) (*ListLaunchProfileMembersOutput, error)
}

var _ ListLaunchProfileMembersAPIClient = (*Client)(nil)

// ListLaunchProfileMembersPaginatorOptions is the paginator options for
// ListLaunchProfileMembers
type ListLaunchProfileMembersPaginatorOptions struct {
	// The max number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLaunchProfileMembersPaginator is a paginator for ListLaunchProfileMembers
type ListLaunchProfileMembersPaginator struct {
	options   ListLaunchProfileMembersPaginatorOptions
	client    ListLaunchProfileMembersAPIClient
	params    *ListLaunchProfileMembersInput
	nextToken *string
	firstPage bool
}

// NewListLaunchProfileMembersPaginator returns a new
// ListLaunchProfileMembersPaginator
func NewListLaunchProfileMembersPaginator(client ListLaunchProfileMembersAPIClient, params *ListLaunchProfileMembersInput, optFns ...func(*ListLaunchProfileMembersPaginatorOptions)) *ListLaunchProfileMembersPaginator {
	if params == nil {
		params = &ListLaunchProfileMembersInput{}
	}

	options := ListLaunchProfileMembersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLaunchProfileMembersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLaunchProfileMembersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLaunchProfileMembers page.
func (p *ListLaunchProfileMembersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLaunchProfileMembersOutput, error) {
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

	result, err := p.client.ListLaunchProfileMembers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLaunchProfileMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLaunchProfileMembers",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array containing metadata of the resources associated with the
// target backup plan.
func (c *Client) ListBackupSelections(ctx context.Context, params *ListBackupSelectionsInput, optFns ...func(*Options)) (*ListBackupSelectionsOutput, error) {
	if params == nil {
		params = &ListBackupSelectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackupSelections", params, optFns, c.addOperationListBackupSelectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackupSelectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupSelectionsInput struct {

	// Uniquely identifies a backup plan.
	//
	// This member is required.
	BackupPlanId *string

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBackupSelectionsOutput struct {

	// An array of backup selection list items containing metadata about each resource
	// in the list.
	BackupSelectionsList []types.BackupSelectionsListMember

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBackupSelectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackupSelections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackupSelections{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBackupSelections"); err != nil {
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
	if err = addOpListBackupSelectionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackupSelections(options.Region), middleware.Before); err != nil {
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

// ListBackupSelectionsAPIClient is a client that implements the
// ListBackupSelections operation.
type ListBackupSelectionsAPIClient interface {
	ListBackupSelections(context.Context, *ListBackupSelectionsInput, ...func(*Options)) (*ListBackupSelectionsOutput, error)
}

var _ ListBackupSelectionsAPIClient = (*Client)(nil)

// ListBackupSelectionsPaginatorOptions is the paginator options for
// ListBackupSelections
type ListBackupSelectionsPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBackupSelectionsPaginator is a paginator for ListBackupSelections
type ListBackupSelectionsPaginator struct {
	options   ListBackupSelectionsPaginatorOptions
	client    ListBackupSelectionsAPIClient
	params    *ListBackupSelectionsInput
	nextToken *string
	firstPage bool
}

// NewListBackupSelectionsPaginator returns a new ListBackupSelectionsPaginator
func NewListBackupSelectionsPaginator(client ListBackupSelectionsAPIClient, params *ListBackupSelectionsInput, optFns ...func(*ListBackupSelectionsPaginatorOptions)) *ListBackupSelectionsPaginator {
	if params == nil {
		params = &ListBackupSelectionsInput{}
	}

	options := ListBackupSelectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBackupSelectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBackupSelectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBackupSelections page.
func (p *ListBackupSelectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBackupSelectionsOutput, error) {
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

	result, err := p.client.ListBackupSelections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBackupSelections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBackupSelections",
	}
}

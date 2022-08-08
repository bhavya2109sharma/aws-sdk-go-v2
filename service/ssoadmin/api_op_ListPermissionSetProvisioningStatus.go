// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the status of the permission set provisioning requests for a specified
// Amazon Web Services SSO instance.
func (c *Client) ListPermissionSetProvisioningStatus(ctx context.Context, params *ListPermissionSetProvisioningStatusInput, optFns ...func(*Options)) (*ListPermissionSetProvisioningStatusOutput, error) {
	if params == nil {
		params = &ListPermissionSetProvisioningStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissionSetProvisioningStatus", params, optFns, c.addOperationListPermissionSetProvisioningStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionSetProvisioningStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionSetProvisioningStatusInput struct {

	// The ARN of the Amazon Web Services SSO instance under which the operation will
	// be executed. For more information about ARNs, see Amazon Resource Names (ARNs)
	// and Amazon Web Services Service Namespaces in the Amazon Web Services General
	// Reference.
	//
	// This member is required.
	InstanceArn *string

	// Filters results based on the passed attribute value.
	Filter *types.OperationStatusFilter

	// The maximum number of results to display for the assignment.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPermissionSetProvisioningStatusOutput struct {

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// The status object for the permission set provisioning operation.
	PermissionSetsProvisioningStatus []types.PermissionSetProvisioningStatusMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPermissionSetProvisioningStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPermissionSetProvisioningStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPermissionSetProvisioningStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListPermissionSetProvisioningStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissionSetProvisioningStatus(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListPermissionSetProvisioningStatusAPIClient is a client that implements the
// ListPermissionSetProvisioningStatus operation.
type ListPermissionSetProvisioningStatusAPIClient interface {
	ListPermissionSetProvisioningStatus(context.Context, *ListPermissionSetProvisioningStatusInput, ...func(*Options)) (*ListPermissionSetProvisioningStatusOutput, error)
}

var _ ListPermissionSetProvisioningStatusAPIClient = (*Client)(nil)

// ListPermissionSetProvisioningStatusPaginatorOptions is the paginator options for
// ListPermissionSetProvisioningStatus
type ListPermissionSetProvisioningStatusPaginatorOptions struct {
	// The maximum number of results to display for the assignment.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPermissionSetProvisioningStatusPaginator is a paginator for
// ListPermissionSetProvisioningStatus
type ListPermissionSetProvisioningStatusPaginator struct {
	options   ListPermissionSetProvisioningStatusPaginatorOptions
	client    ListPermissionSetProvisioningStatusAPIClient
	params    *ListPermissionSetProvisioningStatusInput
	nextToken *string
	firstPage bool
}

// NewListPermissionSetProvisioningStatusPaginator returns a new
// ListPermissionSetProvisioningStatusPaginator
func NewListPermissionSetProvisioningStatusPaginator(client ListPermissionSetProvisioningStatusAPIClient, params *ListPermissionSetProvisioningStatusInput, optFns ...func(*ListPermissionSetProvisioningStatusPaginatorOptions)) *ListPermissionSetProvisioningStatusPaginator {
	if params == nil {
		params = &ListPermissionSetProvisioningStatusInput{}
	}

	options := ListPermissionSetProvisioningStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPermissionSetProvisioningStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPermissionSetProvisioningStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPermissionSetProvisioningStatus page.
func (p *ListPermissionSetProvisioningStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPermissionSetProvisioningStatusOutput, error) {
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

	result, err := p.client.ListPermissionSetProvisioningStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPermissionSetProvisioningStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "ListPermissionSetProvisioningStatus",
	}
}

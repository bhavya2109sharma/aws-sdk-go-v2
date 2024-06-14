// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the principals that you are sharing resources with or that are sharing
// resources with you.
func (c *Client) ListPrincipals(ctx context.Context, params *ListPrincipalsInput, optFns ...func(*Options)) (*ListPrincipalsOutput, error) {
	if params == nil {
		params = &ListPrincipalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPrincipals", params, optFns, c.addOperationListPrincipalsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPrincipalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrincipalsInput struct {

	// Specifies that you want to list information for only resource shares that match
	// the following:
	//
	//   - SELF – principals that your account is sharing resources with
	//
	//   - OTHER-ACCOUNTS – principals that are sharing resources with your account
	//
	// This member is required.
	ResourceOwner types.ResourceOwner

	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	// Specifies that you want to list information for only the listed principals.
	//
	// You can include the following values:
	//
	//   - An Amazon Web Services account ID, for example: 123456789012
	//
	//   - An [Amazon Resource Name (ARN)]of an organization in Organizations, for example:
	//   organizations::123456789012:organization/o-exampleorgid
	//
	//   - An ARN of an organizational unit (OU) in Organizations, for example:
	//   organizations::123456789012:ou/o-exampleorgid/ou-examplerootid-exampleouid123
	//
	//   - An ARN of an IAM role, for example: iam::123456789012:role/rolename
	//
	//   - An ARN of an IAM user, for example: iam::123456789012user/username
	//
	// Not all resource types can be shared with IAM roles and users. For more
	// information, see [Sharing with IAM roles and users]in the Resource Access Manager User Guide.
	//
	// [Amazon Resource Name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	// [Sharing with IAM roles and users]: https://docs.aws.amazon.com/ram/latest/userguide/permissions.html#permissions-rbp-supported-resource-types
	Principals []string

	// Specifies that you want to list principal information for the resource share
	// with the specified [Amazon Resource Name (ARN)].
	//
	// [Amazon Resource Name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	ResourceArn *string

	// Specifies that you want to list information for only principals associated with
	// the resource shares specified by a list the [Amazon Resource Names (ARNs)].
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	ResourceShareArns []string

	// Specifies that you want to list information for only principals associated with
	// resource shares that include the specified resource type.
	//
	// For a list of valid values, query the ListResourceTypes operation.
	ResourceType *string

	noSmithyDocumentSerde
}

type ListPrincipalsOutput struct {

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// An array of objects that contain the details about the principals.
	Principals []types.Principal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPrincipalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPrincipals{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrincipals{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPrincipals"); err != nil {
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
	if err = addOpListPrincipalsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrincipals(options.Region), middleware.Before); err != nil {
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

// ListPrincipalsAPIClient is a client that implements the ListPrincipals
// operation.
type ListPrincipalsAPIClient interface {
	ListPrincipals(context.Context, *ListPrincipalsInput, ...func(*Options)) (*ListPrincipalsOutput, error)
}

var _ ListPrincipalsAPIClient = (*Client)(nil)

// ListPrincipalsPaginatorOptions is the paginator options for ListPrincipals
type ListPrincipalsPaginatorOptions struct {
	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPrincipalsPaginator is a paginator for ListPrincipals
type ListPrincipalsPaginator struct {
	options   ListPrincipalsPaginatorOptions
	client    ListPrincipalsAPIClient
	params    *ListPrincipalsInput
	nextToken *string
	firstPage bool
}

// NewListPrincipalsPaginator returns a new ListPrincipalsPaginator
func NewListPrincipalsPaginator(client ListPrincipalsAPIClient, params *ListPrincipalsInput, optFns ...func(*ListPrincipalsPaginatorOptions)) *ListPrincipalsPaginator {
	if params == nil {
		params = &ListPrincipalsInput{}
	}

	options := ListPrincipalsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPrincipalsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPrincipalsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPrincipals page.
func (p *ListPrincipalsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPrincipalsOutput, error) {
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

	result, err := p.client.ListPrincipals(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPrincipals(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPrincipals",
	}
}

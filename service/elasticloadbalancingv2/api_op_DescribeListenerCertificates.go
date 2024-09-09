// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the default certificate and the certificate list for the specified
// HTTPS or TLS listener.
//
// If the default certificate is also in the certificate list, it appears twice in
// the results (once with IsDefault set to true and once with IsDefault set to
// false).
//
// For more information, see [SSL certificates] in the Application Load Balancers Guide or [Server certificates] in the
// Network Load Balancers Guide.
//
// [Server certificates]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#tls-listener-certificate
// [SSL certificates]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#https-listener-certificates
func (c *Client) DescribeListenerCertificates(ctx context.Context, params *DescribeListenerCertificatesInput, optFns ...func(*Options)) (*DescribeListenerCertificatesOutput, error) {
	if params == nil {
		params = &DescribeListenerCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeListenerCertificates", params, optFns, c.addOperationDescribeListenerCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeListenerCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeListenerCertificatesInput struct {

	// The Amazon Resource Names (ARN) of the listener.
	//
	// This member is required.
	ListenerArn *string

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The maximum number of results to return with this call.
	PageSize *int32

	noSmithyDocumentSerde
}

type DescribeListenerCertificatesOutput struct {

	// Information about the certificates.
	Certificates []types.Certificate

	// If there are additional results, this is the marker for the next set of
	// results. Otherwise, this is null.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeListenerCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeListenerCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeListenerCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeListenerCertificates"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeListenerCertificatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeListenerCertificates(options.Region), middleware.Before); err != nil {
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

// DescribeListenerCertificatesPaginatorOptions is the paginator options for
// DescribeListenerCertificates
type DescribeListenerCertificatesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeListenerCertificatesPaginator is a paginator for
// DescribeListenerCertificates
type DescribeListenerCertificatesPaginator struct {
	options   DescribeListenerCertificatesPaginatorOptions
	client    DescribeListenerCertificatesAPIClient
	params    *DescribeListenerCertificatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeListenerCertificatesPaginator returns a new
// DescribeListenerCertificatesPaginator
func NewDescribeListenerCertificatesPaginator(client DescribeListenerCertificatesAPIClient, params *DescribeListenerCertificatesInput, optFns ...func(*DescribeListenerCertificatesPaginatorOptions)) *DescribeListenerCertificatesPaginator {
	if params == nil {
		params = &DescribeListenerCertificatesInput{}
	}

	options := DescribeListenerCertificatesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeListenerCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeListenerCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeListenerCertificates page.
func (p *DescribeListenerCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeListenerCertificatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeListenerCertificates(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeListenerCertificatesAPIClient is a client that implements the
// DescribeListenerCertificates operation.
type DescribeListenerCertificatesAPIClient interface {
	DescribeListenerCertificates(context.Context, *DescribeListenerCertificatesInput, ...func(*Options)) (*DescribeListenerCertificatesOutput, error)
}

var _ DescribeListenerCertificatesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeListenerCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeListenerCertificates",
	}
}

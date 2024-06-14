// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all Amazon OpenSearch Service-managed VPC endpoints associated with a
// particular domain.
func (c *Client) ListVpcEndpointsForDomain(ctx context.Context, params *ListVpcEndpointsForDomainInput, optFns ...func(*Options)) (*ListVpcEndpointsForDomainOutput, error) {
	if params == nil {
		params = &ListVpcEndpointsForDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVpcEndpointsForDomain", params, optFns, c.addOperationListVpcEndpointsForDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVpcEndpointsForDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVpcEndpointsForDomainInput struct {

	// The name of the domain to list associated VPC endpoints for.
	//
	// This member is required.
	DomainName *string

	// If your initial ListEndpointsForDomain operation returns a nextToken , you can
	// include the returned nextToken in subsequent ListEndpointsForDomain operations,
	// which returns results in the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVpcEndpointsForDomainOutput struct {

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Send the request again
	// using the returned token to retrieve the next page.
	//
	// This member is required.
	NextToken *string

	// Information about each endpoint associated with the domain.
	//
	// This member is required.
	VpcEndpointSummaryList []types.VpcEndpointSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVpcEndpointsForDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVpcEndpointsForDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVpcEndpointsForDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVpcEndpointsForDomain"); err != nil {
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
	if err = addOpListVpcEndpointsForDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVpcEndpointsForDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListVpcEndpointsForDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVpcEndpointsForDomain",
	}
}

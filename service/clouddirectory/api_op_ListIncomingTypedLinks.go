// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of all the incoming TypedLinkSpecifier information for an object. It
// also supports filtering by typed link facet and identity attributes. For more
// information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func (c *Client) ListIncomingTypedLinks(ctx context.Context, params *ListIncomingTypedLinksInput, optFns ...func(*Options)) (*ListIncomingTypedLinksOutput, error) {
	if params == nil {
		params = &ListIncomingTypedLinksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIncomingTypedLinks", params, optFns, c.addOperationListIncomingTypedLinksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIncomingTypedLinksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIncomingTypedLinksInput struct {

	// The Amazon Resource Name (ARN) of the directory where you want to list the
	// typed links.
	//
	// This member is required.
	DirectoryArn *string

	// Reference that identifies the object whose attributes will be listed.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// The consistency level to execute the request at.
	ConsistencyLevel types.ConsistencyLevel

	// Provides range filters for multiple attributes. When providing ranges to typed
	// link selection, any inexact ranges must be specified at the end. Any attributes
	// that do not have a range specified are presumed to match the entire range.
	FilterAttributeRanges []types.TypedLinkAttributeRange

	// Filters are interpreted in the order of the attributes on the typed link facet,
	// not the order in which they are supplied to any API calls.
	FilterTypedLink *types.TypedLinkSchemaAndFacetName

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListIncomingTypedLinksOutput struct {

	// Returns one or more typed link specifiers as output.
	LinkSpecifiers []types.TypedLinkSpecifier

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIncomingTypedLinksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIncomingTypedLinks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIncomingTypedLinks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIncomingTypedLinks"); err != nil {
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
	if err = addOpListIncomingTypedLinksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIncomingTypedLinks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListIncomingTypedLinks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIncomingTypedLinks",
	}
}

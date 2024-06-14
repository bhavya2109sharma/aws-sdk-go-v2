// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists tags for up to 10 health checks or hosted zones.
//
// For information about using tags for cost allocation, see [Using Cost Allocation Tags] in the Billing and
// Cost Management User Guide.
//
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func (c *Client) ListTagsForResources(ctx context.Context, params *ListTagsForResourcesInput, optFns ...func(*Options)) (*ListTagsForResourcesOutput, error) {
	if params == nil {
		params = &ListTagsForResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForResources", params, optFns, c.addOperationListTagsForResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the health checks or hosted
// zones for which you want to list tags.
type ListTagsForResourcesInput struct {

	// A complex type that contains the ResourceId element for each resource for which
	// you want to get a list of tags.
	//
	// This member is required.
	ResourceIds []string

	// The type of the resources.
	//
	//   - The resource type for health checks is healthcheck .
	//
	//   - The resource type for hosted zones is hostedzone .
	//
	// This member is required.
	ResourceType types.TagResourceType

	noSmithyDocumentSerde
}

// A complex type containing tags for the specified resources.
type ListTagsForResourcesOutput struct {

	// A list of ResourceTagSet s containing tags associated with the specified
	// resources.
	//
	// This member is required.
	ResourceTagSets []types.ResourceTagSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTagsForResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListTagsForResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListTagsForResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTagsForResources"); err != nil {
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
	if err = addOpListTagsForResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTagsForResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTagsForResources",
	}
}

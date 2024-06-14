// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Changes the type of the index from one of the following types to the other. For
// more information about indexes and the role they perform in Amazon Web Services
// Resource Explorer, see [Turning on cross-Region search by creating an aggregator index]in the Amazon Web Services Resource Explorer User Guide.
//
//   - AGGREGATOR index type
//
// The index contains information about resources from all Amazon Web Services
//
//	Regions in the Amazon Web Services account in which you've created a Resource
//	Explorer index. Resource information from all other Regions is replicated to
//	this Region's index.
//
// When you change the index type to AGGREGATOR , Resource Explorer turns on
//
//	replication of all discovered resource information from the other Amazon Web
//	Services Regions in your account to this index. You can then, from this Region
//	only, perform resource search queries that span all Amazon Web Services Regions
//	in the Amazon Web Services account. Turning on replication from all other
//	Regions is performed by asynchronous background tasks. You can check the status
//	of the asynchronous tasks by using the GetIndexoperation. When the asynchronous tasks
//	complete, the Status response of that operation changes from UPDATING to
//	ACTIVE . After that, you can start to see results from other Amazon Web
//	Services Regions in query results. However, it can take several hours for
//	replication from all other Regions to complete.
//
// You can have only one aggregator index per Amazon Web Services account. Before
//
//	you can promote a different index to be the aggregator index for the account,
//	you must first demote the existing aggregator index to type LOCAL .
//
//	- LOCAL index type
//
// The index contains information about resources in only the Amazon Web Services
//
//	Region in which the index exists. If an aggregator index in another Region
//	exists, then information in this local index is replicated to the aggregator
//	index.
//
// When you change the index type to LOCAL , Resource Explorer turns off the
//
//	replication of resource information from all other Amazon Web Services Regions
//	in the Amazon Web Services account to this Region. The aggregator index remains
//	in the UPDATING state until all replication with other Regions successfully
//	stops. You can check the status of the asynchronous task by using the GetIndex
//	operation. When Resource Explorer successfully stops all replication with other
//	Regions, the Status response of that operation changes from UPDATING to ACTIVE
//	. Separately, the resource information from other Regions that was previously
//	stored in the index is deleted within 30 days by another background task. Until
//	that asynchronous task completes, some results from other Regions can continue
//	to appear in search results.
//
// After you demote an aggregator index to a local index, you must wait 24 hours
//
//	before you can promote another index to be the new aggregator index for the
//	account.
//
// [Turning on cross-Region search by creating an aggregator index]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html
func (c *Client) UpdateIndexType(ctx context.Context, params *UpdateIndexTypeInput, optFns ...func(*Options)) (*UpdateIndexTypeOutput, error) {
	if params == nil {
		params = &UpdateIndexTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIndexType", params, optFns, c.addOperationUpdateIndexTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIndexTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIndexTypeInput struct {

	// The [Amazon resource name (ARN)] of the index that you want to update.
	//
	// [Amazon resource name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	Arn *string

	// The type of the index. To understand the difference between LOCAL and AGGREGATOR
	// , see [Turning on cross-Region search]in the Amazon Web Services Resource Explorer User Guide.
	//
	// [Turning on cross-Region search]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html
	//
	// This member is required.
	Type types.IndexType

	noSmithyDocumentSerde
}

type UpdateIndexTypeOutput struct {

	// The [Amazon resource name (ARN)] of the index that you updated.
	//
	// [Amazon resource name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	Arn *string

	// The date and timestamp when the index was last updated.
	LastUpdatedAt *time.Time

	// Indicates the state of the request to update the index. This operation is
	// asynchronous. Call the GetIndexoperation to check for changes.
	State types.IndexState

	// Specifies the type of the specified index after the operation completes.
	Type types.IndexType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIndexTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIndexType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIndexType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateIndexType"); err != nil {
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
	if err = addOpUpdateIndexTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIndexType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIndexType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateIndexType",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a batch of resources to a percentage custom line item.
func (c *Client) BatchAssociateResourcesToCustomLineItem(ctx context.Context, params *BatchAssociateResourcesToCustomLineItemInput, optFns ...func(*Options)) (*BatchAssociateResourcesToCustomLineItemOutput, error) {
	if params == nil {
		params = &BatchAssociateResourcesToCustomLineItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchAssociateResourcesToCustomLineItem", params, optFns, c.addOperationBatchAssociateResourcesToCustomLineItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchAssociateResourcesToCustomLineItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateResourcesToCustomLineItemInput struct {

	//  A list containing the ARNs of the resources to be associated.
	//
	// This member is required.
	ResourceArns []string

	//  A percentage custom line item ARN to associate the resources to.
	//
	// This member is required.
	TargetArn *string

	// The billing period range in which the custom line item request will be applied.
	BillingPeriodRange *types.CustomLineItemBillingPeriodRange

	noSmithyDocumentSerde
}

type BatchAssociateResourcesToCustomLineItemOutput struct {

	//  A list of AssociateResourceResponseElement for each resource that failed
	// association to a percentage custom line item.
	FailedAssociatedResources []types.AssociateResourceResponseElement

	//  A list of AssociateResourceResponseElement for each resource that's been
	// associated to a percentage custom line item successfully.
	SuccessfullyAssociatedResources []types.AssociateResourceResponseElement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchAssociateResourcesToCustomLineItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchAssociateResourcesToCustomLineItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchAssociateResourcesToCustomLineItem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchAssociateResourcesToCustomLineItem"); err != nil {
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
	if err = addOpBatchAssociateResourcesToCustomLineItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateResourcesToCustomLineItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchAssociateResourcesToCustomLineItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchAssociateResourcesToCustomLineItem",
	}
}

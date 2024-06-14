// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcampaigns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connectcampaigns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the outbound call config of a campaign. This API is idempotent.
func (c *Client) UpdateCampaignOutboundCallConfig(ctx context.Context, params *UpdateCampaignOutboundCallConfigInput, optFns ...func(*Options)) (*UpdateCampaignOutboundCallConfigOutput, error) {
	if params == nil {
		params = &UpdateCampaignOutboundCallConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCampaignOutboundCallConfig", params, optFns, c.addOperationUpdateCampaignOutboundCallConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCampaignOutboundCallConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// UpdateCampaignOutboundCallConfigRequest
type UpdateCampaignOutboundCallConfigInput struct {

	// Identifier representing a Campaign
	//
	// This member is required.
	Id *string

	// Answering Machine Detection config
	AnswerMachineDetectionConfig *types.AnswerMachineDetectionConfig

	// The identifier of the contact flow for the outbound call.
	ConnectContactFlowId *string

	// The phone number associated with the Amazon Connect instance, in E.164 format.
	// If you do not specify a source phone number, you must specify a queue.
	ConnectSourcePhoneNumber *string

	noSmithyDocumentSerde
}

type UpdateCampaignOutboundCallConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCampaignOutboundCallConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCampaignOutboundCallConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCampaignOutboundCallConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCampaignOutboundCallConfig"); err != nil {
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
	if err = addOpUpdateCampaignOutboundCallConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCampaignOutboundCallConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCampaignOutboundCallConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCampaignOutboundCallConfig",
	}
}

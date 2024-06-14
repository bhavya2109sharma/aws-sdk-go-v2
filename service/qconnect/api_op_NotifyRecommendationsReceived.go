// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified recommendations from the specified assistant's queue of
// newly available recommendations. You can use this API in conjunction with [GetRecommendations]and a
// waitTimeSeconds input for long-polling behavior and avoiding duplicate
// recommendations.
//
// [GetRecommendations]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_GetRecommendations.html
func (c *Client) NotifyRecommendationsReceived(ctx context.Context, params *NotifyRecommendationsReceivedInput, optFns ...func(*Options)) (*NotifyRecommendationsReceivedOutput, error) {
	if params == nil {
		params = &NotifyRecommendationsReceivedInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NotifyRecommendationsReceived", params, optFns, c.addOperationNotifyRecommendationsReceivedMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NotifyRecommendationsReceivedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyRecommendationsReceivedInput struct {

	// The identifier of the Amazon Q in Connect assistant. Can be either the ID or
	// the ARN. URLs cannot contain the ARN.
	//
	// This member is required.
	AssistantId *string

	// The identifiers of the recommendations.
	//
	// This member is required.
	RecommendationIds []string

	// The identifier of the session. Can be either the ID or the ARN. URLs cannot
	// contain the ARN.
	//
	// This member is required.
	SessionId *string

	noSmithyDocumentSerde
}

type NotifyRecommendationsReceivedOutput struct {

	// The identifiers of recommendations that are causing errors.
	Errors []types.NotifyRecommendationsReceivedError

	// The identifiers of the recommendations.
	RecommendationIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNotifyRecommendationsReceivedMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpNotifyRecommendationsReceived{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpNotifyRecommendationsReceived{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "NotifyRecommendationsReceived"); err != nil {
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
	if err = addOpNotifyRecommendationsReceivedValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyRecommendationsReceived(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNotifyRecommendationsReceived(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NotifyRecommendationsReceived",
	}
}

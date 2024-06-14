// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The type of event associated with the data set.
func (c *Client) SendDataSetNotification(ctx context.Context, params *SendDataSetNotificationInput, optFns ...func(*Options)) (*SendDataSetNotificationOutput, error) {
	if params == nil {
		params = &SendDataSetNotificationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendDataSetNotification", params, optFns, c.addOperationSendDataSetNotificationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendDataSetNotificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendDataSetNotificationInput struct {

	// Affected data set of the notification.
	//
	// This member is required.
	DataSetId *string

	// The type of the notification. Describing the kind of event the notification is
	// alerting you to.
	//
	// This member is required.
	Type types.NotificationType

	// Idempotency key for the notification, this key allows us to deduplicate
	// notifications that are sent in quick succession erroneously.
	ClientToken *string

	// Free-form text field for providers to add information about their notifications.
	Comment *string

	// Extra details specific to this notification type.
	Details *types.NotificationDetails

	// Affected scope of this notification such as the underlying resources affected
	// by the notification event.
	Scope *types.ScopeDetails

	noSmithyDocumentSerde
}

type SendDataSetNotificationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendDataSetNotificationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendDataSetNotification{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendDataSetNotification{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendDataSetNotification"); err != nil {
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
	if err = addIdempotencyToken_opSendDataSetNotificationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendDataSetNotificationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendDataSetNotification(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSendDataSetNotification struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendDataSetNotification) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendDataSetNotification) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendDataSetNotificationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendDataSetNotificationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSendDataSetNotificationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSendDataSetNotification{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendDataSetNotification(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendDataSetNotification",
	}
}

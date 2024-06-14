// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the default message type on a configuration set.
//
// Choose the category of SMS messages that you plan to send from this account. If
// you send account-related messages or time-sensitive messages such as one-time
// passcodes, choose Transactional. If you plan to send messages that contain
// marketing material or other promotional content, choose Promotional. This
// setting applies to your entire Amazon Web Services account.
func (c *Client) SetDefaultMessageType(ctx context.Context, params *SetDefaultMessageTypeInput, optFns ...func(*Options)) (*SetDefaultMessageTypeOutput, error) {
	if params == nil {
		params = &SetDefaultMessageTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetDefaultMessageType", params, optFns, c.addOperationSetDefaultMessageTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetDefaultMessageTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetDefaultMessageTypeInput struct {

	// The configuration set to update with a new default message type. This field can
	// be the ConsigurationSetName or ConfigurationSetArn.
	//
	// This member is required.
	ConfigurationSetName *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	//
	// This member is required.
	MessageType types.MessageType

	noSmithyDocumentSerde
}

type SetDefaultMessageTypeOutput struct {

	// The Amazon Resource Name (ARN) of the updated configuration set.
	ConfigurationSetArn *string

	// The name of the configuration set that was updated.
	ConfigurationSetName *string

	// The new default message type of the configuration set.
	MessageType types.MessageType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetDefaultMessageTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSetDefaultMessageType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSetDefaultMessageType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetDefaultMessageType"); err != nil {
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
	if err = addOpSetDefaultMessageTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetDefaultMessageType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetDefaultMessageType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetDefaultMessageType",
	}
}

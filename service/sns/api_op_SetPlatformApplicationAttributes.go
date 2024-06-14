// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the attributes of the platform application object for the supported push
// notification services, such as APNS and GCM (Firebase Cloud Messaging). For more
// information, see [Using Amazon SNS Mobile Push Notifications]. For information on configuring attributes for message
// delivery status, see [Using Amazon SNS Application Attributes for Message Delivery Status].
//
// [Using Amazon SNS Application Attributes for Message Delivery Status]: https://docs.aws.amazon.com/sns/latest/dg/sns-msg-status.html
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func (c *Client) SetPlatformApplicationAttributes(ctx context.Context, params *SetPlatformApplicationAttributesInput, optFns ...func(*Options)) (*SetPlatformApplicationAttributesOutput, error) {
	if params == nil {
		params = &SetPlatformApplicationAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetPlatformApplicationAttributes", params, optFns, c.addOperationSetPlatformApplicationAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetPlatformApplicationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for SetPlatformApplicationAttributes action.
type SetPlatformApplicationAttributesInput struct {

	// A map of the platform application attributes. Attributes in this map include
	// the following:
	//
	//   - PlatformCredential – The credential received from the notification service.
	//
	//   - For ADM, PlatformCredential is client secret.
	//
	//   - For Apple Services using certificate credentials, PlatformCredential is
	//   private key.
	//
	//   - For Apple Services using token credentials, PlatformCredential is signing
	//   key.
	//
	//   - For GCM (Firebase Cloud Messaging) using key credentials, there is no
	//   PlatformPrincipal . The PlatformCredential is API key .
	//
	//   - For GCM (Firebase Cloud Messaging) using token credentials, there is no
	//   PlatformPrincipal . The PlatformCredential is a JSON formatted private key
	//   file. When using the Amazon Web Services CLI, the file must be in string format
	//   and special characters must be ignored. To format the file correctly, Amazon SNS
	//   recommends using the following command: SERVICE_JSON=`jq @json <<< cat
	//   service.json` .
	//
	//   - PlatformPrincipal – The principal received from the notification service.
	//
	//   - For ADM, PlatformPrincipal is client id.
	//
	//   - For Apple Services using certificate credentials, PlatformPrincipal is SSL
	//   certificate.
	//
	//   - For Apple Services using token credentials, PlatformPrincipal is signing key
	//   ID.
	//
	//   - For GCM (Firebase Cloud Messaging), there is no PlatformPrincipal .
	//
	//   - EventEndpointCreated – Topic ARN to which EndpointCreated event
	//   notifications are sent.
	//
	//   - EventEndpointDeleted – Topic ARN to which EndpointDeleted event
	//   notifications are sent.
	//
	//   - EventEndpointUpdated – Topic ARN to which EndpointUpdate event notifications
	//   are sent.
	//
	//   - EventDeliveryFailure – Topic ARN to which DeliveryFailure event
	//   notifications are sent upon Direct Publish delivery failure (permanent) to one
	//   of the application's endpoints.
	//
	//   - SuccessFeedbackRoleArn – IAM role ARN used to give Amazon SNS write access
	//   to use CloudWatch Logs on your behalf.
	//
	//   - FailureFeedbackRoleArn – IAM role ARN used to give Amazon SNS write access
	//   to use CloudWatch Logs on your behalf.
	//
	//   - SuccessFeedbackSampleRate – Sample rate percentage (0-100) of successfully
	//   delivered messages.
	//
	// The following attributes only apply to APNs token-based authentication:
	//
	//   - ApplePlatformTeamID – The identifier that's assigned to your Apple developer
	//   account team.
	//
	//   - ApplePlatformBundleID – The bundle identifier that's assigned to your iOS
	//   app.
	//
	// This member is required.
	Attributes map[string]string

	// PlatformApplicationArn for SetPlatformApplicationAttributes action.
	//
	// This member is required.
	PlatformApplicationArn *string

	noSmithyDocumentSerde
}

type SetPlatformApplicationAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetPlatformApplicationAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetPlatformApplicationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetPlatformApplicationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetPlatformApplicationAttributes"); err != nil {
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
	if err = addOpSetPlatformApplicationAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetPlatformApplicationAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetPlatformApplicationAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetPlatformApplicationAttributes",
	}
}

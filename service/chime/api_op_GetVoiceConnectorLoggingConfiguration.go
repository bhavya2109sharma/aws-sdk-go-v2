// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the logging configuration details for the specified Amazon Chime
// Voice Connector. Shows whether SIP message logs are enabled for sending to
// Amazon CloudWatch Logs.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [GetVoiceConnectorLoggingConfiguration], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by GetVoiceConnectorLoggingConfiguration in the Amazon
// Chime SDK Voice Namespace
//
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
// [GetVoiceConnectorLoggingConfiguration]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_GetVoiceConnectorLoggingConfiguration.html
func (c *Client) GetVoiceConnectorLoggingConfiguration(ctx context.Context, params *GetVoiceConnectorLoggingConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorLoggingConfigurationOutput, error) {
	if params == nil {
		params = &GetVoiceConnectorLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVoiceConnectorLoggingConfiguration", params, optFns, c.addOperationGetVoiceConnectorLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVoiceConnectorLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVoiceConnectorLoggingConfigurationInput struct {

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

type GetVoiceConnectorLoggingConfigurationOutput struct {

	// The logging configuration details.
	LoggingConfiguration *types.LoggingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVoiceConnectorLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVoiceConnectorLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVoiceConnectorLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetVoiceConnectorLoggingConfiguration"); err != nil {
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
	if err = addOpGetVoiceConnectorLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVoiceConnectorLoggingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetVoiceConnectorLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetVoiceConnectorLoggingConfiguration",
	}
}

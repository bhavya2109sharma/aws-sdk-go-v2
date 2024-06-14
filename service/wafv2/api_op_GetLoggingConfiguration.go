// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the LoggingConfiguration for the specified web ACL.
func (c *Client) GetLoggingConfiguration(ctx context.Context, params *GetLoggingConfigurationInput, optFns ...func(*Options)) (*GetLoggingConfigurationOutput, error) {
	if params == nil {
		params = &GetLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLoggingConfiguration", params, optFns, c.addOperationGetLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoggingConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the web ACL for which you want to get the LoggingConfiguration.
	//
	// This member is required.
	ResourceArn *string

	// The owner of the logging configuration, which must be set to CUSTOMER for the
	// configurations that you manage.
	//
	// The log scope SECURITY_LAKE indicates a configuration that is managed through
	// Amazon Security Lake. You can use Security Lake to collect log and event data
	// from various sources for normalization, analysis, and management. For
	// information, see [Collecting data from Amazon Web Services services]in the Amazon Security Lake user guide.
	//
	// Default: CUSTOMER
	//
	// [Collecting data from Amazon Web Services services]: https://docs.aws.amazon.com/security-lake/latest/userguide/internal-sources.html
	LogScope types.LogScope

	// Used to distinguish between various logging options. Currently, there is one
	// option.
	//
	// Default: WAF_LOGS
	LogType types.LogType

	noSmithyDocumentSerde
}

type GetLoggingConfigurationOutput struct {

	// The LoggingConfiguration for the specified web ACL.
	LoggingConfiguration *types.LoggingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetLoggingConfiguration"); err != nil {
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
	if err = addOpGetLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoggingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetLoggingConfiguration",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associate a protect configuration with a configuration set. This replaces the
// configuration sets current protect configuration. A configuration set can only
// be associated with one protect configuration at a time. A protect configuration
// can be associated with multiple configuration sets.
func (c *Client) AssociateProtectConfiguration(ctx context.Context, params *AssociateProtectConfigurationInput, optFns ...func(*Options)) (*AssociateProtectConfigurationOutput, error) {
	if params == nil {
		params = &AssociateProtectConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateProtectConfiguration", params, optFns, c.addOperationAssociateProtectConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateProtectConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateProtectConfigurationInput struct {

	// The name of the ConfigurationSet.
	//
	// This member is required.
	ConfigurationSetName *string

	// The unique identifier for the protect configuration.
	//
	// This member is required.
	ProtectConfigurationId *string

	noSmithyDocumentSerde
}

type AssociateProtectConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the configuration set.
	//
	// This member is required.
	ConfigurationSetArn *string

	// The name of the ConfigurationSet.
	//
	// This member is required.
	ConfigurationSetName *string

	// The Amazon Resource Name (ARN) of the protect configuration.
	//
	// This member is required.
	ProtectConfigurationArn *string

	// The unique identifier for the protect configuration.
	//
	// This member is required.
	ProtectConfigurationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateProtectConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpAssociateProtectConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpAssociateProtectConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateProtectConfiguration"); err != nil {
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
	if err = addOpAssociateProtectConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateProtectConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateProtectConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateProtectConfiguration",
	}
}

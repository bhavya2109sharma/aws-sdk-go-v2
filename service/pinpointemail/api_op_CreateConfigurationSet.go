// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a configuration set. Configuration sets are groups of rules that you can
// apply to the emails you send using Amazon Pinpoint. You apply a configuration
// set to an email by including a reference to the configuration set in the headers
// of the email. When you apply a configuration set to an email, all of the rules
// in that configuration set are applied to the email.
func (c *Client) CreateConfigurationSet(ctx context.Context, params *CreateConfigurationSetInput, optFns ...func(*Options)) (*CreateConfigurationSetOutput, error) {
	if params == nil {
		params = &CreateConfigurationSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfigurationSet", params, optFns, c.addOperationCreateConfigurationSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create a configuration set.
type CreateConfigurationSetInput struct {

	// The name of the configuration set.
	//
	// This member is required.
	ConfigurationSetName *string

	// An object that defines the dedicated IP pool that is used to send emails that
	// you send using the configuration set.
	DeliveryOptions *types.DeliveryOptions

	// An object that defines whether or not Amazon Pinpoint collects reputation
	// metrics for the emails that you send that use the configuration set.
	ReputationOptions *types.ReputationOptions

	// An object that defines whether or not Amazon Pinpoint can send email that you
	// send using the configuration set.
	SendingOptions *types.SendingOptions

	// An array of objects that define the tags (keys and values) that you want to
	// associate with the configuration set.
	Tags []types.Tag

	// An object that defines the open and click tracking options for emails that you
	// send using the configuration set.
	TrackingOptions *types.TrackingOptions

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type CreateConfigurationSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConfigurationSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfigurationSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfigurationSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateConfigurationSet"); err != nil {
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
	if err = addOpCreateConfigurationSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfigurationSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfigurationSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateConfigurationSet",
	}
}

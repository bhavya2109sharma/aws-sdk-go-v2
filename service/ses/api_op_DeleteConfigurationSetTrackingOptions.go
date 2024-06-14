// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an association between a configuration set and a custom domain for open
// and click event tracking.
//
// By default, images and links used for tracking open and click events are hosted
// on domains operated by Amazon SES. You can configure a subdomain of your own to
// handle these events. For information about using custom domains, see the [Amazon SES Developer Guide].
//
// Deleting this kind of association results in emails sent using the specified
// configuration set to capture open and click events using the standard, Amazon
// SES-operated domains.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/configure-custom-open-click-domains.html
func (c *Client) DeleteConfigurationSetTrackingOptions(ctx context.Context, params *DeleteConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*DeleteConfigurationSetTrackingOptionsOutput, error) {
	if params == nil {
		params = &DeleteConfigurationSetTrackingOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConfigurationSetTrackingOptions", params, optFns, c.addOperationDeleteConfigurationSetTrackingOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConfigurationSetTrackingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete open and click tracking options in a
// configuration set.
type DeleteConfigurationSetTrackingOptionsInput struct {

	// The name of the configuration set.
	//
	// This member is required.
	ConfigurationSetName *string

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type DeleteConfigurationSetTrackingOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteConfigurationSetTrackingOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteConfigurationSetTrackingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteConfigurationSetTrackingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteConfigurationSetTrackingOptions"); err != nil {
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
	if err = addOpDeleteConfigurationSetTrackingOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationSetTrackingOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteConfigurationSetTrackingOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteConfigurationSetTrackingOptions",
	}
}

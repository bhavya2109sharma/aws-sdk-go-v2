// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an SSL/TLS certificate that secures traffic for your website. After the
// certificate is created, it is installed on the specified Lightsail instance.
//
// If you provide more than one domain name in the request, at least one name must
// be less than or equal to 63 characters in length.
func (c *Client) SetupInstanceHttps(ctx context.Context, params *SetupInstanceHttpsInput, optFns ...func(*Options)) (*SetupInstanceHttpsOutput, error) {
	if params == nil {
		params = &SetupInstanceHttpsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetupInstanceHttps", params, optFns, c.addOperationSetupInstanceHttpsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetupInstanceHttpsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetupInstanceHttpsInput struct {

	// The certificate authority that issues the SSL/TLS certificate.
	//
	// This member is required.
	CertificateProvider types.CertificateProvider

	// The name of the domain and subdomains that were specified for the SSL/TLS
	// certificate.
	//
	// This member is required.
	DomainNames []string

	// The contact method for SSL/TLS certificate renewal alerts. You can enter one
	// email address.
	//
	// This member is required.
	EmailAddress *string

	// The name of the Lightsail instance.
	//
	// This member is required.
	InstanceName *string

	noSmithyDocumentSerde
}

type SetupInstanceHttpsOutput struct {

	// The available API operations for SetupInstanceHttps .
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetupInstanceHttpsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetupInstanceHttps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetupInstanceHttps{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetupInstanceHttps"); err != nil {
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
	if err = addOpSetupInstanceHttpsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetupInstanceHttps(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetupInstanceHttps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetupInstanceHttps",
	}
}

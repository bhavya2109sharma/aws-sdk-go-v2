// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Network Firewall TLS inspection configuration. A TLS inspection
// configuration contains Certificate Manager certificate associations between and
// the scope configurations that Network Firewall uses to decrypt and re-encrypt
// traffic traveling through your firewall.
//
// After you create a TLS inspection configuration, you can associate it with a
// new firewall policy.
//
// To update the settings for a TLS inspection configuration, use UpdateTLSInspectionConfiguration.
//
// To manage a TLS inspection configuration's tags, use the standard Amazon Web
// Services resource tagging operations, ListTagsForResource, TagResource, and UntagResource.
//
// To retrieve information about TLS inspection configurations, use ListTLSInspectionConfigurations and DescribeTLSInspectionConfiguration.
//
// For more information about TLS inspection configurations, see [Inspecting SSL/TLS traffic with TLS inspection configurations] in the Network
// Firewall Developer Guide.
//
// [Inspecting SSL/TLS traffic with TLS inspection configurations]: https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection.html
func (c *Client) CreateTLSInspectionConfiguration(ctx context.Context, params *CreateTLSInspectionConfigurationInput, optFns ...func(*Options)) (*CreateTLSInspectionConfigurationOutput, error) {
	if params == nil {
		params = &CreateTLSInspectionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTLSInspectionConfiguration", params, optFns, c.addOperationCreateTLSInspectionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTLSInspectionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTLSInspectionConfigurationInput struct {

	// The object that defines a TLS inspection configuration. This, along with TLSInspectionConfigurationResponse,
	// define the TLS inspection configuration. You can retrieve all objects for a TLS
	// inspection configuration by calling DescribeTLSInspectionConfiguration.
	//
	// Network Firewall uses a TLS inspection configuration to decrypt traffic.
	// Network Firewall re-encrypts the traffic before sending it to its destination.
	//
	// To use a TLS inspection configuration, you add it to a new Network Firewall
	// firewall policy, then you apply the firewall policy to a firewall. Network
	// Firewall acts as a proxy service to decrypt and inspect the traffic traveling
	// through your firewalls. You can reference a TLS inspection configuration from
	// more than one firewall policy, and you can use a firewall policy in more than
	// one firewall. For more information about using TLS inspection configurations,
	// see [Inspecting SSL/TLS traffic with TLS inspection configurations]in the Network Firewall Developer Guide.
	//
	// [Inspecting SSL/TLS traffic with TLS inspection configurations]: https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection.html
	//
	// This member is required.
	TLSInspectionConfiguration *types.TLSInspectionConfiguration

	// The descriptive name of the TLS inspection configuration. You can't change the
	// name of a TLS inspection configuration after you create it.
	//
	// This member is required.
	TLSInspectionConfigurationName *string

	// A description of the TLS inspection configuration.
	Description *string

	// A complex type that contains optional Amazon Web Services Key Management
	// Service (KMS) encryption settings for your Network Firewall resources. Your data
	// is encrypted by default with an Amazon Web Services owned key that Amazon Web
	// Services owns and manages for you. You can use either the Amazon Web Services
	// owned key, or provide your own customer managed key. To learn more about KMS
	// encryption of your Network Firewall resources, see [Encryption at rest with Amazon Web Services Key Managment Service]in the Network Firewall
	// Developer Guide.
	//
	// [Encryption at rest with Amazon Web Services Key Managment Service]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-encryption-at-rest.html
	EncryptionConfiguration *types.EncryptionConfiguration

	// The key:value pairs to associate with the resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateTLSInspectionConfigurationOutput struct {

	// The high-level properties of a TLS inspection configuration. This, along with
	// the TLSInspectionConfiguration, define the TLS inspection configuration. You can retrieve all objects for
	// a TLS inspection configuration by calling DescribeTLSInspectionConfiguration.
	//
	// This member is required.
	TLSInspectionConfigurationResponse *types.TLSInspectionConfigurationResponse

	// A token used for optimistic locking. Network Firewall returns a token to your
	// requests that access the TLS inspection configuration. The token marks the state
	// of the TLS inspection configuration resource at the time of the request.
	//
	// To make changes to the TLS inspection configuration, you provide the token in
	// your request. Network Firewall uses the token to ensure that the TLS inspection
	// configuration hasn't changed since you last retrieved it. If it has changed, the
	// operation fails with an InvalidTokenException . If this happens, retrieve the
	// TLS inspection configuration again to get a current copy of it with a current
	// token. Reapply your changes as needed, then try the operation again using the
	// new token.
	//
	// This member is required.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTLSInspectionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateTLSInspectionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateTLSInspectionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTLSInspectionConfiguration"); err != nil {
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
	if err = addOpCreateTLSInspectionConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTLSInspectionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTLSInspectionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTLSInspectionConfiguration",
	}
}

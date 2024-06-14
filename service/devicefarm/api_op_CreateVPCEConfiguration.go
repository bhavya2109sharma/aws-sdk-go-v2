// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a configuration record in Device Farm for your Amazon Virtual Private
// Cloud (VPC) endpoint.
func (c *Client) CreateVPCEConfiguration(ctx context.Context, params *CreateVPCEConfigurationInput, optFns ...func(*Options)) (*CreateVPCEConfigurationOutput, error) {
	if params == nil {
		params = &CreateVPCEConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVPCEConfiguration", params, optFns, c.addOperationCreateVPCEConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVPCEConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVPCEConfigurationInput struct {

	// The DNS name of the service running in your VPC that you want Device Farm to
	// test.
	//
	// This member is required.
	ServiceDnsName *string

	// The friendly name you give to your VPC endpoint configuration, to manage your
	// configurations more easily.
	//
	// This member is required.
	VpceConfigurationName *string

	// The name of the VPC endpoint service running in your AWS account that you want
	// Device Farm to test.
	//
	// This member is required.
	VpceServiceName *string

	// An optional description that provides details about your VPC endpoint
	// configuration.
	VpceConfigurationDescription *string

	noSmithyDocumentSerde
}

type CreateVPCEConfigurationOutput struct {

	// An object that contains information about your VPC endpoint configuration.
	VpceConfiguration *types.VPCEConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVPCEConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateVPCEConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateVPCEConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVPCEConfiguration"); err != nil {
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
	if err = addOpCreateVPCEConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVPCEConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVPCEConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVPCEConfiguration",
	}
}

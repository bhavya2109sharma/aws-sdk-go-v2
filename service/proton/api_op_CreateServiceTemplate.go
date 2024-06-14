// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a service template. The administrator creates a service template to
// define standardized infrastructure and an optional CI/CD service pipeline.
// Developers, in turn, select the service template from Proton. If the selected
// service template includes a service pipeline definition, they provide a link to
// their source code repository. Proton then deploys and manages the infrastructure
// defined by the selected service template. For more information, see [Proton templates]in the
// Proton User Guide.
//
// [Proton templates]: https://docs.aws.amazon.com/proton/latest/userguide/ag-templates.html
func (c *Client) CreateServiceTemplate(ctx context.Context, params *CreateServiceTemplateInput, optFns ...func(*Options)) (*CreateServiceTemplateOutput, error) {
	if params == nil {
		params = &CreateServiceTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServiceTemplate", params, optFns, c.addOperationCreateServiceTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceTemplateInput struct {

	// The name of the service template.
	//
	// This member is required.
	Name *string

	// A description of the service template.
	Description *string

	// The name of the service template as displayed in the developer interface.
	DisplayName *string

	// A customer provided encryption key that's used to encrypt data.
	EncryptionKey *string

	// By default, Proton provides a service pipeline for your service. When this
	// parameter is included, it indicates that an Proton service pipeline isn't
	// provided for your service. After it's included, it can't be changed. For more
	// information, see [Template bundles]in the Proton User Guide.
	//
	// [Template bundles]: https://docs.aws.amazon.com/proton/latest/userguide/ag-template-authoring.html#ag-template-bundles
	PipelineProvisioning types.Provisioning

	// An optional list of metadata items that you can associate with the Proton
	// service template. A tag is a key-value pair.
	//
	// For more information, see [Proton resources and tagging] in the Proton User Guide.
	//
	// [Proton resources and tagging]: https://docs.aws.amazon.com/proton/latest/userguide/resources.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateServiceTemplateOutput struct {

	// The service template detail data that's returned by Proton.
	//
	// This member is required.
	ServiceTemplate *types.ServiceTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateServiceTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateServiceTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateServiceTemplate"); err != nil {
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
	if err = addOpCreateServiceTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateServiceTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateServiceTemplate",
	}
}

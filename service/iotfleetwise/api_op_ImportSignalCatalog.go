// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Creates a signal catalog using your existing VSS formatted content from your
//
// local device.
func (c *Client) ImportSignalCatalog(ctx context.Context, params *ImportSignalCatalogInput, optFns ...func(*Options)) (*ImportSignalCatalogOutput, error) {
	if params == nil {
		params = &ImportSignalCatalogInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportSignalCatalog", params, optFns, c.addOperationImportSignalCatalogMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportSignalCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportSignalCatalogInput struct {

	// The name of the signal catalog to import.
	//
	// This member is required.
	Name *string

	//  A brief description of the signal catalog.
	Description *string

	// Metadata that can be used to manage the signal catalog.
	Tags []types.Tag

	// The contents of the Vehicle Signal Specification (VSS) configuration. VSS is a
	// precise language used to describe and model signals in vehicle networks.
	Vss types.FormattedVss

	noSmithyDocumentSerde
}

type ImportSignalCatalogOutput struct {

	//  The Amazon Resource Name (ARN) of the imported signal catalog.
	//
	// This member is required.
	Arn *string

	//  The name of the imported signal catalog.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportSignalCatalogMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpImportSignalCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpImportSignalCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportSignalCatalog"); err != nil {
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
	if err = addOpImportSignalCatalogValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportSignalCatalog(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportSignalCatalog(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportSignalCatalog",
	}
}

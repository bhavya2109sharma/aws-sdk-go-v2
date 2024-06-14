// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Updates engine-specific attributes on a specified server. The server enters
//
// the MODIFYING state when this operation is in progress. Only one update can
// occur at a time. You can use this command to reset a Chef server's public key (
// CHEF_PIVOTAL_KEY ) or a Puppet server's admin password ( PUPPET_ADMIN_PASSWORD
// ).
//
// This operation is asynchronous.
//
// This operation can only be called for servers in HEALTHY or UNHEALTHY states.
// Otherwise, an InvalidStateException is raised. A ResourceNotFoundException is
// thrown when the server does not exist. A ValidationException is raised when
// parameters of the request are not valid.
func (c *Client) UpdateServerEngineAttributes(ctx context.Context, params *UpdateServerEngineAttributesInput, optFns ...func(*Options)) (*UpdateServerEngineAttributesOutput, error) {
	if params == nil {
		params = &UpdateServerEngineAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServerEngineAttributes", params, optFns, c.addOperationUpdateServerEngineAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServerEngineAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServerEngineAttributesInput struct {

	// The name of the engine attribute to update.
	//
	// This member is required.
	AttributeName *string

	// The name of the server to update.
	//
	// This member is required.
	ServerName *string

	// The value to set for the attribute.
	AttributeValue *string

	noSmithyDocumentSerde
}

type UpdateServerEngineAttributesOutput struct {

	// Contains the response to an UpdateServerEngineAttributes request.
	Server *types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServerEngineAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServerEngineAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServerEngineAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateServerEngineAttributes"); err != nil {
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
	if err = addOpUpdateServerEngineAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServerEngineAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServerEngineAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateServerEngineAttributes",
	}
}

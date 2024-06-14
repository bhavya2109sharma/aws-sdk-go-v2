// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a configured audience model association within a collaboration.
func (c *Client) GetCollaborationConfiguredAudienceModelAssociation(ctx context.Context, params *GetCollaborationConfiguredAudienceModelAssociationInput, optFns ...func(*Options)) (*GetCollaborationConfiguredAudienceModelAssociationOutput, error) {
	if params == nil {
		params = &GetCollaborationConfiguredAudienceModelAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCollaborationConfiguredAudienceModelAssociation", params, optFns, c.addOperationGetCollaborationConfiguredAudienceModelAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCollaborationConfiguredAudienceModelAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCollaborationConfiguredAudienceModelAssociationInput struct {

	// A unique identifier for the collaboration that the configured audience model
	// association belongs to. Accepts a collaboration ID.
	//
	// This member is required.
	CollaborationIdentifier *string

	// A unique identifier for the configured audience model association that you want
	// to retrieve.
	//
	// This member is required.
	ConfiguredAudienceModelAssociationIdentifier *string

	noSmithyDocumentSerde
}

type GetCollaborationConfiguredAudienceModelAssociationOutput struct {

	// The metadata of the configured audience model association.
	//
	// This member is required.
	CollaborationConfiguredAudienceModelAssociation *types.CollaborationConfiguredAudienceModelAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCollaborationConfiguredAudienceModelAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCollaborationConfiguredAudienceModelAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCollaborationConfiguredAudienceModelAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCollaborationConfiguredAudienceModelAssociation"); err != nil {
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
	if err = addOpGetCollaborationConfiguredAudienceModelAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCollaborationConfiguredAudienceModelAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCollaborationConfiguredAudienceModelAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCollaborationConfiguredAudienceModelAssociation",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Shares a Amazon Web Services Systems Manager document (SSM document)publicly or
// privately. If you share a document privately, you must specify the Amazon Web
// Services user IDs for those people who can use the document. If you share a
// document publicly, you must specify All as the account ID.
func (c *Client) ModifyDocumentPermission(ctx context.Context, params *ModifyDocumentPermissionInput, optFns ...func(*Options)) (*ModifyDocumentPermissionOutput, error) {
	if params == nil {
		params = &ModifyDocumentPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDocumentPermission", params, optFns, c.addOperationModifyDocumentPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDocumentPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDocumentPermissionInput struct {

	// The name of the document that you want to share.
	//
	// This member is required.
	Name *string

	// The permission type for the document. The permission type can be Share.
	//
	// This member is required.
	PermissionType types.DocumentPermissionType

	// The Amazon Web Services users that should have access to the document. The
	// account IDs can either be a group of account IDs or All.
	AccountIdsToAdd []string

	// The Amazon Web Services users that should no longer have access to the
	// document. The Amazon Web Services user can either be a group of account IDs or
	// All. This action has a higher priority than AccountIdsToAdd . If you specify an
	// ID to add and the same ID to remove, the system removes access to the document.
	AccountIdsToRemove []string

	// (Optional) The version of the document to share. If it isn't specified, the
	// system choose the Default version to share.
	SharedDocumentVersion *string

	noSmithyDocumentSerde
}

type ModifyDocumentPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDocumentPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyDocumentPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyDocumentPermission{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyDocumentPermission"); err != nil {
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
	if err = addOpModifyDocumentPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDocumentPermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDocumentPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyDocumentPermission",
	}
}

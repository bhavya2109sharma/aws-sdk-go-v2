// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Shares or unshares a connection alias with one account by specifying whether
// that account has permission to associate the connection alias with a directory.
// If the association permission is granted, the connection alias is shared with
// that account. If the association permission is revoked, the connection alias is
// unshared with the account. For more information, see [Cross-Region Redirection for Amazon WorkSpaces].
//
//   - Before performing this operation, call [DescribeConnectionAliases]to make sure that the current state
//     of the connection alias is CREATED .
//
//   - To delete a connection alias that has been shared, the shared account must
//     first disassociate the connection alias from any directories it has been
//     associated with. Then you must unshare the connection alias from the account it
//     has been shared with. You can delete a connection alias only after it is no
//     longer shared with any accounts or associated with any directories.
//
// [DescribeConnectionAliases]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeConnectionAliases.html
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
func (c *Client) UpdateConnectionAliasPermission(ctx context.Context, params *UpdateConnectionAliasPermissionInput, optFns ...func(*Options)) (*UpdateConnectionAliasPermissionOutput, error) {
	if params == nil {
		params = &UpdateConnectionAliasPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConnectionAliasPermission", params, optFns, c.addOperationUpdateConnectionAliasPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConnectionAliasPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConnectionAliasPermissionInput struct {

	// The identifier of the connection alias that you want to update permissions for.
	//
	// This member is required.
	AliasId *string

	// Indicates whether to share or unshare the connection alias with the specified
	// Amazon Web Services account.
	//
	// This member is required.
	ConnectionAliasPermission *types.ConnectionAliasPermission

	noSmithyDocumentSerde
}

type UpdateConnectionAliasPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConnectionAliasPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateConnectionAliasPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateConnectionAliasPermission{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateConnectionAliasPermission"); err != nil {
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
	if err = addOpUpdateConnectionAliasPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConnectionAliasPermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConnectionAliasPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateConnectionAliasPermission",
	}
}

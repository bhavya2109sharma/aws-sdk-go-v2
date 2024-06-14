// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Suspends up to 50 users from a Team or EnterpriseLWA Amazon Chime account. For
// more information about different account types, see [Managing Your Amazon Chime Accounts]in the Amazon Chime
// Administration Guide.
//
// Users suspended from a Team account are disassociated from the account,but they
// can continue to use Amazon Chime as free users. To remove the suspension from
// suspended Team account users, invite them to the Team account again. You can
// use the InviteUsersaction to do so.
//
// Users suspended from an EnterpriseLWA account are immediately signed out of
// Amazon Chime and can no longer sign in. To remove the suspension from suspended
// EnterpriseLWA account users, use the BatchUnsuspendUser action.
//
// To sign out users without suspending them, use the LogoutUser action.
//
// [Managing Your Amazon Chime Accounts]: https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html
func (c *Client) BatchSuspendUser(ctx context.Context, params *BatchSuspendUserInput, optFns ...func(*Options)) (*BatchSuspendUserOutput, error) {
	if params == nil {
		params = &BatchSuspendUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchSuspendUser", params, optFns, c.addOperationBatchSuspendUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchSuspendUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchSuspendUserInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The request containing the user IDs to suspend.
	//
	// This member is required.
	UserIdList []string

	noSmithyDocumentSerde
}

type BatchSuspendUserOutput struct {

	// If the BatchSuspendUser action fails for one or more of the user IDs in the request, a list of
	// the user IDs is returned, along with error codes and error messages.
	UserErrors []types.UserError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchSuspendUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchSuspendUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchSuspendUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchSuspendUser"); err != nil {
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
	if err = addOpBatchSuspendUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchSuspendUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchSuspendUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchSuspendUser",
	}
}

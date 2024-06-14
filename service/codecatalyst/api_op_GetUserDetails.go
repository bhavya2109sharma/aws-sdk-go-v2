// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a user.
func (c *Client) GetUserDetails(ctx context.Context, params *GetUserDetailsInput, optFns ...func(*Options)) (*GetUserDetailsOutput, error) {
	if params == nil {
		params = &GetUserDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUserDetails", params, optFns, c.addOperationGetUserDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUserDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserDetailsInput struct {

	// The system-generated unique ID of the user.
	Id *string

	// The name of the user as displayed in Amazon CodeCatalyst.
	UserName *string

	noSmithyDocumentSerde
}

type GetUserDetailsOutput struct {

	// The friendly name displayed for the user in Amazon CodeCatalyst.
	DisplayName *string

	// The email address provided by the user when they signed up.
	PrimaryEmail *types.EmailAddress

	// The system-generated unique ID of the user.
	UserId *string

	// The name of the user as displayed in Amazon CodeCatalyst.
	UserName *string

	//
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUserDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUserDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUserDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUserDetails"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUserDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUserDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUserDetails",
	}
}

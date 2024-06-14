// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a password for the specified IAM user. A password allows an IAM user to
// access Amazon Web Services services through the Amazon Web Services Management
// Console.
//
// You can use the CLI, the Amazon Web Services API, or the Users page in the IAM
// console to create a password for any IAM user. Use ChangePasswordto update your own existing
// password in the My Security Credentials page in the Amazon Web Services
// Management Console.
//
// For more information about managing passwords, see [Managing passwords] in the IAM User Guide.
//
// [Managing passwords]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html
func (c *Client) CreateLoginProfile(ctx context.Context, params *CreateLoginProfileInput, optFns ...func(*Options)) (*CreateLoginProfileOutput, error) {
	if params == nil {
		params = &CreateLoginProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoginProfile", params, optFns, c.addOperationCreateLoginProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoginProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLoginProfileInput struct {

	// The new password for the user.
	//
	// The [regex pattern] that is used to validate this parameter is a string of characters. That
	// string can include almost any printable ASCII character from the space ( \u0020
	// ) through the end of the ASCII character range ( \u00FF ). You can also include
	// the tab ( \u0009 ), line feed ( \u000A ), and carriage return ( \u000D )
	// characters. Any of these characters are valid in a password. However, many
	// tools, such as the Amazon Web Services Management Console, might restrict the
	// ability to type certain characters because they have special meaning within that
	// tool.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	Password *string

	// The name of the IAM user to create a password for. The user must already exist.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	UserName *string

	// Specifies whether the user is required to set a new password on next sign-in.
	PasswordResetRequired bool

	noSmithyDocumentSerde
}

// Contains the response to a successful CreateLoginProfile request.
type CreateLoginProfileOutput struct {

	// A structure containing the user name and password create date.
	//
	// This member is required.
	LoginProfile *types.LoginProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLoginProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLoginProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLoginProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLoginProfile"); err != nil {
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
	if err = addOpCreateLoginProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoginProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLoginProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLoginProfile",
	}
}

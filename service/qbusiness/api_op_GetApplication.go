// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about an existing Amazon Q Business application.
func (c *Client) GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) {
	if params == nil {
		params = &GetApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApplication", params, optFns, c.addOperationGetApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApplicationInput struct {

	// The identifier of the Amazon Q Business application.
	//
	// This member is required.
	ApplicationId *string

	noSmithyDocumentSerde
}

type GetApplicationOutput struct {

	// The Amazon Resource Name (ARN) of the Amazon Q Business application.
	ApplicationArn *string

	// The identifier of the Amazon Q Business application.
	ApplicationId *string

	// Settings for whether end users can upload files directly during chat.
	AttachmentsConfiguration *types.AppliedAttachmentsConfiguration

	// Settings for auto-subscription behavior for this application. This is only
	// applicable to SAML and OIDC applications.
	AutoSubscriptionConfiguration *types.AutoSubscriptionConfiguration

	// The OIDC client ID for a Amazon Q Business application.
	ClientIdsForOIDC []string

	// The Unix timestamp when the Amazon Q Business application was last updated.
	CreatedAt *time.Time

	// A description for the Amazon Q Business application.
	Description *string

	// The name of the Amazon Q Business application.
	DisplayName *string

	// The identifier of the Amazon Web Services KMS key that is used to encrypt your
	// data. Amazon Q Business doesn't support asymmetric keys.
	EncryptionConfiguration *types.EncryptionConfiguration

	// If the Status field is set to ERROR , the ErrorMessage field contains a
	// description of the error that caused the synchronization to fail.
	Error *types.ErrorDetail

	// The Amazon Resource Name (ARN) of an identity provider being used by an Amazon
	// Q Business application.
	IamIdentityProviderArn *string

	// The Amazon Resource Name (ARN) of the AWS IAM Identity Center instance attached
	// to your Amazon Q Business application.
	IdentityCenterApplicationArn *string

	// The authentication type being used by a Amazon Q Business application.
	IdentityType types.IdentityType

	// Configuration information about chat response personalization. For more
	// information, see [Personalizing chat responses].
	//
	// [Personalizing chat responses]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/personalizing-chat-responses.html
	PersonalizationConfiguration *types.PersonalizationConfiguration

	// Settings for whether end users can create and use Amazon Q Apps in the web
	// experience.
	QAppsConfiguration *types.QAppsConfiguration

	// The Amazon Resource Name (ARN) of the IAM with permissions to access your
	// CloudWatch logs and metrics.
	RoleArn *string

	// The status of the Amazon Q Business application.
	Status types.ApplicationStatus

	// The Unix timestamp when the Amazon Q Business application was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetApplication"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpGetApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetApplication",
	}
}

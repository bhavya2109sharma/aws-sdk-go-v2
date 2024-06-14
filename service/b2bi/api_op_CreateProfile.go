// Code generated by smithy-go-codegen DO NOT EDIT.

package b2bi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/b2bi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a customer profile. You can have up to five customer profiles, each
// representing a distinct private network. A profile is the mechanism used to
// create the concept of a private network.
func (c *Client) CreateProfile(ctx context.Context, params *CreateProfileInput, optFns ...func(*Options)) (*CreateProfileOutput, error) {
	if params == nil {
		params = &CreateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProfile", params, optFns, c.addOperationCreateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProfileInput struct {

	// Specifies the name for the business associated with this profile.
	//
	// This member is required.
	BusinessName *string

	// Specifies whether or not logging is enabled for this profile.
	//
	// This member is required.
	Logging types.Logging

	// Specifies the name of the profile.
	//
	// This member is required.
	Name *string

	// Specifies the phone number associated with the profile.
	//
	// This member is required.
	Phone *string

	// Reserved for future use.
	ClientToken *string

	// Specifies the email address associated with this customer profile.
	Email *string

	// Specifies the key-value pairs assigned to ARNs that you can use to group and
	// search for resources by type. You can attach this metadata to resources
	// (capabilities, partnerships, and so on) for any purpose.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateProfileOutput struct {

	// Returns the name for the business associated with this profile.
	//
	// This member is required.
	BusinessName *string

	// Returns a timestamp representing the time the profile was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Returns the name of the profile, used to identify it.
	//
	// This member is required.
	Name *string

	// Returns the phone number associated with the profile.
	//
	// This member is required.
	Phone *string

	// Returns an Amazon Resource Name (ARN) for the profile.
	//
	// This member is required.
	ProfileArn *string

	// Returns the unique, system-generated identifier for the profile.
	//
	// This member is required.
	ProfileId *string

	// Returns the email address associated with this customer profile.
	Email *string

	// Returns the name of the logging group.
	LogGroupName *string

	// Returns whether or not logging is turned on for this profile.
	Logging types.Logging

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateProfile"); err != nil {
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
	if err = addIdempotencyToken_opCreateProfileMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProfile(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateProfile struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateProfile) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateProfileInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateProfileMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateProfile{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateProfile",
	}
}

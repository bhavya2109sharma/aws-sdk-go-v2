// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new version of the registration and increase the VersionNumber. The
// previous version of the registration becomes read-only.
func (c *Client) CreateRegistrationVersion(ctx context.Context, params *CreateRegistrationVersionInput, optFns ...func(*Options)) (*CreateRegistrationVersionOutput, error) {
	if params == nil {
		params = &CreateRegistrationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRegistrationVersion", params, optFns, c.addOperationCreateRegistrationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRegistrationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegistrationVersionInput struct {

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	noSmithyDocumentSerde
}

type CreateRegistrationVersionOutput struct {

	// The Amazon Resource Name (ARN) for the registration.
	//
	// This member is required.
	RegistrationArn *string

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	// The status of the registration.
	//
	//   - DRAFT : The initial status of a registration version after it’s created.
	//
	//   - SUBMITTED : Your registration has been submitted.
	//
	//   - REVIEWING : Your registration has been accepted and is being reviewed.
	//
	//   - APPROVED : Your registration has been approved.
	//
	//   - DISCARDED : You've abandon this version of their registration to start over
	//   with a new version.
	//
	//   - DENIED : You must fix your registration and resubmit it.
	//
	//   - REVOKED : Your previously approved registration has been revoked.
	//
	//   - ARCHIVED : Your previously approved registration version moves into this
	//   status when a more recently submitted version is approved.
	//
	// This member is required.
	RegistrationVersionStatus types.RegistrationVersionStatus

	// A RegistrationVersionStatusHistory object that contains timestamps for the
	// registration.
	//
	// This member is required.
	RegistrationVersionStatusHistory *types.RegistrationVersionStatusHistory

	// The new version number of the registration.
	//
	// This member is required.
	VersionNumber *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRegistrationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateRegistrationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateRegistrationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRegistrationVersion"); err != nil {
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
	if err = addOpCreateRegistrationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegistrationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRegistrationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRegistrationVersion",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a policy store. A policy store is a container for policy resources.
//
// Although [Cedar supports multiple namespaces], Verified Permissions currently supports only one namespace per
// policy store.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [Cedar supports multiple namespaces]: https://docs.cedarpolicy.com/schema/schema.html#namespace
func (c *Client) CreatePolicyStore(ctx context.Context, params *CreatePolicyStoreInput, optFns ...func(*Options)) (*CreatePolicyStoreOutput, error) {
	if params == nil {
		params = &CreatePolicyStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePolicyStore", params, optFns, c.addOperationCreatePolicyStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePolicyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePolicyStoreInput struct {

	// Specifies the validation setting for this policy store.
	//
	// Currently, the only valid and required value is Mode .
	//
	// We recommend that you turn on STRICT mode only after you define a schema. If a
	// schema doesn't exist, then STRICT mode causes any policy to fail validation,
	// and Verified Permissions rejects the policy. You can turn off validation by
	// using the [UpdatePolicyStore]. Then, when you have a schema defined, use [UpdatePolicyStore] again to turn validation
	// back on.
	//
	// [UpdatePolicyStore]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyStore
	//
	// This member is required.
	ValidationSettings *types.ValidationSettings

	// Specifies a unique, case-sensitive ID that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a [UUID type of value.].
	//
	// If you don't provide this value, then Amazon Web Services generates a random
	// one for you.
	//
	// If you retry the operation with the same ClientToken , but with different
	// parameters, the retry fails with an ConflictException error.
	//
	// Verified Permissions recognizes a ClientToken for eight hours. After eight
	// hours, the next request with the same parameters performs the operation again
	// regardless of the value of ClientToken .
	//
	// [UUID type of value.]: https://wikipedia.org/wiki/Universally_unique_identifier
	ClientToken *string

	// Descriptive text that you can provide to help with identification of the
	// current policy store.
	Description *string

	noSmithyDocumentSerde
}

type CreatePolicyStoreOutput struct {

	// The Amazon Resource Name (ARN) of the new policy store.
	//
	// This member is required.
	Arn *string

	// The date and time the policy store was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The date and time the policy store was last updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The unique ID of the new policy store.
	//
	// This member is required.
	PolicyStoreId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePolicyStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreatePolicyStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreatePolicyStore{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePolicyStore"); err != nil {
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
	if err = addIdempotencyToken_opCreatePolicyStoreMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePolicyStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicyStore(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePolicyStore struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePolicyStore) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePolicyStore) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePolicyStoreInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePolicyStoreInput ")
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
func addIdempotencyToken_opCreatePolicyStoreMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePolicyStore{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePolicyStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePolicyStore",
	}
}

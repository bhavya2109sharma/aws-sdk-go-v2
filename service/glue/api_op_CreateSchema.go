// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new schema set and registers the schema definition. Returns an error
// if the schema set already exists without actually registering the version.
//
// When the schema set is created, a version checkpoint will be set to the first
// version. Compatibility mode "DISABLED" restricts any additional schema versions
// from being added after the first schema version. For all other compatibility
// modes, validation of compatibility settings will be applied only from the second
// version onwards when the RegisterSchemaVersion API is used.
//
// When this API is called without a RegistryId , this will create an entry for a
// "default-registry" in the registry database tables, if it is not already
// present.
func (c *Client) CreateSchema(ctx context.Context, params *CreateSchemaInput, optFns ...func(*Options)) (*CreateSchemaOutput, error) {
	if params == nil {
		params = &CreateSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSchema", params, optFns, c.addOperationCreateSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSchemaInput struct {

	// The data format of the schema definition. Currently AVRO , JSON and PROTOBUF
	// are supported.
	//
	// This member is required.
	DataFormat types.DataFormat

	// Name of the schema to be created of max length of 255, and may only contain
	// letters, numbers, hyphen, underscore, dollar sign, or hash mark. No whitespace.
	//
	// This member is required.
	SchemaName *string

	// The compatibility mode of the schema. The possible values are:
	//
	//   - NONE: No compatibility mode applies. You can use this choice in development
	//   scenarios or if you do not know the compatibility mode that you want to apply to
	//   schemas. Any new version added will be accepted without undergoing a
	//   compatibility check.
	//
	//   - DISABLED: This compatibility choice prevents versioning for a particular
	//   schema. You can use this choice to prevent future versioning of a schema.
	//
	//   - BACKWARD: This compatibility choice is recommended as it allows data
	//   receivers to read both the current and one previous schema version. This means
	//   that for instance, a new schema version cannot drop data fields or change the
	//   type of these fields, so they can't be read by readers using the previous
	//   version.
	//
	//   - BACKWARD_ALL: This compatibility choice allows data receivers to read both
	//   the current and all previous schema versions. You can use this choice when you
	//   need to delete fields or add optional fields, and check compatibility against
	//   all previous schema versions.
	//
	//   - FORWARD: This compatibility choice allows data receivers to read both the
	//   current and one next schema version, but not necessarily later versions. You can
	//   use this choice when you need to add fields or delete optional fields, but only
	//   check compatibility against the last schema version.
	//
	//   - FORWARD_ALL: This compatibility choice allows data receivers to read
	//   written by producers of any new registered schema. You can use this choice when
	//   you need to add fields or delete optional fields, and check compatibility
	//   against all previous schema versions.
	//
	//   - FULL: This compatibility choice allows data receivers to read data written
	//   by producers using the previous or next version of the schema, but not
	//   necessarily earlier or later versions. You can use this choice when you need to
	//   add or remove optional fields, but only check compatibility against the last
	//   schema version.
	//
	//   - FULL_ALL: This compatibility choice allows data receivers to read data
	//   written by producers using all previous schema versions. You can use this choice
	//   when you need to add or remove optional fields, and check compatibility against
	//   all previous schema versions.
	Compatibility types.Compatibility

	// An optional description of the schema. If description is not provided, there
	// will not be any automatic default value for this.
	Description *string

	//  This is a wrapper shape to contain the registry identity fields. If this is
	// not provided, the default registry will be used. The ARN format for the same
	// will be: arn:aws:glue:us-east-2::registry/default-registry:random-5-letter-id .
	RegistryId *types.RegistryId

	// The schema definition using the DataFormat setting for SchemaName .
	SchemaDefinition *string

	// Amazon Web Services tags that contain a key value pair and may be searched by
	// console, command line, or API. If specified, follows the Amazon Web Services
	// tags-on-create pattern.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSchemaOutput struct {

	// The schema compatibility mode.
	Compatibility types.Compatibility

	// The data format of the schema definition. Currently AVRO , JSON and PROTOBUF
	// are supported.
	DataFormat types.DataFormat

	// A description of the schema if specified when created.
	Description *string

	// The latest version of the schema associated with the returned schema definition.
	LatestSchemaVersion *int64

	// The next version of the schema associated with the returned schema definition.
	NextSchemaVersion *int64

	// The Amazon Resource Name (ARN) of the registry.
	RegistryArn *string

	// The name of the registry.
	RegistryName *string

	// The Amazon Resource Name (ARN) of the schema.
	SchemaArn *string

	// The version number of the checkpoint (the last time the compatibility mode was
	// changed).
	SchemaCheckpoint *int64

	// The name of the schema.
	SchemaName *string

	// The status of the schema.
	SchemaStatus types.SchemaStatus

	// The unique identifier of the first schema version.
	SchemaVersionId *string

	// The status of the first schema version created.
	SchemaVersionStatus types.SchemaVersionStatus

	// The tags for the schema.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSchema{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSchema"); err != nil {
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
	if err = addOpCreateSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSchema",
	}
}

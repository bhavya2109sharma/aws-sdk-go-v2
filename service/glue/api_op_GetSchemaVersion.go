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

// Get the specified schema by its unique ID assigned when a version of the schema
// is created or registered. Schema versions in Deleted status will not be included
// in the results.
func (c *Client) GetSchemaVersion(ctx context.Context, params *GetSchemaVersionInput, optFns ...func(*Options)) (*GetSchemaVersionOutput, error) {
	if params == nil {
		params = &GetSchemaVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSchemaVersion", params, optFns, c.addOperationGetSchemaVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSchemaVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSchemaVersionInput struct {

	// This is a wrapper structure to contain schema identity fields. The structure
	// contains:
	//
	//   - SchemaId$SchemaArn: The Amazon Resource Name (ARN) of the schema. Either
	//   SchemaArn or SchemaName and RegistryName has to be provided.
	//
	//   - SchemaId$SchemaName: The name of the schema. Either SchemaArn or SchemaName
	//   and RegistryName has to be provided.
	SchemaId *types.SchemaId

	// The SchemaVersionId of the schema version. This field is required for fetching
	// by schema ID. Either this or the SchemaId wrapper has to be provided.
	SchemaVersionId *string

	// The version number of the schema.
	SchemaVersionNumber *types.SchemaVersionNumber

	noSmithyDocumentSerde
}

type GetSchemaVersionOutput struct {

	// The date and time the schema version was created.
	CreatedTime *string

	// The data format of the schema definition. Currently AVRO , JSON and PROTOBUF
	// are supported.
	DataFormat types.DataFormat

	// The Amazon Resource Name (ARN) of the schema.
	SchemaArn *string

	// The schema definition for the schema ID.
	SchemaDefinition *string

	// The SchemaVersionId of the schema version.
	SchemaVersionId *string

	// The status of the schema version.
	Status types.SchemaVersionStatus

	// The version number of the schema.
	VersionNumber *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSchemaVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSchemaVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSchemaVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSchemaVersion"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSchemaVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSchemaVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSchemaVersion",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Personalize schema from the specified schema string. The
// schema you create must be in Avro JSON format.
//
// Amazon Personalize recognizes three schema variants. Each schema is associated
// with a dataset type and has a set of required field and keywords. If you are
// creating a schema for a dataset in a Domain dataset group, you provide the
// domain of the Domain dataset group. You specify a schema when you call [CreateDataset].
//
// # Related APIs
//
// [ListSchemas]
//
// [DescribeSchema]
//
// [DeleteSchema]
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
// [DescribeSchema]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSchema.html
// [DeleteSchema]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteSchema.html
// [ListSchemas]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSchemas.html
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

	// The name for the schema.
	//
	// This member is required.
	Name *string

	// A schema in Avro JSON format.
	//
	// This member is required.
	Schema *string

	// The domain for the schema. If you are creating a schema for a dataset in a
	// Domain dataset group, specify the domain you chose when you created the Domain
	// dataset group.
	Domain types.Domain

	noSmithyDocumentSerde
}

type CreateSchemaOutput struct {

	// The Amazon Resource Name (ARN) of the created schema.
	SchemaArn *string

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

// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies the input published schema, at the specified version, into the Directory with the
// same name and version as that of the published schema.
func (c *Client) ApplySchema(ctx context.Context, params *ApplySchemaInput, optFns ...func(*Options)) (*ApplySchemaOutput, error) {
	if params == nil {
		params = &ApplySchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ApplySchema", params, optFns, c.addOperationApplySchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ApplySchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ApplySchemaInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory into which the
	// schema is copied. For more information, see arns.
	//
	// This member is required.
	DirectoryArn *string

	// Published schema Amazon Resource Name (ARN) that needs to be copied. For more
	// information, see arns.
	//
	// This member is required.
	PublishedSchemaArn *string

	noSmithyDocumentSerde
}

type ApplySchemaOutput struct {

	// The applied schema ARN that is associated with the copied schema in the Directory. You
	// can use this ARN to describe the schema information applied on this directory.
	// For more information, see arns.
	AppliedSchemaArn *string

	// The ARN that is associated with the Directory. For more information, see arns.
	DirectoryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationApplySchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpApplySchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpApplySchema{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ApplySchema"); err != nil {
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
	if err = addOpApplySchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opApplySchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opApplySchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ApplySchema",
	}
}

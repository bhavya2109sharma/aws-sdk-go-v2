// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new IngestConfiguration resource, used to specify the ingest protocol
// for a stage.
func (c *Client) CreateIngestConfiguration(ctx context.Context, params *CreateIngestConfigurationInput, optFns ...func(*Options)) (*CreateIngestConfigurationOutput, error) {
	if params == nil {
		params = &CreateIngestConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIngestConfiguration", params, optFns, c.addOperationCreateIngestConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIngestConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIngestConfigurationInput struct {

	// Type of ingest protocol that the user employs to broadcast. If this is set to
	// RTMP , insecureIngest must be set to true .
	//
	// This member is required.
	IngestProtocol types.IngestProtocol

	// Application-provided attributes to store in the IngestConfiguration and attach
	// to a stage. Map keys and values can contain UTF-8 encoded text. The maximum
	// length of this field is 1 KB total. This field is exposed to all stage
	// participants and should not be used for personally identifying, confidential, or
	// sensitive information.
	Attributes map[string]string

	// Whether the stage allows insecure RTMP ingest. This must be set to true , if
	// ingestProtocol is set to RTMP . Default: false .
	InsecureIngest bool

	// Optional name that can be specified for the IngestConfiguration being created.
	Name *string

	// ARN of the stage with which the IngestConfiguration is associated.
	StageArn *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	// Customer-assigned name to help identify the participant using the
	// IngestConfiguration; this can be used to link a participant to a user in the
	// customer’s own systems. This can be any UTF-8 encoded text. This field is
	// exposed to all stage participants and should not be used for personally
	// identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

type CreateIngestConfigurationOutput struct {

	// The IngestConfiguration that was created.
	IngestConfiguration *types.IngestConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIngestConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIngestConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIngestConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateIngestConfiguration"); err != nil {
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
	if err = addOpCreateIngestConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIngestConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIngestConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateIngestConfiguration",
	}
}

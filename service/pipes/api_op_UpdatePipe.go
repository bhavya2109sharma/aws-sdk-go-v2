// Code generated by smithy-go-codegen DO NOT EDIT.

package pipes

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Update an existing pipe. When you call UpdatePipe , EventBridge only the updates
// fields you have specified in the request; the rest remain unchanged. The
// exception to this is if you modify any Amazon Web Services-service specific
// fields in the SourceParameters , EnrichmentParameters , or TargetParameters
// objects. For example, DynamoDBStreamParameters or EventBridgeEventBusParameters
// . EventBridge updates the fields in these objects atomically as one and
// overrides existing values. This is by design, and means that if you don't
// specify an optional field in one of these Parameters objects, EventBridge sets
// that field to its system-default value during the update.
//
// For more information about pipes, see [Amazon EventBridge Pipes] in the Amazon EventBridge User Guide.
//
// [Amazon EventBridge Pipes]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html
func (c *Client) UpdatePipe(ctx context.Context, params *UpdatePipeInput, optFns ...func(*Options)) (*UpdatePipeOutput, error) {
	if params == nil {
		params = &UpdatePipeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePipe", params, optFns, c.addOperationUpdatePipeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePipeInput struct {

	// The name of the pipe.
	//
	// This member is required.
	Name *string

	// The ARN of the role that allows the pipe to send data to the target.
	//
	// This member is required.
	RoleArn *string

	// A description of the pipe.
	Description *string

	// The state the pipe should be in.
	DesiredState types.RequestedPipeState

	// The ARN of the enrichment resource.
	Enrichment *string

	// The parameters required to set up enrichment on your pipe.
	EnrichmentParameters *types.PipeEnrichmentParameters

	// The logging configuration settings for the pipe.
	LogConfiguration *types.PipeLogConfigurationParameters

	// The parameters required to set up a source for your pipe.
	SourceParameters *types.UpdatePipeSourceParameters

	// The ARN of the target resource.
	Target *string

	// The parameters required to set up a target for your pipe.
	//
	// For more information about pipe target parameters, including how to use dynamic
	// path parameters, see [Target parameters]in the Amazon EventBridge User Guide.
	//
	// [Target parameters]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html
	TargetParameters *types.PipeTargetParameters

	noSmithyDocumentSerde
}

type UpdatePipeOutput struct {

	// The ARN of the pipe.
	Arn *string

	// The time the pipe was created.
	CreationTime *time.Time

	// The state the pipe is in.
	CurrentState types.PipeState

	// The state the pipe should be in.
	DesiredState types.RequestedPipeState

	// When the pipe was last updated, in [ISO-8601 format] (YYYY-MM-DDThh:mm:ss.sTZD).
	//
	// [ISO-8601 format]: https://www.w3.org/TR/NOTE-datetime
	LastModifiedTime *time.Time

	// The name of the pipe.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePipeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePipe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePipe{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePipe"); err != nil {
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
	if err = addOpUpdatePipeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePipe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePipe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePipe",
	}
}

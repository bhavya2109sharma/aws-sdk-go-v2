// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests that the status of the specified physical or logical pipeline objects
// be updated in the specified pipeline. This update might not occur immediately,
// but is eventually consistent. The status that can be set depends on the type of
// object (for example, DataNode or Activity). You cannot perform this operation on
// FINISHED pipelines and attempting to do so returns InvalidRequestException .
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.SetStatus Content-Length: 100 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-0634701J7KEXAMPLE", "objectIds": ["o-08600941GHJWMBR9E2"],
// "status": "pause"}
//
// x-amzn-RequestId: e83b8ab7-076a-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 0 Date: Mon, 12 Nov 2012 17:50:53 GMT
//
// Unexpected response: 200, OK, undefined
func (c *Client) SetStatus(ctx context.Context, params *SetStatusInput, optFns ...func(*Options)) (*SetStatusOutput, error) {
	if params == nil {
		params = &SetStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetStatus", params, optFns, c.addOperationSetStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetStatus.
type SetStatusInput struct {

	// The IDs of the objects. The corresponding objects can be either physical or
	// components, but not a mix of both types.
	//
	// This member is required.
	ObjectIds []string

	// The ID of the pipeline that contains the objects.
	//
	// This member is required.
	PipelineId *string

	// The status to be set on all the objects specified in objectIds . For components,
	// use PAUSE or RESUME . For instances, use TRY_CANCEL , RERUN , or MARK_FINISHED .
	//
	// This member is required.
	Status *string

	noSmithyDocumentSerde
}

type SetStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetStatus"); err != nil {
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
	if err = addOpSetStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetStatus",
	}
}

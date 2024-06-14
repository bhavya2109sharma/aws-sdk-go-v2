// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Increases or decreases the stream's data retention period by the value that you
// specify. To indicate whether you want to increase or decrease the data retention
// period, specify the Operation parameter in the request body. In the request,
// you must specify either the StreamName or the StreamARN .
//
// This operation requires permission for the KinesisVideo:UpdateDataRetention
// action.
//
// Changing the data retention period affects the data in the stream as follows:
//
//   - If the data retention period is increased, existing data is retained for
//     the new retention period. For example, if the data retention period is increased
//     from one hour to seven hours, all existing data is retained for seven hours.
//
//   - If the data retention period is decreased, existing data is retained for
//     the new retention period. For example, if the data retention period is decreased
//     from seven hours to one hour, all existing data is retained for one hour, and
//     any data older than one hour is deleted immediately.
func (c *Client) UpdateDataRetention(ctx context.Context, params *UpdateDataRetentionInput, optFns ...func(*Options)) (*UpdateDataRetentionOutput, error) {
	if params == nil {
		params = &UpdateDataRetentionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataRetention", params, optFns, c.addOperationUpdateDataRetentionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataRetentionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataRetentionInput struct {

	// The version of the stream whose retention period you want to change. To get the
	// version, call either the DescribeStream or the ListStreams API.
	//
	// This member is required.
	CurrentVersion *string

	// The number of hours to adjust the current retention by. The value you specify
	// is added to or subtracted from the current value, depending on the operation .
	//
	// The minimum value for data retention is 0 and the maximum value is 87600 (ten
	// years).
	//
	// This member is required.
	DataRetentionChangeInHours *int32

	// Indicates whether you want to increase or decrease the retention period.
	//
	// This member is required.
	Operation types.UpdateDataRetentionOperation

	// The Amazon Resource Name (ARN) of the stream whose retention period you want to
	// change.
	StreamARN *string

	// The name of the stream whose retention period you want to change.
	StreamName *string

	noSmithyDocumentSerde
}

type UpdateDataRetentionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDataRetentionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataRetention{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataRetention{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDataRetention"); err != nil {
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
	if err = addOpUpdateDataRetentionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataRetention(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataRetention(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDataRetention",
	}
}

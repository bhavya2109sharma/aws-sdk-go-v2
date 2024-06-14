// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a signal to the specified resource with a success or failure status. You
// can use the SignalResource operation in conjunction with a creation policy or
// update policy. CloudFormation doesn't proceed with a stack creation or update
// until resources receive the required number of signals or the timeout period is
// exceeded. The SignalResource operation is useful in cases where you want to
// send signals from anywhere other than an Amazon EC2 instance.
func (c *Client) SignalResource(ctx context.Context, params *SignalResourceInput, optFns ...func(*Options)) (*SignalResourceOutput, error) {
	if params == nil {
		params = &SignalResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SignalResource", params, optFns, c.addOperationSignalResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SignalResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the SignalResource action.
type SignalResourceInput struct {

	// The logical ID of the resource that you want to signal. The logical ID is the
	// name of the resource that given in the template.
	//
	// This member is required.
	LogicalResourceId *string

	// The stack name or unique stack ID that includes the resource that you want to
	// signal.
	//
	// This member is required.
	StackName *string

	// The status of the signal, which is either success or failure. A failure signal
	// causes CloudFormation to immediately fail the stack creation or update.
	//
	// This member is required.
	Status types.ResourceSignalStatus

	// A unique ID of the signal. When you signal Amazon EC2 instances or Auto Scaling
	// groups, specify the instance ID that you are signaling as the unique ID. If you
	// send multiple signals to a single resource (such as signaling a wait condition),
	// each signal requires a different unique ID.
	//
	// This member is required.
	UniqueId *string

	noSmithyDocumentSerde
}

type SignalResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSignalResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSignalResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSignalResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SignalResource"); err != nil {
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
	if err = addOpSignalResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSignalResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSignalResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SignalResource",
	}
}

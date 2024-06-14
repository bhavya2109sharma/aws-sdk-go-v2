// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets information about a specified RTMP distribution, including the
// distribution configuration.
func (c *Client) GetStreamingDistribution(ctx context.Context, params *GetStreamingDistributionInput, optFns ...func(*Options)) (*GetStreamingDistributionOutput, error) {
	if params == nil {
		params = &GetStreamingDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStreamingDistribution", params, optFns, c.addOperationGetStreamingDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStreamingDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to get a streaming distribution's information.
type GetStreamingDistributionInput struct {

	// The streaming distribution's ID.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// The returned result of the corresponding request.
type GetStreamingDistributionOutput struct {

	// The current version of the streaming distribution's information. For example:
	// E2QWRUHAPOMQZL .
	ETag *string

	// The streaming distribution's information.
	StreamingDistribution *types.StreamingDistribution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStreamingDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetStreamingDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetStreamingDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetStreamingDistribution"); err != nil {
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
	if err = addOpGetStreamingDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStreamingDistribution(options.Region), middleware.Before); err != nil {
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

// GetStreamingDistributionAPIClient is a client that implements the
// GetStreamingDistribution operation.
type GetStreamingDistributionAPIClient interface {
	GetStreamingDistribution(context.Context, *GetStreamingDistributionInput, ...func(*Options)) (*GetStreamingDistributionOutput, error)
}

var _ GetStreamingDistributionAPIClient = (*Client)(nil)

// StreamingDistributionDeployedWaiterOptions are waiter options for
// StreamingDistributionDeployedWaiter
type StreamingDistributionDeployedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// StreamingDistributionDeployedWaiter will use default minimum delay of 60
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, StreamingDistributionDeployedWaiter will use default max delay of
	// 120 seconds. Note that MaxDelay must resolve to value greater than or equal to
	// the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetStreamingDistributionInput, *GetStreamingDistributionOutput, error) (bool, error)
}

// StreamingDistributionDeployedWaiter defines the waiters for
// StreamingDistributionDeployed
type StreamingDistributionDeployedWaiter struct {
	client GetStreamingDistributionAPIClient

	options StreamingDistributionDeployedWaiterOptions
}

// NewStreamingDistributionDeployedWaiter constructs a
// StreamingDistributionDeployedWaiter.
func NewStreamingDistributionDeployedWaiter(client GetStreamingDistributionAPIClient, optFns ...func(*StreamingDistributionDeployedWaiterOptions)) *StreamingDistributionDeployedWaiter {
	options := StreamingDistributionDeployedWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = streamingDistributionDeployedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &StreamingDistributionDeployedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for StreamingDistributionDeployed waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *StreamingDistributionDeployedWaiter) Wait(ctx context.Context, params *GetStreamingDistributionInput, maxWaitDur time.Duration, optFns ...func(*StreamingDistributionDeployedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for StreamingDistributionDeployed
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *StreamingDistributionDeployedWaiter) WaitForOutput(ctx context.Context, params *GetStreamingDistributionInput, maxWaitDur time.Duration, optFns ...func(*StreamingDistributionDeployedWaiterOptions)) (*GetStreamingDistributionOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetStreamingDistribution(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for StreamingDistributionDeployed waiter")
}

func streamingDistributionDeployedStateRetryable(ctx context.Context, input *GetStreamingDistributionInput, output *GetStreamingDistributionOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("StreamingDistribution.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Deployed"
		value, ok := pathValue.(*string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
		}

		if string(*value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetStreamingDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetStreamingDistribution",
	}
}

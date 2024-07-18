// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// Retrieves a certificate from your private CA or one that has been shared with
// you. The ARN of the certificate is returned when you call the [IssueCertificate]action. You must
// specify both the ARN of your private CA and the ARN of the issued certificate
// when calling the GetCertificate action. You can retrieve the certificate if it
// is in the ISSUED state. You can call the [CreateCertificateAuthorityAuditReport]action to create a report that
// contains information about all of the certificates issued and revoked by your
// private CA.
//
// [IssueCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html
// [CreateCertificateAuthorityAuditReport]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html
func (c *Client) GetCertificate(ctx context.Context, params *GetCertificateInput, optFns ...func(*Options)) (*GetCertificateOutput, error) {
	if params == nil {
		params = &GetCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCertificate", params, optFns, c.addOperationGetCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCertificateInput struct {

	// The ARN of the issued certificate. The ARN contains the certificate serial
	// number and must be in the following form:
	//
	//     arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/286535153982981100925020015808220737245
	//
	// This member is required.
	CertificateArn *string

	// The Amazon Resource Name (ARN) that was returned when you called [CreateCertificateAuthority]. This must be
	// of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// [CreateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html
	//
	// This member is required.
	CertificateAuthorityArn *string

	noSmithyDocumentSerde
}

type GetCertificateOutput struct {

	// The base64 PEM-encoded certificate specified by the CertificateArn parameter.
	Certificate *string

	// The base64 PEM-encoded certificate chain that chains up to the root CA
	// certificate that you used to sign your private CA certificate.
	CertificateChain *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCertificate"); err != nil {
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
	if err = addOpGetCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCertificate(options.Region), middleware.Before); err != nil {
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

// CertificateIssuedWaiterOptions are waiter options for CertificateIssuedWaiter
type CertificateIssuedWaiterOptions struct {

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
	// CertificateIssuedWaiter will use default minimum delay of 1 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, CertificateIssuedWaiter will use default max delay of 60 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *GetCertificateInput, *GetCertificateOutput, error) (bool, error)
}

// CertificateIssuedWaiter defines the waiters for CertificateIssued
type CertificateIssuedWaiter struct {
	client GetCertificateAPIClient

	options CertificateIssuedWaiterOptions
}

// NewCertificateIssuedWaiter constructs a CertificateIssuedWaiter.
func NewCertificateIssuedWaiter(client GetCertificateAPIClient, optFns ...func(*CertificateIssuedWaiterOptions)) *CertificateIssuedWaiter {
	options := CertificateIssuedWaiterOptions{}
	options.MinDelay = 1 * time.Second
	options.MaxDelay = 60 * time.Second
	options.Retryable = certificateIssuedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &CertificateIssuedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for CertificateIssued waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *CertificateIssuedWaiter) Wait(ctx context.Context, params *GetCertificateInput, maxWaitDur time.Duration, optFns ...func(*CertificateIssuedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for CertificateIssued waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *CertificateIssuedWaiter) WaitForOutput(ctx context.Context, params *GetCertificateInput, maxWaitDur time.Duration, optFns ...func(*CertificateIssuedWaiterOptions)) (*GetCertificateOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 60 * time.Second
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

		out, err := w.client.GetCertificate(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
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
	return nil, fmt.Errorf("exceeded max wait time for CertificateIssued waiter")
}

func certificateIssuedStateRetryable(ctx context.Context, input *GetCertificateInput, output *GetCertificateOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var errorType *types.RequestInProgressException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "AccessDeniedException" == apiErr.ErrorCode() {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// GetCertificateAPIClient is a client that implements the GetCertificate
// operation.
type GetCertificateAPIClient interface {
	GetCertificate(context.Context, *GetCertificateInput, ...func(*Options)) (*GetCertificateOutput, error)
}

var _ GetCertificateAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCertificate",
	}
}

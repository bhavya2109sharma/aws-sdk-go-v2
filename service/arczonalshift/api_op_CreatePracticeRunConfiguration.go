// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A practice run configuration for zonal autoshift is required when you enable
// zonal autoshift. A practice run configuration includes specifications for
// blocked dates and blocked time windows, and for Amazon CloudWatch alarms that
// you create to use with practice runs. The alarms that you specify are an outcome
// alarm, to monitor application health during practice runs and, optionally, a
// blocking alarm, to block practice runs from starting.
//
// For more information, see [Considerations when you configure zonal autoshift] in the Amazon Route 53 Application Recovery
// Controller Developer Guide.
//
// [Considerations when you configure zonal autoshift]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.considerations.html
func (c *Client) CreatePracticeRunConfiguration(ctx context.Context, params *CreatePracticeRunConfigurationInput, optFns ...func(*Options)) (*CreatePracticeRunConfigurationOutput, error) {
	if params == nil {
		params = &CreatePracticeRunConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePracticeRunConfiguration", params, optFns, c.addOperationCreatePracticeRunConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePracticeRunConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePracticeRunConfigurationInput struct {

	// The outcome alarm for practice runs is a required Amazon CloudWatch alarm that
	// you specify that ends a practice run when the alarm is in an ALARM state.
	//
	// Configure the alarm to monitor the health of your application when traffic is
	// shifted away from an Availability Zone during each weekly practice run. You
	// should configure the alarm to go into an ALARM state if your application is
	// impacted by the zonal shift, and you want to stop the zonal shift, to let
	// traffic for the resource return to the Availability Zone.
	//
	// This member is required.
	OutcomeAlarms []types.ControlCondition

	// The identifier of the resource to shift away traffic for when a practice run
	// starts a zonal shift. The identifier is the Amazon Resource Name (ARN) for the
	// resource.
	//
	// At this time, supported resources are Network Load Balancers and Application
	// Load Balancers with cross-zone load balancing turned off.
	//
	// This member is required.
	ResourceIdentifier *string

	// Optionally, you can block Route 53 ARC from starting practice runs for a
	// resource on specific calendar dates.
	//
	// The format for blocked dates is: YYYY-MM-DD. Keep in mind, when you specify
	// dates, that dates and times for practice runs are in UTC. Separate multiple
	// blocked dates with spaces.
	//
	// For example, if you have an application update scheduled to launch on May 1,
	// 2024, and you don't want practice runs to shift traffic away at that time, you
	// could set a blocked date for 2024-05-01 .
	BlockedDates []string

	// Optionally, you can block Route 53 ARC from starting practice runs for specific
	// windows of days and times.
	//
	// The format for blocked windows is: DAY:HH:SS-DAY:HH:SS. Keep in mind, when you
	// specify dates, that dates and times for practice runs are in UTC. Also, be aware
	// of potential time adjustments that might be required for daylight saving time
	// differences. Separate multiple blocked windows with spaces.
	//
	// For example, say you run business report summaries three days a week. For this
	// scenario, you might set the following recurring days and times as blocked
	// windows, for example: MON-20:30-21:30 WED-20:30-21:30 FRI-20:30-21:30 .
	BlockedWindows []string

	// An Amazon CloudWatch alarm that you can specify for zonal autoshift practice
	// runs. This alarm blocks Route 53 ARC from starting practice run zonal shifts,
	// and ends a practice run that's in progress, when the alarm is in an ALARM
	// state.
	BlockingAlarms []types.ControlCondition

	noSmithyDocumentSerde
}

type CreatePracticeRunConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the resource that you configured the practice
	// run for.
	//
	// This member is required.
	Arn *string

	// The name of the resource that you configured the practice run for.
	//
	// This member is required.
	Name *string

	// A practice run configuration for a resource. Configurations include the outcome
	// alarm that you specify for practice runs, and, optionally, a blocking alarm and
	// blocking dates and windows.
	//
	// This member is required.
	PracticeRunConfiguration *types.PracticeRunConfiguration

	// The status for zonal autoshift for a resource. When you specify the autoshift
	// status as ENABLED , Amazon Web Services shifts traffic away from shifts away
	// application resource traffic from an Availability Zone, on your behalf, when
	// Amazon Web Services determines that there's an issue in the Availability Zone
	// that could potentially affect customers.
	//
	// When you enable zonal autoshift, you must also configure practice runs for the
	// resource.
	//
	// This member is required.
	ZonalAutoshiftStatus types.ZonalAutoshiftStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePracticeRunConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePracticeRunConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePracticeRunConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePracticeRunConfiguration"); err != nil {
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
	if err = addOpCreatePracticeRunConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePracticeRunConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePracticeRunConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePracticeRunConfiguration",
	}
}

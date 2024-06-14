// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a prefetch schedule for a playback configuration. A prefetch schedule
// allows you to tell MediaTailor to fetch and prepare certain ads before an ad
// break happens. For more information about ad prefetching, see [Using ad prefetching]in the
// MediaTailor User Guide.
//
// [Using ad prefetching]: https://docs.aws.amazon.com/mediatailor/latest/ug/prefetching-ads.html
func (c *Client) CreatePrefetchSchedule(ctx context.Context, params *CreatePrefetchScheduleInput, optFns ...func(*Options)) (*CreatePrefetchScheduleOutput, error) {
	if params == nil {
		params = &CreatePrefetchScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePrefetchSchedule", params, optFns, c.addOperationCreatePrefetchScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePrefetchScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePrefetchScheduleInput struct {

	// The configuration settings for MediaTailor's consumption of the prefetched ads
	// from the ad decision server. Each consumption configuration contains an end time
	// and an optional start time that define the consumption window. Prefetch
	// schedules automatically expire no earlier than seven days after the end time.
	//
	// This member is required.
	Consumption *types.PrefetchConsumption

	// The name to assign to the schedule request.
	//
	// This member is required.
	Name *string

	// The name to assign to the playback configuration.
	//
	// This member is required.
	PlaybackConfigurationName *string

	// The configuration settings for retrieval of prefetched ads from the ad decision
	// server. Only one set of prefetched ads will be retrieved and subsequently
	// consumed for each ad break.
	//
	// This member is required.
	Retrieval *types.PrefetchRetrieval

	// An optional stream identifier that MediaTailor uses to prefetch ads for
	// multiple streams that use the same playback configuration. If StreamId is
	// specified, MediaTailor returns all of the prefetch schedules with an exact match
	// on StreamId . If not specified, MediaTailor returns all of the prefetch
	// schedules for the playback configuration, regardless of StreamId .
	StreamId *string

	noSmithyDocumentSerde
}

type CreatePrefetchScheduleOutput struct {

	// The ARN to assign to the prefetch schedule.
	Arn *string

	// The configuration settings for MediaTailor's consumption of the prefetched ads
	// from the ad decision server. Each consumption configuration contains an end time
	// and an optional start time that define the consumption window. Prefetch
	// schedules automatically expire no earlier than seven days after the end time.
	Consumption *types.PrefetchConsumption

	// The name to assign to the prefetch schedule.
	Name *string

	// The name to assign to the playback configuration.
	PlaybackConfigurationName *string

	// The configuration settings for retrieval of prefetched ads from the ad decision
	// server. Only one set of prefetched ads will be retrieved and subsequently
	// consumed for each ad break.
	Retrieval *types.PrefetchRetrieval

	// An optional stream identifier that MediaTailor uses to prefetch ads for
	// multiple streams that use the same playback configuration. If StreamId is
	// specified, MediaTailor returns all of the prefetch schedules with an exact match
	// on StreamId . If not specified, MediaTailor returns all of the prefetch
	// schedules for the playback configuration, regardless of StreamId .
	StreamId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePrefetchScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePrefetchSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePrefetchSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePrefetchSchedule"); err != nil {
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
	if err = addOpCreatePrefetchScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePrefetchSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePrefetchSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePrefetchSchedule",
	}
}

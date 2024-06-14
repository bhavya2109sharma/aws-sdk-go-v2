// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a Capacity Reservation Fleet. For more information, see [Create a Capacity Reservation Fleet] in the Amazon
// EC2 User Guide.
//
// [Create a Capacity Reservation Fleet]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/work-with-cr-fleets.html#create-crfleet
func (c *Client) CreateCapacityReservationFleet(ctx context.Context, params *CreateCapacityReservationFleetInput, optFns ...func(*Options)) (*CreateCapacityReservationFleetOutput, error) {
	if params == nil {
		params = &CreateCapacityReservationFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCapacityReservationFleet", params, optFns, c.addOperationCreateCapacityReservationFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCapacityReservationFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCapacityReservationFleetInput struct {

	// Information about the instance types for which to reserve the capacity.
	//
	// This member is required.
	InstanceTypeSpecifications []types.ReservationFleetInstanceSpecification

	// The total number of capacity units to be reserved by the Capacity Reservation
	// Fleet. This value, together with the instance type weights that you assign to
	// each instance type used by the Fleet determine the number of instances for which
	// the Fleet reserves capacity. Both values are based on units that make sense for
	// your workload. For more information, see [Total target capacity]in the Amazon EC2 User Guide.
	//
	// [Total target capacity]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/crfleet-concepts.html#target-capacity
	//
	// This member is required.
	TotalTargetCapacity *int32

	// The strategy used by the Capacity Reservation Fleet to determine which of the
	// specified instance types to use. Currently, only the prioritized allocation
	// strategy is supported. For more information, see [Allocation strategy]in the Amazon EC2 User Guide.
	//
	// Valid values: prioritized
	//
	// [Allocation strategy]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/crfleet-concepts.html#allocation-strategy
	AllocationStrategy *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see [Ensure Idempotency].
	//
	// [Ensure Idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The date and time at which the Capacity Reservation Fleet expires. When the
	// Capacity Reservation Fleet expires, its state changes to expired and all of the
	// Capacity Reservations in the Fleet expire.
	//
	// The Capacity Reservation Fleet expires within an hour after the specified time.
	// For example, if you specify 5/31/2019 , 13:30:55 , the Capacity Reservation
	// Fleet is guaranteed to expire between 13:30:55 and 14:30:55 on 5/31/2019 .
	EndDate *time.Time

	// Indicates the type of instance launches that the Capacity Reservation Fleet
	// accepts. All Capacity Reservations in the Fleet inherit this instance matching
	// criteria.
	//
	// Currently, Capacity Reservation Fleets support open instance matching criteria
	// only. This means that instances that have matching attributes (instance type,
	// platform, and Availability Zone) run in the Capacity Reservations automatically.
	// Instances do not need to explicitly target a Capacity Reservation Fleet to use
	// its reserved capacity.
	InstanceMatchCriteria types.FleetInstanceMatchCriteria

	// The tags to assign to the Capacity Reservation Fleet. The tags are
	// automatically assigned to the Capacity Reservations in the Fleet.
	TagSpecifications []types.TagSpecification

	// Indicates the tenancy of the Capacity Reservation Fleet. All Capacity
	// Reservations in the Fleet inherit this tenancy. The Capacity Reservation Fleet
	// can have one of the following tenancy settings:
	//
	//   - default - The Capacity Reservation Fleet is created on hardware that is
	//   shared with other Amazon Web Services accounts.
	//
	//   - dedicated - The Capacity Reservations are created on single-tenant hardware
	//   that is dedicated to a single Amazon Web Services account.
	Tenancy types.FleetCapacityReservationTenancy

	noSmithyDocumentSerde
}

type CreateCapacityReservationFleetOutput struct {

	// The allocation strategy used by the Capacity Reservation Fleet.
	AllocationStrategy *string

	// The ID of the Capacity Reservation Fleet.
	CapacityReservationFleetId *string

	// The date and time at which the Capacity Reservation Fleet was created.
	CreateTime *time.Time

	// The date and time at which the Capacity Reservation Fleet expires.
	EndDate *time.Time

	// Information about the individual Capacity Reservations in the Capacity
	// Reservation Fleet.
	FleetCapacityReservations []types.FleetCapacityReservation

	// The instance matching criteria for the Capacity Reservation Fleet.
	InstanceMatchCriteria types.FleetInstanceMatchCriteria

	// The status of the Capacity Reservation Fleet.
	State types.CapacityReservationFleetState

	// The tags assigned to the Capacity Reservation Fleet.
	Tags []types.Tag

	// Indicates the tenancy of Capacity Reservation Fleet.
	Tenancy types.FleetCapacityReservationTenancy

	// The requested capacity units that have been successfully reserved.
	TotalFulfilledCapacity *float64

	// The total number of capacity units for which the Capacity Reservation Fleet
	// reserves capacity.
	TotalTargetCapacity *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCapacityReservationFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateCapacityReservationFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateCapacityReservationFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCapacityReservationFleet"); err != nil {
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
	if err = addIdempotencyToken_opCreateCapacityReservationFleetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCapacityReservationFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCapacityReservationFleet(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCapacityReservationFleet struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCapacityReservationFleet) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCapacityReservationFleet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCapacityReservationFleetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCapacityReservationFleetInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateCapacityReservationFleetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCapacityReservationFleet{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCapacityReservationFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCapacityReservationFleet",
	}
}

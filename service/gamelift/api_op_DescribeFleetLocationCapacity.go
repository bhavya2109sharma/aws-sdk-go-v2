// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the resource capacity settings for a fleet location. The data
// returned includes the current capacity (number of EC2 instances) and some
// scaling settings for the requested fleet location. For a container fleet, this
// operation also returns counts for replica container groups.
//
// Use this operation to retrieve capacity information for a fleet's remote
// location or home Region (you can also retrieve home Region capacity by calling
// DescribeFleetCapacity ).
//
// To retrieve capacity data, identify a fleet and location.
//
// If successful, a FleetCapacity object is returned for the requested fleet
// location.
//
// # Learn more
//
// [Setting up Amazon GameLift fleets]
//
// [GameLift metrics for fleets]
//
// [Setting up Amazon GameLift fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
// [GameLift metrics for fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html#gamelift-metrics-fleet
func (c *Client) DescribeFleetLocationCapacity(ctx context.Context, params *DescribeFleetLocationCapacityInput, optFns ...func(*Options)) (*DescribeFleetLocationCapacityOutput, error) {
	if params == nil {
		params = &DescribeFleetLocationCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetLocationCapacity", params, optFns, c.addOperationDescribeFleetLocationCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetLocationCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetLocationCapacityInput struct {

	// A unique identifier for the fleet to request location capacity for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The fleet location to retrieve capacity information for. Specify a location in
	// the form of an Amazon Web Services Region code, such as us-west-2 .
	//
	// This member is required.
	Location *string

	noSmithyDocumentSerde
}

type DescribeFleetLocationCapacityOutput struct {

	// Resource capacity information for the requested fleet location. Capacity
	// objects are returned only for fleets and locations that currently exist. Changes
	// in desired instance value can take up to 1 minute to be reflected.
	FleetCapacity *types.FleetCapacity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetLocationCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetLocationCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetLocationCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFleetLocationCapacity"); err != nil {
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
	if err = addOpDescribeFleetLocationCapacityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetLocationCapacity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFleetLocationCapacity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFleetLocationCapacity",
	}
}

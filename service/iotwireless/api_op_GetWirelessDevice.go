// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a wireless device.
func (c *Client) GetWirelessDevice(ctx context.Context, params *GetWirelessDeviceInput, optFns ...func(*Options)) (*GetWirelessDeviceOutput, error) {
	if params == nil {
		params = &GetWirelessDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWirelessDevice", params, optFns, c.addOperationGetWirelessDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWirelessDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWirelessDeviceInput struct {

	// The identifier of the wireless device to get.
	//
	// This member is required.
	Identifier *string

	// The type of identifier used in identifier .
	//
	// This member is required.
	IdentifierType types.WirelessDeviceIdType

	noSmithyDocumentSerde
}

type GetWirelessDeviceOutput struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// The description of the resource.
	Description *string

	// The name of the destination to which the device is assigned.
	DestinationName *string

	// The ID of the wireless device.
	Id *string

	// Information about the wireless device.
	LoRaWAN *types.LoRaWANDevice

	// The name of the resource.
	Name *string

	// FPort values for the GNSS, stream, and ClockSync functions of the positioning
	// information.
	Positioning types.PositioningConfigStatus

	// Sidewalk device object.
	Sidewalk *types.SidewalkDevice

	// The ARN of the thing associated with the wireless device.
	ThingArn *string

	// The name of the thing associated with the wireless device. The value is empty
	// if a thing isn't associated with the device.
	ThingName *string

	// The wireless device type.
	Type types.WirelessDeviceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWirelessDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWirelessDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWirelessDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWirelessDevice"); err != nil {
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
	if err = addOpGetWirelessDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWirelessDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWirelessDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWirelessDevice",
	}
}

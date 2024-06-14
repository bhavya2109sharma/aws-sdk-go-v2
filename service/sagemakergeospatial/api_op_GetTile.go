// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakergeospatial

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Gets a web mercator tile for the given Earth Observation job.
func (c *Client) GetTile(ctx context.Context, params *GetTileInput, optFns ...func(*Options)) (*GetTileOutput, error) {
	if params == nil {
		params = &GetTileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTile", params, optFns, c.addOperationGetTileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTileInput struct {

	// The Amazon Resource Name (ARN) of the tile operation.
	//
	// This member is required.
	Arn *string

	// The particular assets or bands to tile.
	//
	// This member is required.
	ImageAssets []string

	// Determines what part of the Earth Observation job to tile. 'INPUT' or 'OUTPUT'
	// are the valid options.
	//
	// This member is required.
	Target types.TargetOptions

	// The x coordinate of the tile input.
	//
	// This member is required.
	X *int32

	// The y coordinate of the tile input.
	//
	// This member is required.
	Y *int32

	// The z coordinate of the tile input.
	//
	// This member is required.
	Z *int32

	// The Amazon Resource Name (ARN) of the IAM role that you specify.
	ExecutionRoleArn *string

	// Determines whether or not to return a valid data mask.
	ImageMask *bool

	// The output data type of the tile operation.
	OutputDataType types.OutputType

	// The data format of the output tile. The formats include .npy, .png and .jpg.
	OutputFormat *string

	// Property filters for the imagery to tile.
	PropertyFilters *string

	// Time range filter applied to imagery to find the images to tile.
	TimeRangeFilter *string

	noSmithyDocumentSerde
}

type GetTileOutput struct {

	// The output binary file.
	//
	// This value conforms to the media type: application/x-binary
	BinaryFile io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTile"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpGetTileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTile",
	}
}

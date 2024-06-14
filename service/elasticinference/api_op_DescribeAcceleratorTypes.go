// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticinference

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticinference/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Describes the accelerator types available in a given region, as well as their
//
// characteristics, such as memory and throughput.
//
// February 15, 2023: Starting April 15, 2023, AWS will not onboard new customers
// to Amazon Elastic Inference (EI), and will help current customers migrate their
// workloads to options that offer better price and performance. After April 15,
// 2023, new customers will not be able to launch instances with Amazon EI
// accelerators in Amazon SageMaker, Amazon ECS, or Amazon EC2. However, customers
// who have used Amazon EI at least once during the past 30-day period are
// considered current customers and will be able to continue using the service.
func (c *Client) DescribeAcceleratorTypes(ctx context.Context, params *DescribeAcceleratorTypesInput, optFns ...func(*Options)) (*DescribeAcceleratorTypesOutput, error) {
	if params == nil {
		params = &DescribeAcceleratorTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAcceleratorTypes", params, optFns, c.addOperationDescribeAcceleratorTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAcceleratorTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAcceleratorTypesInput struct {
	noSmithyDocumentSerde
}

type DescribeAcceleratorTypesOutput struct {

	//  The available accelerator types.
	AcceleratorTypes []types.AcceleratorType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAcceleratorTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAcceleratorTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAcceleratorTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAcceleratorTypes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAcceleratorTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAcceleratorTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAcceleratorTypes",
	}
}

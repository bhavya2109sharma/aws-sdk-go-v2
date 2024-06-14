// Code generated by smithy-go-codegen DO NOT EDIT.

package launchwizard

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/launchwizard/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns details for a given workload and deployment pattern, including the
// available specifications. You can use the [ListWorkloads]operation to discover the available
// workload names and the [ListWorkloadDeploymentPatterns]operation to discover the available deployment pattern
// names of a given workload.
//
// [ListWorkloadDeploymentPatterns]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_ListWorkloadDeploymentPatterns.html
// [ListWorkloads]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_ListWorkloads.html
func (c *Client) GetWorkloadDeploymentPattern(ctx context.Context, params *GetWorkloadDeploymentPatternInput, optFns ...func(*Options)) (*GetWorkloadDeploymentPatternOutput, error) {
	if params == nil {
		params = &GetWorkloadDeploymentPatternInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkloadDeploymentPattern", params, optFns, c.addOperationGetWorkloadDeploymentPatternMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkloadDeploymentPatternOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkloadDeploymentPatternInput struct {

	// The name of the deployment pattern.
	//
	// This member is required.
	DeploymentPatternName *string

	// The name of the workload.
	//
	// This member is required.
	WorkloadName *string

	noSmithyDocumentSerde
}

type GetWorkloadDeploymentPatternOutput struct {

	// Details about the workload deployment pattern.
	WorkloadDeploymentPattern *types.WorkloadDeploymentPatternData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkloadDeploymentPatternMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkloadDeploymentPattern{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkloadDeploymentPattern{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWorkloadDeploymentPattern"); err != nil {
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
	if err = addOpGetWorkloadDeploymentPatternValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkloadDeploymentPattern(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWorkloadDeploymentPattern(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWorkloadDeploymentPattern",
	}
}

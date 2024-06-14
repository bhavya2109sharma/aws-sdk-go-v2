// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets metadata information about the test set generation.
func (c *Client) DescribeTestSetGeneration(ctx context.Context, params *DescribeTestSetGenerationInput, optFns ...func(*Options)) (*DescribeTestSetGenerationOutput, error) {
	if params == nil {
		params = &DescribeTestSetGenerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTestSetGeneration", params, optFns, c.addOperationDescribeTestSetGenerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTestSetGenerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTestSetGenerationInput struct {

	// The unique identifier of the test set generation.
	//
	// This member is required.
	TestSetGenerationId *string

	noSmithyDocumentSerde
}

type DescribeTestSetGenerationOutput struct {

	// The creation date and time for the test set generation.
	CreationDateTime *time.Time

	// The test set description for the test set generation.
	Description *string

	// The reasons the test set generation failed.
	FailureReasons []string

	// The data source of the test set used for the test set generation.
	GenerationDataSource *types.TestSetGenerationDataSource

	// The date and time of the last update for the test set generation.
	LastUpdatedDateTime *time.Time

	//  The roleARN of the test set used for the test set generation.
	RoleArn *string

	// The Amazon S3 storage location for the test set generation.
	StorageLocation *types.TestSetStorageLocation

	// The unique identifier of the test set generation.
	TestSetGenerationId *string

	// The status for the test set generation.
	TestSetGenerationStatus types.TestSetGenerationStatus

	// The unique identifier for the test set created for the generated test set.
	TestSetId *string

	// The test set name for the generated test set.
	TestSetName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTestSetGenerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTestSetGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTestSetGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTestSetGeneration"); err != nil {
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
	if err = addOpDescribeTestSetGenerationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTestSetGeneration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTestSetGeneration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTestSetGeneration",
	}
}

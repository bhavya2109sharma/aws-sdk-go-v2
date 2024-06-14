// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a predictor created using the CreateAutoPredictor operation.
func (c *Client) DescribeAutoPredictor(ctx context.Context, params *DescribeAutoPredictorInput, optFns ...func(*Options)) (*DescribeAutoPredictorOutput, error) {
	if params == nil {
		params = &DescribeAutoPredictorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutoPredictor", params, optFns, c.addOperationDescribeAutoPredictorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutoPredictorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutoPredictorInput struct {

	// The Amazon Resource Name (ARN) of the predictor.
	//
	// This member is required.
	PredictorArn *string

	noSmithyDocumentSerde
}

type DescribeAutoPredictorOutput struct {

	// The timestamp of the CreateAutoPredictor request.
	CreationTime *time.Time

	// The data configuration for your dataset group and any additional datasets.
	DataConfig *types.DataConfig

	// An array of the ARNs of the dataset import jobs used to import training data
	// for the predictor.
	DatasetImportJobArns []string

	// An Key Management Service (KMS) key and an Identity and Access Management (IAM)
	// role that Amazon Forecast can assume to access the key. You can specify this
	// optional object in the CreateDatasetand CreatePredictor requests.
	EncryptionConfig *types.EncryptionConfig

	// The estimated time remaining in minutes for the predictor training job to
	// complete.
	EstimatedTimeRemainingInMinutes *int64

	// Provides the status and ARN of the Predictor Explainability.
	ExplainabilityInfo *types.ExplainabilityInfo

	// An array of dimension (field) names that specify the attributes used to group
	// your time series.
	ForecastDimensions []string

	// The frequency of predictions in a forecast.
	//
	// Valid intervals are Y (Year), M (Month), W (Week), D (Day), H (Hour), 30min (30
	// minutes), 15min (15 minutes), 10min (10 minutes), 5min (5 minutes), and 1min (1
	// minute). For example, "Y" indicates every year and "5min" indicates every five
	// minutes.
	ForecastFrequency *string

	// The number of time-steps that the model predicts. The forecast horizon is also
	// called the prediction length.
	ForecastHorizon *int32

	// The forecast types used during predictor training. Default value is
	// ["0.1","0.5","0.9"].
	ForecastTypes []string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//
	//   - CREATE_PENDING - The CreationTime .
	//
	//   - CREATE_IN_PROGRESS - The current timestamp.
	//
	//   - CREATE_STOPPING - The current timestamp.
	//
	//   - CREATE_STOPPED - When the job stopped.
	//
	//   - ACTIVE or CREATE_FAILED - When the job finished or failed.
	LastModificationTime *time.Time

	// In the event of an error, a message detailing the cause of the error.
	Message *string

	// A object with the Amazon Resource Name (ARN) and status of the monitor resource.
	MonitorInfo *types.MonitorInfo

	// The accuracy metric used to optimize the predictor.
	OptimizationMetric types.OptimizationMetric

	// The Amazon Resource Name (ARN) of the predictor
	PredictorArn *string

	// The name of the predictor.
	PredictorName *string

	// The ARN and state of the reference predictor. This parameter is only valid for
	// retrained or upgraded predictors.
	ReferencePredictorSummary *types.ReferencePredictorSummary

	// The status of the predictor. States include:
	//
	//   - ACTIVE
	//
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//
	//   - CREATE_STOPPING , CREATE_STOPPED
	//
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	Status *string

	// The time boundary Forecast uses when aggregating data.
	TimeAlignmentBoundary *types.TimeAlignmentBoundary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAutoPredictorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAutoPredictor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAutoPredictor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAutoPredictor"); err != nil {
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
	if err = addOpDescribeAutoPredictorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutoPredictor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAutoPredictor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAutoPredictor",
	}
}

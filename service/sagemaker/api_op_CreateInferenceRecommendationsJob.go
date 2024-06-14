// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a recommendation job. You can create either an instance recommendation
// or load test job.
func (c *Client) CreateInferenceRecommendationsJob(ctx context.Context, params *CreateInferenceRecommendationsJobInput, optFns ...func(*Options)) (*CreateInferenceRecommendationsJobOutput, error) {
	if params == nil {
		params = &CreateInferenceRecommendationsJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInferenceRecommendationsJob", params, optFns, c.addOperationCreateInferenceRecommendationsJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInferenceRecommendationsJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInferenceRecommendationsJobInput struct {

	// Provides information about the versioned model package Amazon Resource Name
	// (ARN), the traffic pattern, and endpoint configurations.
	//
	// This member is required.
	InputConfig *types.RecommendationJobInputConfig

	// A name for the recommendation job. The name must be unique within the Amazon
	// Web Services Region and within your Amazon Web Services account. The job name is
	// passed down to the resources created by the recommendation job. The names of
	// resources (such as the model, endpoint configuration, endpoint, and compilation)
	// that are prefixed with the job name are truncated at 40 characters.
	//
	// This member is required.
	JobName *string

	// Defines the type of recommendation job. Specify Default to initiate an instance
	// recommendation and Advanced to initiate a load test. If left unspecified,
	// Amazon SageMaker Inference Recommender will run an instance recommendation (
	// DEFAULT ) job.
	//
	// This member is required.
	JobType types.RecommendationJobType

	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to
	// perform tasks on your behalf.
	//
	// This member is required.
	RoleArn *string

	// Description of the recommendation job.
	JobDescription *string

	// Provides information about the output artifacts and the KMS key to use for
	// Amazon S3 server-side encryption.
	OutputConfig *types.RecommendationJobOutputConfig

	// A set of conditions for stopping a recommendation job. If any of the conditions
	// are met, the job is automatically stopped.
	StoppingConditions *types.RecommendationJobStoppingConditions

	// The metadata that you apply to Amazon Web Services resources to help you
	// categorize and organize them. Each tag consists of a key and a value, both of
	// which you define. For more information, see [Tagging Amazon Web Services Resources]in the Amazon Web Services General
	// Reference.
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateInferenceRecommendationsJobOutput struct {

	// The Amazon Resource Name (ARN) of the recommendation job.
	//
	// This member is required.
	JobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInferenceRecommendationsJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInferenceRecommendationsJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInferenceRecommendationsJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateInferenceRecommendationsJob"); err != nil {
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
	if err = addOpCreateInferenceRecommendationsJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInferenceRecommendationsJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInferenceRecommendationsJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateInferenceRecommendationsJob",
	}
}

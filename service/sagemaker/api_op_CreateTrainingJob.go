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

// Starts a model training job. After training completes, SageMaker saves the
// resulting model artifacts to an Amazon S3 location that you specify.
//
// If you choose to host your model using SageMaker hosting services, you can use
// the resulting model artifacts as part of the model. You can also use the
// artifacts in a machine learning service other than SageMaker, provided that you
// know how to use them for inference.
//
// In the request body, you provide the following:
//
//   - AlgorithmSpecification - Identifies the training algorithm to use.
//
//   - HyperParameters - Specify these algorithm-specific parameters to enable the
//     estimation of model parameters during training. Hyperparameters can be tuned to
//     optimize this learning process. For a list of hyperparameters for each training
//     algorithm provided by SageMaker, see [Algorithms].
//
// Do not include any security-sensitive information including account access IDs,
//
//	secrets or tokens in any hyperparameter field. If the use of security-sensitive
//	credentials are detected, SageMaker will reject your training job request and
//	return an exception error.
//
//	- InputDataConfig - Describes the input required by the training job and the
//	Amazon S3, EFS, or FSx location where it is stored.
//
//	- OutputDataConfig - Identifies the Amazon S3 bucket where you want SageMaker
//	to save the results of model training.
//
//	- ResourceConfig - Identifies the resources, ML compute instances, and ML
//	storage volumes to deploy for model training. In distributed training, you
//	specify more than one instance.
//
//	- EnableManagedSpotTraining - Optimize the cost of training machine learning
//	models by up to 80% by using Amazon EC2 Spot instances. For more information,
//	see [Managed Spot Training].
//
//	- RoleArn - The Amazon Resource Name (ARN) that SageMaker assumes to perform
//	tasks on your behalf during model training.
//
// You must grant this role the necessary permissions so that SageMaker can
//
//	successfully complete model training.
//
//	- StoppingCondition - To help cap training costs, use MaxRuntimeInSeconds to
//	set a time limit for training. Use MaxWaitTimeInSeconds to specify how long a
//	managed spot training job has to complete.
//
//	- Environment - The environment variables to set in the Docker container.
//
//	- RetryStrategy - The number of times to retry the job when the job fails due
//	to an InternalServerError .
//
// For more information about SageMaker, see [How It Works].
//
// [Algorithms]: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
// [How It Works]: https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html
//
// [Managed Spot Training]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-managed-spot-training.html
func (c *Client) CreateTrainingJob(ctx context.Context, params *CreateTrainingJobInput, optFns ...func(*Options)) (*CreateTrainingJobOutput, error) {
	if params == nil {
		params = &CreateTrainingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrainingJob", params, optFns, c.addOperationCreateTrainingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrainingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrainingJobInput struct {

	// The registry path of the Docker image that contains the training algorithm and
	// algorithm-specific metadata, including the input mode. For more information
	// about algorithms provided by SageMaker, see [Algorithms]. For information about providing
	// your own algorithms, see [Using Your Own Algorithms with Amazon SageMaker].
	//
	// [Algorithms]: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
	// [Using Your Own Algorithms with Amazon SageMaker]: https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html
	//
	// This member is required.
	AlgorithmSpecification *types.AlgorithmSpecification

	// Specifies the path to the S3 location where you want to store model artifacts.
	// SageMaker creates subfolders for the artifacts.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// The resources, including the ML compute instances and ML storage volumes, to
	// use for model training.
	//
	// ML storage volumes store model artifacts and incremental states. Training
	// algorithms might also use ML storage volumes for scratch space. If you want
	// SageMaker to use the ML storage volume to store the training data, choose File
	// as the TrainingInputMode in the algorithm specification. For distributed
	// training algorithms, specify an instance count greater than 1.
	//
	// This member is required.
	ResourceConfig *types.ResourceConfig

	// The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to
	// perform tasks on your behalf.
	//
	// During model training, SageMaker needs your permission to read input data from
	// an S3 bucket, download a Docker image that contains training code, write model
	// artifacts to an S3 bucket, write logs to Amazon CloudWatch Logs, and publish
	// metrics to Amazon CloudWatch. You grant permissions for all of these tasks to an
	// IAM role. For more information, see [SageMaker Roles].
	//
	// To be able to pass this role to SageMaker, the caller of this API must have the
	// iam:PassRole permission.
	//
	// [SageMaker Roles]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html
	//
	// This member is required.
	RoleArn *string

	// Specifies a limit to how long a model training job can run. It also specifies
	// how long a managed Spot training job has to complete. When the job reaches the
	// time limit, SageMaker ends the training job. Use this API to cap model training
	// costs.
	//
	// To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delays
	// job termination for 120 seconds. Algorithms can use this 120-second window to
	// save the model artifacts, so the results of training are not lost.
	//
	// This member is required.
	StoppingCondition *types.StoppingCondition

	// The name of the training job. The name must be unique within an Amazon Web
	// Services Region in an Amazon Web Services account.
	//
	// This member is required.
	TrainingJobName *string

	// Contains information about the output location for managed spot training
	// checkpoint data.
	CheckpointConfig *types.CheckpointConfig

	// Configuration information for the Amazon SageMaker Debugger hook parameters,
	// metric and tensor collections, and storage paths. To learn more about how to
	// configure the DebugHookConfig parameter, see [Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job].
	//
	// [Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job]: https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html
	DebugHookConfig *types.DebugHookConfig

	// Configuration information for Amazon SageMaker Debugger rules for debugging
	// output tensors.
	DebugRuleConfigurations []types.DebugRuleConfiguration

	// To encrypt all communications between ML compute instances in distributed
	// training, choose True . Encryption provides greater security for distributed
	// training, but training might take longer. How long it takes depends on the
	// amount of communication between compute instances, especially if you use a deep
	// learning algorithm in distributed training. For more information, see [Protect Communications Between ML Compute Instances in a Distributed Training Job].
	//
	// [Protect Communications Between ML Compute Instances in a Distributed Training Job]: https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html
	EnableInterContainerTrafficEncryption *bool

	// To train models using managed spot training, choose True . Managed spot training
	// provides a fully managed and scalable infrastructure for training machine
	// learning models. this option is useful when training jobs can be interrupted and
	// when there is flexibility when the training job is run.
	//
	// The complete and intermediate results of jobs are stored in an Amazon S3
	// bucket, and can be used as a starting point to train models incrementally.
	// Amazon SageMaker provides metrics and logs in CloudWatch. They can be used to
	// see when managed spot training jobs are running, interrupted, resumed, or
	// completed.
	EnableManagedSpotTraining *bool

	// Isolates the training container. No inbound or outbound network calls can be
	// made, except for calls between peers within a training cluster for distributed
	// training. If you enable network isolation for training jobs that are configured
	// to use a VPC, SageMaker downloads and uploads customer data and model artifacts
	// through the specified VPC, but the training container does not have network
	// access.
	EnableNetworkIsolation *bool

	// The environment variables to set in the Docker container.
	Environment map[string]string

	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//
	// [CreateProcessingJob]
	//
	// [CreateTrainingJob]
	//
	// [CreateTransformJob]
	//
	// [CreateTransformJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html
	// [CreateTrainingJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html
	// [CreateProcessingJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html
	ExperimentConfig *types.ExperimentConfig

	// Algorithm-specific parameters that influence the quality of the model. You set
	// hyperparameters before you start the learning process. For a list of
	// hyperparameters for each training algorithm provided by SageMaker, see [Algorithms].
	//
	// You can specify a maximum of 100 hyperparameters. Each hyperparameter is a
	// key-value pair. Each key and value is limited to 256 characters, as specified by
	// the Length Constraint .
	//
	// Do not include any security-sensitive information including account access IDs,
	// secrets or tokens in any hyperparameter field. If the use of security-sensitive
	// credentials are detected, SageMaker will reject your training job request and
	// return an exception error.
	//
	// [Algorithms]: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
	HyperParameters map[string]string

	// Contains information about the infrastructure health check configuration for
	// the training job.
	InfraCheckConfig *types.InfraCheckConfig

	// An array of Channel objects. Each channel is a named input source.
	// InputDataConfig describes the input data and its location.
	//
	// Algorithms can accept input data from one or more channels. For example, an
	// algorithm might have two channels of input data, training_data and
	// validation_data . The configuration for each channel provides the S3, EFS, or
	// FSx location where the input data is stored. It also provides information about
	// the stored data: the MIME type, compression method, and whether the data is
	// wrapped in RecordIO format.
	//
	// Depending on the input mode that the algorithm supports, SageMaker either
	// copies input data files from an S3 bucket to a local directory in the Docker
	// container, or makes it available as input streams. For example, if you specify
	// an EFS location, input data files are available as input streams. They do not
	// need to be downloaded.
	//
	// Your input must be in the same Amazon Web Services region as your training job.
	InputDataConfig []types.Channel

	// Configuration information for Amazon SageMaker Debugger system monitoring,
	// framework profiling, and storage paths.
	ProfilerConfig *types.ProfilerConfig

	// Configuration information for Amazon SageMaker Debugger rules for profiling
	// system and framework metrics.
	ProfilerRuleConfigurations []types.ProfilerRuleConfiguration

	// Configuration for remote debugging. To learn more about the remote debugging
	// functionality of SageMaker, see [Access a training container through Amazon Web Services Systems Manager (SSM) for remote debugging].
	//
	// [Access a training container through Amazon Web Services Systems Manager (SSM) for remote debugging]: https://docs.aws.amazon.com/sagemaker/latest/dg/train-remote-debugging.html
	RemoteDebugConfig *types.RemoteDebugConfig

	// The number of times to retry the job when the job fails due to an
	// InternalServerError .
	RetryStrategy *types.RetryStrategy

	// Contains information about attribute-based access control (ABAC) for the
	// training job.
	SessionChainingConfig *types.SessionChainingConfig

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see [Tagging Amazon Web Services Resources].
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	// Configuration of storage locations for the Amazon SageMaker Debugger
	// TensorBoard output data.
	TensorBoardOutputConfig *types.TensorBoardOutputConfig

	// A [VpcConfig] object that specifies the VPC that you want your training job to connect to.
	// Control access to and from your training container by configuring the VPC. For
	// more information, see [Protect Training Jobs by Using an Amazon Virtual Private Cloud].
	//
	// [VpcConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html
	// [Protect Training Jobs by Using an Amazon Virtual Private Cloud]: https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateTrainingJobOutput struct {

	// The Amazon Resource Name (ARN) of the training job.
	//
	// This member is required.
	TrainingJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrainingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrainingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrainingJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTrainingJob"); err != nil {
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
	if err = addOpCreateTrainingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrainingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTrainingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTrainingJob",
	}
}

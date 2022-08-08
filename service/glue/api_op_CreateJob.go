// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new job definition.
func (c *Client) CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) {
	if params == nil {
		params = &CreateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJob", params, optFns, c.addOperationCreateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobInput struct {

	// The JobCommand that runs this job.
	//
	// This member is required.
	Command *types.JobCommand

	// The name you assign to this job definition. It must be unique in your account.
	//
	// This member is required.
	Name *string

	// The name or Amazon Resource Name (ARN) of the IAM role associated with this job.
	//
	// This member is required.
	Role *string

	// This parameter is deprecated. Use MaxCapacity instead. The number of Glue data
	// processing units (DPUs) to allocate to this Job. You can allocate a minimum of 2
	// DPUs; the default is 10. A DPU is a relative measure of processing power that
	// consists of 4 vCPUs of compute capacity and 16 GB of memory. For more
	// information, see the Glue pricing page (https://aws.amazon.com/glue/pricing/).
	//
	// Deprecated: This property is deprecated, use MaxCapacity instead.
	AllocatedCapacity int32

	// The representation of a directed acyclic graph on which both the Glue Studio
	// visual component and Glue Studio code generation is based.
	CodeGenConfigurationNodes map[string]types.CodeGenConfigurationNode

	// The connections used for this job.
	Connections *types.ConnectionsList

	// The default arguments for this job. You can specify arguments here that your own
	// job-execution script consumes, as well as arguments that Glue itself consumes.
	// Job arguments may be logged. Do not pass plaintext secrets as arguments.
	// Retrieve secrets from a Glue Connection, Secrets Manager or other secret
	// management mechanism if you intend to keep them within the Job. For information
	// about how to specify and consume your own Job arguments, see the Calling Glue
	// APIs in Python
	// (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html)
	// topic in the developer guide. For information about the key-value pairs that
	// Glue consumes to set up your job, see the Special Parameters Used by Glue
	// (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html)
	// topic in the developer guide.
	DefaultArguments map[string]string

	// Description of the job being defined.
	Description *string

	// Indicates whether the job is run with a standard or flexible execution class.
	// The standard execution-class is ideal for time-sensitive workloads that require
	// fast job startup and dedicated resources. The flexible execution class is
	// appropriate for time-insensitive jobs whose start and completion times may vary.
	// Only jobs with Glue version 3.0 and above and command type glueetl will be
	// allowed to set ExecutionClass to FLEX. The flexible execution class is available
	// for Spark jobs.
	ExecutionClass types.ExecutionClass

	// An ExecutionProperty specifying the maximum number of concurrent runs allowed
	// for this job.
	ExecutionProperty *types.ExecutionProperty

	// Glue version determines the versions of Apache Spark and Python that Glue
	// supports. The Python version indicates the version supported for jobs of type
	// Spark. For more information about the available Glue versions and corresponding
	// Spark and Python versions, see Glue version
	// (https://docs.aws.amazon.com/glue/latest/dg/add-job.html) in the developer
	// guide. Jobs that are created without specifying a Glue version default to Glue
	// 0.9.
	GlueVersion *string

	// This field is reserved for future use.
	LogUri *string

	// For Glue version 1.0 or earlier jobs, using the standard worker type, the number
	// of Glue data processing units (DPUs) that can be allocated when this job runs. A
	// DPU is a relative measure of processing power that consists of 4 vCPUs of
	// compute capacity and 16 GB of memory. For more information, see the Glue pricing
	// page (https://aws.amazon.com/glue/pricing/). Do not set Max Capacity if using
	// WorkerType and NumberOfWorkers. The value that can be allocated for MaxCapacity
	// depends on whether you are running a Python shell job or an Apache Spark ETL
	// job:
	//
	// * When you specify a Python shell job (JobCommand.Name="pythonshell"), you
	// can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	//
	// * When you
	// specify an Apache Spark ETL job (JobCommand.Name="glueetl") or Apache Spark
	// streaming ETL job (JobCommand.Name="gluestreaming"), you can allocate a minimum
	// of 2 DPUs. The default is 10 DPUs. This job type cannot have a fractional DPU
	// allocation.
	//
	// For Glue version 2.0 jobs, you cannot instead specify a Maximum
	// capacity. Instead, you should specify a Worker type and the Number of workers.
	MaxCapacity *float64

	// The maximum number of times to retry this job if it fails.
	MaxRetries int32

	// Non-overridable arguments for this job, specified as name-value pairs.
	NonOverridableArguments map[string]string

	// Specifies configuration properties of a job notification.
	NotificationProperty *types.NotificationProperty

	// The number of workers of a defined workerType that are allocated when a job
	// runs.
	NumberOfWorkers *int32

	// The name of the SecurityConfiguration structure to be used with this job.
	SecurityConfiguration *string

	// The tags to use with this job. You may use tags to limit access to the job. For
	// more information about tags in Glue, see Amazon Web Services Tags in Glue
	// (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html) in the developer
	// guide.
	Tags map[string]string

	// The job timeout in minutes. This is the maximum time that a job run can consume
	// resources before it is terminated and enters TIMEOUT status. The default is
	// 2,880 minutes (48 hours).
	Timeout *int32

	// The type of predefined worker that is allocated when a job runs. Accepts a value
	// of Standard, G.1X, G.2X, or G.025X.
	//
	// * For the Standard worker type, each worker
	// provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	//
	// *
	// For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64
	// GB disk), and provides 1 executor per worker. We recommend this worker type for
	// memory-intensive jobs.
	//
	// * For the G.2X worker type, each worker maps to 2 DPU (8
	// vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. We
	// recommend this worker type for memory-intensive jobs.
	//
	// * For the G.025X worker
	// type, each worker maps to 0.25 DPU (2 vCPU, 4 GB of memory, 64 GB disk), and
	// provides 1 executor per worker. We recommend this worker type for low volume
	// streaming jobs. This worker type is only available for Glue version 3.0
	// streaming jobs.
	WorkerType types.WorkerType

	noSmithyDocumentSerde
}

type CreateJobOutput struct {

	// The unique name that was provided for this job definition.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJob(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opCreateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateJob",
	}
}

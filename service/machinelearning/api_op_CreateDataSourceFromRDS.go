// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a DataSource object from an [Amazon Relational Database Service] (Amazon RDS). A DataSource references data
// that can be used to perform CreateMLModel , CreateEvaluation , or
// CreateBatchPrediction operations.
//
// CreateDataSourceFromRDS is an asynchronous operation. In response to
// CreateDataSourceFromRDS , Amazon Machine Learning (Amazon ML) immediately
// returns and sets the DataSource status to PENDING . After the DataSource is
// created and ready for use, Amazon ML sets the Status parameter to COMPLETED .
// DataSource in the COMPLETED or PENDING state can be used only to perform
// >CreateMLModel >, CreateEvaluation , or CreateBatchPrediction operations.
//
// If Amazon ML cannot accept the input source, it sets the Status parameter to
// FAILED and includes an error message in the Message attribute of the
// GetDataSource operation response.
//
// [Amazon Relational Database Service]: http://aws.amazon.com/rds/
func (c *Client) CreateDataSourceFromRDS(ctx context.Context, params *CreateDataSourceFromRDSInput, optFns ...func(*Options)) (*CreateDataSourceFromRDSOutput, error) {
	if params == nil {
		params = &CreateDataSourceFromRDSInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataSourceFromRDS", params, optFns, c.addOperationCreateDataSourceFromRDSMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataSourceFromRDSOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceFromRDSInput struct {

	// A user-supplied ID that uniquely identifies the DataSource . Typically, an
	// Amazon Resource Number (ARN) becomes the ID for a DataSource .
	//
	// This member is required.
	DataSourceId *string

	// The data specification of an Amazon RDS DataSource :
	//
	//   - DatabaseInformation -
	//
	//   - DatabaseName - The name of the Amazon RDS database.
	//
	//   - InstanceIdentifier - A unique identifier for the Amazon RDS database
	//   instance.
	//
	//   - DatabaseCredentials - AWS Identity and Access Management (IAM) credentials
	//   that are used to connect to the Amazon RDS database.
	//
	//   - ResourceRole - A role (DataPipelineDefaultResourceRole) assumed by an EC2
	//   instance to carry out the copy task from Amazon RDS to Amazon Simple Storage
	//   Service (Amazon S3). For more information, see [Role templates]for data pipelines.
	//
	//   - ServiceRole - A role (DataPipelineDefaultRole) assumed by the AWS Data
	//   Pipeline service to monitor the progress of the copy task from Amazon RDS to
	//   Amazon S3. For more information, see [Role templates]for data pipelines.
	//
	//   - SecurityInfo - The security information to use to access an RDS DB
	//   instance. You need to set up appropriate ingress rules for the security entity
	//   IDs provided to allow access to the Amazon RDS instance. Specify a [ SubnetId
	//   , SecurityGroupIds ] pair for a VPC-based RDS DB instance.
	//
	//   - SelectSqlQuery - A query that is used to retrieve the observation data for
	//   the Datasource .
	//
	//   - S3StagingLocation - The Amazon S3 location for staging Amazon RDS data. The
	//   data retrieved from Amazon RDS using SelectSqlQuery is stored in this location.
	//
	//   - DataSchemaUri - The Amazon S3 location of the DataSchema .
	//
	//   - DataSchema - A JSON string representing the schema. This is not required if
	//   DataSchemaUri is specified.
	//
	//   - DataRearrangement - A JSON string that represents the splitting and
	//   rearrangement requirements for the Datasource .
	//
	// Sample - "{\"splitting\":{\"percentBegin\":10,\"percentEnd\":60}}"
	//
	// [Role templates]: https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-iam-roles.html
	//
	// This member is required.
	RDSData *types.RDSDataSpec

	// The role that Amazon ML assumes on behalf of the user to create and activate a
	// data pipeline in the user's account and copy data using the SelectSqlQuery
	// query from Amazon RDS to Amazon S3.
	//
	// This member is required.
	RoleARN *string

	// The compute statistics for a DataSource . The statistics are generated from the
	// observation data referenced by a DataSource . Amazon ML uses the statistics
	// internally during MLModel training. This parameter must be set to true if the
	// DataSource needs to be used for MLModel training.
	ComputeStatistics bool

	// A user-supplied name or description of the DataSource .
	DataSourceName *string

	noSmithyDocumentSerde
}

//	Represents the output of a CreateDataSourceFromRDS operation, and is an
//
// acknowledgement that Amazon ML received the request.
//
// The CreateDataSourceFromRDS > operation is asynchronous. You can poll for
// updates by using the GetBatchPrediction operation and checking the Status
// parameter. You can inspect the Message when Status shows up as FAILED . You can
// also check the progress of the copy operation by going to the DataPipeline
// console and looking up the pipeline using the pipelineId  from the describe
// call.
type CreateDataSourceFromRDSOutput struct {

	// A user-supplied ID that uniquely identifies the datasource. This value should
	// be identical to the value of the DataSourceID in the request.
	DataSourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataSourceFromRDSMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataSourceFromRDS{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataSourceFromRDS{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataSourceFromRDS"); err != nil {
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
	if err = addOpCreateDataSourceFromRDSValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSourceFromRDS(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataSourceFromRDS(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataSourceFromRDS",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Runs a batch SQL statement over an array of data.
//
// You can run bulk update and insert operations for multiple records using a DML
// statement with different parameter sets. Bulk operations can provide a
// significant performance improvement over individual insert and update
// operations.
//
// If a call isn't part of a transaction because it doesn't include the
// transactionID parameter, changes that result from the call are committed
// automatically.
//
// There isn't a fixed upper limit on the number of parameter sets. However, the
// maximum size of the HTTP request submitted through the Data API is 4 MiB. If the
// request exceeds this limit, the Data API returns an error and doesn't process
// the request. This 4-MiB limit includes the size of the HTTP headers and the JSON
// notation in the request. Thus, the number of parameter sets that you can include
// depends on a combination of factors, such as the size of the SQL statement and
// the size of each parameter set.
//
// The response size limit is 1 MiB. If the call returns more than 1 MiB of
// response data, the call is terminated.
func (c *Client) BatchExecuteStatement(ctx context.Context, params *BatchExecuteStatementInput, optFns ...func(*Options)) (*BatchExecuteStatementOutput, error) {
	if params == nil {
		params = &BatchExecuteStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchExecuteStatement", params, optFns, c.addOperationBatchExecuteStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchExecuteStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters represent the input of a SQL statement over an array of
// data.
type BatchExecuteStatementInput struct {

	// The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.
	//
	// This member is required.
	ResourceArn *string

	// The ARN of the secret that enables access to the DB cluster. Enter the database
	// user name and password for the credentials in the secret.
	//
	// For information about creating the secret, see [Create a database secret].
	//
	// [Create a database secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/create_database_secret.html
	//
	// This member is required.
	SecretArn *string

	// The SQL statement to run. Don't include a semicolon (;) at the end of the SQL
	// statement.
	//
	// This member is required.
	Sql *string

	// The name of the database.
	Database *string

	// The parameter set for the batch operation.
	//
	// The SQL statement is executed as many times as the number of parameter sets
	// provided. To execute a SQL statement with no parameters, use one of the
	// following options:
	//
	//   - Specify one or more empty parameter sets.
	//
	//   - Use the ExecuteStatement operation instead of the BatchExecuteStatement
	//   operation.
	//
	// Array parameters are not supported.
	ParameterSets [][]types.SqlParameter

	// The name of the database schema.
	//
	// Currently, the schema parameter isn't supported.
	Schema *string

	// The identifier of a transaction that was started by using the BeginTransaction
	// operation. Specify the transaction ID of the transaction that you want to
	// include the SQL statement in.
	//
	// If the SQL statement is not part of a transaction, don't set this parameter.
	TransactionId *string

	noSmithyDocumentSerde
}

// The response elements represent the output of a SQL statement over an array of
// data.
type BatchExecuteStatementOutput struct {

	// The execution results of each batch entry.
	UpdateResults []types.UpdateResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchExecuteStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchExecuteStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchExecuteStatement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchExecuteStatement"); err != nil {
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
	if err = addOpBatchExecuteStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchExecuteStatement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchExecuteStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchExecuteStatement",
	}
}

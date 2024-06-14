// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a previously provisioned DB instance. When you delete a DB instance,
// all automated backups for that instance are deleted and can't be recovered.
// However, manual DB snapshots of the DB instance aren't deleted.
//
// If you request a final DB snapshot, the status of the Amazon RDS DB instance is
// deleting until the DB snapshot is created. This operation can't be canceled or
// reverted after it begins. To monitor the status of this operation, use
// DescribeDBInstance .
//
// When a DB instance is in a failure state and has a status of failed ,
// incompatible-restore , or incompatible-network , you can only delete it when you
// skip creation of the final snapshot with the SkipFinalSnapshot parameter.
//
// If the specified DB instance is part of an Amazon Aurora DB cluster, you can't
// delete the DB instance if both of the following conditions are true:
//
//   - The DB cluster is a read replica of another Amazon Aurora DB cluster.
//
//   - The DB instance is the only instance in the DB cluster.
//
// To delete a DB instance in this case, first use the PromoteReadReplicaDBCluster
// operation to promote the DB cluster so that it's no longer a read replica. After
// the promotion completes, use the DeleteDBInstance operation to delete the final
// instance in the DB cluster.
//
// For RDS Custom DB instances, deleting the DB instance permanently deletes the
// EC2 instance and the associated EBS volumes. Make sure that you don't terminate
// or delete these resources before you delete the DB instance. Otherwise, deleting
// the DB instance and creation of the final snapshot might fail.
func (c *Client) DeleteDBInstance(ctx context.Context, params *DeleteDBInstanceInput, optFns ...func(*Options)) (*DeleteDBInstanceOutput, error) {
	if params == nil {
		params = &DeleteDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBInstance", params, optFns, c.addOperationDeleteDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDBInstanceInput struct {

	// The DB instance identifier for the DB instance to be deleted. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//   - Must match the name of an existing DB instance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// Specifies whether to remove automated backups immediately after the DB instance
	// is deleted. This parameter isn't case-sensitive. The default is to remove
	// automated backups immediately after the DB instance is deleted.
	DeleteAutomatedBackups *bool

	// The DBSnapshotIdentifier of the new DBSnapshot created when the
	// SkipFinalSnapshot parameter is disabled.
	//
	// If you enable this parameter and also enable SkipFinalShapshot, the command
	// results in an error.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Constraints:
	//
	//   - Must be 1 to 255 letters or numbers.
	//
	//   - First character must be a letter.
	//
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	//
	//   - Can't be specified when deleting a read replica.
	FinalDBSnapshotIdentifier *string

	// Specifies whether to skip the creation of a final DB snapshot before deleting
	// the instance. If you enable this parameter, RDS doesn't create a DB snapshot. If
	// you don't enable this parameter, RDS creates a DB snapshot before the DB
	// instance is deleted. By default, skip isn't enabled, and the DB snapshot is
	// created.
	//
	// If you don't enable this parameter, you must specify the
	// FinalDBSnapshotIdentifier parameter.
	//
	// When a DB instance is in a failure state and has a status of failed ,
	// incompatible-restore , or incompatible-network , RDS can delete the instance
	// only if you enable this parameter.
	//
	// If you delete a read replica or an RDS Custom instance, you must enable this
	// setting.
	//
	// This setting is required for RDS Custom.
	SkipFinalSnapshot *bool

	noSmithyDocumentSerde
}

type DeleteDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the operations CreateDBInstance
	// , CreateDBInstanceReadReplica , DeleteDBInstance , DescribeDBInstances ,
	// ModifyDBInstance , PromoteReadReplica , RebootDBInstance ,
	// RestoreDBInstanceFromDBSnapshot , RestoreDBInstanceFromS3 ,
	// RestoreDBInstanceToPointInTime , StartDBInstance , and StopDBInstance .
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteDBInstance"); err != nil {
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
	if err = addOpDeleteDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteDBInstance",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the settings of an Amazon Aurora DB cluster or a Multi-AZ DB cluster.
// You can change one or more settings by specifying these parameters and the new
// values in the request. For more information on Amazon Aurora DB clusters, see
// What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see Multi-AZ DB cluster deployments (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) ModifyDBCluster(ctx context.Context, params *ModifyDBClusterInput, optFns ...func(*Options)) (*ModifyDBClusterOutput, error) {
	if params == nil {
		params = &ModifyDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBCluster", params, optFns, c.addOperationModifyDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBClusterInput struct {

	// The DB cluster identifier for the cluster being modified. This parameter isn't
	// case-sensitive. Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB
	// clusters Constraints:
	//   - Must match the identifier of an existing DB cluster.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The amount of storage in gibibytes (GiB) to allocate to each DB instance in the
	// Multi-AZ DB cluster. Valid for Cluster Type: Multi-AZ DB clusters only
	AllocatedStorage *int32

	// Specifies whether engine mode changes from serverless to provisioned are
	// allowed. Valid for Cluster Type: Aurora Serverless v1 DB clusters only
	// Constraints:
	//   - You must allow engine mode changes when specifying a different value for
	//   the EngineMode parameter from the DB cluster's current engine mode.
	AllowEngineModeChange *bool

	// Specifies whether major version upgrades are allowed. Valid for Cluster Type:
	// Aurora DB clusters and Multi-AZ DB clusters Constraints:
	//   - You must allow major version upgrades when specifying a value for the
	//   EngineVersion parameter that is a different major version than the DB
	//   cluster's current version.
	AllowMajorVersionUpgrade *bool

	// Specifies whether the modifications in this request and any pending
	// modifications are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the DB cluster. If this parameter is
	// disabled, changes to the DB cluster are applied during the next maintenance
	// window. Most modifications can be applied immediately or during the next
	// scheduled maintenance window. Some modifications, such as turning on deletion
	// protection and changing the master password, are applied immediately—regardless
	// of when you choose to apply them. By default, this parameter is disabled. Valid
	// for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
	ApplyImmediately *bool

	// Specifies whether minor engine upgrades are applied automatically to the DB
	// cluster during the maintenance window. By default, minor engine upgrades are
	// applied automatically. Valid for Cluster Type: Multi-AZ DB clusters only
	AutoMinorVersionUpgrade *bool

	// The Amazon Resource Name (ARN) of the recovery point in Amazon Web Services
	// Backup.
	AwsBackupRecoveryPointArn *string

	// The target backtrack window, in seconds. To disable backtracking, set this
	// value to 0 . Valid for Cluster Type: Aurora MySQL DB clusters only Default: 0
	// Constraints:
	//   - If specified, this value must be set to a number from 0 to 259,200 (72
	//   hours).
	BacktrackWindow *int64

	// The number of days for which automated backups are retained. Specify a minimum
	// value of 1 . Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
	// Default: 1 Constraints:
	//   - Must be a value from 1 to 35.
	BackupRetentionPeriod *int32

	// The configuration setting for the log types to be enabled for export to
	// CloudWatch Logs for a specific DB cluster. Valid for Cluster Type: Aurora DB
	// clusters and Multi-AZ DB clusters The following values are valid for each DB
	// engine:
	//   - Aurora MySQL - audit | error | general | slowquery
	//   - Aurora PostgreSQL - postgresql
	//   - RDS for MySQL - error | general | slowquery
	//   - RDS for PostgreSQL - postgresql | upgrade
	// For more information about exporting CloudWatch Logs for Amazon RDS, see
	// Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide. For more information about exporting CloudWatch
	// Logs for Amazon Aurora, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide.
	CloudwatchLogsExportConfiguration *types.CloudwatchLogsExportConfiguration

	// Specifies whether to copy all tags from the DB cluster to snapshots of the DB
	// cluster. The default is not to copy them. Valid for Cluster Type: Aurora DB
	// clusters and Multi-AZ DB clusters
	CopyTagsToSnapshot *bool

	// The compute and memory capacity of each DB instance in the Multi-AZ DB cluster,
	// for example db.m6gd.xlarge . Not all DB instance classes are available in all
	// Amazon Web Services Regions, or for all database engines. For the full list of
	// DB instance classes and availability for your engine, see DB Instance Class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. Valid for Cluster Type: Multi-AZ DB clusters only
	DBClusterInstanceClass *string

	// The name of the DB cluster parameter group to use for the DB cluster. Valid for
	// Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
	DBClusterParameterGroupName *string

	// The name of the DB parameter group to apply to all instances of the DB cluster.
	// When you apply a parameter group using the DBInstanceParameterGroupName
	// parameter, the DB cluster isn't rebooted automatically. Also, parameter changes
	// are applied immediately rather than during the next maintenance window. Valid
	// for Cluster Type: Aurora DB clusters only Default: The existing name setting
	// Constraints:
	//   - The DB parameter group must be in the same DB parameter group family as
	//   this DB cluster.
	//   - The DBInstanceParameterGroupName parameter is valid in combination with the
	//   AllowMajorVersionUpgrade parameter for a major version upgrade only.
	DBInstanceParameterGroupName *string

	// Specifies whether the DB cluster has deletion protection enabled. The database
	// can't be deleted when deletion protection is enabled. By default, deletion
	// protection isn't enabled. Valid for Cluster Type: Aurora DB clusters and
	// Multi-AZ DB clusters
	DeletionProtection *bool

	// The Active Directory directory ID to move the DB cluster to. Specify none to
	// remove the cluster from its current domain. The domain must be created prior to
	// this operation. For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html)
	// in the Amazon Aurora User Guide. Valid for Cluster Type: Aurora DB clusters only
	Domain *string

	// The name of the IAM role to use when making API calls to the Directory Service.
	// Valid for Cluster Type: Aurora DB clusters only
	DomainIAMRoleName *string

	// Specifies whether to enable this DB cluster to forward write operations to the
	// primary cluster of a global cluster (Aurora global database). By default, write
	// operations are not allowed on Aurora DB clusters that are secondary clusters in
	// an Aurora global database. You can set this value only on Aurora DB clusters
	// that are members of an Aurora global database. With this parameter enabled, a
	// secondary cluster can forward writes to the current primary cluster, and the
	// resulting changes are replicated back to this cluster. For the primary DB
	// cluster of an Aurora global database, this value is used immediately if the
	// primary is demoted by a global cluster API operation, but it does nothing until
	// then. Valid for Cluster Type: Aurora DB clusters only
	EnableGlobalWriteForwarding *bool

	// Specifies whether to enable the HTTP endpoint for an Aurora Serverless v1 DB
	// cluster. By default, the HTTP endpoint is disabled. When enabled, the HTTP
	// endpoint provides a connectionless web service API for running SQL queries on
	// the Aurora Serverless v1 DB cluster. You can also query your database from
	// inside the RDS console with the query editor. For more information, see Using
	// the Data API for Aurora Serverless v1 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html)
	// in the Amazon Aurora User Guide. Valid for Cluster Type: Aurora DB clusters only
	EnableHttpEndpoint *bool

	// Specifies whether to enable mapping of Amazon Web Services Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping isn't
	// enabled. For more information, see IAM Database Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon Aurora User Guide. Valid for Cluster Type: Aurora DB clusters only
	EnableIAMDatabaseAuthentication *bool

	// Specifies whether read replicas can forward write operations to the writer DB
	// instance in the DB cluster. By default, write operations aren't allowed on
	// reader DB instances. Valid for: Aurora DB clusters only
	EnableLocalWriteForwarding *bool

	// Specifies whether to turn on Performance Insights for the DB cluster. For more
	// information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon RDS User Guide. Valid for Cluster Type: Multi-AZ DB clusters only
	EnablePerformanceInsights *bool

	// The DB engine mode of the DB cluster, either provisioned or serverless . The DB
	// engine mode can be modified only from serverless to provisioned . For more
	// information, see CreateDBCluster (https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBCluster.html)
	// . Valid for Cluster Type: Aurora DB clusters only
	EngineMode *string

	// The version number of the database engine to which you want to upgrade.
	// Changing this parameter results in an outage. The change is applied during the
	// next maintenance window unless ApplyImmediately is enabled. If the cluster that
	// you're modifying has one or more read replicas, all replicas must be running an
	// engine version that's the same or later than the version you specify. To list
	// all of the available engine versions for Aurora MySQL, use the following
	// command: aws rds describe-db-engine-versions --engine aurora-mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for Aurora PostgreSQL, use the following command: aws rds
	// describe-db-engine-versions --engine aurora-postgresql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for RDS for MySQL, use the following command: aws rds
	// describe-db-engine-versions --engine mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for RDS for PostgreSQL, use the following command: aws rds
	// describe-db-engine-versions --engine postgres --query
	// "DBEngineVersions[].EngineVersion" Valid for Cluster Type: Aurora DB clusters
	// and Multi-AZ DB clusters
	EngineVersion *string

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for each DB instance in the Multi-AZ DB cluster. For
	// information about valid IOPS values, see Amazon RDS Provisioned IOPS storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide. Valid for Cluster Type: Multi-AZ DB clusters only
	// Constraints:
	//   - Must be a multiple between .5 and 50 of the storage amount for the DB
	//   cluster.
	Iops *int32

	// Specifies whether to manage the master user password with Amazon Web Services
	// Secrets Manager. If the DB cluster doesn't manage the master user password with
	// Amazon Web Services Secrets Manager, you can turn on this management. In this
	// case, you can't specify MasterUserPassword . If the DB cluster already manages
	// the master user password with Amazon Web Services Secrets Manager, and you
	// specify that the master user password is not managed with Amazon Web Services
	// Secrets Manager, then you must specify MasterUserPassword . In this case, RDS
	// deletes the secret and uses the new password for the master user specified by
	// MasterUserPassword . For more information, see Password management with Amazon
	// Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html)
	// in the Amazon RDS User Guide and Password management with Amazon Web Services
	// Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html)
	// in the Amazon Aurora User Guide. Valid for Cluster Type: Aurora DB clusters and
	// Multi-AZ DB clusters
	ManageMasterUserPassword *bool

	// The new password for the master database user. Valid for Cluster Type: Aurora
	// DB clusters and Multi-AZ DB clusters Constraints:
	//   - Must contain from 8 to 41 characters.
	//   - Can contain any printable ASCII character except "/", """, or "@".
	//   - Can't be specified if ManageMasterUserPassword is turned on.
	MasterUserPassword *string

	// The Amazon Web Services KMS key identifier to encrypt a secret that is
	// automatically generated and managed in Amazon Web Services Secrets Manager. This
	// setting is valid only if both of the following conditions are met:
	//   - The DB cluster doesn't manage the master user password in Amazon Web
	//   Services Secrets Manager. If the DB cluster already manages the master user
	//   password in Amazon Web Services Secrets Manager, you can't change the KMS key
	//   that is used to encrypt the secret.
	//   - You are turning on ManageMasterUserPassword to manage the master user
	//   password in Amazon Web Services Secrets Manager. If you are turning on
	//   ManageMasterUserPassword and don't specify MasterUserSecretKmsKeyId , then the
	//   aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in
	//   a different Amazon Web Services account, then you can't use the
	//   aws/secretsmanager KMS key to encrypt the secret, and you must use a customer
	//   managed KMS key.
	// The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN,
	// or alias name for the KMS key. To use a KMS key in a different Amazon Web
	// Services account, specify the key ARN or alias ARN. There is a default KMS key
	// for your Amazon Web Services account. Your Amazon Web Services account has a
	// different default KMS key for each Amazon Web Services Region. Valid for Cluster
	// Type: Aurora DB clusters and Multi-AZ DB clusters
	MasterUserSecretKmsKeyId *string

	// The interval, in seconds, between points when Enhanced Monitoring metrics are
	// collected for the DB cluster. To turn off collecting Enhanced Monitoring
	// metrics, specify 0 . If MonitoringRoleArn is specified, also set
	// MonitoringInterval to a value other than 0 . Valid for Cluster Type: Multi-AZ DB
	// clusters only Valid Values: 0 | 1 | 5 | 10 | 15 | 30 | 60 Default: 0
	MonitoringInterval *int32

	// The Amazon Resource Name (ARN) for the IAM role that permits RDS to send
	// Enhanced Monitoring metrics to Amazon CloudWatch Logs. An example is
	// arn:aws:iam:123456789012:role/emaccess . For information on creating a
	// monitoring role, see To create an IAM role for Amazon RDS Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html#USER_Monitoring.OS.IAMRole)
	// in the Amazon RDS User Guide. If MonitoringInterval is set to a value other
	// than 0 , supply a MonitoringRoleArn value. Valid for Cluster Type: Multi-AZ DB
	// clusters only
	MonitoringRoleArn *string

	// The network type of the DB cluster. The network type is determined by the
	// DBSubnetGroup specified for the DB cluster. A DBSubnetGroup can support only
	// the IPv4 protocol or the IPv4 and the IPv6 protocols ( DUAL ). For more
	// information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html)
	// in the Amazon Aurora User Guide. Valid for Cluster Type: Aurora DB clusters only
	// Valid Values: IPV4 | DUAL
	NetworkType *string

	// The new DB cluster identifier for the DB cluster when renaming a DB cluster.
	// This value is stored as a lowercase string. Valid for Cluster Type: Aurora DB
	// clusters and Multi-AZ DB clusters Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//   - The first character must be a letter.
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	// Example: my-cluster2
	NewDBClusterIdentifier *string

	// The option group to associate the DB cluster with. DB clusters are associated
	// with a default option group that can't be modified.
	OptionGroupName *string

	// The Amazon Web Services KMS key identifier for encryption of Performance
	// Insights data. The Amazon Web Services KMS key identifier is the key ARN, key
	// ID, alias ARN, or alias name for the KMS key. If you don't specify a value for
	// PerformanceInsightsKMSKeyId , then Amazon RDS uses your default KMS key. There
	// is a default KMS key for your Amazon Web Services account. Your Amazon Web
	// Services account has a different default KMS key for each Amazon Web Services
	// Region. Valid for Cluster Type: Multi-AZ DB clusters only
	PerformanceInsightsKMSKeyId *string

	// The number of days to retain Performance Insights data. Valid for Cluster Type:
	// Multi-AZ DB clusters only Valid Values:
	//   - 7
	//   - month * 31, where month is a number of months from 1-23. Examples: 93 (3
	//   months * 31), 341 (11 months * 31), 589 (19 months * 31)
	//   - 731
	// Default: 7 days If you specify a retention period that isn't valid, such as 94 ,
	// Amazon RDS issues an error.
	PerformanceInsightsRetentionPeriod *int32

	// The port number on which the DB cluster accepts connections. Valid for Cluster
	// Type: Aurora DB clusters only Valid Values: 1150-65535 Default: The same port
	// as the original DB cluster.
	Port *int32

	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. The default is
	// a 30-minute window selected at random from an 8-hour block of time for each
	// Amazon Web Services Region. To view the time blocks available, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow)
	// in the Amazon Aurora User Guide. Valid for Cluster Type: Aurora DB clusters and
	// Multi-AZ DB clusters Constraints:
	//   - Must be in the format hh24:mi-hh24:mi .
	//   - Must be in Universal Coordinated Time (UTC).
	//   - Must not conflict with the preferred maintenance window.
	//   - Must be at least 30 minutes.
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Valid for Cluster Type: Aurora DB clusters and Multi-AZ
	// DB clusters The default is a 30-minute window selected at random from an 8-hour
	// block of time for each Amazon Web Services Region, occurring on a random day of
	// the week. To see the time blocks available, see Adjusting the Preferred DB
	// Cluster Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Constraints:
	//   - Must be in the format ddd:hh24:mi-ddd:hh24:mi .
	//   - Days must be one of Mon | Tue | Wed | Thu | Fri | Sat | Sun .
	//   - Must be in Universal Coordinated Time (UTC).
	//   - Must be at least 30 minutes.
	PreferredMaintenanceWindow *string

	// Specifies whether to rotate the secret managed by Amazon Web Services Secrets
	// Manager for the master user password. This setting is valid only if the master
	// user password is managed by RDS in Amazon Web Services Secrets Manager for the
	// DB cluster. The secret value contains the updated password. For more
	// information, see Password management with Amazon Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html)
	// in the Amazon RDS User Guide and Password management with Amazon Web Services
	// Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html)
	// in the Amazon Aurora User Guide. Valid for Cluster Type: Aurora DB clusters and
	// Multi-AZ DB clusters Constraints:
	//   - You must apply the change immediately when rotating the master user
	//   password.
	RotateMasterUserPassword *bool

	// The scaling properties of the DB cluster. You can only modify scaling
	// properties for DB clusters in serverless DB engine mode. Valid for Cluster
	// Type: Aurora DB clusters only
	ScalingConfiguration *types.ScalingConfiguration

	// Contains the scaling configuration of an Aurora Serverless v2 DB cluster. For
	// more information, see Using Amazon Aurora Serverless v2 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html)
	// in the Amazon Aurora User Guide.
	ServerlessV2ScalingConfiguration *types.ServerlessV2ScalingConfiguration

	// The storage type to associate with the DB cluster. For information on storage
	// types for Aurora DB clusters, see Storage configurations for Amazon Aurora DB
	// clusters (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.StorageReliability.html#aurora-storage-type)
	// . For information on storage types for Multi-AZ DB clusters, see Settings for
	// creating Multi-AZ DB clusters (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/create-multi-az-db-cluster.html#create-multi-az-db-cluster-settings)
	// . When specified for a Multi-AZ DB cluster, a value for the Iops parameter is
	// required. Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
	// Valid Values:
	//   - Aurora DB clusters - aurora | aurora-iopt1
	//   - Multi-AZ DB clusters - io1
	// Default:
	//   - Aurora DB clusters - aurora
	//   - Multi-AZ DB clusters - io1
	StorageType *string

	// A list of EC2 VPC security groups to associate with this DB cluster. Valid for
	// Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type ModifyDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster. For
	// an Amazon Aurora DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , PromoteReadReplicaDBCluster ,
	// RestoreDBClusterFromS3 , RestoreDBClusterFromSnapshot ,
	// RestoreDBClusterToPointInTime , StartDBCluster , and StopDBCluster . For a
	// Multi-AZ DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , RebootDBCluster ,
	// RestoreDBClusterFromSnapshot , and RestoreDBClusterToPointInTime . For more
	// information on Amazon Aurora DB clusters, see What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
	// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
	// see Multi-AZ deployments with two readable standby DB instances (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
	// in the Amazon RDS User Guide.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addModifyDBClusterResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpModifyDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBCluster(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBCluster",
	}
}

type opModifyDBClusterResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opModifyDBClusterResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opModifyDBClusterResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "rds"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "rds"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("rds")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addModifyDBClusterResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opModifyDBClusterResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

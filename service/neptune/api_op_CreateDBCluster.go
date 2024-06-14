// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon Neptune DB cluster.
//
// You can use the ReplicationSourceIdentifier parameter to create the DB cluster
// as a Read Replica of another DB cluster or Amazon Neptune DB instance.
//
// Note that when you create a new cluster using CreateDBCluster directly,
// deletion protection is disabled by default (when you create a new production
// cluster in the console, deletion protection is enabled by default). You can only
// delete a DB cluster if its DeletionProtection field is set to false .
func (c *Client) CreateDBCluster(ctx context.Context, params *CreateDBClusterInput, optFns ...func(*Options)) (*CreateDBClusterOutput, error) {
	if params == nil {
		params = &CreateDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBCluster", params, optFns, c.addOperationCreateDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBClusterInput struct {

	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//   - First character must be a letter.
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster1
	//
	// This member is required.
	DBClusterIdentifier *string

	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values: neptune
	//
	// This member is required.
	Engine *string

	// A list of EC2 Availability Zones that instances in the DB cluster can be
	// created in.
	AvailabilityZones []string

	// The number of days for which automated backups are retained. You must specify a
	// minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	//   - Must be a value from 1 to 35
	BackupRetentionPeriod *int32

	//  (Not supported by Neptune)
	CharacterSetName *string

	//  If set to true , tags are copied to any snapshot of the DB cluster that is
	// created.
	CopyTagsToSnapshot *bool

	//  The name of the DB cluster parameter group to associate with this DB cluster.
	// If this argument is omitted, the default is used.
	//
	// Constraints:
	//
	//   - If supplied, must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string

	// A DB subnet group to associate with this DB cluster.
	//
	// Constraints: Must match the name of an existing DBSubnetGroup. Must not be
	// default.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string

	// The name for your database of up to 64 alpha-numeric characters. If you do not
	// provide a name, Amazon Neptune will not create a database in the DB cluster you
	// are creating.
	DatabaseName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is enabled.
	DeletionProtection *bool

	// A list of the log types that this DB cluster should export to CloudWatch Logs.
	// Valid log types are: audit (to publish audit logs) and slowquery (to publish
	// slow-query logs). See [Publishing Neptune logs to Amazon CloudWatch logs].
	//
	// [Publishing Neptune logs to Amazon CloudWatch logs]: https://docs.aws.amazon.com/neptune/latest/userguide/cloudwatch-logs.html
	EnableCloudwatchLogsExports []string

	// If set to true , enables Amazon Identity and Access Management (IAM)
	// authentication for the entire DB cluster (this cannot be set at an instance
	// level).
	//
	// Default: false .
	EnableIAMDatabaseAuthentication *bool

	// The version number of the database engine to use for the new DB cluster.
	//
	// Example: 1.0.2.1
	EngineVersion *string

	// The ID of the Neptune global database to which this new DB cluster should be
	// added.
	GlobalClusterIdentifier *string

	// The Amazon KMS key identifier for an encrypted DB cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are creating a DB cluster with the same Amazon account that owns the
	// KMS encryption key used to encrypt the new DB cluster, then you can use the KMS
	// key alias instead of the ARN for the KMS encryption key.
	//
	// If an encryption key is not specified in KmsKeyId :
	//
	//   - If ReplicationSourceIdentifier identifies an encrypted source, then Amazon
	//   Neptune will use the encryption key used to encrypt the source. Otherwise,
	//   Amazon Neptune will use your default encryption key.
	//
	//   - If the StorageEncrypted parameter is true and ReplicationSourceIdentifier is
	//   not specified, then Amazon Neptune will use your default encryption key.
	//
	// Amazon KMS creates the default encryption key for your Amazon account. Your
	// Amazon account has a different default encryption key for each Amazon Region.
	//
	// If you create a Read Replica of an encrypted DB cluster in another Amazon
	// Region, you must set KmsKeyId to a KMS key ID that is valid in the destination
	// Amazon Region. This key is used to encrypt the Read Replica in that Amazon
	// Region.
	KmsKeyId *string

	// Not supported by Neptune.
	MasterUserPassword *string

	// Not supported by Neptune.
	MasterUsername *string

	//  (Not supported by Neptune)
	OptionGroupName *string

	// The port number on which the instances in the DB cluster accept connections.
	//
	// Default: 8182
	Port *int32

	// This parameter is not currently supported.
	PreSignedUrl *string

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block of
	// time for each Amazon Region. To see the time blocks available, see [Neptune Maintenance Window]in the
	// Amazon Neptune User Guide.
	//
	// Constraints:
	//
	//   - Must be in the format hh24:mi-hh24:mi .
	//
	//   - Must be in Universal Coordinated Time (UTC).
	//
	//   - Must not conflict with the preferred maintenance window.
	//
	//   - Must be at least 30 minutes.
	//
	// [Neptune Maintenance Window]: https://docs.aws.amazon.com/neptune/latest/userguide/manage-console-maintaining.html#manage-console-maintaining-window
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block of
	// time for each Amazon Region, occurring on a random day of the week. To see the
	// time blocks available, see [Neptune Maintenance Window]in the Amazon Neptune User Guide.
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	//
	// [Neptune Maintenance Window]: https://docs.aws.amazon.com/neptune/latest/userguide/manage-console-maintaining.html#manage-console-maintaining-window
	PreferredMaintenanceWindow *string

	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this
	// DB cluster is created as a Read Replica.
	ReplicationSourceIdentifier *string

	// Contains the scaling configuration of a Neptune Serverless DB cluster.
	//
	// For more information, see [Using Amazon Neptune Serverless] in the Amazon Neptune User Guide.
	//
	// [Using Amazon Neptune Serverless]: https://docs.aws.amazon.com/neptune/latest/userguide/neptune-serverless-using.html
	ServerlessV2ScalingConfiguration *types.ServerlessV2ScalingConfiguration

	// The AWS region the resource is in. The presigned URL will be created with this
	// region, if the PresignURL member is empty set.
	SourceRegion *string

	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool

	// The storage type to associate with the DB cluster.
	//
	// Valid Values:
	//
	//   - standard | iopt1
	//
	// Default:
	//
	//   - standard
	//
	// When you create a Neptune cluster with the storage type set to iopt1 , the
	// storage type is returned in the response. The storage type isn't returned when
	// you set it to standard .
	StorageType *string

	// The tags to assign to the new DB cluster.
	Tags []types.Tag

	// A list of EC2 VPC security groups to associate with this DB cluster.
	VpcSecurityGroupIds []string

	// Used by the SDK's PresignURL autofill customization to specify the region the
	// of the client's request.
	destinationRegion *string

	noSmithyDocumentSerde
}

type CreateDBClusterOutput struct {

	// Contains the details of an Amazon Neptune DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDBCluster"); err != nil {
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
	if err = addCreateDBClusterPresignURLMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpCreateDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBCluster(options.Region), middleware.Before); err != nil {
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

func copyCreateDBClusterInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCreateDBClusterPreSignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	if input.PreSignedUrl == nil || len(*input.PreSignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PreSignedUrl, true, nil
}
func getCreateDBClusterSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCreateDBClusterPreSignedUrl(params interface{}, value string) error {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	input.PreSignedUrl = &value
	return nil
}
func setCreateDBClusterdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}
func addCreateDBClusterPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL: getCreateDBClusterPreSignedUrl,

			GetSourceRegion: getCreateDBClusterSourceRegion,

			CopyInput: copyCreateDBClusterInputForPresign,

			SetDestinationRegion: setCreateDBClusterdestinationRegion,

			SetPresignedURL: setCreateDBClusterPreSignedUrl,
		},
		Presigner: &presignAutoFillCreateDBClusterClient{client: NewPresignClient(New(options))},
	})
}

type presignAutoFillCreateDBClusterClient struct {
	client *PresignClient
}

// PresignURL is a middleware accessor that satisfies URLPresigner interface.
func (c *presignAutoFillCreateDBClusterClient) PresignURL(ctx context.Context, srcRegion string, params interface{}) (*v4.PresignedHTTPRequest, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = srcRegion
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	presignOptFn := WithPresignClientFromClientOptions(optFn)
	return c.client.PresignCreateDBCluster(ctx, input, presignOptFn)
}

func newServiceMetadataMiddleware_opCreateDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDBCluster",
	}
}

// PresignCreateDBCluster is used to generate a presigned HTTP Request which
// contains presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignCreateDBCluster(ctx context.Context, params *CreateDBClusterInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil {
		params = &CreateDBClusterInput{}
	}
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)

	result, _, err := c.client.invokeOperation(ctx, "CreateDBCluster", params, clientOptFns,
		c.client.addOperationCreateDBClusterMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the specified broker.
func (c *Client) DescribeBroker(ctx context.Context, params *DescribeBrokerInput, optFns ...func(*Options)) (*DescribeBrokerOutput, error) {
	if params == nil {
		params = &DescribeBrokerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBroker", params, optFns, c.addOperationDescribeBrokerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBrokerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBrokerInput struct {

	// The unique ID that Amazon MQ generates for the broker.
	//
	// This member is required.
	BrokerId *string

	noSmithyDocumentSerde
}

type DescribeBrokerOutput struct {

	// Actions required for a broker.
	ActionsRequired []types.ActionRequired

	// The authentication strategy used to secure the broker. The default is SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// Enables automatic upgrades to new minor versions for brokers, as new versions
	// are released and supported by Amazon MQ. Automatic upgrades occur during the
	// scheduled maintenance window of the broker or after a manual broker reboot.
	AutoMinorVersionUpgrade *bool

	// The broker's Amazon Resource Name (ARN).
	BrokerArn *string

	// The unique ID that Amazon MQ generates for the broker.
	BrokerId *string

	// A list of information about allocated brokers.
	BrokerInstances []types.BrokerInstance

	// The broker's name. This value must be unique in your Amazon Web Services
	// account account, 1-50 characters long, must contain only letters, numbers,
	// dashes, and underscores, and must not contain white spaces, brackets, wildcard
	// characters, or special characters.
	BrokerName *string

	// The broker's status.
	BrokerState types.BrokerState

	// The list of all revisions for the specified configuration.
	Configurations *types.Configurations

	// The time when the broker was created.
	Created *time.Time

	// The replication details of the data replication-enabled broker. Only returned
	// if dataReplicationMode is set to CRDR.
	DataReplicationMetadata *types.DataReplicationMetadataOutput

	// Describes whether this broker is a part of a data replication pair.
	DataReplicationMode types.DataReplicationMode

	// The broker's deployment mode.
	DeploymentMode types.DeploymentMode

	// Encryption options for the broker.
	EncryptionOptions *types.EncryptionOptions

	// The type of broker engine. Currently, Amazon MQ supports ACTIVEMQ and RABBITMQ.
	EngineType types.EngineType

	// The broker engine's version. For a list of supported engine versions, see [Supported engines].
	//
	// [Supported engines]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker-engine.html
	EngineVersion *string

	// The broker's instance type.
	HostInstanceType *string

	// The metadata of the LDAP server used to authenticate and authorize connections
	// to the broker.
	LdapServerMetadata *types.LdapServerMetadataOutput

	// The list of information about logs currently enabled and pending to be deployed
	// for the specified broker.
	Logs *types.LogsSummary

	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime *types.WeeklyStartTime

	// The authentication strategy that will be applied when the broker is rebooted.
	// The default is SIMPLE.
	PendingAuthenticationStrategy types.AuthenticationStrategy

	// The pending replication details of the data replication-enabled broker. Only
	// returned if pendingDataReplicationMode is set to CRDR.
	PendingDataReplicationMetadata *types.DataReplicationMetadataOutput

	// Describes whether this broker will be a part of a data replication pair after
	// reboot.
	PendingDataReplicationMode types.DataReplicationMode

	// The broker engine version to upgrade to. For a list of supported engine
	// versions, see [Supported engines].
	//
	// [Supported engines]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker-engine.html
	PendingEngineVersion *string

	// The broker's host instance type to upgrade to. For a list of supported instance
	// types, see [Broker instance types].
	//
	// [Broker instance types]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker.html#broker-instance-types
	PendingHostInstanceType *string

	// The metadata of the LDAP server that will be used to authenticate and authorize
	// connections to the broker after it is rebooted.
	PendingLdapServerMetadata *types.LdapServerMetadataOutput

	// The list of pending security groups to authorize connections to brokers.
	PendingSecurityGroups []string

	// Enables connections from applications outside of the VPC that hosts the
	// broker's subnets.
	PubliclyAccessible *bool

	// The list of rules (1 minimum, 125 maximum) that authorize connections to
	// brokers.
	SecurityGroups []string

	// The broker's storage type.
	StorageType types.BrokerStorageType

	// The list of groups that define which subnets and IP ranges the broker can use
	// from different Availability Zones.
	SubnetIds []string

	// The list of all tags associated with this broker.
	Tags map[string]string

	// The list of all broker usernames for the specified broker.
	Users []types.UserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBrokerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBroker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBroker{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBroker"); err != nil {
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
	if err = addOpDescribeBrokerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBroker(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBroker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBroker",
	}
}

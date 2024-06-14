// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disconnects a specific Source Server from Elastic Disaster Recovery. Data
// replication is stopped immediately. All AWS resources created by Elastic
// Disaster Recovery for enabling the replication of the Source Server will be
// terminated / deleted within 90 minutes. You cannot disconnect a Source Server if
// it has a Recovery Instance. If the agent on the Source Server has not been
// prevented from communicating with the Elastic Disaster Recovery service, then it
// will receive a command to uninstall itself (within approximately 10 minutes).
// The following properties of the SourceServer will be changed immediately:
// dataReplicationInfo.dataReplicationState will be set to DISCONNECTED; The
// totalStorageBytes property for each of dataReplicationInfo.replicatedDisks will
// be set to zero; dataReplicationInfo.lagDuration and
// dataReplicationInfo.lagDuration will be nullified.
func (c *Client) DisconnectSourceServer(ctx context.Context, params *DisconnectSourceServerInput, optFns ...func(*Options)) (*DisconnectSourceServerOutput, error) {
	if params == nil {
		params = &DisconnectSourceServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisconnectSourceServer", params, optFns, c.addOperationDisconnectSourceServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisconnectSourceServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisconnectSourceServerInput struct {

	// The ID of the Source Server to disconnect.
	//
	// This member is required.
	SourceServerID *string

	noSmithyDocumentSerde
}

type DisconnectSourceServerOutput struct {

	// The version of the DRS agent installed on the source server
	AgentVersion *string

	// The ARN of the Source Server.
	Arn *string

	// The Data Replication Info of the Source Server.
	DataReplicationInfo *types.DataReplicationInfo

	// The status of the last recovery launch of this Source Server.
	LastLaunchResult types.LastLaunchResult

	// The lifecycle information of this Source Server.
	LifeCycle *types.LifeCycle

	// The ID of the Recovery Instance associated with this Source Server.
	RecoveryInstanceId *string

	// Replication direction of the Source Server.
	ReplicationDirection types.ReplicationDirection

	// For EC2-originated Source Servers which have been failed over and then failed
	// back, this value will mean the ARN of the Source Server on the opposite
	// replication direction.
	ReversedDirectionSourceServerArn *string

	// Source cloud properties of the Source Server.
	SourceCloudProperties *types.SourceCloudProperties

	// ID of the Source Network which is protecting this Source Server's network.
	SourceNetworkID *string

	// The source properties of the Source Server.
	SourceProperties *types.SourceProperties

	// The ID of the Source Server.
	SourceServerID *string

	// The staging area of the source server.
	StagingArea *types.StagingArea

	// The tags associated with the Source Server.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisconnectSourceServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisconnectSourceServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisconnectSourceServer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisconnectSourceServer"); err != nil {
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
	if err = addOpDisconnectSourceServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisconnectSourceServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisconnectSourceServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisconnectSourceServer",
	}
}

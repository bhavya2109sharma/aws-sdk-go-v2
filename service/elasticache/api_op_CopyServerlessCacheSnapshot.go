// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a copy of an existing serverless cache’s snapshot. Available for Redis
// only.
func (c *Client) CopyServerlessCacheSnapshot(ctx context.Context, params *CopyServerlessCacheSnapshotInput, optFns ...func(*Options)) (*CopyServerlessCacheSnapshotOutput, error) {
	if params == nil {
		params = &CopyServerlessCacheSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyServerlessCacheSnapshot", params, optFns, c.addOperationCopyServerlessCacheSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyServerlessCacheSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyServerlessCacheSnapshotInput struct {

	// The identifier of the existing serverless cache’s snapshot to be copied.
	// Available for Redis only.
	//
	// This member is required.
	SourceServerlessCacheSnapshotName *string

	// The identifier for the snapshot to be created. Available for Redis only.
	//
	// This member is required.
	TargetServerlessCacheSnapshotName *string

	// The identifier of the KMS key used to encrypt the target snapshot. Available
	// for Redis only.
	KmsKeyId *string

	// A list of tags to be added to the target snapshot resource. A tag is a
	// key-value pair. Available for Redis only. Default: NULL
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CopyServerlessCacheSnapshotOutput struct {

	// The response for the attempt to copy the serverless cache snapshot. Available
	// for Redis only.
	ServerlessCacheSnapshot *types.ServerlessCacheSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyServerlessCacheSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyServerlessCacheSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyServerlessCacheSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopyServerlessCacheSnapshot"); err != nil {
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
	if err = addOpCopyServerlessCacheSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyServerlessCacheSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopyServerlessCacheSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopyServerlessCacheSnapshot",
	}
}

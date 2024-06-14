// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists an environment's upcoming and in-progress managed actions.
func (c *Client) DescribeEnvironmentManagedActions(ctx context.Context, params *DescribeEnvironmentManagedActionsInput, optFns ...func(*Options)) (*DescribeEnvironmentManagedActionsOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentManagedActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentManagedActions", params, optFns, c.addOperationDescribeEnvironmentManagedActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentManagedActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to list an environment's upcoming and in-progress managed actions.
type DescribeEnvironmentManagedActionsInput struct {

	// The environment ID of the target environment.
	EnvironmentId *string

	// The name of the target environment.
	EnvironmentName *string

	// To show only actions with a particular status, specify a status.
	Status types.ActionStatus

	noSmithyDocumentSerde
}

// The result message containing a list of managed actions.
type DescribeEnvironmentManagedActionsOutput struct {

	// A list of upcoming and in-progress managed actions.
	ManagedActions []types.ManagedAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEnvironmentManagedActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironmentManagedActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironmentManagedActions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEnvironmentManagedActions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentManagedActions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEnvironmentManagedActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEnvironmentManagedActions",
	}
}

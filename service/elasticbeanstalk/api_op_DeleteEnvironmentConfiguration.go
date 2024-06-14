// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the draft configuration associated with the running environment.
//
// Updating a running environment with any configuration changes creates a draft
// configuration set. You can get the draft configuration using DescribeConfigurationSettingswhile the update
// is in progress or if the update fails. The DeploymentStatus for the draft
// configuration indicates whether the deployment is in process or has failed. The
// draft configuration remains in existence until it is deleted with this action.
func (c *Client) DeleteEnvironmentConfiguration(ctx context.Context, params *DeleteEnvironmentConfigurationInput, optFns ...func(*Options)) (*DeleteEnvironmentConfigurationOutput, error) {
	if params == nil {
		params = &DeleteEnvironmentConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEnvironmentConfiguration", params, optFns, c.addOperationDeleteEnvironmentConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEnvironmentConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to delete a draft environment configuration.
type DeleteEnvironmentConfigurationInput struct {

	// The name of the application the environment is associated with.
	//
	// This member is required.
	ApplicationName *string

	// The name of the environment to delete the draft configuration from.
	//
	// This member is required.
	EnvironmentName *string

	noSmithyDocumentSerde
}

type DeleteEnvironmentConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteEnvironmentConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteEnvironmentConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteEnvironmentConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteEnvironmentConfiguration"); err != nil {
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
	if err = addOpDeleteEnvironmentConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEnvironmentConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEnvironmentConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteEnvironmentConfiguration",
	}
}

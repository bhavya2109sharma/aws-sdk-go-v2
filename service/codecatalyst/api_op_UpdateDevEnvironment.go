// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes one or more values for a Dev Environment. Updating certain values of
// the Dev Environment will cause a restart.
func (c *Client) UpdateDevEnvironment(ctx context.Context, params *UpdateDevEnvironmentInput, optFns ...func(*Options)) (*UpdateDevEnvironmentOutput, error) {
	if params == nil {
		params = &UpdateDevEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDevEnvironment", params, optFns, c.addOperationUpdateDevEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDevEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDevEnvironmentInput struct {

	// The system-generated unique ID of the Dev Environment.
	//
	// This member is required.
	Id *string

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The user-specified alias for the Dev Environment. Changing this value will not
	// cause a restart.
	Alias *string

	// A user-specified idempotency token. Idempotency ensures that an API request
	// completes only once. With an idempotent request, if the original request
	// completes successfully, the subsequent retries return the result from the
	// original successful request and have no additional effect.
	ClientToken *string

	// Information about the integrated development environment (IDE) configured for a
	// Dev Environment.
	Ides []types.IdeConfiguration

	// The amount of time the Dev Environment will run without any activity detected
	// before stopping, in minutes. Only whole integers are allowed. Dev Environments
	// consume compute minutes when running.
	//
	// Changing this value will cause a restart of the Dev Environment if it is
	// running.
	InactivityTimeoutMinutes int32

	// The Amazon EC2 instace type to use for the Dev Environment.
	//
	// Changing this value will cause a restart of the Dev Environment if it is
	// running.
	InstanceType types.InstanceType

	noSmithyDocumentSerde
}

type UpdateDevEnvironmentOutput struct {

	// The system-generated unique ID of the Dev Environment.
	//
	// This member is required.
	Id *string

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The user-specified alias for the Dev Environment.
	Alias *string

	// A user-specified idempotency token. Idempotency ensures that an API request
	// completes only once. With an idempotent request, if the original request
	// completes successfully, the subsequent retries return the result from the
	// original successful request and have no additional effect.
	ClientToken *string

	// Information about the integrated development environment (IDE) configured for
	// the Dev Environment.
	Ides []types.IdeConfiguration

	// The amount of time the Dev Environment will run without any activity detected
	// before stopping, in minutes.
	InactivityTimeoutMinutes int32

	// The Amazon EC2 instace type to use for the Dev Environment.
	InstanceType types.InstanceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDevEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDevEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDevEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDevEnvironment"); err != nil {
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
	if err = addOpUpdateDevEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDevEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDevEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDevEnvironment",
	}
}

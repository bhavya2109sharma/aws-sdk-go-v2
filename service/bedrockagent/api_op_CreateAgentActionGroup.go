// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an action group for an agent. An action group represents the actions
// that an agent can carry out for the customer by defining the APIs that an agent
// can call and the logic for calling them.
//
// To allow your agent to request the user for additional information when trying
// to complete a task, add an action group with the parentActionGroupSignature
// field set to AMAZON.UserInput . You must leave the description , apiSchema , and
// actionGroupExecutor fields blank for this action group. During orchestration, if
// your agent determines that it needs to invoke an API in an action group, but
// doesn't have enough information to complete the API request, it will invoke this
// action group instead and return an [Observation]reprompting the user for more information.
//
// [Observation]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_Observation.html
func (c *Client) CreateAgentActionGroup(ctx context.Context, params *CreateAgentActionGroupInput, optFns ...func(*Options)) (*CreateAgentActionGroupOutput, error) {
	if params == nil {
		params = &CreateAgentActionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAgentActionGroup", params, optFns, c.addOperationCreateAgentActionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAgentActionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAgentActionGroupInput struct {

	// The name to give the action group.
	//
	// This member is required.
	ActionGroupName *string

	// The unique identifier of the agent for which to create the action group.
	//
	// This member is required.
	AgentId *string

	// The version of the agent for which to create the action group.
	//
	// This member is required.
	AgentVersion *string

	// The Amazon Resource Name (ARN) of the Lambda function containing the business
	// logic that is carried out upon invoking the action or the custom control method
	// for handling the information elicited from the user.
	ActionGroupExecutor types.ActionGroupExecutor

	// Specifies whether the action group is available for the agent to invoke or not
	// when sending an [InvokeAgent]request.
	//
	// [InvokeAgent]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html
	ActionGroupState types.ActionGroupState

	// Contains either details about the S3 object containing the OpenAPI schema for
	// the action group or the JSON or YAML-formatted payload defining the schema. For
	// more information, see [Action group OpenAPI schemas].
	//
	// [Action group OpenAPI schemas]: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-api-schema.html
	ApiSchema types.APISchema

	// A unique, case-sensitive identifier to ensure that the API request completes no
	// more than one time. If this token matches a previous request, Amazon Bedrock
	// ignores the request, but does not return an error. For more information, see [Ensuring idempotency].
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientToken *string

	// A description of the action group.
	Description *string

	// Contains details about the function schema for the action group or the JSON or
	// YAML-formatted payload defining the schema.
	FunctionSchema types.FunctionSchema

	// To allow your agent to request the user for additional information when trying
	// to complete a task, set this field to AMAZON.UserInput . You must leave the
	// description , apiSchema , and actionGroupExecutor fields blank for this action
	// group.
	//
	// During orchestration, if your agent determines that it needs to invoke an API
	// in an action group, but doesn't have enough information to complete the API
	// request, it will invoke this action group instead and return an [Observation]reprompting the
	// user for more information.
	//
	// [Observation]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_Observation.html
	ParentActionGroupSignature types.ActionGroupSignature

	noSmithyDocumentSerde
}

type CreateAgentActionGroupOutput struct {

	// Contains details about the action group that was created.
	//
	// This member is required.
	AgentActionGroup *types.AgentActionGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAgentActionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAgentActionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAgentActionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAgentActionGroup"); err != nil {
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
	if err = addIdempotencyToken_opCreateAgentActionGroupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAgentActionGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAgentActionGroup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAgentActionGroup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAgentActionGroup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAgentActionGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAgentActionGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAgentActionGroupInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAgentActionGroupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAgentActionGroup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAgentActionGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAgentActionGroup",
	}
}

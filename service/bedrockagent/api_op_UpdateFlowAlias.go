// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies the alias of a flow. Include both fields that you want to keep and
// ones that you want to change. For more information, see [Deploy a flow in Amazon Bedrock]in the Amazon Bedrock
// User Guide.
//
// [Deploy a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-deploy.html
func (c *Client) UpdateFlowAlias(ctx context.Context, params *UpdateFlowAliasInput, optFns ...func(*Options)) (*UpdateFlowAliasOutput, error) {
	if params == nil {
		params = &UpdateFlowAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFlowAlias", params, optFns, c.addOperationUpdateFlowAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFlowAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFlowAliasInput struct {

	// The unique identifier of the alias.
	//
	// This member is required.
	AliasIdentifier *string

	// The unique identifier of the flow.
	//
	// This member is required.
	FlowIdentifier *string

	// The name of the alias.
	//
	// This member is required.
	Name *string

	// Contains information about the version to which to map the alias.
	//
	// This member is required.
	RoutingConfiguration []types.FlowAliasRoutingConfigurationListItem

	// A description for the alias.
	Description *string

	noSmithyDocumentSerde
}

type UpdateFlowAliasOutput struct {

	// The Amazon Resource Name (ARN) of the flow.
	//
	// This member is required.
	Arn *string

	// The time at which the flow was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique identifier of the flow.
	//
	// This member is required.
	FlowId *string

	// The unique identifier of the alias.
	//
	// This member is required.
	Id *string

	// The name of the alias.
	//
	// This member is required.
	Name *string

	// Contains information about the version that the alias is mapped to.
	//
	// This member is required.
	RoutingConfiguration []types.FlowAliasRoutingConfigurationListItem

	// The time at which the alias was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The description of the flow.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFlowAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFlowAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFlowAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFlowAlias"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateFlowAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFlowAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFlowAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFlowAlias",
	}
}

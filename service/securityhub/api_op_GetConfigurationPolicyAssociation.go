// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

//	Returns the association between a configuration and a target account,
//
// organizational unit, or the root. The configuration can be a configuration
// policy or self-managed behavior. Only the Security Hub delegated administrator
// can invoke this operation from the home Region.
func (c *Client) GetConfigurationPolicyAssociation(ctx context.Context, params *GetConfigurationPolicyAssociationInput, optFns ...func(*Options)) (*GetConfigurationPolicyAssociationOutput, error) {
	if params == nil {
		params = &GetConfigurationPolicyAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfigurationPolicyAssociation", params, optFns, c.addOperationGetConfigurationPolicyAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfigurationPolicyAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfigurationPolicyAssociationInput struct {

	//  The target account ID, organizational unit ID, or the root ID to retrieve the
	// association for.
	//
	// This member is required.
	Target types.Target

	noSmithyDocumentSerde
}

type GetConfigurationPolicyAssociationOutput struct {

	//  The current status of the association between the specified target and the
	// configuration.
	AssociationStatus types.ConfigurationPolicyAssociationStatus

	//  The explanation for a FAILED value for AssociationStatus .
	AssociationStatusMessage *string

	//  Indicates whether the association between the specified target and the
	// configuration was directly applied by the Security Hub delegated administrator
	// or inherited from a parent.
	AssociationType types.AssociationType

	//  The universally unique identifier (UUID) of a configuration policy. For
	// self-managed behavior, the value is SELF_MANAGED_SECURITY_HUB .
	ConfigurationPolicyId *string

	//  The target account ID, organizational unit ID, or the root ID for which the
	// association is retrieved.
	TargetId *string

	//  Specifies whether the target is an Amazon Web Services account, organizational
	// unit, or the organization root.
	TargetType types.TargetType

	//  The date and time, in UTC and ISO 8601 format, that the configuration policy
	// association was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConfigurationPolicyAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConfigurationPolicyAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfigurationPolicyAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetConfigurationPolicyAssociation"); err != nil {
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
	if err = addOpGetConfigurationPolicyAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfigurationPolicyAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConfigurationPolicyAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetConfigurationPolicyAssociation",
	}
}

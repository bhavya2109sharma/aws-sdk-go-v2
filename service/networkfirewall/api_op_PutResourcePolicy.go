// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates an IAM policy for your rule group or firewall policy. Use
// this to share rule groups and firewall policies between accounts. This operation
// works in conjunction with the Amazon Web Services Resource Access Manager (RAM)
// service to manage resource sharing for Network Firewall.
//
// Use this operation to create or update a resource policy for your rule group or
// firewall policy. In the policy, you specify the accounts that you want to share
// the resource with and the operations that you want the accounts to be able to
// perform.
//
// When you add an account in the resource policy, you then run the following
// Resource Access Manager (RAM) operations to access and accept the shared rule
// group or firewall policy.
//
// [GetResourceShareInvitations]
//   - - Returns the Amazon Resource Names (ARNs) of the resource share
//     invitations.
//
// [AcceptResourceShareInvitation]
//   - - Accepts the share invitation for a specified resource share.
//
// For additional information about resource sharing using RAM, see [Resource Access Manager User Guide].
//
// [AcceptResourceShareInvitation]: https://docs.aws.amazon.com/ram/latest/APIReference/API_AcceptResourceShareInvitation.html
// [GetResourceShareInvitations]: https://docs.aws.amazon.com/ram/latest/APIReference/API_GetResourceShareInvitations.html
// [Resource Access Manager User Guide]: https://docs.aws.amazon.com/ram/latest/userguide/what-is.html
func (c *Client) PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) {
	if params == nil {
		params = &PutResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResourcePolicy", params, optFns, c.addOperationPutResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourcePolicyInput struct {

	// The IAM policy statement that lists the accounts that you want to share your
	// rule group or firewall policy with and the operations that you want the accounts
	// to be able to perform.
	//
	// For a rule group resource, you can specify the following operations in the
	// Actions section of the statement:
	//
	//   - network-firewall:CreateFirewallPolicy
	//
	//   - network-firewall:UpdateFirewallPolicy
	//
	//   - network-firewall:ListRuleGroups
	//
	// For a firewall policy resource, you can specify the following operations in the
	// Actions section of the statement:
	//
	//   - network-firewall:AssociateFirewallPolicy
	//
	//   - network-firewall:ListFirewallPolicies
	//
	// In the Resource section of the statement, you specify the ARNs for the rule
	// groups and firewall policies that you want to share with the account that you
	// specified in Arn .
	//
	// This member is required.
	Policy *string

	// The Amazon Resource Name (ARN) of the account that you want to share rule
	// groups and firewall policies with.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type PutResourcePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutResourcePolicy"); err != nil {
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
	if err = addOpPutResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourcePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutResourcePolicy",
	}
}

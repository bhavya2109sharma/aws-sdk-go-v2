// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes ActivatedRule objects in a RuleGroup .
//
// You can only insert REGULAR rules into a rule group.
//
// You can have a maximum of ten rules per rule group.
//
// To create and configure a RuleGroup , perform the following steps:
//
//   - Create and update the Rules that you want to include in the RuleGroup . See CreateRule
//     .
//
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of an UpdateRuleGrouprequest.
//
//   - Submit an UpdateRuleGroup request to add Rules to the RuleGroup .
//
//   - Create and update a WebACL that contains the RuleGroup . See CreateWebACL.
//
// If you want to replace one Rule with another, you delete the existing one and
// add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
func (c *Client) UpdateRuleGroup(ctx context.Context, params *UpdateRuleGroupInput, optFns ...func(*Options)) (*UpdateRuleGroupOutput, error) {
	if params == nil {
		params = &UpdateRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRuleGroup", params, optFns, c.addOperationUpdateRuleGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRuleGroupInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// The RuleGroupId of the RuleGroup that you want to update. RuleGroupId is returned by CreateRuleGroup
	// and by ListRuleGroups.
	//
	// This member is required.
	RuleGroupId *string

	// An array of RuleGroupUpdate objects that you want to insert into or delete from
	// a RuleGroup.
	//
	// You can only insert REGULAR rules into a rule group.
	//
	// ActivatedRule|OverrideAction applies only when updating or adding a RuleGroup
	// to a WebACL . In this case you do not use ActivatedRule|Action . For all other
	// update requests, ActivatedRule|Action is used instead of
	// ActivatedRule|OverrideAction .
	//
	// This member is required.
	Updates []types.RuleGroupUpdate

	noSmithyDocumentSerde
}

type UpdateRuleGroupOutput struct {

	// The ChangeToken that you used to submit the UpdateRuleGroup request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRuleGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRuleGroup"); err != nil {
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
	if err = addOpUpdateRuleGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRuleGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRuleGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRuleGroup",
	}
}

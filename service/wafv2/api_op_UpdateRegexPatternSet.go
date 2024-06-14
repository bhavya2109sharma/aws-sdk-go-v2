// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified RegexPatternSet.
//
// This operation completely replaces the mutable specifications that you already
// have for the regex pattern set with the ones that you provide to this call.
//
// To modify a regex pattern set, do the following:
//
//   - Retrieve it by calling GetRegexPatternSet
//
//   - Update its settings as needed
//
//   - Provide the complete regex pattern set specification to this call
//
// # Temporary inconsistencies during updates
//
// When you create or change a web ACL or other WAF resources, the changes take a
// small amount of time to propagate to all areas where the resources are stored.
// The propagation time can be from a few seconds to a number of minutes.
//
// The following are examples of the temporary inconsistencies that you might
// notice during change propagation:
//
//   - After you create a web ACL, if you try to associate it with a resource, you
//     might get an exception indicating that the web ACL is unavailable.
//
//   - After you add a rule group to a web ACL, the new rule group rules might be
//     in effect in one area where the web ACL is used and not in another.
//
//   - After you change a rule action setting, you might see the old action in
//     some places and the new action in others.
//
//   - After you add an IP address to an IP set that is in use in a blocking rule,
//     the new address might be blocked in one area while still allowed in another.
func (c *Client) UpdateRegexPatternSet(ctx context.Context, params *UpdateRegexPatternSetInput, optFns ...func(*Options)) (*UpdateRegexPatternSetOutput, error) {
	if params == nil {
		params = &UpdateRegexPatternSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRegexPatternSet", params, optFns, c.addOperationUpdateRegexPatternSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRegexPatternSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRegexPatternSetInput struct {

	// A unique identifier for the set. This ID is returned in the responses to create
	// and list commands. You provide it to operations like update and delete.
	//
	// This member is required.
	Id *string

	// A token used for optimistic locking. WAF returns a token to your get and list
	// requests, to mark the state of the entity at the time of the request. To make
	// changes to the entity associated with the token, you provide the token to
	// operations like update and delete . WAF uses the token to ensure that no changes
	// have been made to the entity since you last retrieved it. If a change has been
	// made, the update fails with a WAFOptimisticLockException . If this happens,
	// perform another get , and use the new token returned by that operation.
	//
	// This member is required.
	LockToken *string

	// The name of the set. You cannot change the name after you create the set.
	//
	// This member is required.
	Name *string

	//
	//
	// This member is required.
	RegularExpressionList []types.Regex

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito
	// user pool, an App Runner service, or an Amazon Web Services Verified Access
	// instance.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//   - CLI - Specify the Region when you use the CloudFront scope:
	//   --scope=CLOUDFRONT --region=us-east-1 .
	//
	//   - API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// A description of the set that helps with identification.
	Description *string

	noSmithyDocumentSerde
}

type UpdateRegexPatternSetOutput struct {

	// A token used for optimistic locking. WAF returns this token to your update
	// requests. You use NextLockToken in the same manner as you use LockToken .
	NextLockToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRegexPatternSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRegexPatternSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRegexPatternSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRegexPatternSet"); err != nil {
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
	if err = addOpUpdateRegexPatternSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRegexPatternSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRegexPatternSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRegexPatternSet",
	}
}

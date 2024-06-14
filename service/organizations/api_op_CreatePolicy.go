// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a policy of a specified type that you can attach to a root, an
// organizational unit (OU), or an individual Amazon Web Services account.
//
// For more information about policies and their use, see [Managing Organizations policies].
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission.
//
// This operation can be called only from the organization's management account or
// by a member account that is a delegated administrator for an Amazon Web Services
// service.
//
// [Managing Organizations policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html
func (c *Client) CreatePolicy(ctx context.Context, params *CreatePolicyInput, optFns ...func(*Options)) (*CreatePolicyOutput, error) {
	if params == nil {
		params = &CreatePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePolicy", params, optFns, c.addOperationCreatePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePolicyInput struct {

	// The policy text content to add to the new policy. The text that you supply must
	// adhere to the rules of the policy type you specify in the Type parameter.
	//
	// The maximum size of a policy document depends on the policy's type. For more
	// information, see [Maximum and minimum values]in the Organizations User Guide.
	//
	// [Maximum and minimum values]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_limits.html#min-max-values
	//
	// This member is required.
	Content *string

	// An optional description to assign to the policy.
	//
	// This member is required.
	Description *string

	// The friendly name to assign to the policy.
	//
	// The [regex pattern] that is used to validate this parameter is a string of any of the
	// characters in the ASCII character range.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	Name *string

	// The type of policy to create. You can specify one of the following values:
	//
	// [AISERVICES_OPT_OUT_POLICY]
	//
	// [BACKUP_POLICY]
	//
	// [SERVICE_CONTROL_POLICY]
	//
	// [TAG_POLICY]
	//
	// [AISERVICES_OPT_OUT_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html
	// [BACKUP_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html
	// [SERVICE_CONTROL_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html
	// [TAG_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html
	//
	// This member is required.
	Type types.PolicyType

	// A list of tags that you want to attach to the newly created policy. For each
	// tag in the list, you must specify both a tag key and a value. You can set the
	// value to an empty string, but you can't set it to null . For more information
	// about tagging, see [Tagging Organizations resources]in the Organizations User Guide.
	//
	// If any one of the tags is not valid or if you exceed the allowed number of tags
	// for a policy, then the entire request fails and the policy is not created.
	//
	// [Tagging Organizations resources]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreatePolicyOutput struct {

	// A structure that contains details about the newly created policy.
	Policy *types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePolicy"); err != nil {
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
	if err = addOpCreatePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePolicy",
	}
}

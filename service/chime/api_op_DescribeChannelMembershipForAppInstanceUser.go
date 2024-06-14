// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns the details of a channel based on the membership of the specified
//
// AppInstanceUser .
//
// The x-amz-chime-bearer request header is mandatory. Use the AppInstanceUserArn
// of the user that makes the API call as the value in the header.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [DescribeChannelMembershipForAppInstanceUser], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by DescribeChannelMembershipForAppInstanceUser in the
// Amazon Chime SDK Messaging Namespace
//
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
// [DescribeChannelMembershipForAppInstanceUser]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_messaging-chime_DescribeChannelMembershipForAppInstanceUser.html
func (c *Client) DescribeChannelMembershipForAppInstanceUser(ctx context.Context, params *DescribeChannelMembershipForAppInstanceUserInput, optFns ...func(*Options)) (*DescribeChannelMembershipForAppInstanceUserOutput, error) {
	if params == nil {
		params = &DescribeChannelMembershipForAppInstanceUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChannelMembershipForAppInstanceUser", params, optFns, c.addOperationDescribeChannelMembershipForAppInstanceUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChannelMembershipForAppInstanceUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeChannelMembershipForAppInstanceUserInput struct {

	// The ARN of the user in a channel.
	//
	// This member is required.
	AppInstanceUserArn *string

	// The ARN of the channel to which the user belongs.
	//
	// This member is required.
	ChannelArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string

	noSmithyDocumentSerde
}

type DescribeChannelMembershipForAppInstanceUserOutput struct {

	// The channel to which a user belongs.
	ChannelMembership *types.ChannelMembershipForAppInstanceUserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeChannelMembershipForAppInstanceUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChannelMembershipForAppInstanceUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChannelMembershipForAppInstanceUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeChannelMembershipForAppInstanceUser"); err != nil {
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
	if err = addEndpointPrefix_opDescribeChannelMembershipForAppInstanceUserMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeChannelMembershipForAppInstanceUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChannelMembershipForAppInstanceUser(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeChannelMembershipForAppInstanceUserMiddleware struct {
}

func (*endpointPrefix_opDescribeChannelMembershipForAppInstanceUserMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeChannelMembershipForAppInstanceUserMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDescribeChannelMembershipForAppInstanceUserMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDescribeChannelMembershipForAppInstanceUserMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDescribeChannelMembershipForAppInstanceUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeChannelMembershipForAppInstanceUser",
	}
}

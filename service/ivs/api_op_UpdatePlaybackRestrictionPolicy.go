// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a specified playback restriction policy.
func (c *Client) UpdatePlaybackRestrictionPolicy(ctx context.Context, params *UpdatePlaybackRestrictionPolicyInput, optFns ...func(*Options)) (*UpdatePlaybackRestrictionPolicyOutput, error) {
	if params == nil {
		params = &UpdatePlaybackRestrictionPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePlaybackRestrictionPolicy", params, optFns, c.addOperationUpdatePlaybackRestrictionPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePlaybackRestrictionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePlaybackRestrictionPolicyInput struct {

	// ARN of the playback-restriction-policy to be updated.
	//
	// This member is required.
	Arn *string

	// A list of country codes that control geoblocking restriction. Allowed values
	// are the officially assigned [ISO 3166-1 alpha-2]codes. Default: All countries (an empty array).
	//
	// [ISO 3166-1 alpha-2]: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	AllowedCountries []string

	// A list of origin sites that control CORS restriction. Allowed values are the
	// same as valid values of the Origin header defined at [https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin]. Default: All origins (an
	// empty array).
	//
	// [https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
	AllowedOrigins []string

	// Whether channel playback is constrained by origin site. Default: false .
	EnableStrictOriginEnforcement *bool

	// Playback-restriction-policy name. The value does not need to be unique.
	Name *string

	noSmithyDocumentSerde
}

type UpdatePlaybackRestrictionPolicyOutput struct {

	// Object specifying the updated policy.
	PlaybackRestrictionPolicy *types.PlaybackRestrictionPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePlaybackRestrictionPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePlaybackRestrictionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePlaybackRestrictionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePlaybackRestrictionPolicy"); err != nil {
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
	if err = addOpUpdatePlaybackRestrictionPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePlaybackRestrictionPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePlaybackRestrictionPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePlaybackRestrictionPolicy",
	}
}

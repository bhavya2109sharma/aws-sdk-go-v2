// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
	"time"
)

// Gets an Amazon S3 on Outposts bucket. For more information, see [Using Amazon S3 on Outposts] in the Amazon
// S3 User Guide.
//
// If you are using an identity other than the root user of the Amazon Web
// Services account that owns the Outposts bucket, the calling identity must have
// the s3-outposts:GetBucket permissions on the specified Outposts bucket and
// belong to the Outposts bucket owner's account in order to use this action. Only
// users from Outposts bucket owner account with the right permissions can perform
// actions on an Outposts bucket.
//
// If you don't have s3-outposts:GetBucket permissions or you're not using an
// identity that belongs to the bucket owner's account, Amazon S3 returns a 403
// Access Denied error.
//
// The following actions are related to GetBucket for Amazon S3 on Outposts:
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// [PutObject]
//
// [CreateBucket]
//
// [DeleteBucket]
//
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html#API_control_GetBucket_Examples
func (c *Client) GetBucket(ctx context.Context, params *GetBucketInput, optFns ...func(*Options)) (*GetBucketOutput, error) {
	if params == nil {
		params = &GetBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucket", params, optFns, c.addOperationGetBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketInput struct {

	// The Amazon Web Services account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// Specifies the bucket.
	//
	// For using this parameter with Amazon S3 on Outposts with the REST API, you must
	// specify the name and the x-amz-outpost-id as well.
	//
	// For using this parameter with S3 on Outposts with the Amazon Web Services SDK
	// and CLI, you must specify the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/ . For example, to access the bucket
	// reports through Outpost my-outpost owned by account 123456789012 in Region
	// us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	noSmithyDocumentSerde
}

func (in *GetBucketInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.Bucket = in.Bucket
	p.RequiresAccountId = ptr.Bool(true)
}

type GetBucketOutput struct {

	// The Outposts bucket requested.
	Bucket *string

	// The creation date of the Outposts bucket.
	CreationDate *time.Time

	//
	PublicAccessBlockEnabled bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucket{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBucket"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetBucketMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucket(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetBucketUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func (m *GetBucketInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *GetBucketInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opGetBucketMiddleware struct {
}

func (*endpointPrefix_opGetBucketMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetBucketMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*GetBucketInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetBucketMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetBucketMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBucket",
	}
}

func copyGetBucketInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetBucketInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetBucketInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *GetBucketInput) copy() interface{} {
	v := *in
	return &v
}
func getGetBucketARNMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setGetBucketARNMember(input interface{}, v string) error {
	in := input.(*GetBucketInput)
	in.Bucket = &v
	return nil
}
func backFillGetBucketAccountID(input interface{}, v string) error {
	in := input.(*GetBucketInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetBucketUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getGetBucketARNMember,
			BackfillAccountID: backFillGetBucketAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setGetBucketARNMember,
			CopyInput:         copyGetBucketInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

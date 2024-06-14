// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action creates an Amazon S3 on Outposts bucket. To create an S3 bucket,
// see [Create Bucket]in the Amazon S3 API Reference.
//
// Creates a new Outposts bucket. By creating the bucket, you become the bucket
// owner. To create an Outposts bucket, you must have S3 on Outposts. For more
// information, see [Using Amazon S3 on Outposts]in Amazon S3 User Guide.
//
// Not every string is an acceptable bucket name. For information on bucket naming
// restrictions, see [Working with Amazon S3 Buckets].
//
// S3 on Outposts buckets support:
//
//   - Tags
//
//   - LifecycleConfigurations for deleting expired objects
//
// For a complete list of restrictions and Amazon S3 feature limitations on S3 on
// Outposts, see [Amazon S3 on Outposts Restrictions and Limitations].
//
// For an example of the request syntax for Amazon S3 on Outposts that uses the S3
// on Outposts endpoint hostname prefix and x-amz-outpost-id in your API request,
// see the [Examples]section.
//
// The following actions are related to CreateBucket for Amazon S3 on Outposts:
//
// [PutObject]
//
// [GetBucket]
//
// [DeleteBucket]
//
// [CreateAccessPoint]
//
// [PutAccessPointPolicy]
//
// [GetBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html
// [CreateAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html
// [Working with Amazon S3 Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/BucketRestrictions.html#bucketnamingrules
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html
// [Create Bucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [PutAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html#API_control_CreateBucket_Examples
// [Amazon S3 on Outposts Restrictions and Limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OnOutpostsRestrictionsLimitations.html
func (c *Client) CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error) {
	if params == nil {
		params = &CreateBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBucket", params, optFns, c.addOperationCreateBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBucketInput struct {

	// The name of the bucket.
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket.
	//
	// This is not supported by Amazon S3 on Outposts buckets.
	ACL types.BucketCannedACL

	// The configuration information for the bucket.
	//
	// This is not supported by Amazon S3 on Outposts buckets.
	CreateBucketConfiguration *types.CreateBucketConfiguration

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	//
	// This is not supported by Amazon S3 on Outposts buckets.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket.
	//
	// This is not supported by Amazon S3 on Outposts buckets.
	GrantRead *string

	// Allows grantee to read the bucket ACL.
	//
	// This is not supported by Amazon S3 on Outposts buckets.
	GrantReadACP *string

	// Allows grantee to create, overwrite, and delete any object in the bucket.
	//
	// This is not supported by Amazon S3 on Outposts buckets.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket.
	//
	// This is not supported by Amazon S3 on Outposts buckets.
	GrantWriteACP *string

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	//
	// This is not supported by Amazon S3 on Outposts buckets.
	ObjectLockEnabledForBucket bool

	// The ID of the Outposts where the bucket is being created.
	//
	// This ID is required by Amazon S3 on Outposts buckets.
	OutpostId *string

	noSmithyDocumentSerde
}

func (in *CreateBucketInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket
	p.OutpostId = in.OutpostId

}

type CreateBucketOutput struct {

	// The Amazon Resource Name (ARN) of the bucket.
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
	BucketArn *string

	// The location of the bucket.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBucket"); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBucket(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addCreateBucketUpdateEndpoint(stack, options); err != nil {
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
	return nil
}

func (m *CreateBucketInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *CreateBucketInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

func newServiceMetadataMiddleware_opCreateBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBucket",
	}
}

func copyCreateBucketInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateBucketInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateBucketInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *CreateBucketInput) copy() interface{} {
	v := *in
	return &v
}

// getCreateBucketOutpostIDMember returns a pointer to string denoting a provided
// outpost-id member value and a boolean indicating if the input has a modeled
// outpost-id,
func getCreateBucketOutpostIDMember(input interface{}) (*string, bool) {
	in := input.(*CreateBucketInput)
	if in.OutpostId == nil {
		return nil, false
	}
	return in.OutpostId, true
}
func addCreateBucketUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: getCreateBucketOutpostIDMember,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyCreateBucketInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

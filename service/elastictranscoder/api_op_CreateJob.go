// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// When you create a job, Elastic Transcoder returns JSON data that includes the
// values that you specified plus information about the job that is created.
//
// If you have specified more than one output for your jobs (for example, one
// output for the Kindle Fire and another output for the Apple iPhone 4s), you
// currently must use the Elastic Transcoder API to list the jobs (as opposed to
// the AWS Console).
func (c *Client) CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) {
	if params == nil {
		params = &CreateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJob", params, optFns, c.addOperationCreateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The CreateJobRequest structure.
type CreateJobInput struct {

	// The Id of the pipeline that you want Elastic Transcoder to use for transcoding.
	// The pipeline determines several settings, including the Amazon S3 bucket from
	// which Elastic Transcoder gets the files to transcode and the bucket into which
	// Elastic Transcoder puts the transcoded files.
	//
	// This member is required.
	PipelineId *string

	// A section of the request body that provides information about the file that is
	// being transcoded.
	Input *types.JobInput

	// A section of the request body that provides information about the files that
	// are being transcoded.
	Inputs []types.JobInput

	//  A section of the request body that provides information about the transcoded
	// (target) file. We strongly recommend that you use the Outputs syntax instead of
	// the Output syntax.
	Output *types.CreateJobOutput

	// The value, if any, that you want Elastic Transcoder to prepend to the names of
	// all files that this job creates, including output files, thumbnails, and
	// playlists.
	OutputKeyPrefix *string

	//  A section of the request body that provides information about the transcoded
	// (target) files. We recommend that you use the Outputs syntax instead of the
	// Output syntax.
	Outputs []types.CreateJobOutput

	// If you specify a preset in PresetId for which the value of Container is fmp4
	// (Fragmented MP4) or ts (MPEG-TS), Playlists contains information about the
	// master playlists that you want Elastic Transcoder to create.
	//
	// The maximum number of master playlists in a job is 30.
	Playlists []types.CreateJobPlaylist

	// User-defined metadata that you want to associate with an Elastic Transcoder
	// job. You specify metadata in key/value pairs, and you can add up to 10 key/value
	// pairs per job. Elastic Transcoder does not guarantee that key/value pairs are
	// returned in the same order in which you specify them.
	UserMetadata map[string]string

	noSmithyDocumentSerde
}

// The CreateJobResponse structure.
type CreateJobOutput struct {

	// A section of the response body that provides information about the job that is
	// created.
	Job *types.Job

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateJob"); err != nil {
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
	if err = addOpCreateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateJob",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Aggregates on indexed data with search queries pertaining to particular fields.
//
// Requires permission to access the [GetBucketsAggregation] action.
//
// [GetBucketsAggregation]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) GetBucketsAggregation(ctx context.Context, params *GetBucketsAggregationInput, optFns ...func(*Options)) (*GetBucketsAggregationOutput, error) {
	if params == nil {
		params = &GetBucketsAggregationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketsAggregation", params, optFns, c.addOperationGetBucketsAggregationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketsAggregationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketsAggregationInput struct {

	// The aggregation field.
	//
	// This member is required.
	AggregationField *string

	// The basic control of the response shape and the bucket aggregation type to
	// perform.
	//
	// This member is required.
	BucketsAggregationType *types.BucketsAggregationType

	// The search query string.
	//
	// This member is required.
	QueryString *string

	// The name of the index to search.
	IndexName *string

	// The version of the query.
	QueryVersion *string

	noSmithyDocumentSerde
}

type GetBucketsAggregationOutput struct {

	// The main part of the response with a list of buckets. Each bucket contains a
	// keyValue and a count .
	//
	// keyValue : The aggregation field value counted for the particular bucket.
	//
	// count : The number of documents that have that value.
	Buckets []types.Bucket

	// The total number of things that fit the query string criteria.
	TotalCount int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketsAggregationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBucketsAggregation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBucketsAggregation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBucketsAggregation"); err != nil {
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
	if err = addOpGetBucketsAggregationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketsAggregation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBucketsAggregation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBucketsAggregation",
	}
}

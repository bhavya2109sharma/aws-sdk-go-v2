// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon SageMaker Model Card.
//
// For information about how to use model cards, see [Amazon SageMaker Model Card].
//
// [Amazon SageMaker Model Card]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html
func (c *Client) CreateModelCard(ctx context.Context, params *CreateModelCardInput, optFns ...func(*Options)) (*CreateModelCardOutput, error) {
	if params == nil {
		params = &CreateModelCardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModelCard", params, optFns, c.addOperationCreateModelCardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelCardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelCardInput struct {

	// The content of the model card. Content must be in [model card JSON schema] and provided as a string.
	//
	// [model card JSON schema]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema
	//
	// This member is required.
	Content *string

	// The unique name of the model card.
	//
	// This member is required.
	ModelCardName *string

	// The approval status of the model card within your organization. Different
	// organizations might have different criteria for model card review and approval.
	//
	//   - Draft : The model card is a work in progress.
	//
	//   - PendingReview : The model card is pending review.
	//
	//   - Approved : The model card is approved.
	//
	//   - Archived : The model card is archived. No more updates should be made to the
	//   model card, but it can still be exported.
	//
	// This member is required.
	ModelCardStatus types.ModelCardStatus

	// An optional Key Management Service key to encrypt, decrypt, and re-encrypt
	// model card content for regulated workloads with highly sensitive data.
	SecurityConfig *types.ModelCardSecurityConfig

	// Key-value pairs used to manage metadata for model cards.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateModelCardOutput struct {

	// The Amazon Resource Name (ARN) of the successfully created model card.
	//
	// This member is required.
	ModelCardArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelCardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModelCard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModelCard{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateModelCard"); err != nil {
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
	if err = addOpCreateModelCardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModelCard(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateModelCard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateModelCard",
	}
}

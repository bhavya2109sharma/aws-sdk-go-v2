// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// InferICD10CM detects medical conditions as entities listed in a patient record
// and links those entities to normalized concept identifiers in the ICD-10-CM
// knowledge base from the Centers for Disease Control. Amazon Comprehend Medical
// only detects medical entities in English language texts.
func (c *Client) InferICD10CM(ctx context.Context, params *InferICD10CMInput, optFns ...func(*Options)) (*InferICD10CMOutput, error) {
	if params == nil {
		params = &InferICD10CMInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InferICD10CM", params, optFns, c.addOperationInferICD10CMMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InferICD10CMOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InferICD10CMInput struct {

	// The input text used for analysis.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

type InferICD10CMOutput struct {

	// The medical conditions detected in the text linked to ICD-10-CM concepts. If
	// the action is successful, the service sends back an HTTP 200 response, as well
	// as the entities detected.
	//
	// This member is required.
	Entities []types.ICD10CMEntity

	// The version of the model used to analyze the documents, in the format n.n.n You
	// can use this information to track the model used for a particular batch of
	// documents.
	ModelVersion *string

	// If the result of the previous request to InferICD10CM was truncated, include
	// the PaginationToken to fetch the next page of medical condition entities.
	PaginationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInferICD10CMMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInferICD10CM{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInferICD10CM{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InferICD10CM"); err != nil {
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
	if err = addOpInferICD10CMValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInferICD10CM(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInferICD10CM(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InferICD10CM",
	}
}

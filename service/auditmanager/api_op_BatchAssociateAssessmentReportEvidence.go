// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Associates a list of evidence to an assessment report in an Audit Manager
//
// assessment.
func (c *Client) BatchAssociateAssessmentReportEvidence(ctx context.Context, params *BatchAssociateAssessmentReportEvidenceInput, optFns ...func(*Options)) (*BatchAssociateAssessmentReportEvidenceOutput, error) {
	if params == nil {
		params = &BatchAssociateAssessmentReportEvidenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchAssociateAssessmentReportEvidence", params, optFns, c.addOperationBatchAssociateAssessmentReportEvidenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchAssociateAssessmentReportEvidenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateAssessmentReportEvidenceInput struct {

	//  The identifier for the assessment.
	//
	// This member is required.
	AssessmentId *string

	//  The identifier for the folder that the evidence is stored in.
	//
	// This member is required.
	EvidenceFolderId *string

	//  The list of evidence identifiers.
	//
	// This member is required.
	EvidenceIds []string

	noSmithyDocumentSerde
}

type BatchAssociateAssessmentReportEvidenceOutput struct {

	//  A list of errors that the BatchAssociateAssessmentReportEvidence API returned.
	Errors []types.AssessmentReportEvidenceError

	//  The list of evidence identifiers.
	EvidenceIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchAssociateAssessmentReportEvidenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchAssociateAssessmentReportEvidence{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchAssociateAssessmentReportEvidence{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchAssociateAssessmentReportEvidence"); err != nil {
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
	if err = addOpBatchAssociateAssessmentReportEvidenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateAssessmentReportEvidence(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchAssociateAssessmentReportEvidence(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchAssociateAssessmentReportEvidence",
	}
}

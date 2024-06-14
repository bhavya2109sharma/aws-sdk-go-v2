// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an assessment template for the assessment target that is specified by
// the ARN of the assessment target. If the [service-linked role]isn’t already registered, this action
// also creates and registers a service-linked role to grant Amazon Inspector
// access to AWS Services needed to perform security assessments.
//
// [service-linked role]: https://docs.aws.amazon.com/inspector/latest/userguide/inspector_slr.html
func (c *Client) CreateAssessmentTemplate(ctx context.Context, params *CreateAssessmentTemplateInput, optFns ...func(*Options)) (*CreateAssessmentTemplateOutput, error) {
	if params == nil {
		params = &CreateAssessmentTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssessmentTemplate", params, optFns, c.addOperationCreateAssessmentTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssessmentTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssessmentTemplateInput struct {

	// The ARN that specifies the assessment target for which you want to create the
	// assessment template.
	//
	// This member is required.
	AssessmentTargetArn *string

	// The user-defined name that identifies the assessment template that you want to
	// create. You can create several assessment templates for an assessment target.
	// The names of the assessment templates that correspond to a particular assessment
	// target must be unique.
	//
	// This member is required.
	AssessmentTemplateName *string

	// The duration of the assessment run in seconds.
	//
	// This member is required.
	DurationInSeconds *int32

	// The ARNs that specify the rules packages that you want to attach to the
	// assessment template.
	//
	// This member is required.
	RulesPackageArns []string

	// The user-defined attributes that are assigned to every finding that is
	// generated by the assessment run that uses this assessment template. An attribute
	// is a key and value pair (an Attributeobject). Within an assessment template, each key
	// must be unique.
	UserAttributesForFindings []types.Attribute

	noSmithyDocumentSerde
}

type CreateAssessmentTemplateOutput struct {

	// The ARN that specifies the assessment template that is created.
	//
	// This member is required.
	AssessmentTemplateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssessmentTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAssessmentTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAssessmentTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAssessmentTemplate"); err != nil {
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
	if err = addOpCreateAssessmentTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssessmentTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAssessmentTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAssessmentTemplate",
	}
}

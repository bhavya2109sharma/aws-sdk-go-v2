// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about an audience model
func (c *Client) GetAudienceModel(ctx context.Context, params *GetAudienceModelInput, optFns ...func(*Options)) (*GetAudienceModelOutput, error) {
	if params == nil {
		params = &GetAudienceModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAudienceModel", params, optFns, c.addOperationGetAudienceModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAudienceModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAudienceModelInput struct {

	// The Amazon Resource Name (ARN) of the audience model that you are interested in.
	//
	// This member is required.
	AudienceModelArn *string

	noSmithyDocumentSerde
}

type GetAudienceModelOutput struct {

	// The Amazon Resource Name (ARN) of the audience model.
	//
	// This member is required.
	AudienceModelArn *string

	// The time at which the audience model was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the audience model.
	//
	// This member is required.
	Name *string

	// The status of the audience model.
	//
	// This member is required.
	Status types.AudienceModelStatus

	// The Amazon Resource Name (ARN) of the training dataset that was used for this
	// audience model.
	//
	// This member is required.
	TrainingDatasetArn *string

	// The most recent time at which the audience model was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the audience model.
	Description *string

	// The KMS key ARN used for the audience model.
	KmsKeyArn *string

	// Details about the status of the audience model.
	StatusDetails *types.StatusDetails

	// The tags that are assigned to the audience model.
	Tags map[string]string

	// The end date specified for the training window.
	TrainingDataEndTime *time.Time

	// The start date specified for the training window.
	TrainingDataStartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAudienceModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAudienceModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAudienceModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAudienceModel"); err != nil {
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
	if err = addOpGetAudienceModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAudienceModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAudienceModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAudienceModel",
	}
}

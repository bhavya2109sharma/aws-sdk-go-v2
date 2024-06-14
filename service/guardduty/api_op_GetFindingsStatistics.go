// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Amazon GuardDuty findings statistics for the specified detector ID.
//
// There might be regional differences because some flags might not be available
// in all the Regions where GuardDuty is currently supported. For more information,
// see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func (c *Client) GetFindingsStatistics(ctx context.Context, params *GetFindingsStatisticsInput, optFns ...func(*Options)) (*GetFindingsStatisticsOutput, error) {
	if params == nil {
		params = &GetFindingsStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFindingsStatistics", params, optFns, c.addOperationGetFindingsStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFindingsStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingsStatisticsInput struct {

	// The ID of the detector that specifies the GuardDuty service whose findings'
	// statistics you want to retrieve.
	//
	// This member is required.
	DetectorId *string

	// The types of finding statistics to retrieve.
	//
	// This member is required.
	FindingStatisticTypes []types.FindingStatisticType

	// Represents the criteria that is used for querying findings.
	FindingCriteria *types.FindingCriteria

	noSmithyDocumentSerde
}

type GetFindingsStatisticsOutput struct {

	// The finding statistics object.
	//
	// This member is required.
	FindingStatistics *types.FindingStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFindingsStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindingsStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindingsStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFindingsStatistics"); err != nil {
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
	if err = addOpGetFindingsStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindingsStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFindingsStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFindingsStatistics",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves cost and usage metrics for your account. You can specify which cost
// and usage-related metric that you want the request to return. For example, you
// can specify BlendedCosts or UsageQuantity . You can also filter and group your
// data by various dimensions, such as SERVICE or AZ , in a specific time range.
// For a complete list of valid dimensions, see the [GetDimensionValues]operation. Management account
// in an organization in Organizations have access to all member accounts.
//
// For information about filter limitations, see [Quotas and restrictions] in the Billing and Cost
// Management User Guide.
//
// [GetDimensionValues]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetDimensionValues.html
// [Quotas and restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-limits.html
func (c *Client) GetCostAndUsage(ctx context.Context, params *GetCostAndUsageInput, optFns ...func(*Options)) (*GetCostAndUsageOutput, error) {
	if params == nil {
		params = &GetCostAndUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCostAndUsage", params, optFns, c.addOperationGetCostAndUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCostAndUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostAndUsageInput struct {

	// Sets the Amazon Web Services cost granularity to MONTHLY or DAILY , or HOURLY .
	// If Granularity isn't set, the response object doesn't include the Granularity ,
	// either MONTHLY or DAILY , or HOURLY .
	//
	// This member is required.
	Granularity types.Granularity

	// Which metrics are returned in the query. For more information about blended and
	// unblended rates, see [Why does the "blended" annotation appear on some line items in my bill?].
	//
	// Valid values are AmortizedCost , BlendedCost , NetAmortizedCost ,
	// NetUnblendedCost , NormalizedUsageAmount , UnblendedCost , and UsageQuantity .
	//
	// If you return the UsageQuantity metric, the service aggregates all usage
	// numbers without taking into account the units. For example, if you aggregate
	// usageQuantity across all of Amazon EC2, the results aren't meaningful because
	// Amazon EC2 compute hours and data transfer are measured in different units (for
	// example, hours and GB). To get more meaningful UsageQuantity metrics, filter by
	// UsageType or UsageTypeGroups .
	//
	// Metrics is required for GetCostAndUsage requests.
	//
	// [Why does the "blended" annotation appear on some line items in my bill?]: http://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/
	//
	// This member is required.
	Metrics []string

	// Sets the start date and end date for retrieving Amazon Web Services costs. The
	// start date is inclusive, but the end date is exclusive. For example, if start
	// is 2017-01-01 and end is 2017-05-01 , then the cost and usage data is retrieved
	// from 2017-01-01 up to and including 2017-04-30 but not including 2017-05-01 .
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// Filters Amazon Web Services costs by different dimensions. For example, you can
	// specify SERVICE and LINKED_ACCOUNT and get the costs that are associated with
	// that account's usage of that service. You can nest Expression objects to define
	// any combination of dimension filters. For more information, see [Expression].
	//
	// Valid values for MatchOptions for Dimensions are EQUALS and CASE_SENSITIVE .
	//
	// Valid values for MatchOptions for CostCategories and Tags are EQUALS , ABSENT ,
	// and CASE_SENSITIVE . Default values are EQUALS and CASE_SENSITIVE .
	//
	// [Expression]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html
	Filter *types.Expression

	// You can group Amazon Web Services costs using up to two different groups,
	// either dimensions, tag keys, cost categories, or any two group by types.
	//
	// Valid values for the DIMENSION type are AZ , INSTANCE_TYPE , LEGAL_ENTITY_NAME ,
	// INVOICING_ENTITY , LINKED_ACCOUNT , OPERATION , PLATFORM , PURCHASE_TYPE ,
	// SERVICE , TENANCY , RECORD_TYPE , and USAGE_TYPE .
	//
	// When you group by the TAG type and include a valid tag key, you get all tag
	// values, including empty strings.
	GroupBy []types.GroupDefinition

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	noSmithyDocumentSerde
}

type GetCostAndUsageOutput struct {

	// The attributes that apply to a specific dimension value. For example, if the
	// value is a linked account, the attribute is that account name.
	DimensionValueAttributes []types.DimensionValuesWithAttributes

	// The groups that are specified by the Filter or GroupBy parameters in the
	// request.
	GroupDefinitions []types.GroupDefinition

	// The token for the next set of retrievable results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextPageToken *string

	// The time period that's covered by the results in the response.
	ResultsByTime []types.ResultByTime

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCostAndUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCostAndUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCostAndUsage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCostAndUsage"); err != nil {
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
	if err = addOpGetCostAndUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostAndUsage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCostAndUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCostAndUsage",
	}
}

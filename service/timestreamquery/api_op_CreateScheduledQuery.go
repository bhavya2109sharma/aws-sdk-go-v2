// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Create a scheduled query that will be run on your behalf at the configured
//
// schedule. Timestream assumes the execution role provided as part of the
// ScheduledQueryExecutionRoleArn parameter to run the query. You can use the
// NotificationConfiguration parameter to configure notification for your scheduled
// query operations.
func (c *Client) CreateScheduledQuery(ctx context.Context, params *CreateScheduledQueryInput, optFns ...func(*Options)) (*CreateScheduledQueryOutput, error) {
	if params == nil {
		params = &CreateScheduledQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateScheduledQuery", params, optFns, c.addOperationCreateScheduledQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateScheduledQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScheduledQueryInput struct {

	// Configuration for error reporting. Error reports will be generated when a
	// problem is encountered when writing the query results.
	//
	// This member is required.
	ErrorReportConfiguration *types.ErrorReportConfiguration

	// Name of the scheduled query.
	//
	// This member is required.
	Name *string

	// Notification configuration for the scheduled query. A notification is sent by
	// Timestream when a query run finishes, when the state is updated or when you
	// delete it.
	//
	// This member is required.
	NotificationConfiguration *types.NotificationConfiguration

	// The query string to run. Parameter names can be specified in the query string @
	// character followed by an identifier. The named Parameter @scheduled_runtime is
	// reserved and can be used in the query to get the time at which the query is
	// scheduled to run.
	//
	// The timestamp calculated according to the ScheduleConfiguration parameter, will
	// be the value of @scheduled_runtime paramater for each query run. For example,
	// consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For
	// this instance, the @scheduled_runtime parameter is initialized to the timestamp
	// 2021-12-01 00:00:00 when invoking the query.
	//
	// This member is required.
	QueryString *string

	// The schedule configuration for the query.
	//
	// This member is required.
	ScheduleConfiguration *types.ScheduleConfiguration

	// The ARN for the IAM role that Timestream will assume when running the scheduled
	// query.
	//
	// This member is required.
	ScheduledQueryExecutionRoleArn *string

	// Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other
	// words, making the same request repeatedly will produce the same result. Making
	// multiple identical CreateScheduledQuery requests has the same effect as making a
	// single request.
	//
	//   - If CreateScheduledQuery is called without a ClientToken , the Query SDK
	//   generates a ClientToken on your behalf.
	//
	//   - After 8 hours, any request with the same ClientToken is treated as a new
	//   request.
	ClientToken *string

	// The Amazon KMS key used to encrypt the scheduled query resource, at-rest. If
	// the Amazon KMS key is not specified, the scheduled query resource will be
	// encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the
	// key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the
	// name with alias/
	//
	// If ErrorReportConfiguration uses SSE_KMS as encryption type, the same KmsKeyId
	// is used to encrypt the error report at rest.
	KmsKeyId *string

	// A list of key-value pairs to label the scheduled query.
	Tags []types.Tag

	// Configuration used for writing the result of a query.
	TargetConfiguration *types.TargetConfiguration

	noSmithyDocumentSerde
}

type CreateScheduledQueryOutput struct {

	// ARN for the created scheduled query.
	//
	// This member is required.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateScheduledQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateScheduledQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateScheduledQuery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateScheduledQuery"); err != nil {
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
	if err = addOpCreateScheduledQueryDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateScheduledQueryMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateScheduledQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateScheduledQuery(options.Region), middleware.Before); err != nil {
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

func addOpCreateScheduledQueryDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
				opt.EndpointResolverUsedForDiscovery = o.EndpointDiscovery.EndpointResolverUsedForDiscovery
			},
		},
		DiscoverOperation:            c.fetchOpCreateScheduledQueryDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    true,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpCreateScheduledQueryDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*CreateScheduledQueryInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)
	identifierMap["sdk#Region"] = region

	key := fmt.Sprintf("Timestream Query.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	endpoint, err := c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, region, key, opt)
	if err != nil {
		return internalEndpointDiscovery.WeightedAddress{}, err
	}

	weighted, ok := endpoint.GetValidAddress()
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("no valid endpoint address returned by the endpoint discovery api")
	}
	return weighted, nil
}

type idempotencyToken_initializeOpCreateScheduledQuery struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateScheduledQuery) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateScheduledQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateScheduledQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateScheduledQueryInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateScheduledQueryMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateScheduledQuery{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateScheduledQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateScheduledQuery",
	}
}

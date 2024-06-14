// Code generated by smithy-go-codegen DO NOT EDIT.

package mailmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mailmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the details and current status of a specific email archive export job.
func (c *Client) GetArchiveExport(ctx context.Context, params *GetArchiveExportInput, optFns ...func(*Options)) (*GetArchiveExportOutput, error) {
	if params == nil {
		params = &GetArchiveExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetArchiveExport", params, optFns, c.addOperationGetArchiveExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetArchiveExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to retrieve details of a specific archive export job.
type GetArchiveExportInput struct {

	// The identifier of the export job to get details for.
	//
	// This member is required.
	ExportId *string

	noSmithyDocumentSerde
}

// The response containing details of the specified archive export job.
type GetArchiveExportOutput struct {

	// The identifier of the archive the email export was performed from.
	ArchiveId *string

	// Where the exported emails are being delivered.
	ExportDestinationConfiguration types.ExportDestinationConfiguration

	// The criteria used to filter emails included in the export.
	Filters *types.ArchiveFilters

	// The start of the timestamp range the exported emails cover.
	FromTimestamp *time.Time

	// The maximum number of email items included in the export.
	MaxResults *int32

	// The current status of the export job.
	Status *types.ExportStatus

	// The end of the date range the exported emails cover.
	ToTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetArchiveExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetArchiveExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetArchiveExport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetArchiveExport"); err != nil {
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
	if err = addOpGetArchiveExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetArchiveExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetArchiveExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetArchiveExport",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get information about a Changeset.
//
// Deprecated: This method will be discontinued.
func (c *Client) GetChangeset(ctx context.Context, params *GetChangesetInput, optFns ...func(*Options)) (*GetChangesetOutput, error) {
	if params == nil {
		params = &GetChangesetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChangeset", params, optFns, c.addOperationGetChangesetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChangesetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe a changeset.
type GetChangesetInput struct {

	// The unique identifier of the Changeset for which to get data.
	//
	// This member is required.
	ChangesetId *string

	// The unique identifier for the FinSpace Dataset where the Changeset is created.
	//
	// This member is required.
	DatasetId *string

	noSmithyDocumentSerde
}

// The response from a describe changeset operation
type GetChangesetOutput struct {

	// Beginning time from which the Changeset is active. The value is determined as
	// epoch time in milliseconds. For example, the value for Monday, November 1, 2021
	// 12:00:00 PM UTC is specified as 1635768000000.
	ActiveFromTimestamp *int64

	// Time until which the Changeset is active. The value is determined as epoch time
	// in milliseconds. For example, the value for Monday, November 1, 2021 12:00:00 PM
	// UTC is specified as 1635768000000.
	ActiveUntilTimestamp *int64

	// Type that indicates how a Changeset is applied to a Dataset.
	//
	//   - REPLACE – Changeset is considered as a replacement to all prior loaded
	//   Changesets.
	//
	//   - APPEND – Changeset is considered as an addition to the end of all prior
	//   loaded Changesets.
	//
	//   - MODIFY – Changeset is considered as a replacement to a specific prior
	//   ingested Changeset.
	ChangeType types.ChangeType

	// The ARN identifier of the Changeset.
	ChangesetArn *string

	// The unique identifier for a Changeset.
	ChangesetId *string

	// The timestamp at which the Changeset was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreateTime int64

	// The unique identifier for the FinSpace Dataset where the Changeset is created.
	DatasetId *string

	// The structure with error messages.
	ErrorInfo *types.ChangesetErrorInfo

	// Structure of the source file(s).
	FormatParams map[string]string

	// Options that define the location of the data being ingested.
	SourceParams map[string]string

	// The status of Changeset creation operation.
	Status types.IngestionStatus

	// The unique identifier of the updated Changeset.
	UpdatedByChangesetId *string

	// The unique identifier of the Changeset that is being updated.
	UpdatesChangesetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetChangesetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetChangeset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetChangeset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetChangeset"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpGetChangesetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetChangeset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetChangeset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetChangeset",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays the notebook files for the specified workgroup in paginated format.
func (c *Client) ListNotebookMetadata(ctx context.Context, params *ListNotebookMetadataInput, optFns ...func(*Options)) (*ListNotebookMetadataOutput, error) {
	if params == nil {
		params = &ListNotebookMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotebookMetadata", params, optFns, c.addOperationListNotebookMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotebookMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotebookMetadataInput struct {

	// The name of the Spark enabled workgroup to retrieve notebook metadata for.
	//
	// This member is required.
	WorkGroup *string

	// Search filter string.
	Filters *types.FilterDefinition

	// Specifies the maximum number of results to return.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated.
	NextToken *string

	noSmithyDocumentSerde
}

type ListNotebookMetadataOutput struct {

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// The list of notebook metadata for the specified workgroup.
	NotebookMetadataList []types.NotebookMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNotebookMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListNotebookMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNotebookMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNotebookMetadata"); err != nil {
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
	if err = addOpListNotebookMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotebookMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListNotebookMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNotebookMetadata",
	}
}

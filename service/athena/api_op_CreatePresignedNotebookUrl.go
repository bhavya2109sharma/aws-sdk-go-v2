// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets an authentication token and the URL at which the notebook can be accessed.
// During programmatic access, CreatePresignedNotebookUrl must be called every 10
// minutes to refresh the authentication token. For information about granting
// programmatic access, see [Grant programmatic access].
//
// [Grant programmatic access]: https://docs.aws.amazon.com/athena/latest/ug/setting-up.html#setting-up-grant-programmatic-access
func (c *Client) CreatePresignedNotebookUrl(ctx context.Context, params *CreatePresignedNotebookUrlInput, optFns ...func(*Options)) (*CreatePresignedNotebookUrlOutput, error) {
	if params == nil {
		params = &CreatePresignedNotebookUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePresignedNotebookUrl", params, optFns, c.addOperationCreatePresignedNotebookUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePresignedNotebookUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePresignedNotebookUrlInput struct {

	// The session ID.
	//
	// This member is required.
	SessionId *string

	noSmithyDocumentSerde
}

type CreatePresignedNotebookUrlOutput struct {

	// The authentication token for the notebook.
	//
	// This member is required.
	AuthToken *string

	// The UTC epoch time when the authentication token expires.
	//
	// This member is required.
	AuthTokenExpirationTime *int64

	// The URL of the notebook. The URL includes the authentication token and notebook
	// file name and points directly to the opened notebook.
	//
	// This member is required.
	NotebookUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePresignedNotebookUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePresignedNotebookUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePresignedNotebookUrl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePresignedNotebookUrl"); err != nil {
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
	if err = addOpCreatePresignedNotebookUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePresignedNotebookUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePresignedNotebookUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePresignedNotebookUrl",
	}
}

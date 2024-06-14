// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Amazon Q in Connect quick response.
func (c *Client) UpdateQuickResponse(ctx context.Context, params *UpdateQuickResponseInput, optFns ...func(*Options)) (*UpdateQuickResponseOutput, error) {
	if params == nil {
		params = &UpdateQuickResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateQuickResponse", params, optFns, c.addOperationUpdateQuickResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateQuickResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateQuickResponseInput struct {

	// The identifier of the knowledge base. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The identifier of the quick response.
	//
	// This member is required.
	QuickResponseId *string

	// The Amazon Connect contact channels this quick response applies to. The
	// supported contact channel types include Chat .
	Channels []string

	// The updated content of the quick response.
	Content types.QuickResponseDataProvider

	// The media type of the quick response content.
	//
	//   - Use application/x.quickresponse;format=plain for quick response written in
	//   plain text.
	//
	//   - Use application/x.quickresponse;format=markdown for quick response written
	//   in richtext.
	ContentType *string

	// The updated description of the quick response.
	Description *string

	// The updated grouping configuration of the quick response.
	GroupingConfiguration *types.GroupingConfiguration

	// Whether the quick response is active.
	IsActive *bool

	// The language code value for the language in which the quick response is
	// written. The supported language codes include de_DE , en_US , es_ES , fr_FR ,
	// id_ID , it_IT , ja_JP , ko_KR , pt_BR , zh_CN , zh_TW
	Language *string

	// The name of the quick response.
	Name *string

	// Whether to remove the description from the quick response.
	RemoveDescription *bool

	// Whether to remove the grouping configuration of the quick response.
	RemoveGroupingConfiguration *bool

	// Whether to remove the shortcut key of the quick response.
	RemoveShortcutKey *bool

	// The shortcut key of the quick response. The value should be unique across the
	// knowledge base.
	ShortcutKey *string

	noSmithyDocumentSerde
}

type UpdateQuickResponseOutput struct {

	// The quick response.
	QuickResponse *types.QuickResponseData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateQuickResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateQuickResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateQuickResponse{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateQuickResponse"); err != nil {
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
	if err = addOpUpdateQuickResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateQuickResponse(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateQuickResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateQuickResponse",
	}
}

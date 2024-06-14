// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Translates input text from the source language to the target language. For a
// list of available languages and language codes, see [Supported languages].
//
// [Supported languages]: https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html
func (c *Client) TranslateText(ctx context.Context, params *TranslateTextInput, optFns ...func(*Options)) (*TranslateTextOutput, error) {
	if params == nil {
		params = &TranslateTextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TranslateText", params, optFns, c.addOperationTranslateTextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TranslateTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TranslateTextInput struct {

	// The language code for the language of the source text. For a list of language
	// codes, see [Supported languages].
	//
	// To have Amazon Translate determine the source language of your text, you can
	// specify auto in the SourceLanguageCode field. If you specify auto , Amazon
	// Translate will call [Amazon Comprehend]to determine the source language.
	//
	// If you specify auto , you must send the TranslateText request in a region that
	// supports Amazon Comprehend. Otherwise, the request returns an error indicating
	// that autodetect is not supported.
	//
	// [Supported languages]: https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html
	// [Amazon Comprehend]: https://docs.aws.amazon.com/comprehend/latest/dg/comprehend-general.html
	//
	// This member is required.
	SourceLanguageCode *string

	// The language code requested for the language of the target text. For a list of
	// language codes, see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html
	//
	// This member is required.
	TargetLanguageCode *string

	// The text to translate. The text string can be a maximum of 10,000 bytes long.
	// Depending on your character set, this may be fewer than 10,000 characters.
	//
	// This member is required.
	Text *string

	// Settings to configure your translation output. You can configure the following
	// options:
	//
	//   - Brevity: reduces the length of the translated output for most translations.
	//
	//   - Formality: sets the formality level of the output text.
	//
	//   - Profanity: masks profane words and phrases in your translation output.
	Settings *types.TranslationSettings

	// The name of a terminology list file to add to the translation job. This file
	// provides source terms and the desired translation for each term. A terminology
	// list can contain a maximum of 256 terms. You can use one custom terminology
	// resource in your translation request.
	//
	// Use the ListTerminologies operation to get the available terminology lists.
	//
	// For more information about custom terminology lists, see [Custom terminology].
	//
	// [Custom terminology]: https://docs.aws.amazon.com/translate/latest/dg/how-custom-terminology.html
	TerminologyNames []string

	noSmithyDocumentSerde
}

type TranslateTextOutput struct {

	// The language code for the language of the source text.
	//
	// This member is required.
	SourceLanguageCode *string

	// The language code for the language of the target text.
	//
	// This member is required.
	TargetLanguageCode *string

	// The translated text.
	//
	// This member is required.
	TranslatedText *string

	// Optional settings that modify the translation output.
	AppliedSettings *types.TranslationSettings

	// The names of the custom terminologies applied to the input text by Amazon
	// Translate for the translated text response.
	AppliedTerminologies []types.AppliedTerminology

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTranslateTextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTranslateText{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTranslateText{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TranslateText"); err != nil {
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
	if err = addOpTranslateTextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTranslateText(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTranslateText(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TranslateText",
	}
}

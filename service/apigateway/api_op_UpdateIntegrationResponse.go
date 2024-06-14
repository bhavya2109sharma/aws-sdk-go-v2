// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Represents an update integration response.
func (c *Client) UpdateIntegrationResponse(ctx context.Context, params *UpdateIntegrationResponseInput, optFns ...func(*Options)) (*UpdateIntegrationResponseOutput, error) {
	if params == nil {
		params = &UpdateIntegrationResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIntegrationResponse", params, optFns, c.addOperationUpdateIntegrationResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIntegrationResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents an update integration response request.
type UpdateIntegrationResponseInput struct {

	// Specifies an update integration response request's HTTP method.
	//
	// This member is required.
	HttpMethod *string

	// Specifies an update integration response request's resource identifier.
	//
	// This member is required.
	ResourceId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// Specifies an update integration response request's status code.
	//
	// This member is required.
	StatusCode *string

	// For more information about supported patch operations, see [Patch Operations].
	//
	// [Patch Operations]: https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html
	PatchOperations []types.PatchOperation

	noSmithyDocumentSerde
}

// Represents an integration response. The status code must map to an existing
// MethodResponse, and parameters and templates can be used to transform the
// back-end response.
type UpdateIntegrationResponseOutput struct {

	// Specifies how to handle response payload content type conversions. Supported
	// values are CONVERT_TO_BINARY and CONVERT_TO_TEXT , with the following behaviors:
	//
	// If this property is not defined, the response payload will be passed through
	// from the integration response to the method response without modification.
	ContentHandling types.ContentHandlingStrategy

	// A key-value map specifying response parameters that are passed to the method
	// response from the back end. The key is a method response header parameter name
	// and the mapped value is an integration response header value, a static value
	// enclosed within a pair of single quotes, or a JSON expression from the
	// integration response body. The mapping key must match the pattern of
	// method.response.header.{name} , where name is a valid and unique header name.
	// The mapped non-static value must match the pattern of
	// integration.response.header.{name} or
	// integration.response.body.{JSON-expression} , where name is a valid and unique
	// response header name and JSON-expression is a valid JSON expression without the
	// $ prefix.
	ResponseParameters map[string]string

	// Specifies the templates used to transform the integration response body.
	// Response templates are represented as a key/value map, with a content-type as
	// the key and a template as the value.
	ResponseTemplates map[string]string

	// Specifies the regular expression (regex) pattern used to choose an integration
	// response based on the response from the back end. For example, if the success
	// response returns nothing and the error response returns some string, you could
	// use the .+ regex to match error response. However, make sure that the error
	// response does not contain any newline ( \n ) character in such cases. If the
	// back end is an Lambda function, the Lambda function error header is matched. For
	// all other HTTP and Amazon Web Services back ends, the HTTP status code is
	// matched.
	SelectionPattern *string

	// Specifies the status code that is used to map the integration response to an
	// existing MethodResponse.
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIntegrationResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIntegrationResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIntegrationResponse{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateIntegrationResponse"); err != nil {
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
	if err = addOpUpdateIntegrationResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIntegrationResponse(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIntegrationResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateIntegrationResponse",
	}
}

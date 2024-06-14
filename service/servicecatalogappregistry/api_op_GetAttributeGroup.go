// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

//	Retrieves an attribute group by its ARN, ID, or name. The attribute group can
//
// be specified by its ARN, ID, or name.
func (c *Client) GetAttributeGroup(ctx context.Context, params *GetAttributeGroupInput, optFns ...func(*Options)) (*GetAttributeGroupOutput, error) {
	if params == nil {
		params = &GetAttributeGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAttributeGroup", params, optFns, c.addOperationGetAttributeGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAttributeGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAttributeGroupInput struct {

	//  The name, ID, or ARN of the attribute group that holds the attributes to
	// describe the application.
	//
	// This member is required.
	AttributeGroup *string

	noSmithyDocumentSerde
}

type GetAttributeGroupOutput struct {

	// The Amazon resource name (ARN) that specifies the attribute group across
	// services.
	Arn *string

	// A JSON string in the form of nested key-value pairs that represent the
	// attributes in the group and describes an application and its components.
	Attributes *string

	// The service principal that created the attribute group.
	CreatedBy *string

	// The ISO-8601 formatted timestamp of the moment the attribute group was created.
	CreationTime *time.Time

	// The description of the attribute group that the user provides.
	Description *string

	// The identifier of the attribute group.
	Id *string

	// The ISO-8601 formatted timestamp of the moment the attribute group was last
	// updated. This time is the same as the creationTime for a newly created attribute
	// group.
	LastUpdateTime *time.Time

	// The name of the attribute group.
	Name *string

	// Key-value pairs associated with the attribute group.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAttributeGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAttributeGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAttributeGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAttributeGroup"); err != nil {
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
	if err = addOpGetAttributeGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAttributeGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAttributeGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAttributeGroup",
	}
}

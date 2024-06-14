// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the specified association between a service network
// and a service.
func (c *Client) GetServiceNetworkServiceAssociation(ctx context.Context, params *GetServiceNetworkServiceAssociationInput, optFns ...func(*Options)) (*GetServiceNetworkServiceAssociationOutput, error) {
	if params == nil {
		params = &GetServiceNetworkServiceAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceNetworkServiceAssociation", params, optFns, c.addOperationGetServiceNetworkServiceAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceNetworkServiceAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceNetworkServiceAssociationInput struct {

	// The ID or Amazon Resource Name (ARN) of the association.
	//
	// This member is required.
	ServiceNetworkServiceAssociationIdentifier *string

	noSmithyDocumentSerde
}

type GetServiceNetworkServiceAssociationOutput struct {

	// The Amazon Resource Name (ARN) of the association.
	Arn *string

	// The date and time that the association was created, specified in ISO-8601
	// format.
	CreatedAt *time.Time

	// The account that created the association.
	CreatedBy *string

	// The custom domain name of the service.
	CustomDomainName *string

	// The DNS name of the service.
	DnsEntry *types.DnsEntry

	// The failure code.
	FailureCode *string

	// The failure message.
	FailureMessage *string

	// The ID of the service network and service association.
	Id *string

	// The Amazon Resource Name (ARN) of the service.
	ServiceArn *string

	// The ID of the service.
	ServiceId *string

	// The name of the service.
	ServiceName *string

	// The Amazon Resource Name (ARN) of the service network.
	ServiceNetworkArn *string

	// The ID of the service network.
	ServiceNetworkId *string

	// The name of the service network.
	ServiceNetworkName *string

	// The status of the association.
	Status types.ServiceNetworkServiceAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceNetworkServiceAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetServiceNetworkServiceAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetServiceNetworkServiceAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetServiceNetworkServiceAssociation"); err != nil {
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
	if err = addOpGetServiceNetworkServiceAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceNetworkServiceAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceNetworkServiceAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetServiceNetworkServiceAssociation",
	}
}

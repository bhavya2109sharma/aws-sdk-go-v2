// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Amazon Kendra index.
func (c *Client) UpdateIndex(ctx context.Context, params *UpdateIndexInput, optFns ...func(*Options)) (*UpdateIndexOutput, error) {
	if params == nil {
		params = &UpdateIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIndex", params, optFns, c.addOperationUpdateIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIndexInput struct {

	// The identifier of the index you want to update.
	//
	// This member is required.
	Id *string

	// Sets the number of additional document storage and query capacity units that
	// should be used by the index. You can change the capacity of the index up to 5
	// times per day, or make 5 API calls.
	//
	// If you are using extra storage units, you can't reduce the storage capacity
	// below what is required to meet the storage needs for your index.
	CapacityUnits *types.CapacityUnitsConfiguration

	// A new description for the index.
	Description *string

	// The document metadata configuration you want to update for the index. Document
	// metadata are fields or attributes associated with your documents. For example,
	// the company department name associated with each document.
	DocumentMetadataConfigurationUpdates []types.DocumentMetadataConfiguration

	// A new name for the index.
	Name *string

	// An Identity and Access Management (IAM) role that gives Amazon Kendra
	// permission to access Amazon CloudWatch logs and metrics.
	RoleArn *string

	// The user context policy.
	UserContextPolicy types.UserContextPolicy

	// Gets users and groups from IAM Identity Center identity source. To configure
	// this, see [UserGroupResolutionConfiguration]. This is useful for user context filtering, where search results are
	// filtered based on the user or their group access to documents.
	//
	// [UserGroupResolutionConfiguration]: https://docs.aws.amazon.com/kendra/latest/dg/API_UserGroupResolutionConfiguration.html
	UserGroupResolutionConfiguration *types.UserGroupResolutionConfiguration

	// The user token configuration.
	UserTokenConfigurations []types.UserTokenConfiguration

	noSmithyDocumentSerde
}

type UpdateIndexOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateIndex{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateIndex"); err != nil {
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
	if err = addOpUpdateIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIndex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateIndex",
	}
}

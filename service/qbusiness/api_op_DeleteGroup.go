// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a group so that all users and sub groups that belong to the group can
// no longer access documents only available to that group. For example, after
// deleting the group "Summer Interns", all interns who belonged to that group no
// longer see intern-only documents in their chat results.
//
// If you want to delete, update, or replace users or sub groups of a group, you
// need to use the PutGroup operation. For example, if a user in the group
// "Engineering" leaves the engineering team and another user takes their place,
// you provide an updated list of users or sub groups that belong to the
// "Engineering" group when calling PutGroup .
func (c *Client) DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) {
	if params == nil {
		params = &DeleteGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGroup", params, optFns, c.addOperationDeleteGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGroupInput struct {

	// The identifier of the application in which the group mapping belongs.
	//
	// This member is required.
	ApplicationId *string

	// The name of the group you want to delete.
	//
	// This member is required.
	GroupName *string

	// The identifier of the index you want to delete the group from.
	//
	// This member is required.
	IndexId *string

	// The identifier of the data source linked to the group
	//
	// A group can be tied to multiple data sources. You can delete a group from
	// accessing documents in a certain data source. For example, the groups
	// "Research", "Engineering", and "Sales and Marketing" are all tied to the
	// company's documents stored in the data sources Confluence and Salesforce. You
	// want to delete "Research" and "Engineering" groups from Salesforce, so that
	// these groups cannot access customer-related documents stored in Salesforce. Only
	// "Sales and Marketing" should access documents in the Salesforce data source.
	DataSourceId *string

	noSmithyDocumentSerde
}

type DeleteGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteGroup"); err != nil {
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
	if err = addOpDeleteGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteGroup",
	}
}

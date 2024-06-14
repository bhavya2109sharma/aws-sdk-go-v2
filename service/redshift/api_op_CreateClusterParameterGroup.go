// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Redshift parameter group.
//
// Creating parameter groups is independent of creating clusters. You can
// associate a cluster with a parameter group when you create the cluster. You can
// also associate an existing cluster with a parameter group after the cluster is
// created by using ModifyCluster.
//
// Parameters in the parameter group define specific behavior that applies to the
// databases you create on the cluster. For more information about parameters and
// parameter groups, go to [Amazon Redshift Parameter Groups]in the Amazon Redshift Cluster Management Guide.
//
// [Amazon Redshift Parameter Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html
func (c *Client) CreateClusterParameterGroup(ctx context.Context, params *CreateClusterParameterGroupInput, optFns ...func(*Options)) (*CreateClusterParameterGroupOutput, error) {
	if params == nil {
		params = &CreateClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateClusterParameterGroup", params, optFns, c.addOperationCreateClusterParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterParameterGroupInput struct {

	// A description of the parameter group.
	//
	// This member is required.
	Description *string

	// The Amazon Redshift engine version to which the cluster parameter group
	// applies. The cluster engine version determines the set of parameters.
	//
	// To get a list of valid parameter group family names, you can call DescribeClusterParameterGroups. By default,
	// Amazon Redshift returns a list of all the parameter groups that are owned by
	// your Amazon Web Services account, including the default parameter groups for
	// each Amazon Redshift engine version. The parameter group family names associated
	// with the default parameter groups provide you the valid values. For example, a
	// valid family name is "redshift-1.0".
	//
	// This member is required.
	ParameterGroupFamily *string

	// The name of the cluster parameter group.
	//
	// Constraints:
	//
	//   - Must be 1 to 255 alphanumeric characters or hyphens
	//
	//   - First character must be a letter.
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	//   - Must be unique withing your Amazon Web Services account.
	//
	// This value is stored as a lower-case string.
	//
	// This member is required.
	ParameterGroupName *string

	// A list of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateClusterParameterGroupOutput struct {

	// Describes a parameter group.
	ClusterParameterGroup *types.ClusterParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClusterParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateClusterParameterGroup"); err != nil {
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
	if err = addOpCreateClusterParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClusterParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateClusterParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateClusterParameterGroup",
	}
}

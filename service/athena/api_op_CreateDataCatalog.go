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

// Creates (registers) a data catalog with the specified name and properties.
// Catalogs created are visible to all users of the same Amazon Web Services
// account.
func (c *Client) CreateDataCatalog(ctx context.Context, params *CreateDataCatalogInput, optFns ...func(*Options)) (*CreateDataCatalogOutput, error) {
	if params == nil {
		params = &CreateDataCatalogInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataCatalog", params, optFns, c.addOperationCreateDataCatalogMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataCatalogInput struct {

	// The name of the data catalog to create. The catalog name must be unique for the
	// Amazon Web Services account and can use a maximum of 127 alphanumeric,
	// underscore, at sign, or hyphen characters. The remainder of the length
	// constraint of 256 is reserved for use by Athena.
	//
	// This member is required.
	Name *string

	// The type of data catalog to create: LAMBDA for a federated catalog, HIVE for an
	// external hive metastore, or GLUE for an Glue Data Catalog.
	//
	// This member is required.
	Type types.DataCatalogType

	// A description of the data catalog to be created.
	Description *string

	// Specifies the Lambda function or functions to use for creating the data
	// catalog. This is a mapping whose values depend on the catalog type.
	//
	//   - For the HIVE data catalog type, use the following syntax. The
	//   metadata-function parameter is required. The sdk-version parameter is optional
	//   and defaults to the currently supported version.
	//
	// metadata-function=lambda_arn, sdk-version=version_number
	//
	//   - For the LAMBDA data catalog type, use one of the following sets of required
	//   parameters, but not both.
	//
	//   - If you have one Lambda function that processes metadata and another for
	//   reading the actual data, use the following syntax. Both parameters are required.
	//
	// metadata-function=lambda_arn, record-function=lambda_arn
	//
	//   - If you have a composite Lambda function that processes both metadata and
	//   data, use the following syntax to specify your Lambda function.
	//
	// function=lambda_arn
	//
	//   - The GLUE type takes a catalog ID parameter and is required. The catalog_id
	//   is the account ID of the Amazon Web Services account to which the Glue Data
	//   Catalog belongs.
	//
	// catalog-id=catalog_id
	//
	//   - The GLUE data catalog type also applies to the default AwsDataCatalog that
	//   already exists in your account, of which you can have only one and cannot
	//   modify.
	Parameters map[string]string

	// A list of comma separated tags to add to the data catalog that is created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDataCatalogOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataCatalogMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataCatalog"); err != nil {
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
	if err = addOpCreateDataCatalogValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataCatalog(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataCatalog(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataCatalog",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a set of profiles that belong to the same matching group using the
// matchId or profileId . You can also specify the type of matching that you want
// for finding similar profiles using either RULE_BASED_MATCHING or
// ML_BASED_MATCHING .
func (c *Client) GetSimilarProfiles(ctx context.Context, params *GetSimilarProfilesInput, optFns ...func(*Options)) (*GetSimilarProfilesOutput, error) {
	if params == nil {
		params = &GetSimilarProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSimilarProfiles", params, optFns, c.addOperationGetSimilarProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSimilarProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSimilarProfilesInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// Specify the type of matching to get similar profiles for.
	//
	// This member is required.
	MatchType types.MatchType

	// The string indicating the search key to be used.
	//
	// This member is required.
	SearchKey *string

	// The string based on SearchKey to be searched for similar profiles.
	//
	// This member is required.
	SearchValue *string

	// The maximum number of objects returned per page.
	MaxResults *int32

	// The pagination token from the previous GetSimilarProfiles API call.
	NextToken *string

	noSmithyDocumentSerde
}

type GetSimilarProfilesOutput struct {

	// It only has value when the MatchType is ML_BASED_MATCHING .A number between 0
	// and 1, where a higher score means higher similarity. Examining match confidence
	// scores lets you distinguish between groups of similar records in which the
	// system is highly confident (which you may decide to merge), groups of similar
	// records about which the system is uncertain (which you may decide to have
	// reviewed by a human), and groups of similar records that the system deems to be
	// unlikely (which you may decide to reject). Given confidence scores vary as per
	// the data input, it should not be used as an absolute measure of matching
	// quality.
	ConfidenceScore *float64

	// The string matchId that the similar profiles belong to.
	MatchId *string

	// Specify the type of matching to get similar profiles for.
	MatchType types.MatchType

	// The pagination token from the previous GetSimilarProfiles API call.
	NextToken *string

	// Set of profileId s that belong to the same matching group.
	ProfileIds []string

	// The integer rule level that the profiles matched on.
	RuleLevel *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSimilarProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSimilarProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSimilarProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSimilarProfiles"); err != nil {
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
	if err = addOpGetSimilarProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSimilarProfiles(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSimilarProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSimilarProfiles",
	}
}

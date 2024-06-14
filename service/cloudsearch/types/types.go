// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The configured access rules for the domain's document and search endpoints, and
// the current status of those rules.
type AccessPoliciesStatus struct {

	// Access rules for a domain's document or search service endpoints. For more
	// information, see [Configuring Access for a Search Domain]in the Amazon CloudSearch Developer Guide. The maximum size of
	// a policy document is 100 KB.
	//
	// [Configuring Access for a Search Domain]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-access.html
	//
	// This member is required.
	Options *string

	// The status of domain configuration option.
	//
	// This member is required.
	Status *OptionStatus

	noSmithyDocumentSerde
}

// Synonyms, stopwords, and stemming options for an analysis scheme. Includes
// tokenization dictionary for Japanese.
type AnalysisOptions struct {

	// The level of algorithmic stemming to perform: none , minimal , light , or full .
	// The available levels vary depending on the language. For more information, see [Language Specific Text Processing Settings]
	// in the Amazon CloudSearch Developer Guide
	//
	// [Language Specific Text Processing Settings]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/text-processing.html#text-processing-settings
	AlgorithmicStemming AlgorithmicStemming

	// A JSON array that contains a collection of terms, tokens, readings and part of
	// speech for Japanese Tokenizaiton. The Japanese tokenization dictionary enables
	// you to override the default tokenization for selected terms. This is only valid
	// for Japanese language fields.
	JapaneseTokenizationDictionary *string

	// A JSON object that contains a collection of string:value pairs that each map a
	// term to its stem. For example, {"term1": "stem1", "term2": "stem2", "term3":
	// "stem3"} . The stemming dictionary is applied in addition to any algorithmic
	// stemming. This enables you to override the results of the algorithmic stemming
	// to correct specific cases of overstemming or understemming. The maximum size of
	// a stemming dictionary is 500 KB.
	StemmingDictionary *string

	// A JSON array of terms to ignore during indexing and searching. For example,
	// ["a", "an", "the", "of"] . The stopwords dictionary must explicitly list each
	// word you want to ignore. Wildcards and regular expressions are not supported.
	Stopwords *string

	// A JSON object that defines synonym groups and aliases. A synonym group is an
	// array of arrays, where each sub-array is a group of terms where each term in the
	// group is considered a synonym of every other term in the group. The aliases
	// value is an object that contains a collection of string:value pairs where the
	// string specifies a term and the array of values specifies each of the aliases
	// for that term. An alias is considered a synonym of the specified term, but the
	// term is not considered a synonym of the alias. For more information about
	// specifying synonyms, see [Synonyms]in the Amazon CloudSearch Developer Guide.
	//
	// [Synonyms]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-analysis-schemes.html#synonyms
	Synonyms *string

	noSmithyDocumentSerde
}

// Configuration information for an analysis scheme. Each analysis scheme has a
// unique name and specifies the language of the text to be processed. The
// following options can be configured for an analysis scheme: Synonyms , Stopwords
// , StemmingDictionary , JapaneseTokenizationDictionary and AlgorithmicStemming .
type AnalysisScheme struct {

	// An [IETF RFC 4646] language code or mul for multiple languages.
	//
	// [IETF RFC 4646]: http://tools.ietf.org/html/rfc4646
	//
	// This member is required.
	AnalysisSchemeLanguage AnalysisSchemeLanguage

	// Names must begin with a letter and can contain the following characters: a-z
	// (lowercase), 0-9, and _ (underscore).
	//
	// This member is required.
	AnalysisSchemeName *string

	// Synonyms, stopwords, and stemming options for an analysis scheme. Includes
	// tokenization dictionary for Japanese.
	AnalysisOptions *AnalysisOptions

	noSmithyDocumentSerde
}

// The status and configuration of an AnalysisScheme .
type AnalysisSchemeStatus struct {

	// Configuration information for an analysis scheme. Each analysis scheme has a
	// unique name and specifies the language of the text to be processed. The
	// following options can be configured for an analysis scheme: Synonyms , Stopwords
	// , StemmingDictionary , JapaneseTokenizationDictionary and AlgorithmicStemming .
	//
	// This member is required.
	Options *AnalysisScheme

	// The status of domain configuration option.
	//
	// This member is required.
	Status *OptionStatus

	noSmithyDocumentSerde
}

// The status and configuration of the domain's availability options.
type AvailabilityOptionsStatus struct {

	// The availability options configured for the domain.
	//
	// This member is required.
	Options bool

	// The status of domain configuration option.
	//
	// This member is required.
	Status *OptionStatus

	noSmithyDocumentSerde
}

// Options for a field that contains an array of dates. Present if IndexFieldType
// specifies the field is of type date-array . All options are enabled by default.
type DateArrayOptions struct {

	// A value to use for the field if the field isn't specified for a document.
	DefaultValue *string

	// Whether facet information can be returned for the field.
	FacetEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the contents of the field are searchable.
	SearchEnabled *bool

	// A list of source fields to map to the field.
	SourceFields *string

	noSmithyDocumentSerde
}

// Options for a date field. Dates and times are specified in UTC (Coordinated
// Universal Time) according to IETF RFC3339: yyyy-mm-ddT00:00:00Z. Present if
// IndexFieldType specifies the field is of type date . All options are enabled by
// default.
type DateOptions struct {

	// A value to use for the field if the field isn't specified for a document.
	DefaultValue *string

	// Whether facet information can be returned for the field.
	FacetEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the contents of the field are searchable.
	SearchEnabled *bool

	// Whether the field can be used to sort the search results.
	SortEnabled *bool

	// A string that represents the name of an index field. CloudSearch supports
	// regular index fields as well as dynamic fields. A dynamic field's name defines a
	// pattern that begins or ends with a wildcard. Any document fields that don't map
	// to a regular index field but do match a dynamic field's pattern are configured
	// with the dynamic field's indexing options.
	//
	// Regular field names begin with a letter and can contain the following
	// characters: a-z (lowercase), 0-9, and _ (underscore). Dynamic field names must
	// begin or end with a wildcard (*). The wildcard can also be the only character in
	// a dynamic field name. Multiple wildcards, and wildcards embedded within a string
	// are not supported.
	//
	// The name score is reserved and cannot be used as a field name. To reference a
	// document's ID, you can use the name _id .
	SourceField *string

	noSmithyDocumentSerde
}

// Options for a search suggester.
type DocumentSuggesterOptions struct {

	// The name of the index field you want to use for suggestions.
	//
	// This member is required.
	SourceField *string

	// The level of fuzziness allowed when suggesting matches for a string: none , low
	// , or high . With none, the specified string is treated as an exact prefix. With
	// low, suggestions must differ from the specified string by no more than one
	// character. With high, suggestions can differ by up to two characters. The
	// default is none.
	FuzzyMatching SuggesterFuzzyMatching

	// An expression that computes a score for each suggestion to control how they are
	// sorted. The scores are rounded to the nearest integer, with a floor of 0 and a
	// ceiling of 2^31-1. A document's relevance score is not computed for suggestions,
	// so sort expressions cannot reference the _score value. To sort suggestions
	// using a numeric field or existing expression, simply specify the name of the
	// field or expression. If no expression is configured for the suggester, the
	// suggestions are sorted with the closest matches listed first.
	SortExpression *string

	noSmithyDocumentSerde
}

// The domain's endpoint options.
type DomainEndpointOptions struct {

	// Whether the domain is HTTPS only enabled.
	EnforceHTTPS *bool

	// The minimum required TLS version
	TLSSecurityPolicy TLSSecurityPolicy

	noSmithyDocumentSerde
}

// The configuration and status of the domain's endpoint options.
type DomainEndpointOptionsStatus struct {

	// The domain endpoint options configured for the domain.
	//
	// This member is required.
	Options *DomainEndpointOptions

	// The status of the configured domain endpoint options.
	//
	// This member is required.
	Status *OptionStatus

	noSmithyDocumentSerde
}

// The current status of the search domain.
type DomainStatus struct {

	// An internally generated unique identifier for a domain.
	//
	// This member is required.
	DomainId *string

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// True if IndexDocuments needs to be called to activate the current domain configuration.
	//
	// This member is required.
	RequiresIndexDocuments *bool

	// The Amazon Resource Name (ARN) of the search domain. See [Identifiers for IAM Entities] in Using AWS Identity
	// and Access Management for more information.
	//
	// [Identifiers for IAM Entities]: http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?Using_Identifiers.html
	ARN *string

	// True if the search domain is created. It can take several minutes to initialize
	// a domain when CreateDomainis called. Newly created search domains are returned from DescribeDomains with a
	// false value for Created until domain creation is complete.
	Created *bool

	// True if the search domain has been deleted. The system must clean up resources
	// dedicated to the search domain when DeleteDomainis called. Newly deleted search domains are
	// returned from DescribeDomainswith a true value for IsDeleted for several minutes until
	// resource cleanup is complete.
	Deleted *bool

	// The service endpoint for updating documents in a search domain.
	DocService *ServiceEndpoint

	Limits *Limits

	// True if processing is being done to activate the current domain configuration.
	Processing *bool

	// The number of search instances that are available to process search requests.
	SearchInstanceCount *int32

	// The instance type that is being used to process search requests.
	SearchInstanceType *string

	// The number of partitions across which the search index is spread.
	SearchPartitionCount *int32

	// The service endpoint for requesting search results from a search domain.
	SearchService *ServiceEndpoint

	noSmithyDocumentSerde
}

// Options for a field that contains an array of double-precision 64-bit floating
// point values. Present if IndexFieldType specifies the field is of type
// double-array . All options are enabled by default.
type DoubleArrayOptions struct {

	// A value to use for the field if the field isn't specified for a document.
	DefaultValue *float64

	// Whether facet information can be returned for the field.
	FacetEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the contents of the field are searchable.
	SearchEnabled *bool

	// A list of source fields to map to the field.
	SourceFields *string

	noSmithyDocumentSerde
}

// Options for a double-precision 64-bit floating point field. Present if
// IndexFieldType specifies the field is of type double . All options are enabled
// by default.
type DoubleOptions struct {

	// A value to use for the field if the field isn't specified for a document. This
	// can be important if you are using the field in an expression and that field is
	// not present in every document.
	DefaultValue *float64

	// Whether facet information can be returned for the field.
	FacetEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the contents of the field are searchable.
	SearchEnabled *bool

	// Whether the field can be used to sort the search results.
	SortEnabled *bool

	// The name of the source field to map to the field.
	SourceField *string

	noSmithyDocumentSerde
}

// A named expression that can be evaluated at search time. Can be used to sort
// the search results, define other expressions, or return computed information in
// the search results.
type Expression struct {

	// Names must begin with a letter and can contain the following characters: a-z
	// (lowercase), 0-9, and _ (underscore).
	//
	// This member is required.
	ExpressionName *string

	// The expression to evaluate for sorting while processing a search request. The
	// Expression syntax is based on JavaScript expressions. For more information, see [Configuring Expressions]
	// in the Amazon CloudSearch Developer Guide.
	//
	// [Configuring Expressions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-expressions.html
	//
	// This member is required.
	ExpressionValue *string

	noSmithyDocumentSerde
}

// The value of an Expression and its current status.
type ExpressionStatus struct {

	// The expression that is evaluated for sorting while processing a search request.
	//
	// This member is required.
	Options *Expression

	// The status of domain configuration option.
	//
	// This member is required.
	Status *OptionStatus

	noSmithyDocumentSerde
}

// Configuration information for a field in the index, including its name, type,
// and options. The supported options depend on the IndexFieldType.
type IndexField struct {

	// A string that represents the name of an index field. CloudSearch supports
	// regular index fields as well as dynamic fields. A dynamic field's name defines a
	// pattern that begins or ends with a wildcard. Any document fields that don't map
	// to a regular index field but do match a dynamic field's pattern are configured
	// with the dynamic field's indexing options.
	//
	// Regular field names begin with a letter and can contain the following
	// characters: a-z (lowercase), 0-9, and _ (underscore). Dynamic field names must
	// begin or end with a wildcard (*). The wildcard can also be the only character in
	// a dynamic field name. Multiple wildcards, and wildcards embedded within a string
	// are not supported.
	//
	// The name score is reserved and cannot be used as a field name. To reference a
	// document's ID, you can use the name _id .
	//
	// This member is required.
	IndexFieldName *string

	// The type of field. The valid options for a field depend on the field type. For
	// more information about the supported field types, see [Configuring Index Fields]in the Amazon CloudSearch
	// Developer Guide.
	//
	// [Configuring Index Fields]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-index-fields.html
	//
	// This member is required.
	IndexFieldType IndexFieldType

	// Options for a field that contains an array of dates. Present if IndexFieldType
	// specifies the field is of type date-array . All options are enabled by default.
	DateArrayOptions *DateArrayOptions

	// Options for a date field. Dates and times are specified in UTC (Coordinated
	// Universal Time) according to IETF RFC3339: yyyy-mm-ddT00:00:00Z. Present if
	// IndexFieldType specifies the field is of type date . All options are enabled by
	// default.
	DateOptions *DateOptions

	// Options for a field that contains an array of double-precision 64-bit floating
	// point values. Present if IndexFieldType specifies the field is of type
	// double-array . All options are enabled by default.
	DoubleArrayOptions *DoubleArrayOptions

	// Options for a double-precision 64-bit floating point field. Present if
	// IndexFieldType specifies the field is of type double . All options are enabled
	// by default.
	DoubleOptions *DoubleOptions

	// Options for a field that contains an array of 64-bit signed integers. Present
	// if IndexFieldType specifies the field is of type int-array . All options are
	// enabled by default.
	IntArrayOptions *IntArrayOptions

	// Options for a 64-bit signed integer field. Present if IndexFieldType specifies
	// the field is of type int . All options are enabled by default.
	IntOptions *IntOptions

	// Options for a latlon field. A latlon field contains a location stored as a
	// latitude and longitude value pair. Present if IndexFieldType specifies the
	// field is of type latlon . All options are enabled by default.
	LatLonOptions *LatLonOptions

	// Options for a field that contains an array of literal strings. Present if
	// IndexFieldType specifies the field is of type literal-array . All options are
	// enabled by default.
	LiteralArrayOptions *LiteralArrayOptions

	// Options for literal field. Present if IndexFieldType specifies the field is of
	// type literal . All options are enabled by default.
	LiteralOptions *LiteralOptions

	// Options for a field that contains an array of text strings. Present if
	// IndexFieldType specifies the field is of type text-array . A text-array field
	// is always searchable. All options are enabled by default.
	TextArrayOptions *TextArrayOptions

	// Options for text field. Present if IndexFieldType specifies the field is of
	// type text . A text field is always searchable. All options are enabled by
	// default.
	TextOptions *TextOptions

	noSmithyDocumentSerde
}

// The value of an IndexField and its current status.
type IndexFieldStatus struct {

	// Configuration information for a field in the index, including its name, type,
	// and options. The supported options depend on the IndexFieldType.
	//
	// This member is required.
	Options *IndexField

	// The status of domain configuration option.
	//
	// This member is required.
	Status *OptionStatus

	noSmithyDocumentSerde
}

// Options for a field that contains an array of 64-bit signed integers. Present
// if IndexFieldType specifies the field is of type int-array . All options are
// enabled by default.
type IntArrayOptions struct {

	// A value to use for the field if the field isn't specified for a document.
	DefaultValue *int64

	// Whether facet information can be returned for the field.
	FacetEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the contents of the field are searchable.
	SearchEnabled *bool

	// A list of source fields to map to the field.
	SourceFields *string

	noSmithyDocumentSerde
}

// Options for a 64-bit signed integer field. Present if IndexFieldType specifies
// the field is of type int . All options are enabled by default.
type IntOptions struct {

	// A value to use for the field if the field isn't specified for a document. This
	// can be important if you are using the field in an expression and that field is
	// not present in every document.
	DefaultValue *int64

	// Whether facet information can be returned for the field.
	FacetEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the contents of the field are searchable.
	SearchEnabled *bool

	// Whether the field can be used to sort the search results.
	SortEnabled *bool

	// The name of the source field to map to the field.
	SourceField *string

	noSmithyDocumentSerde
}

// Options for a latlon field. A latlon field contains a location stored as a
// latitude and longitude value pair. Present if IndexFieldType specifies the
// field is of type latlon . All options are enabled by default.
type LatLonOptions struct {

	// A value to use for the field if the field isn't specified for a document.
	DefaultValue *string

	// Whether facet information can be returned for the field.
	FacetEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the contents of the field are searchable.
	SearchEnabled *bool

	// Whether the field can be used to sort the search results.
	SortEnabled *bool

	// A string that represents the name of an index field. CloudSearch supports
	// regular index fields as well as dynamic fields. A dynamic field's name defines a
	// pattern that begins or ends with a wildcard. Any document fields that don't map
	// to a regular index field but do match a dynamic field's pattern are configured
	// with the dynamic field's indexing options.
	//
	// Regular field names begin with a letter and can contain the following
	// characters: a-z (lowercase), 0-9, and _ (underscore). Dynamic field names must
	// begin or end with a wildcard (*). The wildcard can also be the only character in
	// a dynamic field name. Multiple wildcards, and wildcards embedded within a string
	// are not supported.
	//
	// The name score is reserved and cannot be used as a field name. To reference a
	// document's ID, you can use the name _id .
	SourceField *string

	noSmithyDocumentSerde
}

type Limits struct {

	// This member is required.
	MaximumPartitionCount *int32

	// This member is required.
	MaximumReplicationCount *int32

	noSmithyDocumentSerde
}

// Options for a field that contains an array of literal strings. Present if
// IndexFieldType specifies the field is of type literal-array . All options are
// enabled by default.
type LiteralArrayOptions struct {

	// A value to use for the field if the field isn't specified for a document.
	DefaultValue *string

	// Whether facet information can be returned for the field.
	FacetEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the contents of the field are searchable.
	SearchEnabled *bool

	// A list of source fields to map to the field.
	SourceFields *string

	noSmithyDocumentSerde
}

// Options for literal field. Present if IndexFieldType specifies the field is of
// type literal . All options are enabled by default.
type LiteralOptions struct {

	// A value to use for the field if the field isn't specified for a document.
	DefaultValue *string

	// Whether facet information can be returned for the field.
	FacetEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the contents of the field are searchable.
	SearchEnabled *bool

	// Whether the field can be used to sort the search results.
	SortEnabled *bool

	// A string that represents the name of an index field. CloudSearch supports
	// regular index fields as well as dynamic fields. A dynamic field's name defines a
	// pattern that begins or ends with a wildcard. Any document fields that don't map
	// to a regular index field but do match a dynamic field's pattern are configured
	// with the dynamic field's indexing options.
	//
	// Regular field names begin with a letter and can contain the following
	// characters: a-z (lowercase), 0-9, and _ (underscore). Dynamic field names must
	// begin or end with a wildcard (*). The wildcard can also be the only character in
	// a dynamic field name. Multiple wildcards, and wildcards embedded within a string
	// are not supported.
	//
	// The name score is reserved and cannot be used as a field name. To reference a
	// document's ID, you can use the name _id .
	SourceField *string

	noSmithyDocumentSerde
}

// The status of domain configuration option.
type OptionStatus struct {

	// A timestamp for when this option was created.
	//
	// This member is required.
	CreationDate *time.Time

	// The state of processing a change to an option. Possible values:
	//
	//   - RequiresIndexDocuments : the option's latest value will not be deployed
	//   until IndexDocumentshas been called and indexing is complete.
	//   - Processing : the option's latest value is in the process of being activated.
	//   - Active : the option's latest value is completely deployed.
	//   - FailedToValidate : the option value is not compatible with the domain's data
	//   and cannot be used to index the data. You must either modify the option value or
	//   update or remove the incompatible documents.
	//
	// This member is required.
	State OptionState

	// A timestamp for when this option was last updated.
	//
	// This member is required.
	UpdateDate *time.Time

	// Indicates that the option will be deleted once processing is complete.
	PendingDeletion *bool

	// A unique integer that indicates when this option was last updated.
	UpdateVersion int32

	noSmithyDocumentSerde
}

// The desired instance type and desired number of replicas of each index
// partition.
type ScalingParameters struct {

	// The instance type that you want to preconfigure for your domain. For example,
	// search.m1.small .
	DesiredInstanceType PartitionInstanceType

	// The number of partitions you want to preconfigure for your domain. Only valid
	// when you select m2.2xlarge as the desired instance type.
	DesiredPartitionCount int32

	// The number of replicas you want to preconfigure for each index partition.
	DesiredReplicationCount int32

	noSmithyDocumentSerde
}

// The status and configuration of a search domain's scaling parameters.
type ScalingParametersStatus struct {

	// The desired instance type and desired number of replicas of each index
	// partition.
	//
	// This member is required.
	Options *ScalingParameters

	// The status of domain configuration option.
	//
	// This member is required.
	Status *OptionStatus

	noSmithyDocumentSerde
}

// The endpoint to which service requests can be submitted.
type ServiceEndpoint struct {

	// The endpoint to which service requests can be submitted. For example,
	// search-imdb-movies-oopcnjfn6ugofer3zx5iadxxca.eu-west-1.cloudsearch.amazonaws.com
	// or
	// doc-imdb-movies-oopcnjfn6ugofer3zx5iadxxca.eu-west-1.cloudsearch.amazonaws.com .
	Endpoint *string

	noSmithyDocumentSerde
}

// Configuration information for a search suggester. Each suggester has a unique
// name and specifies the text field you want to use for suggestions. The following
// options can be configured for a suggester: FuzzyMatching , SortExpression .
type Suggester struct {

	// Options for a search suggester.
	//
	// This member is required.
	DocumentSuggesterOptions *DocumentSuggesterOptions

	// Names must begin with a letter and can contain the following characters: a-z
	// (lowercase), 0-9, and _ (underscore).
	//
	// This member is required.
	SuggesterName *string

	noSmithyDocumentSerde
}

// The value of a Suggester and its current status.
type SuggesterStatus struct {

	// Configuration information for a search suggester. Each suggester has a unique
	// name and specifies the text field you want to use for suggestions. The following
	// options can be configured for a suggester: FuzzyMatching , SortExpression .
	//
	// This member is required.
	Options *Suggester

	// The status of domain configuration option.
	//
	// This member is required.
	Status *OptionStatus

	noSmithyDocumentSerde
}

// Options for a field that contains an array of text strings. Present if
// IndexFieldType specifies the field is of type text-array . A text-array field
// is always searchable. All options are enabled by default.
type TextArrayOptions struct {

	// The name of an analysis scheme for a text-array field.
	AnalysisScheme *string

	// A value to use for the field if the field isn't specified for a document.
	DefaultValue *string

	// Whether highlights can be returned for the field.
	HighlightEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// A list of source fields to map to the field.
	SourceFields *string

	noSmithyDocumentSerde
}

// Options for text field. Present if IndexFieldType specifies the field is of
// type text . A text field is always searchable. All options are enabled by
// default.
type TextOptions struct {

	// The name of an analysis scheme for a text field.
	AnalysisScheme *string

	// A value to use for the field if the field isn't specified for a document.
	DefaultValue *string

	// Whether highlights can be returned for the field.
	HighlightEnabled *bool

	// Whether the contents of the field can be returned in the search results.
	ReturnEnabled *bool

	// Whether the field can be used to sort the search results.
	SortEnabled *bool

	// A string that represents the name of an index field. CloudSearch supports
	// regular index fields as well as dynamic fields. A dynamic field's name defines a
	// pattern that begins or ends with a wildcard. Any document fields that don't map
	// to a regular index field but do match a dynamic field's pattern are configured
	// with the dynamic field's indexing options.
	//
	// Regular field names begin with a letter and can contain the following
	// characters: a-z (lowercase), 0-9, and _ (underscore). Dynamic field names must
	// begin or end with a wildcard (*). The wildcard can also be the only character in
	// a dynamic field name. Multiple wildcards, and wildcards embedded within a string
	// are not supported.
	//
	// The name score is reserved and cannot be used as a field name. To reference a
	// document's ID, you can use the name _id .
	SourceField *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

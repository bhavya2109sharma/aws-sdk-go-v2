// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type InternalServerExceptionReason string

// Enum values for InternalServerExceptionReason
const (
	InternalServerExceptionReasonFailedToGenerateSbom InternalServerExceptionReason = "FAILED_TO_GENERATE_SBOM"
	InternalServerExceptionReasonOther                InternalServerExceptionReason = "OTHER"
)

// Values returns all known values for InternalServerExceptionReason. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InternalServerExceptionReason) Values() []InternalServerExceptionReason {
	return []InternalServerExceptionReason{
		"FAILED_TO_GENERATE_SBOM",
		"OTHER",
	}
}

type OutputFormat string

// Enum values for OutputFormat
const (
	OutputFormatCycloneDx15 OutputFormat = "CYCLONE_DX_1_5"
	OutputFormatInspector   OutputFormat = "INSPECTOR"
)

// Values returns all known values for OutputFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OutputFormat) Values() []OutputFormat {
	return []OutputFormat{
		"CYCLONE_DX_1_5",
		"INSPECTOR",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonUnsupportedSbomType   ValidationExceptionReason = "UNSUPPORTED_SBOM_TYPE"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UNKNOWN_OPERATION",
		"CANNOT_PARSE",
		"FIELD_VALIDATION_FAILED",
		"UNSUPPORTED_SBOM_TYPE",
		"OTHER",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Details about the progress of a batch load task.
type BatchLoadProgressReport struct {

	//
	BytesMetered int64

	//
	FileFailures int64

	//
	ParseFailures int64

	//
	RecordIngestionFailures int64

	//
	RecordsIngested int64

	//
	RecordsProcessed int64

	noSmithyDocumentSerde
}

// Details about a batch load task.
type BatchLoadTask struct {

	// The time when the Timestream batch load task was created.
	CreationTime *time.Time

	// Database name for the database into which a batch load task loads data.
	DatabaseName *string

	// The time when the Timestream batch load task was last updated.
	LastUpdatedTime *time.Time

	//
	ResumableUntil *time.Time

	// Table name for the table into which a batch load task loads data.
	TableName *string

	// The ID of the batch load task.
	TaskId *string

	// Status of the batch load task.
	TaskStatus BatchLoadStatus

	noSmithyDocumentSerde
}

// Details about a batch load task.
type BatchLoadTaskDescription struct {

	// The time when the Timestream batch load task was created.
	CreationTime *time.Time

	// Data model configuration for a batch load task. This contains details about
	// where a data model for a batch load task is stored.
	DataModelConfiguration *DataModelConfiguration

	// Configuration details about the data source for a batch load task.
	DataSourceConfiguration *DataSourceConfiguration

	//
	ErrorMessage *string

	// The time when the Timestream batch load task was last updated.
	LastUpdatedTime *time.Time

	//
	ProgressReport *BatchLoadProgressReport

	//
	RecordVersion int64

	// Report configuration for a batch load task. This contains details about where
	// error reports are stored.
	ReportConfiguration *ReportConfiguration

	//
	ResumableUntil *time.Time

	//
	TargetDatabaseName *string

	//
	TargetTableName *string

	// The ID of the batch load task.
	TaskId *string

	// Status of the batch load task.
	TaskStatus BatchLoadStatus

	noSmithyDocumentSerde
}

// A delimited data format where the column separator can be a comma and the
// record separator is a newline character.
type CsvConfiguration struct {

	// Column separator can be one of comma (','), pipe ('|), semicolon (';'),
	// tab('/t'), or blank space (' ').
	ColumnSeparator *string

	// Escape character can be one of
	EscapeChar *string

	// Can be blank space (' ').
	NullValue *string

	// Can be single quote (') or double quote (").
	QuoteChar *string

	// Specifies to trim leading and trailing white space.
	TrimWhiteSpace *bool

	noSmithyDocumentSerde
}

// A top-level container for a table. Databases and tables are the fundamental
// management concepts in Amazon Timestream. All tables in a database are encrypted
// with the same KMS key.
type Database struct {

	// The Amazon Resource Name that uniquely identifies this database.
	Arn *string

	// The time when the database was created, calculated from the Unix epoch time.
	CreationTime *time.Time

	// The name of the Timestream database.
	DatabaseName *string

	// The identifier of the KMS key used to encrypt the data stored in the database.
	KmsKeyId *string

	//  The last time that this database was updated.
	LastUpdatedTime *time.Time

	// The total number of tables found within a Timestream database.
	TableCount int64

	noSmithyDocumentSerde
}

// Data model for a batch load task.
type DataModel struct {

	// Source to target mappings for dimensions.
	//
	// This member is required.
	DimensionMappings []DimensionMapping

	//
	MeasureNameColumn *string

	// Source to target mappings for measures.
	MixedMeasureMappings []MixedMeasureMapping

	// Source to target mappings for multi-measure records.
	MultiMeasureMappings *MultiMeasureMappings

	// Source column to be mapped to time.
	TimeColumn *string

	//  The granularity of the timestamp unit. It indicates if the time value is in
	// seconds, milliseconds, nanoseconds, or other supported values. Default is
	// MILLISECONDS .
	TimeUnit TimeUnit

	noSmithyDocumentSerde
}

type DataModelConfiguration struct {

	//
	DataModel *DataModel

	//
	DataModelS3Configuration *DataModelS3Configuration

	noSmithyDocumentSerde
}

type DataModelS3Configuration struct {

	//
	BucketName *string

	//
	ObjectKey *string

	noSmithyDocumentSerde
}

// Defines configuration details about the data source.
type DataSourceConfiguration struct {

	// This is currently CSV.
	//
	// This member is required.
	DataFormat BatchLoadDataFormat

	// Configuration of an S3 location for a file which contains data to load.
	//
	// This member is required.
	DataSourceS3Configuration *DataSourceS3Configuration

	// A delimited data format where the column separator can be a comma and the
	// record separator is a newline character.
	CsvConfiguration *CsvConfiguration

	noSmithyDocumentSerde
}

type DataSourceS3Configuration struct {

	// The bucket name of the customer S3 bucket.
	//
	// This member is required.
	BucketName *string

	//
	ObjectKeyPrefix *string

	noSmithyDocumentSerde
}

// Represents the metadata attributes of the time series. For example, the name
// and Availability Zone of an EC2 instance or the name of the manufacturer of a
// wind turbine are dimensions.
type Dimension struct {

	//  Dimension represents the metadata attributes of the time series. For example,
	// the name and Availability Zone of an EC2 instance or the name of the
	// manufacturer of a wind turbine are dimensions.
	//
	// For constraints on dimension names, see [Naming Constraints].
	//
	// [Naming Constraints]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html#limits.naming
	//
	// This member is required.
	Name *string

	// The value of the dimension.
	//
	// This member is required.
	Value *string

	// The data type of the dimension for the time-series data point.
	DimensionValueType DimensionValueType

	noSmithyDocumentSerde
}

type DimensionMapping struct {

	//
	DestinationColumn *string

	//
	SourceColumn *string

	noSmithyDocumentSerde
}

// Represents an available endpoint against which to make API calls against, as
// well as the TTL for that endpoint.
type Endpoint struct {

	// An endpoint address.
	//
	// This member is required.
	Address *string

	// The TTL for the endpoint, in minutes.
	//
	// This member is required.
	CachePeriodInMinutes int64

	noSmithyDocumentSerde
}

// The location to write error reports for records rejected, asynchronously,
// during magnetic store writes.
type MagneticStoreRejectedDataLocation struct {

	// Configuration of an S3 location to write error reports for records rejected,
	// asynchronously, during magnetic store writes.
	S3Configuration *S3Configuration

	noSmithyDocumentSerde
}

// The set of properties on a table for configuring magnetic store writes.
type MagneticStoreWriteProperties struct {

	// A flag to enable magnetic store writes.
	//
	// This member is required.
	EnableMagneticStoreWrites *bool

	// The location to write error reports for records rejected asynchronously during
	// magnetic store writes.
	MagneticStoreRejectedDataLocation *MagneticStoreRejectedDataLocation

	noSmithyDocumentSerde
}

//	Represents the data attribute of the time series. For example, the CPU
//
// utilization of an EC2 instance or the RPM of a wind turbine are measures.
// MeasureValue has both name and value.
//
// MeasureValue is only allowed for type MULTI . Using MULTI type, you can pass
// multiple data attributes associated with the same time series in a single record
type MeasureValue struct {

	//  The name of the MeasureValue.
	//
	// For constraints on MeasureValue names, see [Naming Constraints] in the Amazon Timestream Developer
	// Guide.
	//
	// [Naming Constraints]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html#limits.naming
	//
	// This member is required.
	Name *string

	// Contains the data type of the MeasureValue for the time-series data point.
	//
	// This member is required.
	Type MeasureValueType

	//  The value for the MeasureValue. For information, see [Data types].
	//
	// [Data types]: https://docs.aws.amazon.com/timestream/latest/developerguide/writes.html#writes.data-types
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type MixedMeasureMapping struct {

	//
	//
	// This member is required.
	MeasureValueType MeasureValueType

	//
	MeasureName *string

	//
	MultiMeasureAttributeMappings []MultiMeasureAttributeMapping

	//
	SourceColumn *string

	//
	TargetMeasureName *string

	noSmithyDocumentSerde
}

type MultiMeasureAttributeMapping struct {

	//
	//
	// This member is required.
	SourceColumn *string

	//
	MeasureValueType ScalarMeasureValueType

	//
	TargetMultiMeasureAttributeName *string

	noSmithyDocumentSerde
}

type MultiMeasureMappings struct {

	//
	//
	// This member is required.
	MultiMeasureAttributeMappings []MultiMeasureAttributeMapping

	//
	TargetMultiMeasureName *string

	noSmithyDocumentSerde
}

//	An attribute used in partitioning data in a table. A dimension key partitions
//
// data using the values of the dimension specified by the dimension-name as
// partition key, while a measure key partitions data using measure names (values
// of the 'measure_name' column).
type PartitionKey struct {

	//  The type of the partition key. Options are DIMENSION (dimension key) and
	// MEASURE (measure key).
	//
	// This member is required.
	Type PartitionKeyType

	//  The level of enforcement for the specification of a dimension key in ingested
	// records. Options are REQUIRED (dimension key must be specified) and OPTIONAL
	// (dimension key does not have to be specified).
	EnforcementInRecord PartitionKeyEnforcementLevel

	//  The name of the attribute used for a dimension key.
	Name *string

	noSmithyDocumentSerde
}

// Represents a time-series data point being written into Timestream. Each record
// contains an array of dimensions. Dimensions represent the metadata attributes of
// a time-series data point, such as the instance name or Availability Zone of an
// EC2 instance. A record also contains the measure name, which is the name of the
// measure being collected (for example, the CPU utilization of an EC2 instance).
// Additionally, a record contains the measure value and the value type, which is
// the data type of the measure value. Also, the record contains the timestamp of
// when the measure was collected and the timestamp unit, which represents the
// granularity of the timestamp.
//
// Records have a Version field, which is a 64-bit long that you can use for
// updating data points. Writes of a duplicate record with the same dimension,
// timestamp, and measure name but different measure value will only succeed if the
// Version attribute of the record in the write request is higher than that of the
// existing record. Timestream defaults to a Version of 1 for records without the
// Version field.
type Record struct {

	// Contains the list of dimensions for time-series data points.
	Dimensions []Dimension

	// Measure represents the data attribute of the time series. For example, the CPU
	// utilization of an EC2 instance or the RPM of a wind turbine are measures.
	MeasureName *string

	//  Contains the measure value for the time-series data point.
	MeasureValue *string

	//  Contains the data type of the measure value for the time-series data point.
	// Default type is DOUBLE . For more information, see [Data types].
	//
	// [Data types]: https://docs.aws.amazon.com/timestream/latest/developerguide/writes.html#writes.data-types
	MeasureValueType MeasureValueType

	//  Contains the list of MeasureValue for time-series data points.
	//
	// This is only allowed for type MULTI . For scalar values, use MeasureValue
	// attribute of the record directly.
	MeasureValues []MeasureValue

	//  Contains the time at which the measure value for the data point was collected.
	// The time value plus the unit provides the time elapsed since the epoch. For
	// example, if the time value is 12345 and the unit is ms , then 12345 ms have
	// elapsed since the epoch.
	Time *string

	//  The granularity of the timestamp unit. It indicates if the time value is in
	// seconds, milliseconds, nanoseconds, or other supported values. Default is
	// MILLISECONDS .
	TimeUnit TimeUnit

	// 64-bit attribute used for record updates. Write requests for duplicate data
	// with a higher version number will update the existing measure value and version.
	// In cases where the measure value is the same, Version will still be updated.
	// Default value is 1 .
	//
	// Version must be 1 or greater, or you will receive a ValidationException error.
	Version *int64

	noSmithyDocumentSerde
}

// Information on the records ingested by this request.
type RecordsIngested struct {

	// Count of records ingested into the magnetic store.
	MagneticStore int32

	// Count of records ingested into the memory store.
	MemoryStore int32

	// Total count of successfully ingested records.
	Total int32

	noSmithyDocumentSerde
}

//	Represents records that were not successfully inserted into Timestream due to
//
// data validation issues that must be resolved before reinserting time-series data
// into the system.
type RejectedRecord struct {

	// The existing version of the record. This value is populated in scenarios where
	// an identical record exists with a higher version than the version in the write
	// request.
	ExistingVersion *int64

	//  The reason why a record was not successfully inserted into Timestream.
	// Possible causes of failure include:
	//
	//   - Records with duplicate data where there are multiple records with the same
	//   dimensions, timestamps, and measure names but:
	//
	//   - Measure values are different
	//
	//   - Version is not present in the request, or the value of version in the new
	//   record is equal to or lower than the existing value
	//
	// If Timestream rejects data for this case, the ExistingVersion field in the
	//   RejectedRecords response will indicate the current record’s version. To force
	//   an update, you can resend the request with a version for the record set to a
	//   value greater than the ExistingVersion .
	//
	//   - Records with timestamps that lie outside the retention duration of the
	//   memory store.
	//
	// When the retention window is updated, you will receive a RejectedRecords
	//   exception if you immediately try to ingest data within the new window. To avoid
	//   a RejectedRecords exception, wait until the duration of the new window to
	//   ingest new data. For further information, see [Best Practices for Configuring Timestream]and [the explanation of how storage works in Timestream].
	//
	//   - Records with dimensions or measures that exceed the Timestream defined
	//   limits.
	//
	// For more information, see [Access Management] in the Timestream Developer Guide.
	//
	// [Access Management]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
	// [Best Practices for Configuring Timestream]: https://docs.aws.amazon.com/timestream/latest/developerguide/best-practices.html#configuration
	// [the explanation of how storage works in Timestream]: https://docs.aws.amazon.com/timestream/latest/developerguide/storage.html
	Reason *string

	//  The index of the record in the input request for WriteRecords. Indexes begin
	// with 0.
	RecordIndex int32

	noSmithyDocumentSerde
}

// Report configuration for a batch load task. This contains details about where
// error reports are stored.
type ReportConfiguration struct {

	// Configuration of an S3 location to write error reports and events for a batch
	// load.
	ReportS3Configuration *ReportS3Configuration

	noSmithyDocumentSerde
}

type ReportS3Configuration struct {

	//
	//
	// This member is required.
	BucketName *string

	//
	EncryptionOption S3EncryptionOption

	//
	KmsKeyId *string

	//
	ObjectKeyPrefix *string

	noSmithyDocumentSerde
}

// Retention properties contain the duration for which your time-series data must
// be stored in the magnetic store and the memory store.
type RetentionProperties struct {

	// The duration for which data must be stored in the magnetic store.
	//
	// This member is required.
	MagneticStoreRetentionPeriodInDays *int64

	// The duration for which data must be stored in the memory store.
	//
	// This member is required.
	MemoryStoreRetentionPeriodInHours *int64

	noSmithyDocumentSerde
}

// The configuration that specifies an S3 location.
type S3Configuration struct {

	// The bucket name of the customer S3 bucket.
	BucketName *string

	// The encryption option for the customer S3 location. Options are S3 server-side
	// encryption with an S3 managed key or Amazon Web Services managed key.
	EncryptionOption S3EncryptionOption

	// The KMS key ID for the customer S3 location when encrypting with an Amazon Web
	// Services managed key.
	KmsKeyId *string

	// The object key preview for the customer S3 location.
	ObjectKeyPrefix *string

	noSmithyDocumentSerde
}

// A Schema specifies the expected data model of the table.
type Schema struct {

	// A non-empty list of partition keys defining the attributes used to partition
	// the table data. The order of the list determines the partition hierarchy. The
	// name and type of each partition key as well as the partition key order cannot be
	// changed after the table is created. However, the enforcement level of each
	// partition key can be changed.
	CompositePartitionKey []PartitionKey

	noSmithyDocumentSerde
}

// Represents a database table in Timestream. Tables contain one or more related
// time series. You can modify the retention duration of the memory store and the
// magnetic store for a table.
type Table struct {

	// The Amazon Resource Name that uniquely identifies this table.
	Arn *string

	// The time when the Timestream table was created.
	CreationTime *time.Time

	// The name of the Timestream database that contains this table.
	DatabaseName *string

	// The time when the Timestream table was last updated.
	LastUpdatedTime *time.Time

	// Contains properties to set on the table when enabling magnetic store writes.
	MagneticStoreWriteProperties *MagneticStoreWriteProperties

	// The retention duration for the memory store and magnetic store.
	RetentionProperties *RetentionProperties

	//  The schema of the table.
	Schema *Schema

	// The name of the Timestream table.
	TableName *string

	// The current state of the table:
	//
	//   - DELETING - The table is being deleted.
	//
	//   - ACTIVE - The table is ready for use.
	TableStatus TableStatus

	noSmithyDocumentSerde
}

//	A tag is a label that you assign to a Timestream database and/or table. Each
//
// tag consists of a key and an optional value, both of which you define. With
// tags, you can categorize databases and/or tables, for example, by purpose,
// owner, or environment.
type Tag struct {

	//  The key of the tag. Tag keys are case sensitive.
	//
	// This member is required.
	Key *string

	//  The value of the tag. Tag values are case-sensitive and can be null.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Object specifying a configuration for individual participant recording.
type AutoParticipantRecordingConfiguration struct {

	// ARN of the StorageConfiguration resource to use for individual participant recording. Default: ""
	// (empty string, no storage configuration is specified). Individual participant
	// recording cannot be started unless a storage configuration is specified, when a Stage
	// is created or updated.
	//
	// This member is required.
	StorageConfigurationArn *string

	// Types of media to be recorded. Default: AUDIO_VIDEO .
	MediaTypes []ParticipantRecordingMediaType

	noSmithyDocumentSerde
}

// Object specifying a channel as a destination.
type ChannelDestinationConfiguration struct {

	// ARN of the channel to use for broadcasting. The channel and stage resources
	// must be in the same AWS account and region. The channel must be offline (not
	// broadcasting).
	//
	// This member is required.
	ChannelArn *string

	// ARN of the EncoderConfiguration resource. The encoder configuration and stage resources must be in
	// the same AWS account and region.
	EncoderConfigurationArn *string

	noSmithyDocumentSerde
}

// Object specifying a Composition resource.
type Composition struct {

	// ARN of the Composition resource.
	//
	// This member is required.
	Arn *string

	// Array of Destination objects. A Composition can contain either one destination (
	// channel or s3 ) or two (one channel and one s3 ).
	//
	// This member is required.
	Destinations []Destination

	// Layout object to configure composition parameters.
	//
	// This member is required.
	Layout *LayoutConfiguration

	// ARN of the stage used as input
	//
	// This member is required.
	StageArn *string

	// State of the Composition.
	//
	// This member is required.
	State CompositionState

	// UTC time of the Composition end. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	EndTime *time.Time

	// UTC time of the Composition start. This is an ISO 8601 timestamp; note that
	// this is returned as a string.
	StartTime *time.Time

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary information about a Composition.
type CompositionSummary struct {

	// ARN of the Composition resource.
	//
	// This member is required.
	Arn *string

	// Array of Destination objects.
	//
	// This member is required.
	Destinations []DestinationSummary

	// ARN of the attached stage.
	//
	// This member is required.
	StageArn *string

	// State of the Composition resource.
	//
	// This member is required.
	State CompositionState

	// UTC time of the Composition end. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	EndTime *time.Time

	// UTC time of the Composition start. This is an ISO 8601 timestamp; note that
	// this is returned as a string.
	StartTime *time.Time

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	noSmithyDocumentSerde
}

// Object specifying the status of a Destination.
type Destination struct {

	// Configuration used to create this destination.
	//
	// This member is required.
	Configuration *DestinationConfiguration

	// Unique identifier for this destination, assigned by IVS.
	//
	// This member is required.
	Id *string

	// State of the Composition Destination.
	//
	// This member is required.
	State DestinationState

	// Optional details regarding the status of the destination.
	Detail *DestinationDetail

	// UTC time of the destination end. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	EndTime *time.Time

	// UTC time of the destination start. This is an ISO 8601 timestamp; note that
	// this is returned as a string.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Complex data type that defines destination-configuration objects.
type DestinationConfiguration struct {

	// An IVS channel to be used for broadcasting, for server-side composition. Either
	// a channel or an s3 must be specified.
	Channel *ChannelDestinationConfiguration

	// Name that can be specified to help identify the destination.
	Name *string

	// An S3 storage configuration to be used for recording video data. Either a
	// channel or an s3 must be specified.
	S3 *S3DestinationConfiguration

	noSmithyDocumentSerde
}

// Complex data type that defines destination-detail objects.
type DestinationDetail struct {

	// An S3 detail object to return information about the S3 destination.
	S3 *S3Detail

	noSmithyDocumentSerde
}

// Summary information about a Destination.
type DestinationSummary struct {

	// Unique identifier for this destination, assigned by IVS.
	//
	// This member is required.
	Id *string

	// State of the Composition Destination.
	//
	// This member is required.
	State DestinationState

	// UTC time of the destination end. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	EndTime *time.Time

	// UTC time of the destination start. This is an ISO 8601 timestamp; note that
	// this is returned as a string.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Settings for transcoding.
type EncoderConfiguration struct {

	// ARN of the EncoderConfiguration resource.
	//
	// This member is required.
	Arn *string

	// Optional name to identify the resource.
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	// Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30
	// fps
	Video *Video

	noSmithyDocumentSerde
}

// Summary information about an EncoderConfiguration.
type EncoderConfigurationSummary struct {

	// ARN of the EncoderConfiguration resource.
	//
	// This member is required.
	Arn *string

	// Optional name to identify the resource.
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	noSmithyDocumentSerde
}

// An occurrence during a stage session.
type Event struct {

	// If the event is an error event, the error code is provided to give insight into
	// the specific error that occurred. If the event is not an error event, this field
	// is null. INSUFFICIENT_CAPABILITIES indicates that the participant tried to take
	// an action that the participant’s token is not allowed to do. For more
	// information about participant capabilities, see the capabilities field in CreateParticipantToken.
	// QUOTA_EXCEEDED indicates that the number of participants who want to
	// publish/subscribe to a stage exceeds the quota; for more information, see [Service Quotas].
	// PUBLISHER_NOT_FOUND indicates that the participant tried to subscribe to a
	// publisher that doesn’t exist.
	//
	// [Service Quotas]: https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/service-quotas.html
	ErrorCode EventErrorCode

	// ISO 8601 timestamp (returned as a string) for when the event occurred.
	EventTime *time.Time

	// The name of the event.
	Name EventName

	// Unique identifier for the participant who triggered the event. This is assigned
	// by IVS.
	ParticipantId *string

	// Unique identifier for the remote participant. For a subscribe event, this is
	// the publisher. For a publish or join event, this is null. This is assigned by
	// IVS.
	RemoteParticipantId *string

	noSmithyDocumentSerde
}

// Configuration information specific to Grid layout, for server-side composition.
// See "Layouts" in [Server-Side Composition].
//
// [Server-Side Composition]: https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/server-side-composition.html
type GridConfiguration struct {

	// This attribute name identifies the featured slot. A participant with this
	// attribute set to "true" (as a string value) in ParticipantTokenConfiguration is placed in the featured slot.
	// Default: "" (no featured participant).
	FeaturedParticipantAttribute *string

	// Specifies the spacing between participant tiles in pixels. Default: 2 .
	GridGap int32

	// Determines whether to omit participants with stopped video in the composition.
	// Default: false .
	OmitStoppedVideo bool

	// Sets the non-featured participant display mode, to control the aspect ratio of
	// video tiles. VIDEO is 16:9, SQUARE is 1:1, and PORTRAIT is 3:4. Default: VIDEO .
	VideoAspectRatio VideoAspectRatio

	// Defines how video content fits within the participant tile: FILL (stretched),
	// COVER (cropped), or CONTAIN (letterboxed). When not set, videoFillMode defaults
	// to COVER fill mode for participants in the grid and to CONTAIN fill mode for
	// featured participants.
	VideoFillMode VideoFillMode

	noSmithyDocumentSerde
}

// Object specifying an ingest configuration.
type IngestConfiguration struct {

	// Ingest configuration ARN.
	//
	// This member is required.
	Arn *string

	// Type of ingest protocol that the user employs for broadcasting.
	//
	// This member is required.
	IngestProtocol IngestProtocol

	// ID of the participant within the stage.
	//
	// This member is required.
	ParticipantId *string

	// ARN of the stage with which the IngestConfiguration is associated.
	//
	// This member is required.
	StageArn *string

	// State of the ingest configuration. It is ACTIVE if a publisher currently is
	// publishing to the stage associated with the ingest configuration.
	//
	// This member is required.
	State IngestConfigurationState

	// Ingest-key value for the RTMP(S) protocol.
	//
	// This member is required.
	StreamKey *string

	// Application-provided attributes to to store in the IngestConfiguration and
	// attach to a stage. Map keys and values can contain UTF-8 encoded text. The
	// maximum length of this field is 1 KB total. This field is exposed to all stage
	// participants and should not be used for personally identifying, confidential, or
	// sensitive information.
	Attributes map[string]string

	// Ingest name
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	// Customer-assigned name to help identify the participant using the
	// IngestConfiguration; this can be used to link a participant to a user in the
	// customer’s own systems. This can be any UTF-8 encoded text. This field is
	// exposed to all stage participants and should not be used for personally
	// identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Summary information about an IngestConfiguration.
type IngestConfigurationSummary struct {

	// Ingest configuration ARN.
	//
	// This member is required.
	Arn *string

	// Type of ingest protocol that the user employs for broadcasting.
	//
	// This member is required.
	IngestProtocol IngestProtocol

	// ID of the participant within the stage.
	//
	// This member is required.
	ParticipantId *string

	// ARN of the stage with which the IngestConfiguration is associated.
	//
	// This member is required.
	StageArn *string

	// State of the ingest configuration. It is ACTIVE if a publisher currently is
	// publishing to the stage associated with the ingest configuration.
	//
	// This member is required.
	State IngestConfigurationState

	// Ingest name.
	Name *string

	// Customer-assigned name to help identify the participant using the
	// IngestConfiguration; this can be used to link a participant to a user in the
	// customer’s own systems. This can be any UTF-8 encoded text. This field is
	// exposed to all stage participants and should not be used for personally
	// identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Configuration information of supported layouts for server-side composition.
type LayoutConfiguration struct {

	// Configuration related to grid layout. Default: Grid layout.
	Grid *GridConfiguration

	// Configuration related to PiP layout.
	Pip *PipConfiguration

	noSmithyDocumentSerde
}

// Object describing a participant that has joined a stage.
type Participant struct {

	// Application-provided attributes to encode into the token and attach to a stage.
	// Map keys and values can contain UTF-8 encoded text. The maximum length of this
	// field is 1 KB total. This field is exposed to all stage participants and should
	// not be used for personally identifying, confidential, or sensitive information.
	Attributes map[string]string

	// The participant’s browser.
	BrowserName *string

	// The participant’s browser version.
	BrowserVersion *string

	// ISO 8601 timestamp (returned as a string) when the participant first joined the
	// stage session.
	FirstJoinTime *time.Time

	// The participant’s Internet Service Provider.
	IspName *string

	// The participant’s operating system.
	OsName *string

	// The participant’s operating system version.
	OsVersion *string

	// Unique identifier for this participant, assigned by IVS.
	ParticipantId *string

	// Type of ingest protocol that the participant employs for broadcasting.
	Protocol ParticipantProtocol

	// Whether the participant ever published to the stage session.
	Published bool

	// Name of the S3 bucket to where the participant is being recorded, if individual
	// participant recording is enabled, or "" (empty string), if recording is not
	// enabled.
	RecordingS3BucketName *string

	// S3 prefix of the S3 bucket where the participant is being recorded, if
	// individual participant recording is enabled, or "" (empty string), if recording
	// is not enabled.
	RecordingS3Prefix *string

	// The participant’s recording state.
	RecordingState ParticipantRecordingState

	// The participant’s SDK version.
	SdkVersion *string

	// Whether the participant is connected to or disconnected from the stage.
	State ParticipantState

	// Customer-assigned name to help identify the token; this can be used to link a
	// participant to a user in the customer’s own systems. This can be any UTF-8
	// encoded text. This field is exposed to all stage participants and should not be
	// used for personally identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Summary object describing a participant that has joined a stage.
type ParticipantSummary struct {

	// ISO 8601 timestamp (returned as a string) when the participant first joined the
	// stage session.
	FirstJoinTime *time.Time

	// Unique identifier for this participant, assigned by IVS.
	ParticipantId *string

	// Whether the participant ever published to the stage session.
	Published bool

	// The participant’s recording state.
	RecordingState ParticipantRecordingState

	// Whether the participant is connected to or disconnected from the stage.
	State ParticipantState

	// Customer-assigned name to help identify the token; this can be used to link a
	// participant to a user in the customer’s own systems. This can be any UTF-8
	// encoded text. This field is exposed to all stage participants and should not be
	// used for personally identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Object specifying a participant token in a stage.
//
// Important: Treat tokens as opaque; i.e., do not build functionality based on
// token contents. The format of tokens could change in the future.
type ParticipantToken struct {

	// Application-provided attributes to encode into the token and attach to a stage.
	// This field is exposed to all stage participants and should not be used for
	// personally identifying, confidential, or sensitive information.
	Attributes map[string]string

	// Set of capabilities that the user is allowed to perform in the stage.
	Capabilities []ParticipantTokenCapability

	// Duration (in minutes), after which the participant token expires. Default: 720
	// (12 hours).
	Duration *int32

	// ISO 8601 timestamp (returned as a string) for when this token expires.
	ExpirationTime *time.Time

	// Unique identifier for this participant token, assigned by IVS.
	ParticipantId *string

	// The issued client token, encrypted.
	Token *string

	// Customer-assigned name to help identify the token; this can be used to link a
	// participant to a user in the customer’s own systems. This can be any UTF-8
	// encoded text. This field is exposed to all stage participants and should not be
	// used for personally identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Object specifying a participant token configuration in a stage.
type ParticipantTokenConfiguration struct {

	// Application-provided attributes to encode into the corresponding participant
	// token and attach to a stage. Map keys and values can contain UTF-8 encoded text.
	// The maximum length of this field is 1 KB total. This field is exposed to all
	// stage participants and should not be used for personally identifying,
	// confidential, or sensitive information.
	Attributes map[string]string

	// Set of capabilities that the user is allowed to perform in the stage.
	Capabilities []ParticipantTokenCapability

	// Duration (in minutes), after which the corresponding participant token expires.
	// Default: 720 (12 hours).
	Duration *int32

	// Customer-assigned name to help identify the token; this can be used to link a
	// participant to a user in the customer’s own systems. This can be any UTF-8
	// encoded text. This field is exposed to all stage participants and should not be
	// used for personally identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

// Configuration information specific to Picture-in-Picture (PiP) layout, for [server-side composition].
//
// [server-side composition]: https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/server-side-composition.html
type PipConfiguration struct {

	// This attribute name identifies the featured slot. A participant with this
	// attribute set to "true" (as a string value) in ParticipantTokenConfiguration is placed in the featured slot.
	// Default: "" (no featured participant).
	FeaturedParticipantAttribute *string

	// Specifies the spacing between participant tiles in pixels. Default: 0 .
	GridGap int32

	// Determines whether to omit participants with stopped video in the composition.
	// Default: false .
	OmitStoppedVideo bool

	// Defines PiP behavior when all participants have left: STATIC (maintains
	// original position/size) or DYNAMIC (expands to full composition). Default:
	// STATIC .
	PipBehavior PipBehavior

	// Specifies the height of the PiP window in pixels. When this is not set
	// explicitly, pipHeight ’s value will be based on the size of the composition and
	// the aspect ratio of the participant’s video.
	PipHeight *int32

	// Sets the PiP window’s offset position in pixels from the closest edges
	// determined by PipPosition . Default: 0 .
	PipOffset int32

	// Specifies the participant for the PiP window. A participant with this attribute
	// set to "true" (as a string value) in ParticipantTokenConfiguration is placed in the PiP slot. Default: ""
	// (no PiP participant).
	PipParticipantAttribute *string

	// Determines the corner position of the PiP window. Default: BOTTOM_RIGHT .
	PipPosition PipPosition

	// Specifies the width of the PiP window in pixels. When this is not set
	// explicitly, pipWidth ’s value will be based on the size of the composition and
	// the aspect ratio of the participant’s video.
	PipWidth *int32

	// Defines how video content fits within the participant tile: FILL (stretched),
	// COVER (cropped), or CONTAIN (letterboxed). Default: COVER .
	VideoFillMode VideoFillMode

	noSmithyDocumentSerde
}

// Object specifying a public key used to sign stage participant tokens.
type PublicKey struct {

	// Public key ARN.
	Arn *string

	// The public key fingerprint, a short string used to identify or verify the full
	// public key.
	Fingerprint *string

	// Public key name.
	Name *string

	// Public key material.
	PublicKeyMaterial *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary information about a public key.
type PublicKeySummary struct {

	// Public key ARN.
	Arn *string

	// Public key name.
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	noSmithyDocumentSerde
}

// An object representing a configuration to record a stage stream.
type RecordingConfiguration struct {

	// The recording format for storing a recording in Amazon S3.
	Format RecordingConfigurationFormat

	noSmithyDocumentSerde
}

// A complex type that describes an S3 location where recorded videos will be
// stored.
type S3DestinationConfiguration struct {

	// ARNs of the EncoderConfiguration resource. The encoder configuration and stage resources must be in
	// the same AWS account and region.
	//
	// This member is required.
	EncoderConfigurationArns []string

	// ARN of the StorageConfiguration where recorded videos will be stored.
	//
	// This member is required.
	StorageConfigurationArn *string

	// Array of maps, each of the form string:string (key:value) . This is an optional
	// customer specification, currently used only to specify the recording format for
	// storing a recording in Amazon S3.
	RecordingConfiguration *RecordingConfiguration

	noSmithyDocumentSerde
}

// Complex data type that defines S3Detail objects.
type S3Detail struct {

	// The S3 bucket prefix under which the recording is stored.
	//
	// This member is required.
	RecordingPrefix *string

	noSmithyDocumentSerde
}

// A complex type that describes an S3 location where recorded videos will be
// stored.
type S3StorageConfiguration struct {

	// Location (S3 bucket name) where recorded videos will be stored. Note that the
	// StorageConfiguration and S3 bucket must be in the same region as the
	// Composition.
	//
	// This member is required.
	BucketName *string

	noSmithyDocumentSerde
}

// Object specifying a stage.
type Stage struct {

	// Stage ARN.
	//
	// This member is required.
	Arn *string

	// ID of the active session within the stage.
	ActiveSessionId *string

	// Configuration object for individual participant recording, attached to the
	// stage.
	AutoParticipantRecordingConfiguration *AutoParticipantRecordingConfiguration

	// Summary information about various endpoints for a stage.
	Endpoints *StageEndpoints

	// Stage name.
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary information about various endpoints for a stage. We recommend that you
// cache these values at stage creation; the values can be cached for up to 14
// days.
type StageEndpoints struct {

	// Events endpoint.
	Events *string

	// The endpoint to be used for IVS real-time streaming using the RTMP protocol.
	Rtmp *string

	// The endpoint to be used for IVS real-time streaming using the RTMPS protocol.
	Rtmps *string

	// The endpoint to be used for IVS real-time streaming using the WHIP protocol.
	Whip *string

	noSmithyDocumentSerde
}

// A stage session begins when the first participant joins a stage and ends after
// the last participant leaves the stage. A stage session helps with debugging
// stages by grouping events and participants into shorter periods of time (i.e., a
// session), which is helpful when stages are used over long periods of time.
type StageSession struct {

	// ISO 8601 timestamp (returned as a string) when the stage session ended. This is
	// null if the stage is active.
	EndTime *time.Time

	// ID of the session within the stage.
	SessionId *string

	//  ISO 8601 timestamp (returned as a string) when this stage session began.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Summary information about a stage session.
type StageSessionSummary struct {

	// ISO 8601 timestamp (returned as a string) when the stage session ended. This is
	// null if the stage is active.
	EndTime *time.Time

	// ID of the session within the stage.
	SessionId *string

	//  ISO 8601 timestamp (returned as a string) when this stage session began.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Summary information about a stage.
type StageSummary struct {

	// Stage ARN.
	//
	// This member is required.
	Arn *string

	// ID of the active session within the stage.
	ActiveSessionId *string

	// Stage name.
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	noSmithyDocumentSerde
}

// A complex type that describes a location where recorded videos will be stored.
type StorageConfiguration struct {

	// ARN of the storage configuration.
	//
	// This member is required.
	Arn *string

	// Name of the storage configuration.
	Name *string

	// An S3 destination configuration where recorded videos will be stored.
	S3 *S3StorageConfiguration

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary information about a storage configuration.
type StorageConfigurationSummary struct {

	// ARN of the storage configuration.
	//
	// This member is required.
	Arn *string

	// Name of the storage configuration.
	Name *string

	// An S3 destination configuration where recorded videos will be stored.
	S3 *S3StorageConfiguration

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) . See [Best practices and strategies] in Tagging AWS Resources and Tag Editor for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS has no constraints on tags beyond what is documented
	// there.
	//
	// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
	Tags map[string]string

	noSmithyDocumentSerde
}

// Settings for video.
type Video struct {

	// Bitrate for generated output, in bps. Default: 2500000.
	Bitrate *int32

	// Video frame rate, in fps. Default: 30.
	Framerate *float32

	// Video-resolution height. Note that the maximum value is determined by width
	// times height , such that the maximum total pixels is 2073600 (1920x1080 or
	// 1080x1920). Default: 720.
	Height *int32

	// Video-resolution width. Note that the maximum value is determined by width
	// times height , such that the maximum total pixels is 2073600 (1920x1080 or
	// 1080x1920). Default: 1280.
	Width *int32

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

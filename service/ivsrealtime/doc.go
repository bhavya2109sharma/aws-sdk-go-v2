// Code generated by smithy-go-codegen DO NOT EDIT.

// Package ivsrealtime provides the API client, operations, and parameter types
// for Amazon Interactive Video Service RealTime.
//
// The Amazon Interactive Video Service (IVS) real-time API is REST compatible,
// using a standard HTTP API and an AWS EventBridge event stream for responses.
// JSON is used for both requests and responses, including errors.
//
// Key Concepts
//
//   - Stage — A virtual space where participants can exchange video in real time.
//
//   - Participant token — A token that authenticates a participant when they join
//     a stage.
//
//   - Participant object — Represents participants (people) in the stage and
//     contains information about them. When a token is created, it includes a
//     participant ID; when a participant uses that token to join a stage, the
//     participant is associated with that participant ID. There is a 1:1 mapping
//     between participant tokens and participants.
//
// For server-side composition:
//
//   - Composition process — Composites participants of a stage into a single
//     video and forwards it to a set of outputs (e.g., IVS channels). Composition
//     operations support this process.
//
//   - Composition — Controls the look of the outputs, including how participants
//     are positioned in the video.
//
// For more information about your IVS live stream, also see [Getting Started with Amazon IVS Real-Time Streaming].
//
// # Tagging
//
// A tag is a metadata label that you assign to an AWS resource. A tag comprises a
// key and a value, both set by you. For example, you might set a tag as
// topic:nature to label a particular video category. See [Best practices and strategies] in Tagging AWS
// Resources and Tag Editor for details, including restrictions that apply to tags
// and "Tag naming limits and requirements"; Amazon IVS stages has no
// service-specific constraints beyond what is documented there.
//
// Tags can help you identify and organize your AWS resources. For example, you
// can use the same tag for different resources to indicate that they are related.
// You can also use tags to manage access (see [Access Tags]).
//
// The Amazon IVS real-time API has these tag-related operations: TagResource, UntagResource, and ListTagsForResource. The
// following resource supports tagging: Stage.
//
// At most 50 tags can be applied to a resource.
//
// [Getting Started with Amazon IVS Real-Time Streaming]: https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/getting-started.html
// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
// [Access Tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
package ivsrealtime

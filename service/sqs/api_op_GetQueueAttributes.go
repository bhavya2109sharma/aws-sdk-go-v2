// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets attributes for the specified queue.
//
// To determine whether a queue is [FIFO], you can check whether QueueName ends with the
// .fifo suffix.
//
// [FIFO]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html
func (c *Client) GetQueueAttributes(ctx context.Context, params *GetQueueAttributesInput, optFns ...func(*Options)) (*GetQueueAttributesOutput, error) {
	if params == nil {
		params = &GetQueueAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueueAttributes", params, optFns, c.addOperationGetQueueAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueueAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueueAttributesInput struct {

	// The URL of the Amazon SQS queue whose attribute information is retrieved.
	//
	// Queue URLs and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	// A list of attributes for which to retrieve information.
	//
	// The AttributeNames parameter is optional, but if you don't specify values for
	// this parameter, the request returns empty results.
	//
	// In the future, new attributes might be added. If you write code that calls this
	// action, we recommend that you structure your code so that it can handle new
	// attributes gracefully.
	//
	// The following attributes are supported:
	//
	// The ApproximateNumberOfMessagesDelayed , ApproximateNumberOfMessagesNotVisible ,
	// and ApproximateNumberOfMessages metrics may not achieve consistency until at
	// least 1 minute after the producers stop sending messages. This period is
	// required for the queue metadata to reach eventual consistency.
	//
	//   - All – Returns all values.
	//
	//   - ApproximateNumberOfMessages – Returns the approximate number of messages
	//   available for retrieval from the queue.
	//
	//   - ApproximateNumberOfMessagesDelayed – Returns the approximate number of
	//   messages in the queue that are delayed and not available for reading
	//   immediately. This can happen when the queue is configured as a delay queue or
	//   when a message has been sent with a delay parameter.
	//
	//   - ApproximateNumberOfMessagesNotVisible – Returns the approximate number of
	//   messages that are in flight. Messages are considered to be in flight if they
	//   have been sent to a client but have not yet been deleted or have not yet reached
	//   the end of their visibility window.
	//
	//   - CreatedTimestamp – Returns the time when the queue was created in seconds ([epoch time]
	//   ).
	//
	//   - DelaySeconds – Returns the default delay on the queue in seconds.
	//
	//   - LastModifiedTimestamp – Returns the time when the queue was last changed in
	//   seconds ([epoch time] ).
	//
	//   - MaximumMessageSize – Returns the limit of how many bytes a message can
	//   contain before Amazon SQS rejects it.
	//
	//   - MessageRetentionPeriod – Returns the length of time, in seconds, for which
	//   Amazon SQS retains a message. When you change a queue's attributes, the change
	//   can take up to 60 seconds for most of the attributes to propagate throughout the
	//   Amazon SQS system. Changes made to the MessageRetentionPeriod attribute can
	//   take up to 15 minutes and will impact existing messages in the queue potentially
	//   causing them to be expired and deleted if the MessageRetentionPeriod is
	//   reduced below the age of existing messages.
	//
	//   - Policy – Returns the policy of the queue.
	//
	//   - QueueArn – Returns the Amazon resource name (ARN) of the queue.
	//
	//   - ReceiveMessageWaitTimeSeconds – Returns the length of time, in seconds, for
	//   which the ReceiveMessage action waits for a message to arrive.
	//
	//   - VisibilityTimeout – Returns the visibility timeout for the queue. For more
	//   information about the visibility timeout, see [Visibility Timeout]in the Amazon SQS Developer
	//   Guide.
	//
	// The following attributes apply only to [dead-letter queues:]
	//
	//   - RedrivePolicy – The string that includes the parameters for the dead-letter
	//   queue functionality of the source queue as a JSON object. The parameters are as
	//   follows:
	//
	//   - deadLetterTargetArn – The Amazon Resource Name (ARN) of the dead-letter
	//   queue to which Amazon SQS moves messages after the value of maxReceiveCount is
	//   exceeded.
	//
	//   - maxReceiveCount – The number of times a message is delivered to the source
	//   queue before being moved to the dead-letter queue. Default: 10. When the
	//   ReceiveCount for a message exceeds the maxReceiveCount for a queue, Amazon SQS
	//   moves the message to the dead-letter-queue.
	//
	//   - RedriveAllowPolicy – The string that includes the parameters for the
	//   permissions for the dead-letter queue redrive permission and which source queues
	//   can specify dead-letter queues as a JSON object. The parameters are as follows:
	//
	//   - redrivePermission – The permission type that defines which source queues can
	//   specify the current queue as the dead-letter queue. Valid values are:
	//
	//   - allowAll – (Default) Any source queues in this Amazon Web Services account
	//   in the same Region can specify this queue as the dead-letter queue.
	//
	//   - denyAll – No source queues can specify this queue as the dead-letter queue.
	//
	//   - byQueue – Only queues specified by the sourceQueueArns parameter can specify
	//   this queue as the dead-letter queue.
	//
	//   - sourceQueueArns – The Amazon Resource Names (ARN)s of the source queues that
	//   can specify this queue as the dead-letter queue and redrive messages. You can
	//   specify this parameter only when the redrivePermission parameter is set to
	//   byQueue . You can specify up to 10 source queue ARNs. To allow more than 10
	//   source queues to specify dead-letter queues, set the redrivePermission
	//   parameter to allowAll .
	//
	// The dead-letter queue of a FIFO queue must also be a FIFO queue. Similarly, the
	// dead-letter queue of a standard queue must also be a standard queue.
	//
	// The following attributes apply only to [server-side-encryption]:
	//
	//   - KmsMasterKeyId – Returns the ID of an Amazon Web Services managed customer
	//   master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms].
	//
	//   - KmsDataKeyReusePeriodSeconds – Returns the length of time, in seconds, for
	//   which Amazon SQS can reuse a data key to encrypt or decrypt messages before
	//   calling KMS again. For more information, see [How Does the Data Key Reuse Period Work?].
	//
	//   - SqsManagedSseEnabled – Returns information about whether the queue is using
	//   SSE-SQS encryption using SQS owned encryption keys. Only one server-side
	//   encryption option is supported per queue (for example, [SSE-KMS]or [SSE-SQS]).
	//
	// The following attributes apply only to [FIFO (first-in-first-out) queues]:
	//
	//   - FifoQueue – Returns information about whether the queue is FIFO. For more
	//   information, see [FIFO queue logic]in the Amazon SQS Developer Guide.
	//
	// To determine whether a queue is [FIFO], you can check whether QueueName ends with the
	//   .fifo suffix.
	//
	//   - ContentBasedDeduplication – Returns whether content-based deduplication is
	//   enabled for the queue. For more information, see [Exactly-once processing]in the Amazon SQS Developer
	//   Guide.
	//
	// The following attributes apply only to [high throughput for FIFO queues]:
	//
	//   - DeduplicationScope – Specifies whether message deduplication occurs at the
	//   message group or queue level. Valid values are messageGroup and queue .
	//
	//   - FifoThroughputLimit – Specifies whether the FIFO queue throughput quota
	//   applies to the entire queue or per message group. Valid values are perQueue
	//   and perMessageGroupId . The perMessageGroupId value is allowed only when the
	//   value for DeduplicationScope is messageGroup .
	//
	// To enable high throughput for FIFO queues, do the following:
	//
	//   - Set DeduplicationScope to messageGroup .
	//
	//   - Set FifoThroughputLimit to perMessageGroupId .
	//
	// If you set these attributes to anything other than the values shown for
	// enabling high throughput, normal throughput is in effect and deduplication
	// occurs as specified.
	//
	// For information on throughput quotas, see [Quotas related to messages] in the Amazon SQS Developer Guide.
	//
	// [SSE-KMS]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html
	// [epoch time]: http://en.wikipedia.org/wiki/Unix_time
	// [How Does the Data Key Reuse Period Work?]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work
	// [SSE-SQS]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html
	// [high throughput for FIFO queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html
	// [dead-letter queues:]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html
	// [Exactly-once processing]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-exactly-once-processing.html
	// [FIFO]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html
	// [Quotas related to messages]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html
	// [Visibility Timeout]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html
	// [Key Terms]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms
	// [server-side-encryption]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html
	// [FIFO queue logic]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-understanding-logic.html
	// [FIFO (first-in-first-out) queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html
	AttributeNames []types.QueueAttributeName

	noSmithyDocumentSerde
}

// A list of returned queue attributes.
type GetQueueAttributesOutput struct {

	// A map of attributes to their respective values.
	Attributes map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQueueAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetQueueAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetQueueAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetQueueAttributes"); err != nil {
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
	if err = addOpGetQueueAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueueAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetQueueAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetQueueAttributes",
	}
}

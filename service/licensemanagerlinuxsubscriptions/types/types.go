// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A filter object that is used to return more specific results from a describe
// operation. Filters can be used to match a set of resources by specific criteria.
type Filter struct {

	// The type of name to filter by.
	Name *string

	// An operator for filtering results.
	Operator Operator

	// One or more values for the name to filter by.
	Values []string

	noSmithyDocumentSerde
}

// Details discovered information about a running instance using Linux
// subscriptions.
type Instance struct {

	// The account ID which owns the instance.
	AccountID *string

	// The AMI ID used to launch the instance.
	AmiId *string

	// Indicates that you have two different license subscriptions for the same
	// software on your instance.
	DualSubscription *string

	// The instance ID of the resource.
	InstanceID *string

	// The instance type of the resource.
	InstanceType *string

	// The time in which the last discovery updated the instance details.
	LastUpdatedTime *string

	// The operating system software version that runs on your instance.
	OsVersion *string

	// The product code for the instance. For more information, see [Usage operation values] in the License
	// Manager User Guide .
	//
	// [Usage operation values]: https://docs.aws.amazon.com/license-manager/latest/userguide/linux-subscriptions-usage-operation.html
	ProductCode []string

	// The Region the instance is running in.
	Region *string

	// Indicates that your instance uses a BYOL license subscription from a
	// third-party Linux subscription provider that you've registered with License
	// Manager.
	RegisteredWithSubscriptionProvider *string

	// The status of the instance.
	Status *string

	// The name of the license subscription that the instance uses.
	SubscriptionName *string

	// The timestamp when you registered the third-party Linux subscription provider
	// for the subscription that the instance uses.
	SubscriptionProviderCreateTime *string

	// The timestamp from the last time that the instance synced with the registered
	// third-party Linux subscription provider.
	SubscriptionProviderUpdateTime *string

	// The usage operation of the instance. For more information, see For more
	// information, see [Usage operation values]in the License Manager User Guide.
	//
	// [Usage operation values]: https://docs.aws.amazon.com/license-manager/latest/userguide/linux-subscriptions-usage-operation.html
	UsageOperation *string

	noSmithyDocumentSerde
}

// Lists the settings defined for discovering Linux subscriptions.
type LinuxSubscriptionsDiscoverySettings struct {

	// Details if you have enabled resource discovery across your accounts in
	// Organizations.
	//
	// This member is required.
	OrganizationIntegration OrganizationIntegration

	// The Regions in which to discover data for Linux subscriptions.
	//
	// This member is required.
	SourceRegions []string

	noSmithyDocumentSerde
}

// A third-party provider for operating system (OS) platform software and license
// subscriptions, such as Red Hat. When you register a third-party Linux
// subscription provider, License Manager can get subscription data from the
// registered provider.
type RegisteredSubscriptionProvider struct {

	// The timestamp from the last time that License Manager accessed third-party
	// subscription data for your account from your registered Linux subscription
	// provider.
	LastSuccessfulDataRetrievalTime *string

	// The Amazon Resource Name (ARN) of the Secrets Manager secret that stores your
	// registered Linux subscription provider access token. For RHEL account
	// subscriptions, this is the offline token.
	SecretArn *string

	// The Amazon Resource Name (ARN) of the Linux subscription provider resource that
	// you registered.
	SubscriptionProviderArn *string

	// A supported third-party Linux subscription provider. License Manager currently
	// supports Red Hat subscriptions.
	SubscriptionProviderSource SubscriptionProviderSource

	// Indicates the status of your registered Linux subscription provider access
	// token from the last time License Manager retrieved subscription data. For RHEL
	// account subscriptions, this is the status of the offline token.
	SubscriptionProviderStatus SubscriptionProviderStatus

	// A detailed message that's associated with your BYOL subscription provider token
	// status.
	SubscriptionProviderStatusMessage *string

	noSmithyDocumentSerde
}

// An object which details a discovered Linux subscription.
type Subscription struct {

	// The total amount of running instances using this subscription.
	InstanceCount *int64

	// The name of the subscription.
	Name *string

	// The type of subscription. The type can be subscription-included with Amazon
	// EC2, Bring Your Own Subscription model (BYOS), or from the Amazon Web Services
	// Marketplace. Certain subscriptions may use licensing from the Amazon Web
	// Services Marketplace as well as OS licensing from Amazon EC2 or BYOS.
	Type *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

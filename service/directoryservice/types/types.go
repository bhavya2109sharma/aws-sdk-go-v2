// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Represents a named directory attribute.
type Attribute struct {

	// The name of the attribute.
	Name *string

	// The value of the attribute.
	Value *string

	noSmithyDocumentSerde
}

// Information about the certificate.
type Certificate struct {

	// The identifier of the certificate.
	CertificateId *string

	// A ClientCertAuthSettings object that contains client certificate authentication
	// settings.
	ClientCertAuthSettings *ClientCertAuthSettings

	// The common name for the certificate.
	CommonName *string

	// The date and time when the certificate will expire.
	ExpiryDateTime *time.Time

	// The date and time that the certificate was registered.
	RegisteredDateTime *time.Time

	// The state of the certificate.
	State CertificateState

	// Describes a state change for the certificate.
	StateReason *string

	// The function that the registered certificate performs. Valid values include
	// ClientLDAPS or ClientCertAuth . The default value is ClientLDAPS .
	Type CertificateType

	noSmithyDocumentSerde
}

// Contains general information about a certificate.
type CertificateInfo struct {

	// The identifier of the certificate.
	CertificateId *string

	// The common name for the certificate.
	CommonName *string

	// The date and time when the certificate will expire.
	ExpiryDateTime *time.Time

	// The state of the certificate.
	State CertificateState

	// The function that the registered certificate performs. Valid values include
	// ClientLDAPS or ClientCertAuth . The default value is ClientLDAPS .
	Type CertificateType

	noSmithyDocumentSerde
}

// Contains information about a client authentication method for a directory.
type ClientAuthenticationSettingInfo struct {

	// The date and time when the status of the client authentication type was last
	// updated.
	LastUpdatedDateTime *time.Time

	// Whether the client authentication type is enabled or disabled for the specified
	// directory.
	Status ClientAuthenticationStatus

	// The type of client authentication for the specified directory. If no type is
	// specified, a list of all client authentication types that are supported for the
	// directory is retrieved.
	Type ClientAuthenticationType

	noSmithyDocumentSerde
}

// Contains information about the client certificate authentication settings for
// the RegisterCertificate and DescribeCertificate operations.
type ClientCertAuthSettings struct {

	// Specifies the URL of the default OCSP server used to check for revocation
	// status. A secondary value to any OCSP address found in the AIA extension of the
	// user certificate.
	OCSPUrl *string

	noSmithyDocumentSerde
}

// Contains information about a computer account in a directory.
type Computer struct {

	// An array of Attribute objects containing the LDAP attributes that belong to the
	// computer account.
	ComputerAttributes []Attribute

	// The identifier of the computer.
	ComputerId *string

	// The computer name.
	ComputerName *string

	noSmithyDocumentSerde
}

// Points to a remote domain with which you are setting up a trust relationship.
// Conditional forwarders are required in order to set up a trust relationship with
// another domain.
type ConditionalForwarder struct {

	// The IP addresses of the remote DNS server associated with RemoteDomainName.
	// This is the IP address of the DNS server that your conditional forwarder points
	// to.
	DnsIpAddrs []string

	// The fully qualified domain name (FQDN) of the remote domains pointed to by the
	// conditional forwarder.
	RemoteDomainName *string

	// The replication scope of the conditional forwarder. The only allowed value is
	// Domain , which will replicate the conditional forwarder to all of the domain
	// controllers for your Amazon Web Services directory.
	ReplicationScope ReplicationScope

	noSmithyDocumentSerde
}

// Contains information for the ConnectDirectory operation when an AD Connector
// directory is being created.
type DirectoryConnectSettings struct {

	// A list of one or more IP addresses of DNS servers or domain controllers in your
	// self-managed directory.
	//
	// This member is required.
	CustomerDnsIps []string

	// The user name of an account in your self-managed directory that is used to
	// connect to the directory. This account must have the following permissions:
	//   - Read users and groups
	//   - Create computer objects
	//   - Join computers to the domain
	//
	// This member is required.
	CustomerUserName *string

	// A list of subnet identifiers in the VPC in which the AD Connector is created.
	//
	// This member is required.
	SubnetIds []string

	// The identifier of the VPC in which the AD Connector is created.
	//
	// This member is required.
	VpcId *string

	noSmithyDocumentSerde
}

// Contains information about an AD Connector directory.
type DirectoryConnectSettingsDescription struct {

	// A list of the Availability Zones that the directory is in.
	AvailabilityZones []string

	// The IP addresses of the AD Connector servers.
	ConnectIps []string

	// The user name of the service account in your self-managed directory.
	CustomerUserName *string

	// The security group identifier for the AD Connector directory.
	SecurityGroupId *string

	// A list of subnet identifiers in the VPC that the AD Connector is in.
	SubnetIds []string

	// The identifier of the VPC that the AD Connector is in.
	VpcId *string

	noSmithyDocumentSerde
}

// Contains information about an Directory Service directory.
type DirectoryDescription struct {

	// The access URL for the directory, such as http://.awsapps.com . If no alias has
	// been created for the directory, is the directory identifier, such as
	// d-XXXXXXXXXX .
	AccessUrl *string

	// The alias for the directory. If no alias has been created for the directory,
	// the alias is the directory identifier, such as d-XXXXXXXXXX .
	Alias *string

	// A DirectoryConnectSettingsDescription object that contains additional
	// information about an AD Connector directory. This member is only present if the
	// directory is an AD Connector directory.
	ConnectSettings *DirectoryConnectSettingsDescription

	// The description for the directory.
	Description *string

	// The desired number of domain controllers in the directory if the directory is
	// Microsoft AD.
	DesiredNumberOfDomainControllers *int32

	// The directory identifier.
	DirectoryId *string

	// The IP addresses of the DNS servers for the directory. For a Simple AD or
	// Microsoft AD directory, these are the IP addresses of the Simple AD or Microsoft
	// AD directory servers. For an AD Connector directory, these are the IP addresses
	// of the DNS servers or domain controllers in your self-managed directory to which
	// the AD Connector is connected.
	DnsIpAddrs []string

	// The edition associated with this directory.
	Edition DirectoryEdition

	// Specifies when the directory was created.
	LaunchTime *time.Time

	// The fully qualified name of the directory.
	Name *string

	// The operating system (OS) version of the directory.
	OsVersion OSVersion

	// Describes the Managed Microsoft AD directory in the directory owner account.
	OwnerDirectoryDescription *OwnerDirectoryDescription

	// A RadiusSettings object that contains information about the RADIUS server
	// configured for this directory.
	RadiusSettings *RadiusSettings

	// The status of the RADIUS MFA server connection.
	RadiusStatus RadiusStatus

	// Lists the Regions where the directory has replicated.
	RegionsInfo *RegionsInfo

	// The method used when sharing a directory to determine whether the directory
	// should be shared within your Amazon Web Services organization ( ORGANIZATIONS )
	// or with any Amazon Web Services account by sending a shared directory request (
	// HANDSHAKE ).
	ShareMethod ShareMethod

	// A directory share request that is sent by the directory owner to the directory
	// consumer. The request includes a typed message to help the directory consumer
	// administrator determine whether to approve or reject the share invitation.
	ShareNotes *string

	// Current directory status of the shared Managed Microsoft AD directory.
	ShareStatus ShareStatus

	// The short name of the directory.
	ShortName *string

	// The directory size.
	Size DirectorySize

	// Indicates if single sign-on is enabled for the directory. For more information,
	// see EnableSso and DisableSso .
	SsoEnabled bool

	// The current stage of the directory.
	Stage DirectoryStage

	// The date and time that the stage was last updated.
	StageLastUpdatedDateTime *time.Time

	// Additional information about the directory stage.
	StageReason *string

	// The directory size.
	Type DirectoryType

	// A DirectoryVpcSettingsDescription object that contains additional information
	// about a directory. This member is only present if the directory is a Simple AD
	// or Managed Microsoft AD directory.
	VpcSettings *DirectoryVpcSettingsDescription

	noSmithyDocumentSerde
}

// Contains directory limit information for a Region.
type DirectoryLimits struct {

	// The current number of cloud directories in the Region.
	CloudOnlyDirectoriesCurrentCount *int32

	// The maximum number of cloud directories allowed in the Region.
	CloudOnlyDirectoriesLimit *int32

	// Indicates if the cloud directory limit has been reached.
	CloudOnlyDirectoriesLimitReached bool

	// The current number of Managed Microsoft AD directories in the region.
	CloudOnlyMicrosoftADCurrentCount *int32

	// The maximum number of Managed Microsoft AD directories allowed in the region.
	CloudOnlyMicrosoftADLimit *int32

	// Indicates if the Managed Microsoft AD directory limit has been reached.
	CloudOnlyMicrosoftADLimitReached bool

	// The current number of connected directories in the Region.
	ConnectedDirectoriesCurrentCount *int32

	// The maximum number of connected directories allowed in the Region.
	ConnectedDirectoriesLimit *int32

	// Indicates if the connected directory limit has been reached.
	ConnectedDirectoriesLimitReached bool

	noSmithyDocumentSerde
}

// Contains VPC information for the CreateDirectory or CreateMicrosoftAD operation.
type DirectoryVpcSettings struct {

	// The identifiers of the subnets for the directory servers. The two subnets must
	// be in different Availability Zones. Directory Service creates a directory server
	// and a DNS server in each of these subnets.
	//
	// This member is required.
	SubnetIds []string

	// The identifier of the VPC in which to create the directory.
	//
	// This member is required.
	VpcId *string

	noSmithyDocumentSerde
}

// Contains information about the directory.
type DirectoryVpcSettingsDescription struct {

	// The list of Availability Zones that the directory is in.
	AvailabilityZones []string

	// The domain controller security group identifier for the directory.
	SecurityGroupId *string

	// The identifiers of the subnets for the directory servers.
	SubnetIds []string

	// The identifier of the VPC that the directory is in.
	VpcId *string

	noSmithyDocumentSerde
}

// Contains information about the domain controllers for a specified directory.
type DomainController struct {

	// The Availability Zone where the domain controller is located.
	AvailabilityZone *string

	// Identifier of the directory where the domain controller resides.
	DirectoryId *string

	// The IP address of the domain controller.
	DnsIpAddr *string

	// Identifies a specific domain controller in the directory.
	DomainControllerId *string

	// Specifies when the domain controller was created.
	LaunchTime *time.Time

	// The status of the domain controller.
	Status DomainControllerStatus

	// The date and time that the status was last updated.
	StatusLastUpdatedDateTime *time.Time

	// A description of the domain controller state.
	StatusReason *string

	// Identifier of the subnet in the VPC that contains the domain controller.
	SubnetId *string

	// The identifier of the VPC that contains the domain controller.
	VpcId *string

	noSmithyDocumentSerde
}

// Information about Amazon SNS topic and Directory Service directory associations.
type EventTopic struct {

	// The date and time of when you associated your directory with the Amazon SNS
	// topic.
	CreatedDateTime *time.Time

	// The Directory ID of an Directory Service directory that will publish status
	// messages to an Amazon SNS topic.
	DirectoryId *string

	// The topic registration status.
	Status TopicStatus

	// The Amazon SNS topic ARN (Amazon Resource Name).
	TopicArn *string

	// The name of an Amazon SNS topic the receives status messages from the directory.
	TopicName *string

	noSmithyDocumentSerde
}

// IP address block. This is often the address block of the DNS server used for
// your self-managed domain.
type IpRoute struct {

	// IP address block using CIDR format, for example 10.0.0.0/24. This is often the
	// address block of the DNS server used for your self-managed domain. For a single
	// IP address use a CIDR address block with /32. For example 10.0.0.0/32.
	CidrIp *string

	// Description of the address block.
	Description *string

	noSmithyDocumentSerde
}

// Information about one or more IP address blocks.
type IpRouteInfo struct {

	// The date and time the address block was added to the directory.
	AddedDateTime *time.Time

	// IP address block in the IpRoute .
	CidrIp *string

	// Description of the IpRouteInfo .
	Description *string

	// Identifier (ID) of the directory associated with the IP addresses.
	DirectoryId *string

	// The status of the IP address block.
	IpRouteStatusMsg IpRouteStatusMsg

	// The reason for the IpRouteStatusMsg.
	IpRouteStatusReason *string

	noSmithyDocumentSerde
}

// Contains general information about the LDAPS settings.
type LDAPSSettingInfo struct {

	// The state of the LDAPS settings.
	LDAPSStatus LDAPSStatus

	// Describes a state change for LDAPS.
	LDAPSStatusReason *string

	// The date and time when the LDAPS settings were last updated.
	LastUpdatedDateTime *time.Time

	noSmithyDocumentSerde
}

// Represents a log subscription, which tracks real-time data from a chosen log
// group to a specified destination.
type LogSubscription struct {

	// Identifier (ID) of the directory that you want to associate with the log
	// subscription.
	DirectoryId *string

	// The name of the log group.
	LogGroupName *string

	// The date and time that the log subscription was created.
	SubscriptionCreatedDateTime *time.Time

	noSmithyDocumentSerde
}

// OS version that the directory needs to be updated to.
type OSUpdateSettings struct {

	// OS version that the directory needs to be updated to.
	OSVersion OSVersion

	noSmithyDocumentSerde
}

// Describes the directory owner account details that have been shared to the
// directory consumer account.
type OwnerDirectoryDescription struct {

	// Identifier of the directory owner account.
	AccountId *string

	// Identifier of the Managed Microsoft AD directory in the directory owner account.
	DirectoryId *string

	// IP address of the directory’s domain controllers.
	DnsIpAddrs []string

	// A RadiusSettings object that contains information about the RADIUS server.
	RadiusSettings *RadiusSettings

	// Information about the status of the RADIUS server.
	RadiusStatus RadiusStatus

	// Information about the VPC settings for the directory.
	VpcSettings *DirectoryVpcSettingsDescription

	noSmithyDocumentSerde
}

// Contains information about a Remote Authentication Dial In User Service
// (RADIUS) server.
type RadiusSettings struct {

	// The protocol specified for your RADIUS endpoints.
	AuthenticationProtocol RadiusAuthenticationProtocol

	// Not currently used.
	DisplayLabel *string

	// The port that your RADIUS server is using for communications. Your self-managed
	// network must allow inbound traffic over this port from the Directory Service
	// servers.
	RadiusPort *int32

	// The maximum number of times that communication with the RADIUS server is
	// attempted.
	RadiusRetries int32

	// An array of strings that contains the fully qualified domain name (FQDN) or IP
	// addresses of the RADIUS server endpoints, or the FQDN or IP addresses of your
	// RADIUS server load balancer.
	RadiusServers []string

	// The amount of time, in seconds, to wait for the RADIUS server to respond.
	RadiusTimeout *int32

	// Required for enabling RADIUS on the directory.
	SharedSecret *string

	// Not currently used.
	UseSameUsername bool

	noSmithyDocumentSerde
}

// The replicated Region information for a directory.
type RegionDescription struct {

	// The desired number of domain controllers in the specified Region for the
	// specified directory.
	DesiredNumberOfDomainControllers *int32

	// The identifier of the directory.
	DirectoryId *string

	// The date and time that the Region description was last updated.
	LastUpdatedDateTime *time.Time

	// Specifies when the Region replication began.
	LaunchTime *time.Time

	// The name of the Region. For example, us-east-1 .
	RegionName *string

	// Specifies whether the Region is the primary Region or an additional Region.
	RegionType RegionType

	// The status of the replication process for the specified Region.
	Status DirectoryStage

	// The date and time that the Region status was last updated.
	StatusLastUpdatedDateTime *time.Time

	// Contains VPC information for the CreateDirectory or CreateMicrosoftAD operation.
	VpcSettings *DirectoryVpcSettings

	noSmithyDocumentSerde
}

// Provides information about the Regions that are configured for multi-Region
// replication.
type RegionsInfo struct {

	// Lists the Regions where the directory has been replicated, excluding the
	// primary Region.
	AdditionalRegions []string

	// The Region where the Managed Microsoft AD directory was originally created.
	PrimaryRegion *string

	noSmithyDocumentSerde
}

// Information about a schema extension.
type SchemaExtensionInfo struct {

	// A description of the schema extension.
	Description *string

	// The identifier of the directory to which the schema extension is applied.
	DirectoryId *string

	// The date and time that the schema extension was completed.
	EndDateTime *time.Time

	// The identifier of the schema extension.
	SchemaExtensionId *string

	// The current status of the schema extension.
	SchemaExtensionStatus SchemaExtensionStatus

	// The reason for the SchemaExtensionStatus .
	SchemaExtensionStatusReason *string

	// The date and time that the schema extension started being applied to the
	// directory.
	StartDateTime *time.Time

	noSmithyDocumentSerde
}

// Contains information about the configurable settings for a directory.
type Setting struct {

	// The name of the directory setting. For example: TLS_1_0
	//
	// This member is required.
	Name *string

	// The value of the directory setting for which to retrieve information. For
	// example, for TLS_1_0 , the valid values are: Enable and Disable .
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Contains information about the specified configurable setting for a directory.
type SettingEntry struct {

	// The valid range of values for the directory setting. These values depend on the
	// DataType of your directory.
	AllowedValues *string

	// The value of the directory setting that is applied to the directory.
	AppliedValue *string

	// The data type of a directory setting. This is used to define the AllowedValues
	// of a setting. For example a data type can be Boolean , DurationInSeconds , or
	// Enum .
	DataType *string

	// The date and time when the request to update a directory setting was last
	// submitted.
	LastRequestedDateTime *time.Time

	// The date and time when the directory setting was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the directory setting. For example: TLS_1_0
	Name *string

	// Details about the status of the request to update the directory setting. If the
	// directory setting is deployed in more than one region, status is returned for
	// the request in each region where the setting is deployed.
	RequestDetailedStatus map[string]DirectoryConfigurationStatus

	// The overall status of the request to update the directory setting request. If
	// the directory setting is deployed in more than one region, and the request fails
	// in any region, the overall status is Failed .
	RequestStatus DirectoryConfigurationStatus

	// The last status message for the directory status request.
	RequestStatusMessage *string

	// The value that was last requested for the directory setting.
	RequestedValue *string

	// The type, or category, of a directory setting. Similar settings have the same
	// type. For example, Protocol , Cipher , or Certificate-Based Authentication .
	Type *string

	noSmithyDocumentSerde
}

// Details about the shared directory in the directory owner account for which the
// share request in the directory consumer account has been accepted.
type SharedDirectory struct {

	// The date and time that the shared directory was created.
	CreatedDateTime *time.Time

	// The date and time that the shared directory was last updated.
	LastUpdatedDateTime *time.Time

	// Identifier of the directory owner account, which contains the directory that
	// has been shared to the consumer account.
	OwnerAccountId *string

	// Identifier of the directory in the directory owner account.
	OwnerDirectoryId *string

	// The method used when sharing a directory to determine whether the directory
	// should be shared within your Amazon Web Services organization ( ORGANIZATIONS )
	// or with any Amazon Web Services account by sending a shared directory request (
	// HANDSHAKE ).
	ShareMethod ShareMethod

	// A directory share request that is sent by the directory owner to the directory
	// consumer. The request includes a typed message to help the directory consumer
	// administrator determine whether to approve or reject the share invitation.
	ShareNotes *string

	// Current directory status of the shared Managed Microsoft AD directory.
	ShareStatus ShareStatus

	// Identifier of the directory consumer account that has access to the shared
	// directory ( OwnerDirectoryId ) in the directory owner account.
	SharedAccountId *string

	// Identifier of the shared directory in the directory consumer account. This
	// identifier is different for each directory owner account.
	SharedDirectoryId *string

	noSmithyDocumentSerde
}

// Identifier that contains details about the directory consumer account.
type ShareTarget struct {

	// Identifier of the directory consumer account.
	//
	// This member is required.
	Id *string

	// Type of identifier to be used in the Id field.
	//
	// This member is required.
	Type TargetType

	noSmithyDocumentSerde
}

// Describes a directory snapshot.
type Snapshot struct {

	// The directory identifier.
	DirectoryId *string

	// The descriptive name of the snapshot.
	Name *string

	// The snapshot identifier.
	SnapshotId *string

	// The date and time that the snapshot was taken.
	StartTime *time.Time

	// The snapshot status.
	Status SnapshotStatus

	// The snapshot type.
	Type SnapshotType

	noSmithyDocumentSerde
}

// Contains manual snapshot limit information for a directory.
type SnapshotLimits struct {

	// The current number of manual snapshots of the directory.
	ManualSnapshotsCurrentCount *int32

	// The maximum number of manual snapshots allowed.
	ManualSnapshotsLimit *int32

	// Indicates if the manual snapshot limit has been reached.
	ManualSnapshotsLimitReached bool

	noSmithyDocumentSerde
}

// Metadata assigned to a directory consisting of a key-value pair.
type Tag struct {

	// Required name of the tag. The string value can be Unicode characters and cannot
	// be prefixed with "aws:". The string can contain only the set of Unicode letters,
	// digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex:
	// "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	//
	// This member is required.
	Key *string

	// The optional value of the tag. The string value can be Unicode characters. The
	// string can contain only the set of Unicode letters, digits, white-space, '_',
	// '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Describes a trust relationship between an Managed Microsoft AD directory and an
// external domain.
type Trust struct {

	// The date and time that the trust relationship was created.
	CreatedDateTime *time.Time

	// The Directory ID of the Amazon Web Services directory involved in the trust
	// relationship.
	DirectoryId *string

	// The date and time that the trust relationship was last updated.
	LastUpdatedDateTime *time.Time

	// The Fully Qualified Domain Name (FQDN) of the external domain involved in the
	// trust relationship.
	RemoteDomainName *string

	// Current state of selective authentication for the trust.
	SelectiveAuth SelectiveAuth

	// The date and time that the TrustState was last updated.
	StateLastUpdatedDateTime *time.Time

	// The trust relationship direction.
	TrustDirection TrustDirection

	// The unique ID of the trust relationship.
	TrustId *string

	// The trust relationship state.
	TrustState TrustState

	// The reason for the TrustState.
	TrustStateReason *string

	// The trust relationship type. Forest is the default.
	TrustType TrustType

	noSmithyDocumentSerde
}

// Identifier that contains details about the directory consumer account with whom
// the directory is being unshared.
type UnshareTarget struct {

	// Identifier of the directory consumer account.
	//
	// This member is required.
	Id *string

	// Type of identifier to be used in the Id field.
	//
	// This member is required.
	Type TargetType

	noSmithyDocumentSerde
}

// An entry of update information related to a requested update type.
type UpdateInfoEntry struct {

	// This specifies if the update was initiated by the customer or by the service
	// team.
	InitiatedBy *string

	// The last updated date and time of a particular directory setting.
	LastUpdatedDateTime *time.Time

	// The new value of the target setting.
	NewValue *UpdateValue

	// The old value of the target setting.
	PreviousValue *UpdateValue

	// The name of the Region.
	Region *string

	// The start time of the UpdateDirectorySetup for the particular type.
	StartTime *time.Time

	// The status of the update performed on the directory.
	Status UpdateStatus

	// The reason for the current status of the update type activity.
	StatusReason *string

	noSmithyDocumentSerde
}

// The value for a given type of UpdateSettings .
type UpdateValue struct {

	// The OS update related settings.
	OSUpdateSettings *OSUpdateSettings

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

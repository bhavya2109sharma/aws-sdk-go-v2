// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An accelerator is a complex type that includes one or more listeners that
// process inbound connections and then direct traffic to one or more endpoint
// groups, each of which includes endpoints, such as load balancers.
type Accelerator struct {

	// The Amazon Resource Name (ARN) of the accelerator.
	AcceleratorArn *string

	// The date and time that the accelerator was created.
	CreatedTime *time.Time

	// The Domain Name System (DNS) name that Global Accelerator creates that points
	// to an accelerator's static IPv4 addresses.
	//
	// The naming convention for the DNS name for an accelerator is the following: A
	// lowercase letter a, followed by a 16-bit random hex string, followed by
	// .awsglobalaccelerator.com. For example:
	// a1234567890abcdef.awsglobalaccelerator.com.
	//
	// If you have a dual-stack accelerator, you also have a second DNS name,
	// DualStackDnsName , that points to both the A record and the AAAA record for all
	// four static addresses for the accelerator: two IPv4 addresses and two IPv6
	// addresses.
	//
	// For more information about the default DNS name, see [Support for DNS addressing in Global Accelerator] in the Global Accelerator
	// Developer Guide.
	//
	// [Support for DNS addressing in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/dns-addressing-custom-domains.dns-addressing.html
	DnsName *string

	// The Domain Name System (DNS) name that Global Accelerator creates that points
	// to a dual-stack accelerator's four static IP addresses: two IPv4 addresses and
	// two IPv6 addresses.
	//
	// The naming convention for the dual-stack DNS name is the following: A lowercase
	// letter a, followed by a 16-bit random hex string, followed by
	// .dualstack.awsglobalaccelerator.com. For example:
	// a1234567890abcdef.dualstack.awsglobalaccelerator.com.
	//
	// Note: Global Accelerator also assigns a default DNS name, DnsName , to your
	// accelerator that points just to the static IPv4 addresses.
	//
	// For more information, see [Support for DNS addressing in Global Accelerator] in the Global Accelerator Developer Guide.
	//
	// [Support for DNS addressing in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-accelerators.html#about-accelerators.dns-addressing
	DualStackDnsName *string

	// Indicates whether the accelerator is enabled. The value is true or false. The
	// default value is true.
	//
	// If the value is set to true, the accelerator cannot be deleted. If set to
	// false, accelerator can be deleted.
	Enabled *bool

	// A history of changes that you make to an accelerator in Global Accelerator.
	Events []AcceleratorEvent

	// The IP address type that an accelerator supports. For a standard accelerator,
	// the value can be IPV4 or DUAL_STACK.
	IpAddressType IpAddressType

	// The static IP addresses that Global Accelerator associates with the accelerator.
	IpSets []IpSet

	// The date and time that the accelerator was last modified.
	LastModifiedTime *time.Time

	// The name of the accelerator. The name must contain only alphanumeric characters
	// or hyphens (-), and must not begin or end with a hyphen.
	Name *string

	// Describes the deployment status of the accelerator.
	Status AcceleratorStatus

	noSmithyDocumentSerde
}

// Attributes of an accelerator.
type AcceleratorAttributes struct {

	// Indicates whether flow logs are enabled. The default value is false. If the
	// value is true, FlowLogsS3Bucket and FlowLogsS3Prefix must be specified.
	//
	// For more information, see [Flow logs] in the Global Accelerator Developer Guide.
	//
	// [Flow logs]: https://docs.aws.amazon.com/global-accelerator/latest/dg/monitoring-global-accelerator.flow-logs.html
	FlowLogsEnabled *bool

	// The name of the Amazon S3 bucket for the flow logs. Attribute is required if
	// FlowLogsEnabled is true . The bucket must exist and have a bucket policy that
	// grants Global Accelerator permission to write to the bucket.
	FlowLogsS3Bucket *string

	// The prefix for the location in the Amazon S3 bucket for the flow logs.
	// Attribute is required if FlowLogsEnabled is true .
	//
	// If you specify slash (/) for the S3 bucket prefix, the log file bucket folder
	// structure will include a double slash (//), like the following:
	//
	// s3-bucket_name//AWSLogs/aws_account_id
	FlowLogsS3Prefix *string

	noSmithyDocumentSerde
}

// A complex type that contains a Timestamp value and Message for changes that you
// make to an accelerator in Global Accelerator. Messages stored here provide
// progress or error information when you update an accelerator from IPv4 to
// dual-stack, or from dual-stack to IPv4. Global Accelerator stores a maximum of
// ten event messages.
type AcceleratorEvent struct {

	// A string that contains an Event message describing changes or errors when you
	// update an accelerator in Global Accelerator from IPv4 to dual-stack, or
	// dual-stack to IPv4.
	Message *string

	// A timestamp for when you update an accelerator in Global Accelerator from IPv4
	// to dual-stack, or dual-stack to IPv4.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// A cross-account attachment in Global Accelerator. A cross-account attachment
// specifies the principals who have permission to work with resources in your
// account, which you also list in the attachment.
type Attachment struct {

	// The Amazon Resource Name (ARN) of the cross-account attachment.
	AttachmentArn *string

	// The date and time that the cross-account attachment was created.
	CreatedTime *time.Time

	// The date and time that the cross-account attachment was last modified.
	LastModifiedTime *time.Time

	// The name of the cross-account attachment.
	Name *string

	// The principals included in the cross-account attachment.
	Principals []string

	// The resources included in the cross-account attachment.
	Resources []Resource

	noSmithyDocumentSerde
}

// Information about an IP address range that is provisioned for use with your
// Amazon Web Services resources through bring your own IP address (BYOIP).
//
// The following describes each BYOIP State that your IP address range can be in.
//
//   - PENDING_PROVISIONING — You’ve submitted a request to provision an IP
//     address range but it is not yet provisioned with Global Accelerator.
//
//   - READY — The address range is provisioned with Global Accelerator and can be
//     advertised.
//
//   - PENDING_ADVERTISING — You’ve submitted a request for Global Accelerator to
//     advertise an address range but it is not yet being advertised.
//
//   - ADVERTISING — The address range is being advertised by Global Accelerator.
//
//   - PENDING_WITHDRAWING — You’ve submitted a request to withdraw an address
//     range from being advertised but it is still being advertised by Global
//     Accelerator.
//
//   - PENDING_DEPROVISIONING — You’ve submitted a request to deprovision an
//     address range from Global Accelerator but it is still provisioned.
//
//   - DEPROVISIONED — The address range is deprovisioned from Global Accelerator.
//
//   - FAILED_PROVISION — The request to provision the address range from Global
//     Accelerator was not successful. Please make sure that you provide all of the
//     correct information, and try again. If the request fails a second time, contact
//     Amazon Web Services support.
//
//   - FAILED_ADVERTISING — The request for Global Accelerator to advertise the
//     address range was not successful. Please make sure that you provide all of the
//     correct information, and try again. If the request fails a second time, contact
//     Amazon Web Services support.
//
//   - FAILED_WITHDRAW — The request to withdraw the address range from
//     advertising by Global Accelerator was not successful. Please make sure that you
//     provide all of the correct information, and try again. If the request fails a
//     second time, contact Amazon Web Services support.
//
//   - FAILED_DEPROVISION — The request to deprovision the address range from
//     Global Accelerator was not successful. Please make sure that you provide all of
//     the correct information, and try again. If the request fails a second time,
//     contact Amazon Web Services support.
type ByoipCidr struct {

	// The address range, in CIDR notation.
	//
	// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
	//
	// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
	Cidr *string

	// A history of status changes for an IP address range that you bring to Global
	// Accelerator through bring your own IP address (BYOIP).
	Events []ByoipCidrEvent

	// The state of the address pool.
	State ByoipCidrState

	noSmithyDocumentSerde
}

// A complex type that contains a Message and a Timestamp value for changes that
// you make in the status of an IP address range that you bring to Global
// Accelerator through bring your own IP address (BYOIP).
type ByoipCidrEvent struct {

	// A string that contains an Event message describing changes that you make in the
	// status of an IP address range that you bring to Global Accelerator through bring
	// your own IP address (BYOIP).
	Message *string

	// A timestamp for when you make a status change for an IP address range that you
	// bring to Global Accelerator through bring your own IP address (BYOIP).
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// Provides authorization for Amazon to bring a specific IP address range to a
// specific Amazon Web Services account using bring your own IP addresses (BYOIP).
//
// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
//
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
type CidrAuthorizationContext struct {

	// The plain-text authorization message for the prefix and account.
	//
	// This member is required.
	Message *string

	// The signed authorization message for the prefix and account.
	//
	// This member is required.
	Signature *string

	noSmithyDocumentSerde
}

// An endpoint (Amazon Web Services resource) or an IP address range, in CIDR
// format, that is listed in a cross-account attachment. A cross-account resource
// can be added to an accelerator by specified principals, which are also listed in
// the attachment.
//
// For more information, see [Working with cross-account attachments and resources in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Working with cross-account attachments and resources in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html
type CrossAccountResource struct {

	// The Amazon Resource Name (ARN) of the cross-account attachment that specifies
	// the resources (endpoints or CIDR range) that can be added to accelerators and
	// principals that have permission to add them.
	AttachmentArn *string

	// An IP address range, in CIDR format, that is specified as an Amazon Web
	// Services resource. The address must be provisioned and advertised in Global
	// Accelerator by following the bring your own IP address (BYOIP) process for
	// Global Accelerator.
	//
	// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
	//
	// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
	Cidr *string

	// The endpoint ID for the endpoint that is listed in a cross-account attachment
	// and can be added to an accelerator by specified principals.
	EndpointId *string

	noSmithyDocumentSerde
}

// Attributes of a custom routing accelerator.
type CustomRoutingAccelerator struct {

	// The Amazon Resource Name (ARN) of the custom routing accelerator.
	AcceleratorArn *string

	// The date and time that the accelerator was created.
	CreatedTime *time.Time

	// The Domain Name System (DNS) name that Global Accelerator creates that points
	// to an accelerator's static IPv4 addresses.
	//
	// The naming convention for the DNS name is the following: A lowercase letter a,
	// followed by a 16-bit random hex string, followed by .awsglobalaccelerator.com.
	// For example: a1234567890abcdef.awsglobalaccelerator.com.
	//
	// If you have a dual-stack accelerator, you also have a second DNS name,
	// DualStackDnsName , that points to both the A record and the AAAA record for all
	// four static addresses for the accelerator: two IPv4 addresses and two IPv6
	// addresses.
	//
	// For more information about the default DNS name, see [Support for DNS addressing in Global Accelerator] in the Global Accelerator
	// Developer Guide.
	//
	// [Support for DNS addressing in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/dns-addressing-custom-domains.dns-addressing.html
	DnsName *string

	// Indicates whether the accelerator is enabled. The value is true or false. The
	// default value is true.
	//
	// If the value is set to true, the accelerator cannot be deleted. If set to
	// false, accelerator can be deleted.
	Enabled *bool

	// The IP address type that an accelerator supports. For a custom routing
	// accelerator, the value must be IPV4.
	IpAddressType IpAddressType

	// The static IP addresses that Global Accelerator associates with the accelerator.
	IpSets []IpSet

	// The date and time that the accelerator was last modified.
	LastModifiedTime *time.Time

	// The name of the accelerator. The name must contain only alphanumeric characters
	// or hyphens (-), and must not begin or end with a hyphen.
	Name *string

	// Describes the deployment status of the accelerator.
	Status CustomRoutingAcceleratorStatus

	noSmithyDocumentSerde
}

// Attributes of a custom routing accelerator.
type CustomRoutingAcceleratorAttributes struct {

	// Indicates whether flow logs are enabled. The default value is false. If the
	// value is true, FlowLogsS3Bucket and FlowLogsS3Prefix must be specified.
	//
	// For more information, see [Flow logs] in the Global Accelerator Developer Guide.
	//
	// [Flow logs]: https://docs.aws.amazon.com/global-accelerator/latest/dg/monitoring-global-accelerator.flow-logs.html
	FlowLogsEnabled *bool

	// The name of the Amazon S3 bucket for the flow logs. Attribute is required if
	// FlowLogsEnabled is true . The bucket must exist and have a bucket policy that
	// grants Global Accelerator permission to write to the bucket.
	FlowLogsS3Bucket *string

	// The prefix for the location in the Amazon S3 bucket for the flow logs.
	// Attribute is required if FlowLogsEnabled is true .
	//
	// If you don’t specify a prefix, the flow logs are stored in the root of the
	// bucket. If you specify slash (/) for the S3 bucket prefix, the log file bucket
	// folder structure will include a double slash (//), like the following:
	//
	// DOC-EXAMPLE-BUCKET//AWSLogs/aws_account_id
	FlowLogsS3Prefix *string

	noSmithyDocumentSerde
}

// For a custom routing accelerator, sets the port range and protocol for all
// endpoints (virtual private cloud subnets) in an endpoint group to accept client
// traffic on.
type CustomRoutingDestinationConfiguration struct {

	// The first port, inclusive, in the range of ports for the endpoint group that is
	// associated with a custom routing accelerator.
	//
	// This member is required.
	FromPort *int32

	// The protocol for the endpoint group that is associated with a custom routing
	// accelerator. The protocol can be either TCP or UDP.
	//
	// This member is required.
	Protocols []CustomRoutingProtocol

	// The last port, inclusive, in the range of ports for the endpoint group that is
	// associated with a custom routing accelerator.
	//
	// This member is required.
	ToPort *int32

	noSmithyDocumentSerde
}

// For a custom routing accelerator, describes the port range and protocol for all
// endpoints (virtual private cloud subnets) in an endpoint group to accept client
// traffic on.
type CustomRoutingDestinationDescription struct {

	// The first port, inclusive, in the range of ports for the endpoint group that is
	// associated with a custom routing accelerator.
	FromPort *int32

	// The protocol for the endpoint group that is associated with a custom routing
	// accelerator. The protocol can be either TCP or UDP.
	Protocols []Protocol

	// The last port, inclusive, in the range of ports for the endpoint group that is
	// associated with a custom routing accelerator.
	ToPort *int32

	noSmithyDocumentSerde
}

// The list of endpoint objects. For custom routing, this is a list of virtual
// private cloud (VPC) subnet IDs.
type CustomRoutingEndpointConfiguration struct {

	// The Amazon Resource Name (ARN) of the cross-account attachment that specifies
	// the endpoints (resources) that can be added to accelerators and principals that
	// have permission to add the endpoints.
	AttachmentArn *string

	// An ID for the endpoint. For custom routing accelerators, this is the virtual
	// private cloud (VPC) subnet ID.
	EndpointId *string

	noSmithyDocumentSerde
}

// A complex type for an endpoint for a custom routing accelerator. Each endpoint
// group can include one or more endpoints, which are virtual private cloud (VPC)
// subnets.
type CustomRoutingEndpointDescription struct {

	// An ID for the endpoint. For custom routing accelerators, this is the virtual
	// private cloud (VPC) subnet ID.
	EndpointId *string

	noSmithyDocumentSerde
}

// A complex type for the endpoint group for a custom routing accelerator. An
// Amazon Web Services Region can have only one endpoint group for a specific
// listener.
type CustomRoutingEndpointGroup struct {

	// For a custom routing accelerator, describes the port range and protocol for all
	// endpoints (virtual private cloud subnets) in an endpoint group to accept client
	// traffic on.
	DestinationDescriptions []CustomRoutingDestinationDescription

	// For a custom routing accelerator, describes the endpoints (virtual private
	// cloud subnets) in an endpoint group to accept client traffic on.
	EndpointDescriptions []CustomRoutingEndpointDescription

	// The Amazon Resource Name (ARN) of the endpoint group.
	EndpointGroupArn *string

	// The Amazon Web Services Region where the endpoint group is located.
	EndpointGroupRegion *string

	noSmithyDocumentSerde
}

// A complex type for a listener for a custom routing accelerator.
type CustomRoutingListener struct {

	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string

	// The port range to support for connections from clients to your accelerator.
	//
	// Separately, you set port ranges for endpoints. For more information, see [About endpoints for custom routing accelerators].
	//
	// [About endpoints for custom routing accelerators]: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-custom-routing-endpoints.html
	PortRanges []PortRange

	noSmithyDocumentSerde
}

// The port mappings for a specified endpoint IP address (destination).
type DestinationPortMapping struct {

	// The Amazon Resource Name (ARN) of the custom routing accelerator that you have
	// port mappings for.
	AcceleratorArn *string

	// The IP address/port combinations (sockets) that map to a given destination
	// socket address.
	AcceleratorSocketAddresses []SocketAddress

	// The endpoint IP address/port combination for traffic received on the
	// accelerator socket address.
	DestinationSocketAddress *SocketAddress

	// Indicates whether or not a port mapping destination can receive traffic. The
	// value is either ALLOW, if traffic is allowed to the destination, or DENY, if
	// traffic is not allowed to the destination.
	DestinationTrafficState CustomRoutingDestinationTrafficState

	// The Amazon Resource Name (ARN) of the endpoint group.
	EndpointGroupArn *string

	// The Amazon Web Services Region for the endpoint group.
	EndpointGroupRegion *string

	// The ID for the virtual private cloud (VPC) subnet.
	EndpointId *string

	// The IP address type that an accelerator supports. For a custom routing
	// accelerator, the value must be IPV4.
	IpAddressType IpAddressType

	noSmithyDocumentSerde
}

// A complex type for endpoints. A resource must be valid and active when you add
// it as an endpoint.
type EndpointConfiguration struct {

	// The Amazon Resource Name (ARN) of the cross-account attachment that specifies
	// the endpoints (resources) that can be added to accelerators and principals that
	// have permission to add the endpoints.
	AttachmentArn *string

	// Indicates whether client IP address preservation is enabled for an endpoint.
	// The value is true or false. The default value is true for Application Load
	// Balancer endpoints.
	//
	// If the value is set to true, the client's IP address is preserved in the
	// X-Forwarded-For request header as traffic travels to applications on the
	// endpoint fronted by the accelerator.
	//
	// Client IP address preservation is supported, in specific Amazon Web Services
	// Regions, for endpoints that are Application Load Balancers, Amazon EC2
	// instances, and Network Load Balancers with security groups. IMPORTANT: You
	// cannot use client IP address preservation with Network Load Balancers with TLS
	// listeners.
	//
	// For more information, see [Preserve client IP addresses in Global Accelerator] in the Global Accelerator Developer Guide.
	//
	// [Preserve client IP addresses in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/preserve-client-ip-address.html
	ClientIPPreservationEnabled *bool

	// An ID for the endpoint. If the endpoint is a Network Load Balancer or
	// Application Load Balancer, this is the Amazon Resource Name (ARN) of the
	// resource. If the endpoint is an Elastic IP address, this is the Elastic IP
	// address allocation ID. For Amazon EC2 instances, this is the EC2 instance ID. A
	// resource must be valid and active when you add it as an endpoint.
	//
	// For cross-account endpoints, this must be the ARN of the resource.
	EndpointId *string

	// The weight associated with the endpoint. When you add weights to endpoints, you
	// configure Global Accelerator to route traffic based on proportions that you
	// specify. For example, you might specify endpoint weights of 4, 5, 5, and 6
	// (sum=20). The result is that 4/20 of your traffic, on average, is routed to the
	// first endpoint, 5/20 is routed both to the second and third endpoints, and 6/20
	// is routed to the last endpoint. For more information, see [Endpoint weights]in the Global
	// Accelerator Developer Guide.
	//
	// [Endpoint weights]: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints-endpoint-weights.html
	Weight *int32

	noSmithyDocumentSerde
}

// A complex type for an endpoint. Each endpoint group can include one or more
// endpoints, such as load balancers.
type EndpointDescription struct {

	// Indicates whether client IP address preservation is enabled for an endpoint.
	// The value is true or false. The default value is true for Application Load
	// Balancers endpoints.
	//
	// If the value is set to true, the client's IP address is preserved in the
	// X-Forwarded-For request header as traffic travels to applications on the
	// endpoint fronted by the accelerator.
	//
	// Client IP address preservation is supported, in specific Amazon Web Services
	// Regions, for endpoints that are Application Load Balancers, Amazon EC2
	// instances, and Network Load Balancers with security groups. IMPORTANT: You
	// cannot use client IP address preservation with Network Load Balancers with TLS
	// listeners.
	//
	// For more information, see [Preserve client IP addresses in Global Accelerator] in the Global Accelerator Developer Guide.
	//
	// [Preserve client IP addresses in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/preserve-client-ip-address.html
	ClientIPPreservationEnabled *bool

	// An ID for the endpoint. If the endpoint is a Network Load Balancer or
	// Application Load Balancer, this is the Amazon Resource Name (ARN) of the
	// resource. If the endpoint is an Elastic IP address, this is the Elastic IP
	// address allocation ID. For Amazon EC2 instances, this is the EC2 instance ID.
	//
	// An Application Load Balancer can be either internal or internet-facing.
	EndpointId *string

	// Returns a null result.
	HealthReason *string

	// The health status of the endpoint.
	HealthState HealthState

	// The weight associated with the endpoint. When you add weights to endpoints, you
	// configure Global Accelerator to route traffic based on proportions that you
	// specify. For example, you might specify endpoint weights of 4, 5, 5, and 6
	// (sum=20). The result is that 4/20 of your traffic, on average, is routed to the
	// first endpoint, 5/20 is routed both to the second and third endpoints, and 6/20
	// is routed to the last endpoint. For more information, see [Endpoint weights]in the Global
	// Accelerator Developer Guide.
	//
	// [Endpoint weights]: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints-endpoint-weights.html
	Weight *int32

	noSmithyDocumentSerde
}

// A complex type for the endpoint group. An Amazon Web Services Region can have
// only one endpoint group for a specific listener.
type EndpointGroup struct {

	// The list of endpoint objects.
	EndpointDescriptions []EndpointDescription

	// The Amazon Resource Name (ARN) of the endpoint group.
	EndpointGroupArn *string

	// The Amazon Web Services Region where the endpoint group is located.
	EndpointGroupRegion *string

	// The time—10 seconds or 30 seconds—between health checks for each endpoint. The
	// default value is 30.
	HealthCheckIntervalSeconds *int32

	// If the protocol is HTTP/S, then this value provides the ping path that Global
	// Accelerator uses for the destination on the endpoints for health checks. The
	// default is slash (/).
	HealthCheckPath *string

	// The port that Global Accelerator uses to perform health checks on endpoints
	// that are part of this endpoint group.
	//
	// The default port is the port for the listener that this endpoint group is
	// associated with. If the listener port is a list, Global Accelerator uses the
	// first specified port in the list of ports.
	HealthCheckPort *int32

	// The protocol that Global Accelerator uses to perform health checks on endpoints
	// that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol HealthCheckProtocol

	// Allows you to override the destination ports used to route traffic to an
	// endpoint. Using a port override lets you map a list of external destination
	// ports (that your users send traffic to) to a list of internal destination ports
	// that you want an application endpoint to receive traffic on.
	PortOverrides []PortOverride

	// The number of consecutive health checks required to set the state of a healthy
	// endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default
	// value is 3.
	ThresholdCount *int32

	// The percentage of traffic to send to an Amazon Web Services Region. Additional
	// traffic is distributed to other endpoint groups for this listener.
	//
	// Use this action to increase (dial up) or decrease (dial down) traffic to a
	// specific Region. The percentage is applied to the traffic that would otherwise
	// have been routed to the Region based on optimal routing.
	//
	// The default value is 100.
	TrafficDialPercentage *float32

	noSmithyDocumentSerde
}

// A complex type for an endpoint. Specifies information about the endpoint to
// remove from the endpoint group.
type EndpointIdentifier struct {

	// An ID for the endpoint. If the endpoint is a Network Load Balancer or
	// Application Load Balancer, this is the Amazon Resource Name (ARN) of the
	// resource. If the endpoint is an Elastic IP address, this is the Elastic IP
	// address allocation ID. For Amazon EC2 instances, this is the EC2 instance ID.
	//
	// An Application Load Balancer can be either internal or internet-facing.
	//
	// This member is required.
	EndpointId *string

	// Indicates whether client IP address preservation is enabled for an endpoint.
	// The value is true or false.
	//
	// If the value is set to true, the client's IP address is preserved in the
	// X-Forwarded-For request header as traffic travels to applications on the
	// endpoint fronted by the accelerator.
	ClientIPPreservationEnabled *bool

	noSmithyDocumentSerde
}

// A complex type for the set of IP addresses for an accelerator.
type IpSet struct {

	// The types of IP addresses included in this IP set.
	IpAddressFamily IpAddressFamily

	// The array of IP addresses in the IP address set. An IP address set can have a
	// maximum of two IP addresses.
	IpAddresses []string

	// IpFamily is deprecated and has been replaced by IpAddressFamily.
	//
	// Deprecated: IpFamily has been replaced by IpAddressFamily
	IpFamily *string

	noSmithyDocumentSerde
}

// A complex type for a listener.
type Listener struct {

	// Client affinity lets you direct all requests from a user to the same endpoint,
	// if you have stateful applications, regardless of the port and protocol of the
	// client request. Client affinity gives you control over whether to always route
	// each client to the same specific endpoint.
	//
	// Global Accelerator uses a consistent-flow hashing algorithm to choose the
	// optimal endpoint for a connection. If client affinity is NONE , Global
	// Accelerator uses the "five-tuple" (5-tuple) properties—source IP address, source
	// port, destination IP address, destination port, and protocol—to select the hash
	// value, and then chooses the best endpoint. However, with this setting, if
	// someone uses different ports to connect to Global Accelerator, their connections
	// might not be always routed to the same endpoint because the hash value changes.
	//
	// If you want a given client to always be routed to the same endpoint, set client
	// affinity to SOURCE_IP instead. When you use the SOURCE_IP setting, Global
	// Accelerator uses the "two-tuple" (2-tuple) properties— source (client) IP
	// address and destination IP address—to select the hash value.
	//
	// The default value is NONE .
	ClientAffinity ClientAffinity

	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string

	// The list of port ranges for the connections from clients to the accelerator.
	PortRanges []PortRange

	// The protocol for the connections from clients to the accelerator.
	Protocol Protocol

	noSmithyDocumentSerde
}

// Returns the ports and associated IP addresses and ports of Amazon EC2 instances
// in your virtual private cloud (VPC) subnets. Custom routing is a port mapping
// protocol in Global Accelerator that statically associates port ranges with VPC
// subnets, which allows Global Accelerator to route to specific instances and
// ports within one or more subnets.
type PortMapping struct {

	// The accelerator port.
	AcceleratorPort *int32

	// The EC2 instance IP address and port number in the virtual private cloud (VPC)
	// subnet.
	DestinationSocketAddress *SocketAddress

	// Indicates whether or not a port mapping destination can receive traffic. The
	// value is either ALLOW, if traffic is allowed to the destination, or DENY, if
	// traffic is not allowed to the destination.
	DestinationTrafficState CustomRoutingDestinationTrafficState

	// The Amazon Resource Name (ARN) of the endpoint group.
	EndpointGroupArn *string

	// The IP address of the VPC subnet (the subnet ID).
	EndpointId *string

	// The protocols supported by the endpoint group.
	Protocols []CustomRoutingProtocol

	noSmithyDocumentSerde
}

// Override specific listener ports used to route traffic to endpoints that are
// part of an endpoint group. For example, you can create a port override in which
// the listener receives user traffic on ports 80 and 443, but your accelerator
// routes that traffic to ports 1080 and 1443, respectively, on the endpoints.
//
// For more information, see [Overriding listener ports] in the Global Accelerator Developer Guide.
//
// [Overriding listener ports]: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoint-groups-port-override.html
type PortOverride struct {

	// The endpoint port that you want a listener port to be mapped to. This is the
	// port on the endpoint, such as the Application Load Balancer or Amazon EC2
	// instance.
	EndpointPort *int32

	// The listener port that you want to map to a specific endpoint port. This is the
	// port that user traffic arrives to the Global Accelerator on.
	ListenerPort *int32

	noSmithyDocumentSerde
}

// A complex type for a range of ports for a listener.
type PortRange struct {

	// The first port in the range of ports, inclusive.
	FromPort *int32

	// The last port in the range of ports, inclusive.
	ToPort *int32

	noSmithyDocumentSerde
}

// A resource is one of the following: the ARN for an Amazon Web Services resource
// that is supported by Global Accelerator to be added as an endpoint, or a CIDR
// range that specifies a bring your own IP (BYOIP) address pool.
type Resource struct {

	// An IP address range, in CIDR format, that is specified as resource. The address
	// must be provisioned and advertised in Global Accelerator by following the bring
	// your own IP address (BYOIP) process for Global Accelerator
	//
	// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
	//
	// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
	Cidr *string

	// The endpoint ID for the endpoint that is specified as a Amazon Web Services
	// resource.
	//
	// An endpoint ID for the cross-account feature is the ARN of an Amazon Web
	// Services resource, such as a Network Load Balancer, that Global Accelerator
	// supports as an endpoint for an accelerator.
	EndpointId *string

	// The Amazon Web Services Region where a shared endpoint resource is located.
	Region *string

	noSmithyDocumentSerde
}

// An IP address/port combination.
type SocketAddress struct {

	// The IP address for the socket address.
	IpAddress *string

	// The port for the socket address.
	Port *int32

	noSmithyDocumentSerde
}

// A complex type that contains a Tag key and Tag value.
type Tag struct {

	// A string that contains a Tag key.
	//
	// This member is required.
	Key *string

	// A string that contains a Tag value.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

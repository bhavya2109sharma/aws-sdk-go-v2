// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Instantiates an auto-scaling virtual server based on the selected file transfer
// protocol in Amazon Web Services. When you make updates to your file transfer
// protocol-enabled server or when you work with users, use the service-generated
// ServerId property that is assigned to the newly created server.
func (c *Client) CreateServer(ctx context.Context, params *CreateServerInput, optFns ...func(*Options)) (*CreateServerOutput, error) {
	if params == nil {
		params = &CreateServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServer", params, optFns, c.addOperationCreateServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServerInput struct {

	// The Amazon Resource Name (ARN) of the Certificate Manager (ACM) certificate.
	// Required when Protocols is set to FTPS .
	//
	// To request a new public certificate, see [Request a public certificate] in the Certificate Manager User Guide.
	//
	// To import an existing certificate into ACM, see [Importing certificates into ACM] in the Certificate Manager
	// User Guide.
	//
	// To request a private certificate to use FTPS through private IP addresses, see [Request a private certificate]
	// in the Certificate Manager User Guide.
	//
	// Certificates with the following cryptographic algorithms and key sizes are
	// supported:
	//
	//   - 2048-bit RSA (RSA_2048)
	//
	//   - 4096-bit RSA (RSA_4096)
	//
	//   - Elliptic Prime Curve 256 bit (EC_prime256v1)
	//
	//   - Elliptic Prime Curve 384 bit (EC_secp384r1)
	//
	//   - Elliptic Prime Curve 521 bit (EC_secp521r1)
	//
	// The certificate must be a valid SSL/TLS X.509 version 3 certificate with FQDN
	// or IP address specified and information about the issuer.
	//
	// [Request a public certificate]: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-public.html
	// [Request a private certificate]: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-private.html
	// [Importing certificates into ACM]: https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html
	Certificate *string

	// The domain of the storage system that is used for file transfers. There are two
	// domains available: Amazon Simple Storage Service (Amazon S3) and Amazon Elastic
	// File System (Amazon EFS). The default value is S3.
	//
	// After the server is created, the domain cannot be changed.
	Domain types.Domain

	// The virtual private cloud (VPC) endpoint settings that are configured for your
	// server. When you host your endpoint within your VPC, you can make your endpoint
	// accessible only to resources within your VPC, or you can attach Elastic IP
	// addresses and make your endpoint accessible to clients over the internet. Your
	// VPC's default security groups are automatically assigned to your endpoint.
	EndpointDetails *types.EndpointDetails

	// The type of endpoint that you want your server to use. You can choose to make
	// your server's endpoint publicly accessible (PUBLIC) or host it inside your VPC.
	// With an endpoint that is hosted in a VPC, you can restrict access to your server
	// and resources only within your VPC or choose to make it internet facing by
	// attaching Elastic IP addresses directly to it.
	//
	// After May 19, 2021, you won't be able to create a server using
	// EndpointType=VPC_ENDPOINT in your Amazon Web Services account if your account
	// hasn't already done so before May 19, 2021. If you have already created servers
	// with EndpointType=VPC_ENDPOINT in your Amazon Web Services account on or before
	// May 19, 2021, you will not be affected. After this date, use EndpointType = VPC .
	//
	// For more information, see
	// https://docs.aws.amazon.com/transfer/latest/userguide/create-server-in-vpc.html#deprecate-vpc-endpoint.
	//
	// It is recommended that you use VPC as the EndpointType . With this endpoint
	// type, you have the option to directly associate up to three Elastic IPv4
	// addresses (BYO IP included) with your server's endpoint and use VPC security
	// groups to restrict traffic by the client's public IP address. This is not
	// possible with EndpointType set to VPC_ENDPOINT .
	EndpointType types.EndpointType

	// The RSA, ECDSA, or ED25519 private key to use for your SFTP-enabled server. You
	// can add multiple host keys, in case you want to rotate keys, or have a set of
	// active keys that use different algorithms.
	//
	// Use the following command to generate an RSA 2048 bit key with no passphrase:
	//
	// ssh-keygen -t rsa -b 2048 -N "" -m PEM -f my-new-server-key .
	//
	// Use a minimum value of 2048 for the -b option. You can create a stronger key by
	// using 3072 or 4096.
	//
	// Use the following command to generate an ECDSA 256 bit key with no passphrase:
	//
	// ssh-keygen -t ecdsa -b 256 -N "" -m PEM -f my-new-server-key .
	//
	// Valid values for the -b option for ECDSA are 256, 384, and 521.
	//
	// Use the following command to generate an ED25519 key with no passphrase:
	//
	// ssh-keygen -t ed25519 -N "" -f my-new-server-key .
	//
	// For all of these commands, you can replace my-new-server-key with a string of
	// your choice.
	//
	// If you aren't planning to migrate existing users from an existing SFTP-enabled
	// server to a new server, don't update the host key. Accidentally changing a
	// server's host key can be disruptive.
	//
	// For more information, see [Manage host keys for your SFTP-enabled server] in the Transfer Family User Guide.
	//
	// [Manage host keys for your SFTP-enabled server]: https://docs.aws.amazon.com/transfer/latest/userguide/edit-server-config.html#configuring-servers-change-host-key
	HostKey *string

	// Required when IdentityProviderType is set to AWS_DIRECTORY_SERVICE , Amazon Web
	// Services_LAMBDA or API_GATEWAY . Accepts an array containing all of the
	// information required to use a directory in AWS_DIRECTORY_SERVICE or invoke a
	// customer-supplied authentication API, including the API Gateway URL. Not
	// required when IdentityProviderType is set to SERVICE_MANAGED .
	IdentityProviderDetails *types.IdentityProviderDetails

	// The mode of authentication for a server. The default value is SERVICE_MANAGED ,
	// which allows you to store and access user credentials within the Transfer Family
	// service.
	//
	// Use AWS_DIRECTORY_SERVICE to provide access to Active Directory groups in
	// Directory Service for Microsoft Active Directory or Microsoft Active Directory
	// in your on-premises environment or in Amazon Web Services using AD Connector.
	// This option also requires you to provide a Directory ID by using the
	// IdentityProviderDetails parameter.
	//
	// Use the API_GATEWAY value to integrate with an identity provider of your
	// choosing. The API_GATEWAY setting requires you to provide an Amazon API Gateway
	// endpoint URL to call for authentication by using the IdentityProviderDetails
	// parameter.
	//
	// Use the AWS_LAMBDA value to directly use an Lambda function as your identity
	// provider. If you choose this value, you must specify the ARN for the Lambda
	// function in the Function parameter for the IdentityProviderDetails data type.
	IdentityProviderType types.IdentityProviderType

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or
	// Amazon EFSevents. When set, you can view user activity in your CloudWatch logs.
	LoggingRole *string

	// Specifies a string to display when users connect to a server. This string is
	// displayed after the user authenticates.
	//
	// The SFTP protocol does not support post-authentication display banners.
	PostAuthenticationLoginBanner *string

	// Specifies a string to display when users connect to a server. This string is
	// displayed before the user authenticates. For example, the following banner
	// displays details about using the system:
	//
	//     This system is for the use of authorized users only. Individuals using this
	//     computer system without authority, or in excess of their authority, are subject
	//     to having all of their activities on this system monitored and recorded by
	//     system personnel.
	PreAuthenticationLoginBanner *string

	// The protocol settings that are configured for your server.
	//
	//   - To indicate passive mode (for FTP and FTPS protocols), use the PassiveIp
	//   parameter. Enter a single dotted-quad IPv4 address, such as the external IP
	//   address of a firewall, router, or load balancer.
	//
	//   - To ignore the error that is generated when the client attempts to use the
	//   SETSTAT command on a file that you are uploading to an Amazon S3 bucket, use
	//   the SetStatOption parameter. To have the Transfer Family server ignore the
	//   SETSTAT command and upload files without needing to make any changes to your
	//   SFTP client, set the value to ENABLE_NO_OP . If you set the SetStatOption
	//   parameter to ENABLE_NO_OP , Transfer Family generates a log entry to Amazon
	//   CloudWatch Logs, so that you can determine when the client is making a SETSTAT
	//   call.
	//
	//   - To determine whether your Transfer Family server resumes recent, negotiated
	//   sessions through a unique session ID, use the TlsSessionResumptionMode
	//   parameter.
	//
	//   - As2Transports indicates the transport method for the AS2 messages.
	//   Currently, only HTTP is supported.
	ProtocolDetails *types.ProtocolDetails

	// Specifies the file transfer protocol or protocols over which your file transfer
	// protocol client can connect to your server's endpoint. The available protocols
	// are:
	//
	//   - SFTP (Secure Shell (SSH) File Transfer Protocol): File transfer over SSH
	//
	//   - FTPS (File Transfer Protocol Secure): File transfer with TLS encryption
	//
	//   - FTP (File Transfer Protocol): Unencrypted file transfer
	//
	//   - AS2 (Applicability Statement 2): used for transporting structured
	//   business-to-business data
	//
	//   - If you select FTPS , you must choose a certificate stored in Certificate
	//   Manager (ACM) which is used to identify your server when clients connect to it
	//   over FTPS.
	//
	//   - If Protocol includes either FTP or FTPS , then the EndpointType must be VPC
	//   and the IdentityProviderType must be either AWS_DIRECTORY_SERVICE , AWS_LAMBDA
	//   , or API_GATEWAY .
	//
	//   - If Protocol includes FTP , then AddressAllocationIds cannot be associated.
	//
	//   - If Protocol is set only to SFTP , the EndpointType can be set to PUBLIC and
	//   the IdentityProviderType can be set any of the supported identity types:
	//   SERVICE_MANAGED , AWS_DIRECTORY_SERVICE , AWS_LAMBDA , or API_GATEWAY .
	//
	//   - If Protocol includes AS2 , then the EndpointType must be VPC , and domain
	//   must be Amazon S3.
	Protocols []types.Protocol

	// Specifies whether or not performance for your Amazon S3 directories is
	// optimized. This is disabled by default.
	//
	// By default, home directory mappings have a TYPE of DIRECTORY . If you enable
	// this option, you would then need to explicitly set the HomeDirectoryMapEntry Type
	// to FILE if you want a mapping to have a file target.
	S3StorageOptions *types.S3StorageOptions

	// Specifies the name of the security policy for the server.
	SecurityPolicyName *string

	// Specifies the log groups to which your server logs are sent.
	//
	// To specify a log group, you must provide the ARN for an existing log group. In
	// this case, the format of the log group is as follows:
	//
	//     arn:aws:logs:region-name:amazon-account-id:log-group:log-group-name:*
	//
	// For example, arn:aws:logs:us-east-1:111122223333:log-group:mytestgroup:*
	//
	// If you have previously specified a log group for a server, you can clear it,
	// and in effect turn off structured logging, by providing an empty value for this
	// parameter in an update-server call. For example:
	//
	//     update-server --server-id s-1234567890abcdef0 --structured-log-destinations
	StructuredLogDestinations []string

	// Key-value pairs that can be used to group and search for servers.
	Tags []types.Tag

	// Specifies the workflow ID for the workflow to assign and the execution role
	// that's used for executing the workflow.
	//
	// In addition to a workflow to execute when a file is uploaded completely,
	// WorkflowDetails can also contain a workflow ID (and execution role) for a
	// workflow to execute on partial upload. A partial upload occurs when the server
	// session disconnects while the file is still being uploaded.
	WorkflowDetails *types.WorkflowDetails

	noSmithyDocumentSerde
}

type CreateServerOutput struct {

	// The service-assigned identifier of the server that is created.
	//
	// This member is required.
	ServerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateServer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateServer"); err != nil {
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
	if err = addOpCreateServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateServer",
	}
}

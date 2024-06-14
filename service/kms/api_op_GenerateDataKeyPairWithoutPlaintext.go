// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a unique asymmetric data key pair for use outside of KMS. This
// operation returns a plaintext public key and a copy of the private key that is
// encrypted under the symmetric encryption KMS key you specify. Unlike GenerateDataKeyPair, this
// operation does not return a plaintext private key. The bytes in the keys are
// random; they are not related to the caller or to the KMS key that is used to
// encrypt the private key.
//
// You can use the public key that GenerateDataKeyPairWithoutPlaintext returns to
// encrypt data or verify a signature outside of KMS. Then, store the encrypted
// private key with the data. When you are ready to decrypt data or sign a message,
// you can use the Decryptoperation to decrypt the encrypted private key.
//
// To generate a data key pair, you must specify a symmetric encryption KMS key to
// encrypt the private key in a data key pair. You cannot use an asymmetric KMS key
// or a KMS key in a custom key store. To get the type and origin of your KMS key,
// use the DescribeKeyoperation.
//
// Use the KeyPairSpec parameter to choose an RSA or Elliptic Curve (ECC) data key
// pair. In China Regions, you can also choose an SM2 data key pair. KMS recommends
// that you use ECC key pairs for signing, and use RSA and SM2 key pairs for either
// encryption or signing, but not both. However, KMS cannot enforce any
// restrictions on the use of data key pairs outside of KMS.
//
// GenerateDataKeyPairWithoutPlaintext returns a unique data key pair for each
// request. The bytes in the key are not related to the caller or KMS key that is
// used to encrypt the private key. The public key is a DER-encoded X.509
// SubjectPublicKeyInfo, as specified in [RFC 5280].
//
// You can use an optional encryption context to add additional security to the
// encryption operation. If you specify an EncryptionContext , you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the
// encrypted data key. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException . For more information, see [Encryption Context] in the Key Management
// Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateDataKeyPairWithoutPlaintext] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// # GenerateDataKeyWithoutPlaintext
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [RFC 5280]: https://tools.ietf.org/html/rfc5280
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context
// [kms:GenerateDataKeyPairWithoutPlaintext]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html
func (c *Client) GenerateDataKeyPairWithoutPlaintext(ctx context.Context, params *GenerateDataKeyPairWithoutPlaintextInput, optFns ...func(*Options)) (*GenerateDataKeyPairWithoutPlaintextOutput, error) {
	if params == nil {
		params = &GenerateDataKeyPairWithoutPlaintextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateDataKeyPairWithoutPlaintext", params, optFns, c.addOperationGenerateDataKeyPairWithoutPlaintextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateDataKeyPairWithoutPlaintextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateDataKeyPairWithoutPlaintextInput struct {

	// Specifies the symmetric encryption KMS key that encrypts the private key in the
	// data key pair. You cannot specify an asymmetric KMS key or a KMS key in a custom
	// key store. To get the type and origin of your KMS key, use the DescribeKeyoperation.
	//
	// To specify a KMS key, use its key ID, key ARN, alias name, or alias ARN. When
	// using an alias name, prefix it with "alias/" . To specify a KMS key in a
	// different Amazon Web Services account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Alias name: alias/ExampleAlias
	//
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey. To get the alias name
	// and alias ARN, use ListAliases.
	//
	// This member is required.
	KeyId *string

	// Determines the type of data key pair that is generated.
	//
	// The KMS rule that restricts the use of asymmetric RSA and SM2 KMS keys to
	// encrypt and decrypt or to sign and verify (but not both), and the rule that
	// permits you to use ECC KMS keys only to sign and verify, are not effective on
	// data key pairs, which are used outside of KMS. The SM2 key spec is only
	// available in China Regions.
	//
	// This member is required.
	KeyPairSpec types.DataKeyPairSpec

	// Checks if your request will succeed. DryRun is an optional parameter.
	//
	// To learn more about how to use this parameter, see [Testing your KMS API calls] in the Key Management
	// Service Developer Guide.
	//
	// [Testing your KMS API calls]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-dryrun.html
	DryRun *bool

	// Specifies the encryption context that will be used when encrypting the private
	// key in the data key pair.
	//
	// Do not include confidential or sensitive information in this field. This field
	// may be displayed in plaintext in CloudTrail logs and other output.
	//
	// An encryption context is a collection of non-secret key-value pairs that
	// represent additional authenticated data. When you use an encryption context to
	// encrypt data, you must specify the same (an exact case-sensitive match)
	// encryption context to decrypt the data. An encryption context is supported only
	// on operations with symmetric encryption KMS keys. On operations with symmetric
	// encryption KMS keys, an encryption context is optional, but it is strongly
	// recommended.
	//
	// For more information, see [Encryption context] in the Key Management Service Developer Guide.
	//
	// [Encryption context]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context
	EncryptionContext map[string]string

	// A list of grant tokens.
	//
	// Use a grant token when your permission to call this operation comes from a new
	// grant that has not yet achieved eventual consistency. For more information, see [Grant token]
	// and [Using a grant token]in the Key Management Service Developer Guide.
	//
	// [Grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token
	// [Using a grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token
	GrantTokens []string

	noSmithyDocumentSerde
}

type GenerateDataKeyPairWithoutPlaintextOutput struct {

	// The Amazon Resource Name ([key ARN] ) of the KMS key that encrypted the private key.
	//
	// [key ARN]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN
	KeyId *string

	// The type of data key pair that was generated.
	KeyPairSpec types.DataKeyPairSpec

	// The encrypted copy of the private key. When you use the HTTP API or the Amazon
	// Web Services CLI, the value is Base64-encoded. Otherwise, it is not
	// Base64-encoded.
	PrivateKeyCiphertextBlob []byte

	// The public key (in plaintext). When you use the HTTP API or the Amazon Web
	// Services CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	PublicKey []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateDataKeyPairWithoutPlaintextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateDataKeyPairWithoutPlaintext{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateDataKeyPairWithoutPlaintext{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GenerateDataKeyPairWithoutPlaintext"); err != nil {
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
	if err = addOpGenerateDataKeyPairWithoutPlaintextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateDataKeyPairWithoutPlaintext(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateDataKeyPairWithoutPlaintext(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GenerateDataKeyPairWithoutPlaintext",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies Backup Vault Lock to a backup vault, preventing attempts to delete any
// recovery point stored in or created in a backup vault. Vault Lock also prevents
// attempts to update the lifecycle policy that controls the retention period of
// any recovery point currently stored in a backup vault. If specified, Vault Lock
// enforces a minimum and maximum retention period for future backup and copy jobs
// that target a backup vault.
//
// Backup Vault Lock has been assessed by Cohasset Associates for use in
// environments that are subject to SEC 17a-4, CFTC, and FINRA regulations. For
// more information about how Backup Vault Lock relates to these regulations, see
// the Cohasset Associates Compliance Assessment.
func (c *Client) PutBackupVaultLockConfiguration(ctx context.Context, params *PutBackupVaultLockConfigurationInput, optFns ...func(*Options)) (*PutBackupVaultLockConfigurationOutput, error) {
	if params == nil {
		params = &PutBackupVaultLockConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBackupVaultLockConfiguration", params, optFns, c.addOperationPutBackupVaultLockConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBackupVaultLockConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBackupVaultLockConfigurationInput struct {

	// The Backup Vault Lock configuration that specifies the name of the backup vault
	// it protects.
	//
	// This member is required.
	BackupVaultName *string

	// The Backup Vault Lock configuration that specifies the number of days before
	// the lock date. For example, setting ChangeableForDays to 30 on Jan. 1, 2022 at
	// 8pm UTC will set the lock date to Jan. 31, 2022 at 8pm UTC.
	//
	// Backup enforces a 72-hour cooling-off period before Vault Lock takes effect and
	// becomes immutable. Therefore, you must set ChangeableForDays to 3 or greater.
	//
	// Before the lock date, you can delete Vault Lock from the vault using
	// DeleteBackupVaultLockConfiguration or change the Vault Lock configuration using
	// PutBackupVaultLockConfiguration . On and after the lock date, the Vault Lock
	// becomes immutable and cannot be changed or deleted.
	//
	// If this parameter is not specified, you can delete Vault Lock from the vault
	// using DeleteBackupVaultLockConfiguration or change the Vault Lock configuration
	// using PutBackupVaultLockConfiguration at any time.
	ChangeableForDays *int64

	// The Backup Vault Lock configuration that specifies the maximum retention period
	// that the vault retains its recovery points. This setting can be useful if, for
	// example, your organization's policies require you to destroy certain data after
	// retaining it for four years (1460 days).
	//
	// If this parameter is not included, Vault Lock does not enforce a maximum
	// retention period on the recovery points in the vault. If this parameter is
	// included without a value, Vault Lock will not enforce a maximum retention
	// period.
	//
	// If this parameter is specified, any backup or copy job to the vault must have a
	// lifecycle policy with a retention period equal to or shorter than the maximum
	// retention period. If the job's retention period is longer than that maximum
	// retention period, then the vault fails the backup or copy job, and you should
	// either modify your lifecycle settings or use a different vault. The longest
	// maximum retention period you can specify is 36500 days (approximately 100
	// years). Recovery points already saved in the vault prior to Vault Lock are not
	// affected.
	MaxRetentionDays *int64

	// The Backup Vault Lock configuration that specifies the minimum retention period
	// that the vault retains its recovery points. This setting can be useful if, for
	// example, your organization's policies require you to retain certain data for at
	// least seven years (2555 days).
	//
	// If this parameter is not specified, Vault Lock will not enforce a minimum
	// retention period.
	//
	// If this parameter is specified, any backup or copy job to the vault must have a
	// lifecycle policy with a retention period equal to or longer than the minimum
	// retention period. If the job's retention period is shorter than that minimum
	// retention period, then the vault fails that backup or copy job, and you should
	// either modify your lifecycle settings or use a different vault. The shortest
	// minimum retention period you can specify is 1 day. Recovery points already saved
	// in the vault prior to Vault Lock are not affected.
	MinRetentionDays *int64

	noSmithyDocumentSerde
}

type PutBackupVaultLockConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBackupVaultLockConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutBackupVaultLockConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutBackupVaultLockConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBackupVaultLockConfiguration"); err != nil {
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
	if err = addOpPutBackupVaultLockConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBackupVaultLockConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutBackupVaultLockConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBackupVaultLockConfiguration",
	}
}

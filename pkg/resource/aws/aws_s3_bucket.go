// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/helpers"
)

const AwsS3BucketResourceType = "aws_s3_bucket"

type AwsS3Bucket struct {
	AccelerationStatus       *string           `cty:"acceleration_status" computed:"true"`
	Acl                      *string           `cty:"acl" diff:"-"`
	Arn                      *string           `cty:"arn" computed:"true"`
	Bucket                   *string           `cty:"bucket" computed:"true"`
	BucketDomainName         *string           `cty:"bucket_domain_name" computed:"true"`
	BucketPrefix             *string           `cty:"bucket_prefix"`
	BucketRegionalDomainName *string           `cty:"bucket_regional_domain_name" computed:"true"`
	ForceDestroy             *bool             `cty:"force_destroy" diff:"-"`
	HostedZoneId             *string           `cty:"hosted_zone_id" computed:"true"`
	Id                       string            `cty:"id" computed:"true"`
	Policy                   *string           `cty:"policy" jsonstring:"true"`
	Region                   *string           `cty:"region" computed:"true"`
	RequestPayer             *string           `cty:"request_payer" computed:"true"`
	Tags                     map[string]string `cty:"tags"`
	WebsiteDomain            *string           `cty:"website_domain" computed:"true"`
	WebsiteEndpoint          *string           `cty:"website_endpoint" computed:"true"`
	CorsRule                 *[]struct {
		AllowedHeaders []string `cty:"allowed_headers"`
		AllowedMethods []string `cty:"allowed_methods"`
		AllowedOrigins []string `cty:"allowed_origins"`
		ExposeHeaders  []string `cty:"expose_headers"`
		MaxAgeSeconds  *int     `cty:"max_age_seconds"`
	} `cty:"cors_rule"`
	Grant *[]struct {
		Id          string   `cty:"id"`
		Permissions []string `cty:"permissions"`
		Type        *string  `cty:"type"`
		Uri         *string  `cty:"uri"`
	} `cty:"grant"`
	LifecycleRule *[]struct {
		AbortIncompleteMultipartUploadDays *int              `cty:"abort_incomplete_multipart_upload_days"`
		Enabled                            *bool             `cty:"enabled"`
		Id                                 string            `cty:"id" computed:"true"`
		Prefix                             *string           `cty:"prefix"`
		Tags                               map[string]string `cty:"tags"`
		Expiration                         *[]struct {
			Date                      *string `cty:"date"`
			Days                      *int    `cty:"days"`
			ExpiredObjectDeleteMarker *bool   `cty:"expired_object_delete_marker"`
		} `cty:"expiration"`
		NoncurrentVersionExpiration *[]struct {
			Days *int `cty:"days"`
		} `cty:"noncurrent_version_expiration"`
		NoncurrentVersionTransition *[]struct {
			Days         *int    `cty:"days"`
			StorageClass *string `cty:"storage_class"`
		} `cty:"noncurrent_version_transition"`
		Transition *[]struct {
			Date         *string `cty:"date"`
			Days         *int    `cty:"days"`
			StorageClass *string `cty:"storage_class"`
		} `cty:"transition"`
	} `cty:"lifecycle_rule"`
	Logging *[]struct {
		TargetBucket *string `cty:"target_bucket"`
		TargetPrefix *string `cty:"target_prefix"`
	} `cty:"logging"`
	ObjectLockConfiguration *[]struct {
		ObjectLockEnabled *string `cty:"object_lock_enabled"`
		Rule              *[]struct {
			DefaultRetention *[]struct {
				Days  *int    `cty:"days"`
				Mode  *string `cty:"mode"`
				Years *int    `cty:"years"`
			} `cty:"default_retention"`
		} `cty:"rule"`
	} `cty:"object_lock_configuration"`
	ReplicationConfiguration *[]struct {
		Role  *string `cty:"role"`
		Rules *[]struct {
			Id          string  `cty:"id"`
			Prefix      *string `cty:"prefix"`
			Priority    *int    `cty:"priority"`
			Status      *string `cty:"status"`
			Destination *[]struct {
				AccountId                *string `cty:"account_id"`
				Bucket                   *string `cty:"bucket"`
				ReplicaKmsKeyId          *string `cty:"replica_kms_key_id"`
				StorageClass             *string `cty:"storage_class"`
				AccessControlTranslation *[]struct {
					Owner *string `cty:"owner"`
				} `cty:"access_control_translation"`
			} `cty:"destination"`
			Filter *[]struct {
				Prefix *string           `cty:"prefix"`
				Tags   map[string]string `cty:"tags"`
			} `cty:"filter"`
			SourceSelectionCriteria *[]struct {
				SseKmsEncryptedObjects *[]struct {
					Enabled *bool `cty:"enabled"`
				} `cty:"sse_kms_encrypted_objects"`
			} `cty:"source_selection_criteria"`
		} `cty:"rules"`
	} `cty:"replication_configuration"`
	ServerSideEncryptionConfiguration *[]struct {
		Rule *[]struct {
			ApplyServerSideEncryptionByDefault *[]struct {
				KmsMasterKeyId *string `cty:"kms_master_key_id"`
				SseAlgorithm   *string `cty:"sse_algorithm"`
			} `cty:"apply_server_side_encryption_by_default"`
		} `cty:"rule"`
	} `cty:"server_side_encryption_configuration"`
	Versioning *[]struct {
		Enabled   *bool `cty:"enabled"`
		MfaDelete *bool `cty:"mfa_delete"`
	} `cty:"versioning"`
	Website *[]struct {
		ErrorDocument         *string `cty:"error_document"`
		IndexDocument         *string `cty:"index_document"`
		RedirectAllRequestsTo *string `cty:"redirect_all_requests_to"`
		RoutingRules          *string `cty:"routing_rules"`
	} `cty:"website"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsS3Bucket) TerraformId() string {
	return r.Id
}

func (r *AwsS3Bucket) TerraformType() string {
	return AwsS3BucketResourceType
}

func (r *AwsS3Bucket) CtyValue() *cty.Value {
	return r.CtyVal
}

func awsS3BucketNormalizer(val *map[string]interface{}) {
	helpers.SafeDelete(val, []string{"acl"})
	helpers.SafeDelete(val, []string{"force_destroy"})
}

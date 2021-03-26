// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"
)

const AwsIamAccessKeyResourceType = "aws_iam_access_key"

type AwsIamAccessKey struct {
	EncryptedSecret   *string    `cty:"encrypted_secret" computed:"true"`
	Id                string     `cty:"id" computed:"true"`
	KeyFingerprint    *string    `cty:"key_fingerprint" computed:"true"`
	PgpKey            *string    `cty:"pgp_key"`
	Secret            *string    `cty:"secret" computed:"true"`
	SesSmtpPasswordV4 *string    `cty:"ses_smtp_password_v4" computed:"true"`
	Status            *string    `cty:"status" computed:"true"`
	User              *string    `cty:"user"`
	CtyVal            *cty.Value `diff:"-"`
}

func (r *AwsIamAccessKey) TerraformId() string {
	return r.Id
}

func (r *AwsIamAccessKey) TerraformType() string {
	return AwsIamAccessKeyResourceType
}

func (r *AwsIamAccessKey) CtyValue() *cty.Value {
	return r.CtyVal
}

func awsIamAccessKeyNormalizer(val *map[string]interface{}) {
}

// GENERATED, DO NOT EDIT THIS FILE
package aws

import "github.com/zclconf/go-cty/cty"

const AwsIamRolePolicyAttachmentResourceType = "aws_iam_role_policy_attachment"

type AwsIamRolePolicyAttachment struct {
	Id        string     `cty:"id" computed:"true"`
	PolicyArn *string    `cty:"policy_arn"`
	Role      *string    `cty:"role"`
	CtyVal    *cty.Value `diff:"-"`
}

func (r *AwsIamRolePolicyAttachment) TerraformId() string {
	return r.Id
}

func (r *AwsIamRolePolicyAttachment) TerraformType() string {
	return AwsIamRolePolicyAttachmentResourceType
}

func (r *AwsIamRolePolicyAttachment) CtyValue() *cty.Value {
	return r.CtyVal
}

func awsIamRolePolicyAttachmentNormalizer(val *map[string]interface{}) {
}

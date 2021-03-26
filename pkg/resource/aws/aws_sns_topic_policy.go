// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"
)

const AwsSnsTopicPolicyResourceType = "aws_sns_topic_policy"

type AwsSnsTopicPolicy struct {
	Arn    *string    `cty:"arn"`
	Id     string     `cty:"id" computed:"true"`
	Policy *string    `cty:"policy" jsonstring:"true"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsSnsTopicPolicy) TerraformId() string {
	return r.Id
}

func (r *AwsSnsTopicPolicy) TerraformType() string {
	return AwsSnsTopicPolicyResourceType
}

func (r *AwsSnsTopicPolicy) CtyValue() *cty.Value {
	return r.CtyVal
}

func awsSnsTopicPolicyNormalizer(val *map[string]interface{}) {
}

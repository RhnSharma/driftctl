package aws

import "github.com/cloudskiff/driftctl/pkg/resource"

func InitNormalizers() {
	resource.AddNormalizer(AwsAmiResourceType, awsAmiNormalizer)
	resource.AddNormalizer(AwsCloudfrontDistributionResourceType, awsCloudfrontDistributionNormalizer)
	resource.AddNormalizer(AwsDbInstanceResourceType, awsDbInstanceDistributionNormalizer)
	resource.AddNormalizer(AwsDbSubnetGroupResourceType, awsDbSubnetGroupNormalizer)
	resource.AddNormalizer(AwsDefaultRouteTableResourceType, awsDefaultRouteTableNormalizer)
	resource.AddNormalizer(AwsDefaultSecurityGroupResourceType, awsDefaultSecurityGroupNormalizer)
	resource.AddNormalizer(AwsDefaultSubnetResourceType, awsDefaultSubnetNormalizer)
	resource.AddNormalizer(AwsDefaultVpcResourceType, awsDefaultVpcNormalizer)
	resource.AddNormalizer(AwsDynamodbTableResourceType, awsDynamodbTableNormalizer)
	resource.AddNormalizer(AwsEbsSnapshotResourceType, awsEbsSnapshotNormalizer)
	resource.AddNormalizer(AwsEbsVolumeResourceType, awsEbsVolumeNormalizer)
	resource.AddNormalizer(AwsEcrRepositoryResourceType, awsEcrRepositoryNormalizer)
	resource.AddNormalizer(AwsEipResourceType, awsEipNormalizer)
	resource.AddNormalizer(AwsEipAssociationResourceType, awsEipAssociationNormalizer)
	resource.AddNormalizer(AwsIamAccessKeyResourceType, awsIamAccessKeyNormalizer)
	resource.AddNormalizer(AwsIamPolicyResourceType, awsIamPolicyNormalizer)
	resource.AddNormalizer(AwsIamPolicyAttachmentResourceType, awsIamPolicyAttachmentNormalizer)
	resource.AddNormalizer(AwsIamRoleResourceType, awsIamRoleNormalizer)
	resource.AddNormalizer(AwsIamRolePolicyResourceType, awsIamRolePolicyNormalizer)
	resource.AddNormalizer(AwsIamRolePolicyAttachmentResourceType, awsIamRolePolicyAttachmentNormalizer)
	resource.AddNormalizer(AwsIamUserResourceType, awsIamUserNormalizer)
	resource.AddNormalizer(AwsIamUserPolicyResourceType, awsIamUserPolicyNormalizer)
	resource.AddNormalizer(AwsIamUserPolicyAttachmentResourceType, awsIamUserPolicyAttachmentNormalizer)
	resource.AddNormalizer(AwsInstanceResourceType, awsInstanceNormalizer)
	resource.AddNormalizer(AwsInternetGatewayResourceType, awsInternetGatewayNormalizer)
	resource.AddNormalizer(AwsKeyPairResourceType, awsKeyPairNormalizer)
	resource.AddNormalizer(AwsKmsAliasResourceType, awsKmsAliasNormalizer)
	resource.AddNormalizer(AwsKmsKeyResourceType, awsKmsKeyNormalizer)
	resource.AddNormalizer(AwsLambdaEventSourceMappingResourceType, awsLambdaEventSourceMappingNormalizer)
	resource.AddNormalizer(AwsLambdaFunctionResourceType, awsLambdaFunctionNormalizer)
	resource.AddNormalizer(AwsNatGatewayResourceType, awsNatGatewayNormalizer)
	resource.AddNormalizer(AwsRouteResourceType, awsRouteNormalizer)
	resource.AddNormalizer(AwsRoute53HealthCheckResourceType, awsRoute53HealthCheckNormalizer)
	resource.AddNormalizer(AwsRoute53RecordResourceType, awsRoute53RecordNormalizer)
	resource.AddNormalizer(AwsRoute53ZoneResourceType, awsRoute53ZoneNormalizer)
	resource.AddNormalizer(AwsRouteTableResourceType, awsRouteTableNormalizer)
	resource.AddNormalizer(AwsRouteTableAssociationResourceType, awsRouteTableAssociationNormalizer)
	resource.AddNormalizer(AwsS3BucketResourceType, awsS3BucketNormalizer)
	resource.AddNormalizer(AwsS3BucketAnalyticsConfigurationResourceType, awsS3BucketAnalyticsConfigurationNormalizer)
	resource.AddNormalizer(AwsS3BucketInventoryResourceType, awsS3BucketInventoryNormalizer)
	resource.AddNormalizer(AwsS3BucketMetricResourceType, awsS3BucketMetricNormalizer)
	resource.AddNormalizer(AwsS3BucketNotificationResourceType, awsS3BucketNotificationNormalizer)
	resource.AddNormalizer(AwsS3BucketPolicyResourceType, awsS3BucketPolicyNormalizer)
	resource.AddNormalizer(AwsSecurityGroupResourceType, awsSecurityGroupNormalizer)
	resource.AddNormalizer(AwsSecurityGroupRuleResourceType, awsSecurityGroupRuleNormalizer)
	resource.AddNormalizer(AwsSnsTopicResourceType, awsSnsTopicNormalizer)
	resource.AddNormalizer(AwsSnsTopicPolicyResourceType, awsSnsTopicPolicyNormalizer)
	resource.AddNormalizer(AwsSnsTopicSubscriptionResourceType, awsSnsTopicSubscriptionNormalizer)
	resource.AddNormalizer(AwsSqsQueueResourceType, awsSqsQueueNormalizer)
	resource.AddNormalizer(AwsSqsQueuePolicyResourceType, awsSqsQueuePolicyNormalizer)
	resource.AddNormalizer(AwsSubnetResourceType, awsSubnetNormalizer)
	resource.AddNormalizer(AwsVpcResourceType, awsVpcNormalizer)

}

package aws

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const AwsRoute53RecordResourceType = "aws_route53_record"

func initAwsRoute53RecordMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsRoute53RecordResourceType, func(res *resource.AbstractResource) {
		val := res.Attrs
		val.DeleteIfDefault("health_check_id")
		val.DeleteIfDefault("set_identifier")
		val.DeleteIfDefault("ttl")
		val.SafeDelete([]string{"name"})
		val.SafeDelete([]string{"allow_overwrite"})
	})
	resourceSchemaRepository.SetHumanReadableAttributesFunc(AwsRoute53RecordResourceType, func(res *resource.AbstractResource) map[string]string {
		val := res.Attrs
		attrs := make(map[string]string)
		if fqdn := val.GetString("fqdn"); fqdn != nil && *fqdn != "" {
			attrs["Fqdn"] = *fqdn
		}
		if ty := val.GetString("type"); ty != nil && *ty != "" {
			attrs["Type"] = *ty
		}
		if zoneID := val.GetString("zone_id"); zoneID != nil && *zoneID != "" {
			attrs["ZoneId"] = *zoneID
		}
		return attrs
	})
}

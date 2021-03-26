// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/helpers"
)

const AwsRouteTableResourceType = "aws_route_table"

type AwsRouteTable struct {
	Id              string    `cty:"id" computed:"true"`
	OwnerId         *string   `cty:"owner_id" computed:"true"`
	PropagatingVgws *[]string `cty:"propagating_vgws" computed:"true"` // Could be null in state
	Route           *[]struct {
		CidrBlock              *string `cty:"cidr_block"`
		EgressOnlyGatewayId    *string `cty:"egress_only_gateway_id"`
		GatewayId              *string `cty:"gateway_id"`
		InstanceId             *string `cty:"instance_id"`
		Ipv6CidrBlock          *string `cty:"ipv6_cidr_block"`
		LocalGatewayId         *string `cty:"local_gateway_id"`
		NatGatewayId           *string `cty:"nat_gateway_id"`
		NetworkInterfaceId     *string `cty:"network_interface_id"`
		TransitGatewayId       *string `cty:"transit_gateway_id"`
		VpcEndpointId          *string `cty:"vpc_endpoint_id"`
		VpcPeeringConnectionId *string `cty:"vpc_peering_connection_id"`
	} `cty:"route" computed:"true" diff:"-"`
	Tags   map[string]string `cty:"tags"`
	VpcId  *string           `cty:"vpc_id"`
	CtyVal *cty.Value        `diff:"-"`
}

func (r *AwsRouteTable) TerraformId() string {
	return r.Id
}

func (r *AwsRouteTable) TerraformType() string {
	return AwsRouteTableResourceType
}

func (r *AwsRouteTable) CtyValue() *cty.Value {
	return r.CtyVal
}

func awsRouteTableNormalizer(val *map[string]interface{}) {
	helpers.SafeDelete(val, []string{"route"})
}

package remote

import (
	"github.com/cloudskiff/driftctl/pkg/alerter"
	"github.com/cloudskiff/driftctl/pkg/remote/aws"
	"github.com/cloudskiff/driftctl/pkg/remote/github"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/terraform"
	"github.com/pkg/errors"
)

var supportedRemotes = []string{
	aws.RemoteAWSTerraform,
	github.RemoteGithubTerraform,
}

func IsSupported(remote string) bool {
	for _, r := range supportedRemotes {
		if r == remote {
			return true
		}
	}
	return false
}

func Activate(remote string, alerter *alerter.Alerter, providerLibrary *terraform.ProviderLibrary, supplierLibrary *resource.SupplierLibrary) error {
	switch remote {
	case aws.RemoteAWSTerraform:
		return aws.Init(alerter, providerLibrary, supplierLibrary)
	case github.RemoteGithubTerraform:
		return github.Init(alerter, providerLibrary, supplierLibrary)
	default:
		return errors.Errorf("unsupported remote '%s'", remote)
	}
}

func GetSupportedRemotes() []string {
	return supportedRemotes
}

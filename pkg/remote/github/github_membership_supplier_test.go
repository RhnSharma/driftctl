package github

import (
	"context"
	"testing"

	"github.com/cloudskiff/driftctl/pkg/parallel"
	"github.com/cloudskiff/driftctl/pkg/resource"
	resourcegithub "github.com/cloudskiff/driftctl/pkg/resource/github"
	"github.com/cloudskiff/driftctl/pkg/terraform"
	"github.com/cloudskiff/driftctl/test"
	"github.com/cloudskiff/driftctl/test/goldenfile"
	dritftctlmocks "github.com/cloudskiff/driftctl/test/mocks"
	testresource "github.com/cloudskiff/driftctl/test/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGithubMembershipupplier_Resources(t *testing.T) {
	cases := []struct {
		test    string
		dirName string
		mocks   func(client *MockGithubRepository)
		err     error
	}{
		{
			test:    "no members",
			dirName: "github_membership_empty",
			mocks: func(client *MockGithubRepository) {
				client.On("ListMembership").Return([]string{}, nil)
			},
			err: nil,
		},
		{
			test:    "Multiple membership with admin and member roles",
			dirName: "github_membership_multiple",
			mocks: func(client *MockGithubRepository) {
				client.On("ListMembership").Return([]string{
					"driftctl-test:driftctl-acceptance-tester",
					"driftctl-test:eliecharra",
				}, nil)
			},
			err: nil,
		},
	}
	for _, c := range cases {
		shouldUpdate := c.dirName == *goldenfile.Update

		providerLibrary := terraform.NewProviderLibrary()
		supplierLibrary := resource.NewSupplierLibrary()

		repo := testresource.InitFakeSchemaRepository(terraform.GITHUB, "4.4.0")
		resourcegithub.InitResourcesMetadata(repo)
		factory := terraform.NewTerraformResourceFactory(repo)

		deserializer := resource.NewDeserializer(factory)

		mockedRepo := MockGithubRepository{}
		c.mocks(&mockedRepo)

		if shouldUpdate {
			provider, err := InitTestGithubProvider(providerLibrary)
			if err != nil {
				t.Fatal(err)
			}

			supplierLibrary.AddSupplier(NewGithubMembershipSupplier(provider, &mockedRepo, deserializer))
		}

		t.Run(c.test, func(tt *testing.T) {
			provider := dritftctlmocks.NewMockedGoldenTFProvider(c.dirName, providerLibrary.Provider(terraform.GITHUB), shouldUpdate)
			s := &GithubMembershipSupplier{
				provider,
				deserializer,
				&mockedRepo,
				terraform.NewParallelResourceReader(parallel.NewParallelRunner(context.TODO(), 10)),
			}
			got, err := s.Resources()
			assert.Equal(tt, c.err, err)

			mock.AssertExpectationsForObjects(tt)
			test.CtyTestDiff(got, c.dirName, provider, deserializer, shouldUpdate, tt)
		})
	}
}

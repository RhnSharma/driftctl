package github

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GithubRepository interface {
	ListRepositories() ([]repository, error)
}

type GithubGraphQLClient interface {
	Query(ctx context.Context, q interface{}, variables map[string]interface{}) error
}

type githubRepository struct {
	client GithubGraphQLClient
	ctx    context.Context
	config githubConfig
}

type repository struct {
	Name string
}

func NewGithubRepository(config githubConfig) *githubRepository {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)
	oauthClient := oauth2.NewClient(ctx, ts)

	repo := &githubRepository{
		client: githubv4.NewClient(oauthClient),
		ctx:    context.Background(),
		config: config,
	}

	return repo
}

func (r *githubRepository) ListRepositories() ([]repository, error) {
	if r.config.Organization != "" {
		return r.listRepoForOrg()
	}
	return r.listRepoForOwner()
}

type pageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}

type listRepoForOrgQuery struct {
	Organization struct {
		Repositories struct {
			Nodes    []repository
			PageInfo pageInfo
		} `graphql:"repositories(first: 100, after: $cursor)"`
	} `graphql:"organization(login: $org)"`
}

func (r *githubRepository) listRepoForOrg() ([]repository, error) {
	query := listRepoForOrgQuery{}
	variables := map[string]interface{}{
		"org":    (githubv4.String)(r.config.Organization),
		"cursor": (*githubv4.String)(nil),
	}
	var results []repository
	for {
		err := r.client.Query(r.ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		results = append(results, query.Organization.Repositories.Nodes...)
		if !query.Organization.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Organization.Repositories.PageInfo.EndCursor)
	}
	return results, nil
}

type listRepoForOwnerQuery struct {
	Viewer struct {
		Repositories struct {
			Nodes    []repository
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"repositories(first: 100, after: $cursor)"`
	}
}

func (r githubRepository) listRepoForOwner() ([]repository, error) {
	query := listRepoForOwnerQuery{}
	variables := map[string]interface{}{
		"cursor": (*githubv4.String)(nil),
	}
	var results []repository
	for {
		err := r.client.Query(r.ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		results = append(results, query.Viewer.Repositories.Nodes...)
		if !query.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Viewer.Repositories.PageInfo.EndCursor)
	}
	return results, nil
}

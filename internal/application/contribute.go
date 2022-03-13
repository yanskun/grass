package application

import (
	"context"

	"github.com/google/go-github/v43/github"
)

type (
	ContributeAppSrv interface {
		Get(ctx context.Context) (bool, error)
	}

	contributeAppSrv struct {
		githubClient github.Client
	}
)

func NewContributeAppSrv(githubClient github.Client) ContributeAppSrv {
	return &contributeAppSrv{
		githubClient: githubClient,
	}
}

func (c *contributeAppSrv) Get(ctx context.Context) (bool, error) {
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := c.githubClient.Repositories.ListByOrg(ctx, "github", opt)
	if err != nil {
		return false, err
	}

	result := len(repos) > 0

	return result, nil
}

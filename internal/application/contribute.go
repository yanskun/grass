package application

import (
	"context"

	"github.com/google/go-github/v43/github"
)

type (
	ContributeAppSrv interface {
		Get(ctx context.Context, cmd ContributeAppSrvCmd) (bool, error)
	}

	ContributeAppSrvCmd struct {
		username string
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

func (c *contributeAppSrv) Get(
	ctx context.Context,
	cmd ContributeAppSrvCmd,
) (bool, error) {
	opt := &github.ListOptions{}
	// repos, _, err := c.githubClient.Repositories.ListByOrg(ctx, "github", opt)
	orgs, _, err := c.githubClient.Organizations.List(ctx, cmd.username, opt)
	if err != nil {
		return false, err
	}

	orgNames := make([]string, len(orgs))
	for i, org := range orgs {
		orgNames[i] = org.GetLogin()
	}

	return len(orgNames) > 0, nil
}

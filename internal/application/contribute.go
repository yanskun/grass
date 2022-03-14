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
	orgNames, err := c.getOrgNames(ctx, cmd)
	if err != nil {
		return false, err
	}

	awners := make([]string, len(orgNames)+1)
	awners = append(awners, orgNames...)
	awners = append(awners, cmd.username)

	return len(orgNames) > 1, nil
}

func (c *contributeAppSrv) getOrgNames(
	ctx context.Context,
	cmd ContributeAppSrvCmd,
) ([]string, error) {
	opt := &github.ListOptions{}
	orgs, _, err := c.githubClient.Organizations.List(ctx, cmd.username, opt)
	if err != nil {
		return nil, err
	}

	orgNames := make([]string, len(orgs))
	for i, org := range orgs {
		orgNames[i] = org.GetLogin()
	}

	return orgNames, nil
}

func (c *contributeAppSrv) getAllRepos(
	ctx context.Context,
	cmd ContributeAppSrvCmd,
	orgNames []string,
) ([]*github.Repository, error) {
	oopt := &github.RepositoryListByOrgOptions{}

	repos := make([]*github.Repository, 0)
	for _, orgName := range orgNames {
		ors, _, err := c.githubClient.Repositories.ListByOrg(ctx, orgName, oopt)
		if err != nil {
			return nil, err
		}

		repos = append(repos, ors...)
	}

	uopt := &github.RepositoryListOptions{}
	urepos, _, err := c.githubClient.Repositories.List(ctx, cmd.username, uopt)
	if err != nil {
		return nil, err
	}

	repos = append(repos, urepos...)

	return repos, nil
}

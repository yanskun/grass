package application

import (
	"context"
	"time"

	"github.com/google/go-github/v43/github"
)

type (
	ContributeAppSrv interface {
		Get(ctx context.Context, cmd ContributeAppSrvCmd) (bool, error)
	}

	ContributeAppSrvCmd struct {
		Username string
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

	awners := make([]string, 0)
	awners = append(awners, orgNames...)

	repos, err := c.getAllRepos(ctx, cmd, awners)
	if err != nil {
		return false, err
	}

	isCommitted, err := c.findTodayCommit(ctx, cmd, repos)
	if err != nil {
		return false, err
	}

	return isCommitted, nil
}

func (c *contributeAppSrv) getOrgNames(
	ctx context.Context,
	cmd ContributeAppSrvCmd,
) ([]string, error) {
	opt := &github.ListOptions{}
	orgs, _, err := c.githubClient.Organizations.List(ctx, cmd.Username, opt)
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
	urepos, _, err := c.githubClient.Repositories.List(ctx, cmd.Username, uopt)
	if err != nil {
		return nil, err
	}

	repos = append(repos, urepos...)

	return repos, nil
}

func (c *contributeAppSrv) findTodayCommit(
	ctx context.Context,
	cmd ContributeAppSrvCmd,
	repos []*github.Repository,
) (bool, error) {
	var isCommitted bool
	for _, repo := range repos {

		today := time.Now()
		yesterday := today.AddDate(0, 0, -1)
		opt := &github.CommitsListOptions{
			Author: cmd.Username,
			Until:  today,
			Since:  yesterday,
		}
		owner := repo.GetOwner().GetLogin()
		contributors, _, err :=
			c.githubClient.Repositories.ListCommits(
				ctx,
				owner,
				repo.GetName(),
				opt,
			)

		if err != nil {
			return false, err
		}

		if len(contributors) > 0 {
			isCommitted = true
			break
		}
	}

	return isCommitted, nil
}

package testdata

import (
	"github.com/google/go-github/v43/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
)

func MockGitHubClient() *github.Client {
	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatch(
			mock.GetUsersOrgsByUsername,
			[]github.Organization{
				{
					Login: github.String("japan"),
				},
			},
		),
		mock.WithRequestMatch(
			mock.GetOrgsReposByOrg,
			[]github.Repository{
				{
					Name: github.String("Mt_Fuji"),
				},
			},
		),
		mock.WithRequestMatch(
			mock.GetUsersReposByUsername,
			[]github.Repository{
				{
					Name: github.String("skytree"),
				},
				{
					Name: github.String("kokugikan"),
				},
			},
		),
	)

	return github.NewClient(mockedHTTPClient)
}

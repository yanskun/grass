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
					Login: github.String("g-boys"),
				},
			},
		),
	)

	return github.NewClient(mockedHTTPClient)
}

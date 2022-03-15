package testdata

import (
	"time"

	"github.com/google/go-github/v43/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
)

func MockGitHubClient() *github.Client {
	now := time.Now()

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
		mock.WithRequestMatch(
			mock.GetReposCommitsByOwnerByRepo,
			[]github.RepositoryCommit{
				{
					Author: &github.User{
						Login: github.String("tokyo"),
					},
					Commit: &github.Commit{
						Author: &github.CommitAuthor{
							Date:  &now,
							Name:  github.String("tokyo"),
							Login: github.String("tokyo"),
						},
					},
				},
			},
		),
	)

	return github.NewClient(mockedHTTPClient)
}

func MockGitHubClientWithEmpty() *github.Client {
	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatch(
			mock.GetReposCommitsByOwnerByRepo,
			[]github.RepositoryCommit{},
		),
	)

	return github.NewClient(mockedHTTPClient)
}

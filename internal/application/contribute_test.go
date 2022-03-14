package application

import (
	"context"
	"testing"

	"github.com/google/go-github/v43/github"
	"github.com/stretchr/testify/assert"
	"github.com/yasudanaoya/grass/test/testdata"
)

func Test_contributeAppSrv_getOrgNames(t *testing.T) {
	client := testdata.MockGitHubClient()
	var s []string

	type fields struct {
		githubClient github.Client
	}
	type args struct {
		ctx context.Context
		cmd ContributeAppSrvCmd
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "Success",
			fields: fields{
				githubClient: *client,
			},
			args: args{
				ctx: context.Background(),
				cmd: ContributeAppSrvCmd{
					username: "king",
				},
			},
			want:      []string{"g-boys"},
			assertion: assert.NoError,
		},
		{
			name: "Error",
			fields: fields{
				githubClient: *client,
			},
			args: args{
				ctx: context.Background(),
			},
			want:      s,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &contributeAppSrv{
				githubClient: tt.fields.githubClient,
			}
			got, err := c.getOrgNames(tt.args.ctx, tt.args.cmd)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

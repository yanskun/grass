package application

import (
	"context"
	"testing"

	"github.com/google/go-github/v43/github"
	"github.com/stretchr/testify/assert"
)

func Test_contributeAppSrv_Get(t *testing.T) {
	type fields struct {
		githubClient github.Client
	}
	type args struct {
		ctx context.Context
		cmd ContributeAppSrvCmd
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Success",
			fields: fields{
				githubClient: *github.NewClient(nil),
			},
			args: args{
				ctx: context.Background(),
				cmd: ContributeAppSrvCmd{
					username: "octocat",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &contributeAppSrv{
				githubClient: tt.fields.githubClient,
			}
			got, err := c.Get(tt.args.ctx, tt.args.cmd)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

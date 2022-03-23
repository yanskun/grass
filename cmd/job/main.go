package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/v43/github"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"github.com/yasudanaoya/grass/internal/application"
	"golang.org/x/oauth2"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			ctx := context.Background()
			err := godotenv.Load()

			if err != nil {
				log.Fatal("Error loading .env file")
			}
			token := os.Getenv("TOKEN")
			url := os.Getenv("WEBHOOK_URL")

			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)

			srv := application.NewContributeAppSrv(*client)
			cmd := application.ContributeAppSrvCmd{
				Username: "yasudanaya",
			}

			result, err := srv.Get(ctx, cmd)
			if err != nil {
				log.Fatal(err)
			}

			var message string

			if result {
				message = "今日は commit しました。"
			} else {
				message = "今日まだ commit していません。"
			}
			byte, err := json.Marshal(SlackMessage{Text: message})

			http.Post(
				url,
				"aplication/json",
				bytes.NewBufferString(string(byte)),
			)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			println("Hello World!")
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
			_ = os.Getenv("GITHUB_TOKEN")

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

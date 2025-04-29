package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Super2FA",
		Usage: "A CLI Tool that helps you to manage your 2FA tokens.",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

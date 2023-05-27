package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	var flagApiKey string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}
	configData := NewConfig(homeDir, ".gptsh")

	app := &cli.App{
		Name:                 "gptsh",
		Usage:                "Use GPT from shell",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "setApiKey",
				Usage:       "Set Openai API Key",
				Destination: &flagApiKey,
			},
		},
		Action: func(ctx *cli.Context) error {
			if len(flagApiKey) != 0 {
				fmt.Println("Work in progress")
				return nil
			} else {
				apiKey := configData.getOpenaiApiKey()
				if len(apiKey) == 0 {
					_, err = configData.askOpenaiApiKey()
					if err != nil {
						log.Fatal(err)
					}
				}

				openai := NewOpenai(configData)

				prompt := strings.Join(ctx.Args().Slice(), " ")
				completionResponse := openai.chatCompletion(prompt)
				fmt.Print(completionResponse)

				return nil
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

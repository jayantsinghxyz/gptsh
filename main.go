package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var flagApiKey string
	
	app := &cli.App{
	Name:  "gptsh",
	Usage: "Use GPT from shell",
	EnableBashCompletion: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "setApiKey",
			Usage:       "Set Openai API Key",
			Destination: &flagApiKey,
		},
	},
	Action: func(ctx *cli.Context) error {
		if (len(flagApiKey) != 0) {
			fmt.Println("Work in progress")
			return nil
			} else {
				homeDir, err := os.UserHomeDir()
				if (err != nil) {
					fmt.Println(err)
				}
				
				configData := NewConfig(homeDir, ".gptsh")
				apiKey, err := configData.getOpenaiApiKey()
				if err != nil {
					if (errors.Is(err, os.ErrNotExist)) {
						apiKey, err = configData.askOpenaiApiKey()
						if err != nil {
							log.Fatal(err)
						}
					} else {
						log.Fatal(err)
					}
				}
				
				fmt.Println(flagApiKey, apiKey)
				fmt.Println("ctx.App.Name", ctx.Args().Slice(), flagApiKey)
				return nil	
			}
		},
	}
	
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
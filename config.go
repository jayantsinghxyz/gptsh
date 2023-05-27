package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	config map[string]string
	path   string
}

func NewConfig(path string, filename string) *config {
	configPath := fmt.Sprintf("%s/%s", path, filename)
	configData, err := godotenv.Read(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
		}
	}
	return &config{config: configData, path: configPath}
}

func (ctx config) getOpenaiApiKey() (key string) {
	apiKey := ctx.config["OPENAI_API_KEY"]
	return apiKey
}

func (ctx config) askOpenaiApiKey() (key string, err error) {
	fmt.Print("Openai API Key: ")
	var apiKey string
	fmt.Scan(&apiKey)
	err = ctx.setOpenaiApiKey(apiKey)
	if err != nil {
		return "", err
	}
	return ctx.getOpenaiApiKey(), nil
}

func (ctx config) setOpenaiApiKey(key string) error {
	ctx.config["OPENAI_API_KEY"] = key
	err := ctx.writeConfig()
	return err
}

func (ctx config) writeConfig() error {
	configString, err := godotenv.Marshal(ctx.config)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(ctx.path, []byte(configString), 0644)
	return err
}

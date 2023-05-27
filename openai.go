package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type openai struct {
	Config *config
}

type openaiChatComletionPostBody struct {
	Model       string                         `json:"model"`
	Messages    []openaiChatCompletionMessages `json:"messages"`
	Temperature float32                        `json:"temperature"`
}

type openaiChatCompletionMessages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openaiChatCompletionChoices struct {
	Message openaiChatCompletionMessages
}

type openaiChatCompletionResponse struct {
	Id      string                        `json:"id"`
	Object  string                        `json:"object"`
	Choices []openaiChatCompletionChoices `json:"choices"`
}

func NewOpenai(config *config) *openai {
	return &openai{
		Config: config,
	}
}

func (ctx *openai) chatCompletion(prompt string) string {
	bodyData := openaiChatComletionPostBody{
		Model:       "gpt-3.5-turbo",
		Temperature: 0.7,
		Messages:    []openaiChatCompletionMessages{{Role: "user", Content: prompt}},
	}
	body, _ := json.Marshal(bodyData)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ctx.Config.getOpenaiApiKey()))
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 401 {
		log.Fatal("Invalid API Key")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respData openaiChatCompletionResponse
	json.Unmarshal(respBody, &respData)
	return respData.Choices[0].Message.Content
}

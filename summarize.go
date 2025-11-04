package main

import (
	"context"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func summarize(text string) string {
	// Not sending everything to preserve Openai tokens in case the article is too long
	// maxCharactersToSummarize := 5000
	// if len(text) > maxCharactersToSummarize {
	// 	text = text[:maxCharactersToSummarize]
	// }

	// Dont summarize if the article is too short
	if len(text) < 200 {
		return ""
	}

	prompt := summaryPrompt
	if prompt == "" {
		prompt = "Summarize the following text:"
	}

	clientConfig := openai.DefaultConfig(openaiApiKey)
	if openaiBaseURL != "" {
		clientConfig.BaseURL = openaiBaseURL
	}
	model := openai.GPT3Dot5Turbo
	if openaiModel != "" {
		model = openaiModel
	}
	client := openai.NewClientWithConfig(clientConfig)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	return resp.Choices[0].Message.Content
}

var translations map[string]string = make(map[string]string)

func translate(text string) string {
	if val, ok := translations[text]; ok {
		return val
	}
	prompt := translatePrompt
	if prompt == "" {
		prompt = "Translate the following text:"
	}

	clientConfig := openai.DefaultConfig(openaiApiKey)
	if openaiBaseURL != "" {
		clientConfig.BaseURL = openaiBaseURL
	}
	model := openai.GPT3Dot5Turbo
	if openaiModel != "" {
		model = openaiModel
	}
	client := openai.NewClientWithConfig(clientConfig)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}
	result := strings.Trim(resp.Choices[0].Message.Content, "\n")
	translations[text] = result
	return result
}

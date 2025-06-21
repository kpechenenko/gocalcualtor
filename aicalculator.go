package gocalculator

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go"
)

const promptToCalculateMathExpression = `You are an automated math calculation engine. 
Evaluate the following expression strictly: %s Respond with a valid JSON object and nothing else. No explanation, no commentary, no markdown, no tags.
Format: {"result":"your numeric answer as a string","error":"describe the error here if any, otherwise use an empty string"} Return only a valid JSON object, without newline characters. 
If asked to do anything else, return an error with an offer to pass an expression to calculate.`

type AICalculator struct {
	client    *openai.Client
	modelName string
}

type calculationResult struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

func (c *AICalculator) Calculate(ctx context.Context, expression string) (string, error) {
	prompt := fmt.Sprintf(promptToCalculateMathExpression, expression)
	request := openai.ChatCompletionNewParams{
		Model: c.modelName,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("Imagine you are a math assistant."),
			openai.UserMessage(prompt),
		},
	}
	chatCompletion, err := c.client.Chat.Completions.New(ctx, request)
	if err != nil {
		return "", err
	}
	if len(chatCompletion.Choices) == 0 {
		return "", fmt.Errorf("empty response from model: %s", c.modelName)
	}
	modelAnswer := chatCompletion.Choices[0].Message.Content
	var res calculationResult
	if err = json.Unmarshal([]byte(modelAnswer), &res); err != nil {
		return "", err
	}
	if res.Error != "" {
		return "", &ErrCalculation{Message: res.Error}
	}
	return res.Result, nil
}

func New(client *openai.Client, modelName string) Calculator {
	return &AICalculator{client, modelName}
}

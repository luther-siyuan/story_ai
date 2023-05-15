package chatgpt

import (
	"context"
	"fmt"
)

var Cache CacheInterface

func init() {
	Cache = GetSessionCache()
}

func Completions(session, msg string) (string, error) {
	fmt.Println("gpt chart start ----->")
	ctx := context.Background()

	resourceName := "laiye-openai"
	deploymentName := "chatgpt-test"
	apiVersion := "2023-03-15-preview"
	accessToken := "93dea89212dc43ae91f657fc8d4d514c"

	client := New(resourceName, deploymentName, apiVersion, accessToken)

	ms := Cache.GetMsg(session)
	if len(ms) == 0 {
		ms = append(ms, ChatMessage{
			Role:    "system",
			Content: "you are a helpful chatbot",
		})
	}

	ms = append(ms, ChatMessage{
		Role:    "user",
		Content: msg,
	})

	req := ChatRequest{
		Messages: ms,
	}
	resp, err := client.ChatCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return err.Error(), err
	}

	fmt.Println("gpt chat success ----->")

	ms = append(ms, resp.Choices[0].Message)
	Cache.SetMsg(session, ms)
	return "gpt 回答：" + resp.Choices[0].Message.Content, nil
}

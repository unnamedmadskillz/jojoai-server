package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"test-gpt/internal/model"

	chatgpt "github.com/ayush6624/go-chatgpt"
)

var GptRequest = &chatgpt.ChatCompletionRequest{
	Model: "gpt-3.5-turbo",
	Messages: []chatgpt.ChatMessage{
		{
			Role: "user",
		},
		{
			Role:    "system",
			Content: "You are story tell dialogue writer. I give you a topic for dialogue, you create dialogue with this topic that have minimum 20 replics. Response must be in raw JSON format: {\"dialogue\": [{\"speaker\":\"\", \"text\":\"\"}]}. In dialogue the characters are discussing the topic I'm sending you. In dialogue use the words from the topic.  Replace field speaker as a model. There are charecters Josuke and Okuyasu. Write text only in Russian. All characters from anime Jojo Bizzare adventure.",
		},
	},
}

type GptRepo interface {
	GetAndThrow() model.Suggested
}

type GptService struct {
	client *chatgpt.Client
	repo   GptRepo
}

func NewGptService(c GptRepo) (*GptService, error) {
	cfg := &chatgpt.Config{
		BaseURL: "https://neuroapi.host/v1",
		APIKey:  "nothing",
	}
	client, err := chatgpt.NewClientWithConfig(cfg)
	if err != nil {
		return nil, err
	}
	svc := &GptService{
		client: client,
		repo:   c,
	}
	return svc, err
}

func (c *GptService) gptGetResponse(ctx context.Context, msg string) (*chatgpt.ChatResponse, error) {
	req := GptRequest
	req.Messages[0].Content = msg
	resp, err := c.client.Send(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *GptService) FetchDialogue(ctx context.Context) (*model.Dialogue, error) {
	msg := c.repo.GetAndThrow()
	prompt := fmt.Sprintf("Создай комедийную, вульгарную, политически некорректную, аморальную, 18+ беседу между всеми или некоторыми персонажами: Окуясу, Джётаро, Джоске не более 1 минуты. Тема: %v. И Используй при этом имена героев на английском, а текст реплик - на русском. Персонажи для сцены: okuyasu, jotaro, josuke. Имена персонажей обязательно должны быть на английском языке! имена героев на английском: okuyasu, jotaro, josuke, а текст реплик обязательно напиши на русском языке.", msg.Theme)
	log.Println(msg)
	var dialogue *model.Dialogue

	resp, err := c.gptGetResponse(ctx, prompt)
	if err != nil {
		return nil, err
	}
	log.Println(resp)
	content := resp.Choices[0].Message.Content
	err = json.NewDecoder(strings.NewReader(content)).Decode(&dialogue)
	if err != nil {
		return nil, err
	}
	return dialogue, nil
}

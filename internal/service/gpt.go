package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"test-gpt/internal/model"
	"time"

	chatgpt "github.com/ayush6624/go-chatgpt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var GptRequest = &chatgpt.ChatCompletionRequest{
	Model: "gpt-3.5-turbo",
	Messages: []chatgpt.ChatMessage{
		{
			Role: "user",
		},
		{
			Role:    "system",
			Content: "You are story tell dialogue writer. I give you a topic for dialogue, you create dialogue with this topic that have maximum 8 replics. Response must be in raw JSON format: {\"dialogue\": [{\"speaker\":\"\", \"text\":\"\"}]}. In dialogue the characters are discussing the topic I'm sending you. In dialogue use the words from the topic.  Replace field speaker as a model. There are charecters Josuke and Okuyasu. Write text only in Russian. All characters from anime Jojo Bizzare adventure.",
		},
	},
}

type GptRepo interface {
	GetSuggestedAndThrow() model.Suggested
	PutCompletedDialogue(d *model.ReplicDB) (*mongo.InsertOneResult, error)
}

type GptService struct {
	client *chatgpt.Client
	repo   GptRepo
}

func NewGptService(g GptRepo) *GptService {
	cfg := &chatgpt.Config{
		BaseURL: "https://neuroapi.host/v1",
		APIKey:  "nothing",
	}
	client, err := chatgpt.NewClientWithConfig(cfg)
	if err != nil {
		log.Fatal("can't create gpt client")
	}
	svc := &GptService{
		client: client,
		repo:   g,
	}
	return svc
}

func (g *GptService) gptGetResponse(ctx context.Context, msg string) (*chatgpt.ChatResponse, error) {
	req := GptRequest
	req.Messages[0].Content = msg
	resp, err := g.client.Send(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GptService) FetchDialogue(ctx context.Context) (*model.Dialogue, error) {
	msg := g.repo.GetSuggestedAndThrow()
	prompt := fmt.Sprintf("Создай комедийнdую беседу между всеми или некоторыми персонажами: Окуясу, Джётаро, Джоске не более 1 минуты. Тема: %v. И Используй при этом имена героев на английском, а текст реплик - на русском. Персонажи для сцены: okuyasu, jotaro, josuke. Имена персонажей обязательно должны быть на английском языке! имена героев на английском: okuyasu, jotaro, josuke, а текст реплик обязательно напиши на русском языке.", msg.Theme)
	log.Println(msg)
	var dialogue *model.Dialogue

	resp, err := g.gptGetResponse(ctx, prompt)
	if err != nil {
		return nil, err
	}
	log.Println(resp)
	content := resp.Choices[0].Message.Content
	err = json.NewDecoder(strings.NewReader(content)).Decode(&dialogue)
	if err != nil {

		return nil, fmt.Errorf("failed to create dialogue with content \"%v\":%v", content, err)
	}
	return dialogue, nil
}

func (g *GptService) MakeDialogueData(m *model.Dialogue) *model.ReplicDB {
	var replicData []model.ReplicRow
	for i, v := range m.Replic {
		var r model.ReplicRow
		r.Name = v.Speaker
		r.Text = v.Utterance
		r.Order = i
		r.Path = uuid.NewString()
		replicData = append(replicData, r)
	}
	rdb := &model.ReplicDB{ID: primitive.NewObjectID(), Data: replicData, CreatedAt: time.Now()}
	g.repo.PutCompletedDialogue(rdb)
	return rdb
}

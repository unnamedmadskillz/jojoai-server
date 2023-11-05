package endpoint

import (
	"context"
	"test-gpt/internal/model"

	"github.com/gofiber/fiber/v2"
)

type GptService interface {
	FetchDialogue(ctx context.Context) (*model.Dialogue, error)
	MakeDialogueData(m *model.Dialogue) *model.ReplicDB
}

type TTSService interface {
	MakeTTS()
}

type Endpoint struct {
	gptsvc GptService
	tts    TTSService
}

func NewEndpoint(g GptService, tts TTSService) *Endpoint {
	return &Endpoint{
		gptsvc: g,
		tts:    tts,
	}
}

func (e *Endpoint) CreateChat(fctx *fiber.Ctx) error {
	ctx := context.Background()
	d, err := e.gptsvc.FetchDialogue(ctx)
	data := e.gptsvc.MakeDialogueData(d)
	e.tts.MakeTTS()
	if err != nil {
		return err
	}
	err = fctx.JSON(data)
	if err != nil {
		return err
	}
	return nil
}

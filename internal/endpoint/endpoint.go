package endpoint

import (
	"context"
	"test-gpt/internal/model"

	"github.com/gofiber/fiber/v2"
)

type GptService interface {
	FetchDialogue(ctx context.Context) (*model.Dialogue, error)
}

type Endpoint struct {
	gptsvc GptService
}

func NewEndpoint(g GptService) *Endpoint {
	return &Endpoint{
		gptsvc: g,
	}
}

func (e *Endpoint) CreateChat(fctx *fiber.Ctx) error {
	ctx := context.Background()
	d, err := e.gptsvc.FetchDialogue(ctx)
	if err != nil {
		return err
	}
	err = fctx.JSON(d)
	if err != nil {
		return err
	}
	return nil
}

package service

import (
	"fmt"
	"log"
	"sync"
	"test-gpt/internal/model"

	"github.com/gofiber/fiber/v2"
)

type TTSRepo interface {
	GetDialogue() *model.ReplicDB
}

type TTSService struct {
	gptservice GptService
	httpclient *fiber.Client
	repo       TTSRepo
}

func NewTTSService(g GptService, t TTSRepo) *TTSService {
	return &TTSService{
		httpclient: &fiber.Client{},
		gptservice: g,
		repo:       t,
	}
}

func (tts *TTSService) MakeTTS() {
	var wg sync.WaitGroup
	dialogue := tts.repo.GetDialogue()
	for _, v := range dialogue.Data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tts.sendTTS(v.Name, v.Text, v.Path)
		}()
		wg.Wait()
	}

}

func (tts *TTSService) sendTTS(speaker string, text string, id string) []error {
	reqStr := fmt.Sprintf(`{"model_name": "%v", "tts_text": "%v", "output_file_path": "%v"}`, speaker, text, id)
	body := []byte(reqStr)
	req := tts.httpclient.Post("http://0.0.0.0:8000/generate").Body(body)
	statusCode, body, errs := req.Bytes()
	if errs != nil {
		log.Printf("tts request error %v", errs)
		return errs
	}
	log.Println(statusCode, string(body))
	return nil
}

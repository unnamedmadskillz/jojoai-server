package server

import (
	"fmt"
	"test-gpt/internal/config"
	"test-gpt/internal/endpoint"
	"test-gpt/internal/repository"
	"test-gpt/internal/service"
	mongodb "test-gpt/pkg/database"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	port     string
	router   *fiber.App
	endpoint *endpoint.Endpoint
	gpt      *service.GptService
	tts      *service.TTSService
}

func NewServer() (*Server, error) {
	cfg := config.NewCofnig()
	fmt.Println(cfg)
	mongoclient, err := mongodb.NewClient("mongodb+srv://numbx666:vkjpRxePuU0xBzAm@cluster0.bvwoae5.mongodb.net/?retryWrites=true&w=majority", "numbx666", "vkjpRxePuU0xBzAm")
	if err != nil {
		return nil, fmt.Errorf("can't create mongodb client %v", err)
	}
	s := &Server{}
	s.port = cfg.ServerConfig.Port
	s.gpt = service.NewGptService(repository.NewGptRepo(mongoclient.Database("core")))
	s.tts = service.NewTTSService(*s.gpt, repository.NewTTSRepo(mongoclient.Database("core")))
	s.router = fiber.New()
	s.endpoint = endpoint.NewEndpoint(s.gpt, s.tts)
	s.router.Get("/get", s.endpoint.CreateChat)
	return s, nil
}

func (s *Server) RunServer() error {
	err := s.router.Listen(s.port)
	if err != nil {
		return err
	}
	return nil
}

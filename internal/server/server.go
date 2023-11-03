package server

import (
	"test-gpt/internal/endpoint"
	"test-gpt/internal/repository"
	"test-gpt/internal/service"
	mongodb "test-gpt/pkg/database"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	router   *fiber.App
	endpoint *endpoint.Endpoint
	gpt      *service.GptService
}

func NewServer() (*Server, error) {
	var mongoclient, err = mongodb.NewClient("mongodb+srv://numbx666:vkjpRxePuU0xBzAm@cluster0.bvwoae5.mongodb.net/?retryWrites=true&w=majority", "numbx666", "vkjpRxePuU0xBzAm")
	s := &Server{}
	s.router = fiber.New()
	s.gpt, err = service.NewGptService(repository.NewGptRepo(mongoclient.Database("core")))
	if err != nil {
		return nil, err
	}
	s.endpoint = endpoint.NewEndpoint(s.gpt)
	s.router.Get("/get", s.endpoint.CreateChat)
	return s, nil
}

func (s *Server) RunServer() error {
	err := s.router.Listen(":8080")
	if err != nil {
		return err
	}
	return nil
}

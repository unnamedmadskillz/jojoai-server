package app

import (
	"test-gpt/internal/server"
)

type App struct {
	server *server.Server
}

func NewApp() (*App, error) {
	var err error

	a := &App{}
	a.server, err = server.NewServer()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	return a.server.RunServer()
}

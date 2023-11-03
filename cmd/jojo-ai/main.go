package main

import (
	"log"
	"test-gpt/internal/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Println(err)
	}
	log.Fatal(a.Run())
}

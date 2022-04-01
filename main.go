package main

import (
	"log"

	"github.com/pascallohrer/petstore/pkg/router"
	"github.com/pascallohrer/petstore/pkg/storage"
)

func main() {
	repository := storage.NewMemoryPetStorage()
	app := router.NewRouter(repository)

	exitError := make(chan error)
	go func() {
		exitError <- app.Listen("0.0.0.0:8080")
	}()

	if err := <-exitError; err != nil {
		log.Println(err)
	}
}

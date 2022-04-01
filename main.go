package main

import (
	"log"

	"github.com/pascallohrer/petstore/pkg/router"
)

func main() {
	app := router.NewRouter()

	exitError := make(chan error)
	go func() {
		exitError <- app.Listen("0.0.0.0:8080")
	}()

	if err := <-exitError; err != nil {
		log.Println(err)
	}
}

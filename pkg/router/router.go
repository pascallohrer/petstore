package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pascallohrer/petstore/pkg/handlers"
)

type Route struct {
	Method  string
	Path    string
	Handler func(handlers.PetStorage) func(*fiber.Ctx) error
}

func NewRouter(storage handlers.PetStorage) *fiber.App {
	app := fiber.New()
	for _, route := range routes {
		app.Add(route.Method, route.Path, route.Handler(storage))
	}
	return app
}

var routes = []Route{
	{
		Method:  "GET",
		Path:    "/pet/:petId",
		Handler: handlers.GetPetByIdHandler,
	},
	{
		Method:  "POST",
		Path:    "/pet",
		Handler: handlers.AddPetHandler,
	},
	{
		Method:  "DELETE",
		Path:    "/pet/:petId",
		Handler: handlers.DeletePetHandler,
	},
}

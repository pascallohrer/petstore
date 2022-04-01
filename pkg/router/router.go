package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pascallohrer/petstore/pkg/handlers"
)

type Route struct {
	Method  string
	Path    string
	Handler func(*fiber.Ctx) error
}

func NewRouter() *fiber.App {
	app := fiber.New()
	for _, route := range routes {
		app.Add(route.Method, route.Path, route.Handler)
	}
	return app
}

var routes = []Route{
	{
		Method:  "GET",
		Path:    "/pet/:petId",
		Handler: handlers.GetPetById,
	},
	{
		Method:  "POST",
		Path:    "/pet",
		Handler: handlers.AddPet,
	},
	{
		Method:  "DELETE",
		Path:    "/pet/:petId",
		Handler: handlers.DeletePet,
	},
}

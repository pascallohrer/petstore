package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pascallohrer/petstore/pkg/entities"
	"github.com/pascallohrer/petstore/pkg/storage"
)

func GetPetById(ctx *fiber.Ctx) error {
	petId, err := strconv.Atoi(ctx.Params("petId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	pet, err := storage.GetPetById(int64(petId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return ctx.JSON(pet)
}

func AddPet(ctx *fiber.Ctx) error {
	var newPet entities.Pet
	if err := ctx.BodyParser(&newPet); err != nil || !newPet.IsValid() {
		return ctx.Status(fiber.StatusMethodNotAllowed).SendString(err.Error())
	}
	petId := storage.AddPet(newPet)
	// Even though the spec doesn't specify it, returning the newly added ID just makes sense
	return ctx.JSON(fmt.Sprintf("{'petId': %d}", petId))
}

func DeletePet(ctx *fiber.Ctx) error {
	petId, err := strconv.Atoi(ctx.Params("petId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := storage.DeletePet(int64(petId)); err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

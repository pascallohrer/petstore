package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pascallohrer/petstore/pkg/entities"
)

type PetStorage interface {
	GetPetById(int64) (entities.Pet, error)
	AddPet(entities.Pet) int64
	DeletePet(int64) error
}

type addPetResponse struct {
	PetId int64 `json:"petId"`
}

func GetPetByIdHandler(storage PetStorage) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
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
}

func AddPetHandler(storage PetStorage) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var newPet entities.Pet
		if err := ctx.BodyParser(&newPet); err != nil {
			return ctx.Status(fiber.StatusMethodNotAllowed).SendString(err.Error())
		}
		if !newPet.IsValid() {
			return ctx.Status(fiber.StatusMethodNotAllowed).SendString("Required values missing")
		}
		petId := storage.AddPet(newPet)
		// Even though the spec doesn't specify it, returning the newly added ID just makes sense
		return ctx.JSON(addPetResponse{petId})
	}
}

func DeletePetHandler(storage PetStorage) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		petId, err := strconv.Atoi(ctx.Params("petId"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		if err := storage.DeletePet(int64(petId)); err != nil {
			return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return ctx.SendStatus(fiber.StatusNoContent)
	}
}

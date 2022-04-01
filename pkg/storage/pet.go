package storage

import (
	"fmt"

	"github.com/pascallohrer/petstore/pkg/entities"
)

var pets map[int64]entities.Pet
var nextId int64 = 0

func GetPetById(petId int64) (entities.Pet, error) {
	pet, exists := pets[petId]
	if !exists {
		return entities.Pet{}, fmt.Errorf("petId not found")
	}
	return pet, nil
}

func DeletePet(petId int64) error {
	_, exists := pets[petId]
	if !exists {
		return fmt.Errorf("petId not found")
	}
	delete(pets, petId)
	return nil
}

func AddPet(newPet entities.Pet) int64 {
	pets[nextId] = newPet
	nextId++
	return nextId
}

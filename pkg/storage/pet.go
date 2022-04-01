package storage

import (
	"fmt"

	"github.com/pascallohrer/petstore/pkg/entities"
)

type MemoryPetStorage struct {
	pets   map[int64]entities.Pet
	nextId int64
}

func NewMemoryPetStorage() *MemoryPetStorage {
	return &MemoryPetStorage{
		pets:   make(map[int64]entities.Pet),
		nextId: 1,
	}
}

func (m *MemoryPetStorage) GetPetById(petId int64) (entities.Pet, error) {
	pet, exists := m.pets[petId]
	if !exists {
		return entities.Pet{}, fmt.Errorf("petId not found")
	}
	return pet, nil
}

func (m *MemoryPetStorage) DeletePet(petId int64) error {
	_, exists := m.pets[petId]
	if !exists {
		return fmt.Errorf("petId not found")
	}
	delete(m.pets, petId)
	return nil
}

func (m *MemoryPetStorage) AddPet(newPet entities.Pet) int64 {
	newPet.Id = m.nextId
	m.pets[m.nextId] = newPet
	m.nextId++
	return m.nextId - 1
}

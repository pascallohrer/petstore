package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pascallohrer/petstore/pkg/entities"
	"github.com/pascallohrer/petstore/pkg/storage"
	"gotest.tools/assert"
)

func TestAddPet(t *testing.T) {
	app := initializeRouter()
	t.Run("empty body not allowed", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/pet", nil)
		res, err := app.Test(req)
		assert.NilError(t, err, "request failed: %s", err)
		assert.Equal(t, res.StatusCode, fiber.StatusMethodNotAllowed)
	})
	t.Run("name required", func(t *testing.T) {
		newPet := `{"photoUrls": ["example.com"]}`
		req := httptest.NewRequest("POST", "/pet", strings.NewReader(newPet))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)
		assert.NilError(t, err, "request failed: %s", err)
		assert.Equal(t, res.StatusCode, fiber.StatusMethodNotAllowed)
	})
	t.Run("name must not be empty", func(t *testing.T) {
		newPet := `{"name": "", "photoUrls": ["example.com"]}`
		req := httptest.NewRequest("POST", "/pet", strings.NewReader(newPet))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)
		assert.NilError(t, err, "request failed: %s", err)
		assert.Equal(t, res.StatusCode, fiber.StatusMethodNotAllowed)
	})
	t.Run("photos required", func(t *testing.T) {
		newPet := `{"name": "kitty"}`
		req := httptest.NewRequest("POST", "/pet", strings.NewReader(newPet))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)
		assert.NilError(t, err, "request failed: %s", err)
		assert.Equal(t, res.StatusCode, fiber.StatusMethodNotAllowed)
	})
	t.Run("photos must be non-empty", func(t *testing.T) {
		newPet := `{"name": "kitty", "photoUrls": []}`
		req := httptest.NewRequest("POST", "/pet", strings.NewReader(newPet))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)
		assert.NilError(t, err, "request failed: %s", err)
		assert.Equal(t, res.StatusCode, fiber.StatusMethodNotAllowed)
	})
	t.Run("name and photos included", func(t *testing.T) {
		newPet := `{"name": "kitty", "photoUrls": ["example.com"]}`
		req := httptest.NewRequest("POST", "/pet", strings.NewReader(newPet))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)
		t.Run("request successful", func(t *testing.T) {
			assert.NilError(t, err, "request failed: %s", err)
			assert.Equal(t, res.StatusCode, fiber.StatusOK)
		})
		t.Run("correct ID", func(t *testing.T) {
			var result struct {
				PetId int64 `json:"petId"`
			}
			err = json.NewDecoder(res.Body).Decode(&result)
			assert.Equal(t, result.PetId, int64(6)) // initial data is 5 pets -> next ID = 6
		})
	})
}

func TestGetPet(t *testing.T) {
	app := initializeRouter()
	t.Run("invalid ID not found", func(t *testing.T) {
		const iterations = 20
		rand.Seed(time.Now().UnixNano())
		for range [iterations]bool{} {
			id := rand.Int63() + 6 // definitely too large
			t.Run(fmt.Sprintf("ID %d", id), func(t *testing.T) {
				req := httptest.NewRequest("GET", fmt.Sprintf("/pet/%d", id), nil)
				res, err := app.Test(req)
				assert.NilError(t, err)
				assert.Equal(t, res.StatusCode, fiber.StatusNotFound)
			})
			t.Run(fmt.Sprintf("ID %d", -id), func(t *testing.T) {
				req := httptest.NewRequest("GET", fmt.Sprintf("/pet/%d", -id), nil)
				res, err := app.Test(req)
				assert.NilError(t, err)
				assert.Equal(t, res.StatusCode, fiber.StatusNotFound)
			})
		}
	})
	t.Run("initial pets retrievable", func(t *testing.T) {
		for index, pet := range initialPets {
			req := httptest.NewRequest("GET", fmt.Sprintf("/pet/%d", index+1), nil)
			res, err := app.Test(req)
			assert.NilError(t, err)
			assert.Equal(t, res.StatusCode, fiber.StatusOK)
			var result entities.Pet
			json.NewDecoder(res.Body).Decode(&result)
			pet.Id = int64(index + 1)
			assert.DeepEqual(t, result, pet)
		}
	})
	t.Run("added pet can be fully retrieved", func(t *testing.T) {
		newPet := entities.Pet{
			Category: entities.Category{
				Id:   21,
				Name: "Dog",
			},
			Name: "Bello",
			PhotoUrls: []string{
				"example.com/dog.jpg",
			},
			Tags: []entities.Tag{
				{
					Id:   3,
					Name: "tame",
				},
				{
					Id:   5,
					Name: "carnivore",
				},
			},
			Status: "sold",
		}
		encodedPet, err := json.Marshal(newPet)
		if err != nil {
			t.Fatalf("Please fix your test: %s", err)
		}
		req := httptest.NewRequest("POST", "/pet", bytes.NewReader(encodedPet))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)
		assert.NilError(t, err)
		var newId struct {
			PetId int64 `json:"petId"`
		}
		json.NewDecoder(res.Body).Decode(&newId)
		newPet.Id = newId.PetId
		req = httptest.NewRequest("GET", fmt.Sprintf("/pet/%d", newId.PetId), nil)
		res, err = app.Test(req)
		assert.NilError(t, err)
		assert.Equal(t, res.StatusCode, fiber.StatusOK)
		var result entities.Pet
		json.NewDecoder(res.Body).Decode(&result)
		assert.DeepEqual(t, result, newPet)
	})
}

func TestDeletePet(t *testing.T) {
	const iterations = 10
	rand.Seed(time.Now().UnixNano())
	for range [iterations]bool{} {
		id := rand.Int63n(5) + 1 // should always be valid
		t.Run(fmt.Sprintf("ID %d", id), func(t *testing.T) {
			app := initializeRouter()
			req := httptest.NewRequest("GET", fmt.Sprintf("/pet/%d", id), nil)
			res, err := app.Test(req)
			assert.NilError(t, err)
			if res.StatusCode == fiber.StatusNotFound {
				t.Skipf("Unexpected invalid ID %d", id)
			}
			t.Run("deletion reported successful", func(t *testing.T) {
				req = httptest.NewRequest("DELETE", fmt.Sprintf("/pet/%d", id), nil)
				res, err = app.Test(req)
				assert.NilError(t, err)
				assert.Equal(t, res.StatusCode, fiber.StatusNoContent)
			})
			t.Run("pet actually gone", func(t *testing.T) {
				req = httptest.NewRequest("GET", fmt.Sprintf("/pet/%d", id), nil)
				res, err = app.Test(req)
				assert.NilError(t, err)
				assert.Equal(t, res.StatusCode, fiber.StatusNotFound)
			})
		})
	}
}

func initializeRouter() *fiber.App {
	repository := storage.NewMemoryPetStorage()
	for _, pet := range initialPets {
		repository.AddPet(pet)
	}
	return NewRouter(repository)
}

var initialPets = []entities.Pet{
	{
		Name: "kitty",
		PhotoUrls: []string{
			"example.com/pic1.jpg",
			"example.com/pic2.jpg",
		},
	},
	{
		Name: "fluffy",
		PhotoUrls: []string{
			"example.com/pic1.jpg",
			"example.com/pic2.jpg",
		},
	},
	{
		Name: "bruno",
		PhotoUrls: []string{
			"example.com/pic1.jpg",
			"example.com/pic2.jpg",
		},
	},
	{
		Name: "doggo",
		PhotoUrls: []string{
			"example.com/pic1.jpg",
			"example.com/pic2.jpg",
		},
	},
	{
		Name: "babe",
		PhotoUrls: []string{
			"example.com/pic1.jpg",
			"example.com/pic2.jpg",
		},
	},
}

package main

import (
	"github.com/gofiber/fiber/v2"

	_ "github.com/go-sql-driver/mysql"
)

var s *PetShopService

// InjectService do inject the service object for handler
func InjectService(injectedService *PetShopService) {
	s = injectedService
}

// GetPetHandler handle get a pet by id operation
func GetPetHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	pet, err := s.GetPetByID(id)
	if err != nil {
		return err
	}

	return c.JSON(pet)
}

// GetPetListHandler handle get a list of pets
func GetPetListHandler(c *fiber.Ctx) error {
	pets, err := s.GetPetList()
	if err != nil {
		return err
	}

	return c.JSON(pets)
}

// AddPetHandler handle add a pet to the store
func AddPetHandler(c *fiber.Ctx) error {
	type AddPetDTO struct {
		Name string `json:"name"`
		Age  int32  `json:"age"`
	}

	dto := new(AddPetDTO)
	err := c.BodyParser(dto)
	if err != nil {
		return err
	}

	pet, err := s.AddPet(dto.Name, dto.Age)
	if err != nil {
		return err
	}

	return c.JSON(pet)
}

// ChangePetNameHandler handle change a pet's name operation
func ChangePetNameHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	type ChangePetNameDTO struct {
		Name string `json:"name"`
	}
	dto := new(ChangePetNameDTO)
	err := c.BodyParser(dto)
	if err != nil {
		return err
	}

	err = s.ChangePetName(id, dto.Name)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

// SellPetHandler handle selling pet
func SellPetHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	err := s.SellPet(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

// ReturnPetHandler handle returning pet
func ReturnPetHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	err := s.ReturnPet(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

package main

import (
	"time"

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

	return c.JSON(toPetDTO(pet))
}

// GetPetListHandler handle get a list of pets
func GetPetListHandler(c *fiber.Ctx) error {
	pets, err := s.GetPetList()
	if err != nil {
		return err
	}

	return c.JSON(toPetDTOs(pets))
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

	return c.JSON(toPetDTO(pet))
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

// PetDTO a data transfer struct
type PetDTO struct {
	ID      string     `json:"id"`
	Name    string     `json:"name"`
	Age     int32      `json:"age"`
	AddedAt time.Time  `json:"added_at"`
	IsSold  bool       `json:"is_sold"`
	SoldAt  *time.Time `json:"sold_at"`
}

func toPetDTO(p *Pet) *PetDTO {
	return &PetDTO{
		ID:      p.id.String(),
		Name:    p.name,
		Age:     p.age,
		AddedAt: p.addedAt,
		IsSold:  p.isSold,
		SoldAt:  p.soldAt,
	}
}

func toPetDTOs(pets []*Pet) []*PetDTO {
	petDTOs := make([]*PetDTO, 0, 5)
	for i := 0; i < len(pets); i++ {
		petDTOs = append(petDTOs, toPetDTO(pets[i]))
	}
	return petDTOs
}

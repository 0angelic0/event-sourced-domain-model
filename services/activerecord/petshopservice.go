package main

import (
	"time"

	"gorm.io/gorm"
)

// PetShopService is a stateless service object implementing transaction scripts
type PetShopService struct {
	db *gorm.DB
}

// NewPetShopService create new instance of a PetShopService
func NewPetShopService(db *gorm.DB) *PetShopService {
	return &PetShopService{
		db: db,
	}
}

// GetPetByID return a pet by id
func (s *PetShopService) GetPetByID(id string) (*PetORM, error) {
	var pet PetORM
	result := s.db.First(&pet, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pet, nil
}

// GetPetList return a list of pets
func (s *PetShopService) GetPetList() ([]PetORM, error) {
	var pets []PetORM
	result := s.db.Find(&pets)
	if result.Error != nil {
		return nil, result.Error
	}

	return pets, nil
}

// AddPet add a new pet to the shop
func (s *PetShopService) AddPet(name string, age int32) (*PetORM, error) {
	pet := NewPetORM(name, age)

	result := s.db.Create(&pet)
	if result.Error != nil {
		return nil, result.Error
	}

	return pet, nil
}

// ChangePetName change a pet's name
func (s *PetShopService) ChangePetName(id string, name string) error {
	pet, err := s.GetPetByID(id)
	if err != nil {
		return err
	}

	pet.Name = name

	result := s.db.Save(&pet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SellPet sell a pet
func (s *PetShopService) SellPet(id string) error {
	pet, err := s.GetPetByID(id)
	if err != nil {
		return err
	}

	pet.IsSold = true
	now := time.Now()
	pet.SoldAt = &now

	result := s.db.Save(&pet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ReturnPet return a pet to the shop
func (s *PetShopService) ReturnPet(id string) error {
	pet, err := s.GetPetByID(id)
	if err != nil {
		return err
	}

	pet.IsSold = false
	pet.SoldAt = nil

	result := s.db.Save(&pet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

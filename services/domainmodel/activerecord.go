package main

import (
	"time"

	"github.com/google/uuid"
)

// PetORM a orm struct mapping to a table in database, implementing active record pattern
type PetORM struct {
	ID      string
	Name    string
	Age     int32
	AddedAt time.Time
	IsSold  bool
	SoldAt  *time.Time
	Version uint32
}

// TableName overrides the table name used by PetORM
func (PetORM) TableName() string {
	return "pets"
}

// NewPetORM a PetORM's constructor
func NewPetORM(name string, age int32) *PetORM {
	return &PetORM{
		ID:      uuid.New().String(),
		Name:    name,
		Age:     age,
		AddedAt: time.Now(),
		IsSold:  false,
		SoldAt:  nil,
		Version: 1,
	}
}

func toPet(petORM *PetORM) *Pet {
	return &Pet{
		id:      uuid.MustParse(petORM.ID),
		name:    petORM.Name,
		age:     petORM.Age,
		addedAt: petORM.AddedAt,
		isSold:  petORM.IsSold,
		soldAt:  petORM.SoldAt,
		version: petORM.Version,
	}
}

func toPets(petORMs []PetORM) []*Pet {
	pets := make([]*Pet, 0, 5)
	for i := 0; i < len(petORMs); i++ {
		pets = append(pets, toPet(&petORMs[i]))
	}
	return pets
}

func toPetORM(pet *Pet) *PetORM {
	return &PetORM{
		ID:      pet.id.String(),
		Name:    pet.name,
		Age:     pet.age,
		AddedAt: pet.addedAt,
		IsSold:  pet.isSold,
		SoldAt:  pet.soldAt,
		Version: pet.version,
	}
}

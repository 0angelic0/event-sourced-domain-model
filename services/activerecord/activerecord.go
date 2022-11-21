package main

import (
	"time"

	"github.com/google/uuid"
)

// PetORM a orm struct mapping to a table in database, implementing active record pattern
type PetORM struct {
	ID      string     `json:"id"`
	Name    string     `json:"name"`
	Age     int32      `json:"age"`
	AddedAt time.Time  `json:"added_at"`
	IsSold  bool       `json:"is_sold"`
	SoldAt  *time.Time `json:"sold_at"`
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
	}
}

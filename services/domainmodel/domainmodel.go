package main

import (
	"time"

	"github.com/google/uuid"
)

// Pet represents a pet
type Pet struct {
	id      uuid.UUID
	name    string
	age     int32
	addedAt time.Time
	isSold  bool
	soldAt  *time.Time
	version uint32
}

// NewPet is a Pet's constructor
func NewPet(name string, age int32) *Pet {
	return &Pet{
		id:      uuid.New(),
		name:    name,
		age:     age,
		addedAt: time.Now(),
		isSold:  false,
		soldAt:  nil,
		version: 0,
	}
}

// ChangeName change a pet's name
func (p *Pet) ChangeName(name string) {
	p.name = name
}

// Sell pet from the pet shop
func (p *Pet) Sell() {
	p.isSold = true
	now := time.Now()
	p.soldAt = &now
}

// Return pet to the pet shop
func (p *Pet) Return() {
	p.isSold = false
	p.soldAt = nil
}

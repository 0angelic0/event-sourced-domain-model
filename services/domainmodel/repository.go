package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PetRepository interface to retrieve pets from data store
type PetRepository interface {
	FindByID(id string) (*Pet, error)
	FindAll() ([]*Pet, error)
	Save(p *Pet) error
}

// PetRepositoryMySQL implement mysql data store
type PetRepositoryMySQL struct {
	db *gorm.DB
}

// NewPetRepositoryMySQL create new instance of PetRepositoryMySQL
func NewPetRepositoryMySQL() *PetRepositoryMySQL {
	dsn := "root:rootpassword@/domainmodel?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to DB")

	return &PetRepositoryMySQL{
		db: db,
	}
}

// FindByID find a pet by id
func (r *PetRepositoryMySQL) FindByID(id string) (*Pet, error) {
	var petORM PetORM
	result := r.db.First(&petORM, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return toPet(&petORM), nil
}

// FindAll find all pets
func (r *PetRepositoryMySQL) FindAll() ([]*Pet, error) {
	var petORMs []PetORM
	result := r.db.Find(&petORMs)
	if result.Error != nil {
		return nil, result.Error
	}

	return toPets(petORMs), nil
}

// Save do save a pet
func (r *PetRepositoryMySQL) Save(p *Pet) error {
	petORM := toPetORM(p)
	currentVersion := petORM.Version
	petORM.Version = currentVersion + 1 // increment version for saving
	result := r.db.Where("version = ?", currentVersion).Save(petORM)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return fmt.Errorf("concurrency check: invalid record version, cannot save")
	}
	return nil
}

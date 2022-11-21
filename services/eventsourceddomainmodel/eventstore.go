package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PetEventStore interface
type PetEventStore interface {
	FetchByID(id string) ([]PetEvent, error)
	FetchAll() (map[string][]PetEvent, error)
	Append(id string, newEvents []PetEvent, expectedVersion uint32) error
}

// PetEventStoreMySQL is an event store implementation using MySQL
type PetEventStoreMySQL struct {
	db *gorm.DB
}

// NewPetEventStoreMySQL is a constructor
func NewPetEventStoreMySQL() *PetEventStoreMySQL {
	dsn := "root:rootpassword@/eventsourceddomainmodel?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to DB")

	return &PetEventStoreMySQL{
		db: db,
	}
}

// FetchByID fetches events by aggregate id
func (es *PetEventStoreMySQL) FetchByID(id string) ([]PetEvent, error) {
	var petORMs []PetORM
	result := es.db.Order("event_id").Find(&petORMs, "aggr_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return toPetEvents(petORMs), nil
}

// FetchAll fetches all events by all aggregate id (for demo purpose only, do not do this in production)
func (es *PetEventStoreMySQL) FetchAll() (map[string][]PetEvent, error) {
	var petORMs []PetORM
	result := es.db.Order("aggr_id, event_id").Find(&petORMs)
	if result.Error != nil {
		return nil, result.Error
	}

	return toMapPetEvents(petORMs), nil
}

// Append is appending the events by aggregate id and also do concurrency check
func (es *PetEventStoreMySQL) Append(id string, newEvents []PetEvent, expectedVersion uint32) error {
	// Concurrency Check
	var orm PetORM
	result := es.db.Limit(1).Order("event_id desc").Where("aggr_id = ?", id).Find(&orm)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 0 {
		// existing aggregate
		if orm.EventID != expectedVersion {
			return fmt.Errorf("concurrency check: invalid aggregate version, cannot append")
		}
	}

	// Append New Records
	petORMs := make([]PetORM, 0, 5)
	for i := 0; i < len(newEvents); i++ {
		e := newEvents[i]
		petORM := toPetORM(e)
		petORMs = append(petORMs, *petORM)
	}

	result = es.db.Create(&petORMs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// PetShopService is a stateless service object implementing transaction scripts
type PetShopService struct {
	db *sql.DB
}

// NewPetShopService create new instance of a PetShopService
func NewPetShopService(db *sql.DB) *PetShopService {
	return &PetShopService{
		db: db,
	}
}

// GetPetByID return a pet by id
func (s *PetShopService) GetPetByID(id string) (map[string]interface{}, error) {
	rows, err := s.db.Query("SELECT name, age, added_at, is_sold, sold_at FROM pets WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	if rows.Next() == false {
		return nil, nil
	}

	var name string
	var age int32
	var addedAt time.Time
	var isSold bool
	var soldAt sql.NullTime
	err = rows.Scan(&name, &age, &addedAt, &isSold, &soldAt)
	if err != nil {
		return nil, err
	}

	var sSoldAt *string
	if isSold {
		temp := soldAt.Time.Format(time.RFC3339Nano)
		sSoldAt = &temp
	}

	return map[string]interface{}{
		"id":       id,
		"name":     name,
		"age":      age,
		"added_at": addedAt,
		"is_sold":  isSold,
		"sold_at":  sSoldAt,
	}, nil
}

// GetPetList return a list of pets
func (s *PetShopService) GetPetList() ([]map[string]interface{}, error) {
	rows, err := s.db.Query("SELECT id, name, age, added_at, is_sold, sold_at FROM pets")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0, 5)

	for rows.Next() {

		var id string
		var name string
		var age int32
		var addedAt time.Time
		var isSold bool
		var soldAt sql.NullTime
		err = rows.Scan(&id, &name, &age, &addedAt, &isSold, &soldAt)
		if err != nil {
			return nil, err
		}

		var sSoldAt *string
		if isSold {
			temp := soldAt.Time.Format(time.RFC3339Nano)
			sSoldAt = &temp
		}

		result = append(result, map[string]interface{}{
			"id":       id,
			"name":     name,
			"age":      age,
			"added_at": addedAt,
			"is_sold":  isSold,
			"sold_at":  sSoldAt,
		})
	}

	return result, nil
}

// AddPet add a new pet to the shop
func (s *PetShopService) AddPet(name string, age int32) (map[string]interface{}, error) {
	id := uuid.New()
	addedAt := time.Now()

	result, err := s.db.Exec("INSERT INTO pets(id, name, age, added_at) VALUES(?, ?, ?, ?)", id, name, age, addedAt)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected != 1 {
		return nil, fmt.Errorf("rowsAffected is not 1")
	}

	return map[string]interface{}{
		"id":       id,
		"name":     name,
		"age":      age,
		"added_at": addedAt,
		"is_sold":  false,
		"sold_at":  nil,
	}, nil
}

// ChangePetName change a pet's name
func (s *PetShopService) ChangePetName(id string, name string) error {
	result, err := s.db.Exec("UPDATE pets SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected is not 1")
	}
	return nil
}

// SellPet sell a pet
func (s *PetShopService) SellPet(id string) error {
	soldAt := time.Now()
	result, err := s.db.Exec("UPDATE pets SET is_sold = true, sold_at = ? WHERE id = ?", soldAt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected is not 1")
	}
	return nil
}

// ReturnPet return a pet to the shop
func (s *PetShopService) ReturnPet(id string) error {
	result, err := s.db.Exec("UPDATE pets SET is_sold = false, sold_at = NULL WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected is not 1")
	}
	return nil
}

package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// ConnectDB is for connecting to a db
func ConnectDB() error {
	var err error
	db, err = sql.Open("mysql", "root:rootpassword@/transactionscript?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	log.Println("Connected to DB")
	return nil
}

// GetPetHandler handle get a pet by id operation
func GetPetHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	rows, err := db.Query("SELECT name, age, added_at, is_sold, sold_at FROM pets WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		return err
	}

	if rows.Next() == false {
		return fiber.ErrNotFound
	}

	var name string
	var age int32
	var addedAt time.Time
	var isSold bool
	var soldAt sql.NullTime
	err = rows.Scan(&name, &age, &addedAt, &isSold, &soldAt)
	if err != nil {
		return err
	}

	var sSoldAt *string
	if isSold {
		temp := soldAt.Time.Format(time.RFC3339Nano)
		sSoldAt = &temp
	}

	return c.JSON(fiber.Map{
		"id":       id,
		"name":     name,
		"age":      age,
		"added_at": addedAt,
		"is_sold":  isSold,
		"sold_at":  sSoldAt,
	})
}

// GetPetListHandler handle get a list of pets
func GetPetListHandler(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, name, age, added_at, is_sold, sold_at FROM pets")
	defer rows.Close()
	if err != nil {
		return err
	}

	response := make([]fiber.Map, 0, 5)

	for rows.Next() {

		var id string
		var name string
		var age int32
		var addedAt time.Time
		var isSold bool
		var soldAt sql.NullTime
		err = rows.Scan(&id, &name, &age, &addedAt, &isSold, &soldAt)
		if err != nil {
			return err
		}

		var sSoldAt *string
		if isSold {
			temp := soldAt.Time.Format(time.RFC3339Nano)
			sSoldAt = &temp
		}

		response = append(response, fiber.Map{
			"id":       id,
			"name":     name,
			"age":      age,
			"added_at": addedAt,
			"is_sold":  isSold,
			"sold_at":  sSoldAt,
		})
	}

	return c.JSON(response)
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

	id := uuid.New()
	addedAt := time.Now()

	result, err := db.Exec("INSERT INTO pets(id, name, age, added_at) VALUES(?, ?, ?, ?)", id, dto.Name, dto.Age, addedAt)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fiber.NewError(fiber.StatusInternalServerError, "rowsAffected is not 1")
	}

	return c.JSON(fiber.Map{
		"id":       id,
		"name":     dto.Name,
		"age":      dto.Age,
		"added_at": addedAt,
		"is_sold":  false,
		"sold_at":  nil,
	})
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

	result, err := db.Exec("UPDATE pets SET name = ? WHERE id = ?", dto.Name, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fiber.NewError(fiber.StatusInternalServerError, "rowsAffected is not 1")
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

// SellPetHandler handle selling pet
func SellPetHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	soldAt := time.Now()
	result, err := db.Exec("UPDATE pets SET is_sold = true, sold_at = ? WHERE id = ?", soldAt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fiber.NewError(fiber.StatusInternalServerError, "rowsAffected is not 1")
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

// ReturnPetHandler handle returning pet
func ReturnPetHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := db.Exec("UPDATE pets SET is_sold = false, sold_at = NULL WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fiber.NewError(fiber.StatusInternalServerError, "rowsAffected is not 1")
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

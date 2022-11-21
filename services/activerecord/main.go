// Package main is a main program
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	log.Println("App started")
	db := ConnectDB()
	service := NewPetShopService(db)
	InjectService(service)

	app := fiber.New()

	app.Get("/pets/:id", GetPetHandler)
	app.Get("/pets", GetPetListHandler)
	app.Post("/pets", AddPetHandler)
	app.Post("/pets/:id/change-name", ChangePetNameHandler)
	app.Post("/pets/:id/sell", SellPetHandler)
	app.Post("/pets/:id/return", ReturnPetHandler)

	log.Fatal(app.Listen(":3000"))
}

// ConnectDB is for connecting to a db
func ConnectDB() *gorm.DB {
	dsn := "root:rootpassword@/transactionscript?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to DB")
	return db
}

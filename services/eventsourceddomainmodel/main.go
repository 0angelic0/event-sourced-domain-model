// Package main is a main program
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("App started")
	repo := NewPetRepositoryEventStore()
	service := NewPetShopService(repo)
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

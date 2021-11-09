package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/vapopov/ethidx"
	_ "github.com/vapopov/ethidx/docs"
)

func main() {
	api, err := ethidx.NewApi()
	if err!= nil {
		log.Fatalf("can't init application api: %v", err)
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.Handler)

	app.Get("/groups", api.Groups)
	app.Get("/groups/:id", api.GroupsById)
	app.Get("/indexes/:id", api.Indexes)
	app.Get("/blocks/:id", api.BlockInfo)

	log.Fatal(app.Listen(":3000"))
}

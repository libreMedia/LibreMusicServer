package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ok = readDb()

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/dirList", func(c *fiber.Ctx) error {
		// fmt.Print(readDb()[1111].album)
		fmt.Print("%T", readDb())
		return c.JSON(ok)
	})

	app.Get("/artist/:artist", func(c *fiber.Ctx) error {
		artist := c.Params("artist")
		fmt.Println(artist)
		return c.JSON(artist)
	})
	//TODO make go to DB
	app.Static("/music", "../music", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	// fmt.Print("did a thing")
	app.Listen(":9000")
}

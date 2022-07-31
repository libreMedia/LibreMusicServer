package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/dirList", func(c *fiber.Ctx) error {
		// fmt.Print(readDb()[1111].album)
		fmt.Print("%T", readDb())
		ok := readDb()
		return c.JSON(ok)
	})

	//TODO make go to DB
	app.Static("/music", "../music", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	fmt.Print("did a thing")
	app.Listen(":9000")
}

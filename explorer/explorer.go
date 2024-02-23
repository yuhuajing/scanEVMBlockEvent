package explorer

import (
	fiber "github.com/gofiber/fiber/v2"
	html "github.com/gofiber/template/html/v2"
	"log"
	"main/core/database"
)

func Explorer() {
	app := fiber.New(fiber.Config{
		Views: html.New("C:\\Users\\user\\Downloads\\scanEVMBlockEvent\\explorer\\views", ".html"),
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("form", nil)
	})
	app.Get("/data", dealquery)
	log.Fatal(app.Listen(":3000"))
}

func dealquery(c *fiber.Ctx) error {
	resdata := database.QueryStartBlock()
	return c.Render("index", fiber.Map{
		"Data": resdata,
	})
}

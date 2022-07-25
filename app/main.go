package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/ekkinox/manticore-demo/app/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
)

//go:embed views/*
var vfs embed.FS

func main() {

	// viper
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	// fiber
	views := html.NewFileSystem(http.FS(vfs), ".html")
	views.Delims("[[", "]]")

	fiberApp := fiber.New(fiber.Config{
		Views: views,
	})

	fiberApp.Use(logger.New())

	// manticore client
	manticoreClient, err := client.InitManticoreClient()
	if err != nil {
		log.Printf("error: %v", err)
	}

	// main endpoint
	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index", fiber.Map{})
	})

	// search endpoint
	fiberApp.Get("/search", func(c *fiber.Ctx) error {
		var results *client.ManticoreClientSearchResult

		expression := c.Query("expression", "")
		if expression != "" {
			results, err = manticoreClient.FindByExpression("movies", expression)
		} else {
			results, err = manticoreClient.FindAll("movies")
		}

		return c.JSON(results)
	})

	log.Fatal(fiberApp.Listen(":8080"))
}

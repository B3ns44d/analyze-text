package main

import (
	"log"
	"strings"

	"github.com/B3ns44d/analyze-text/config"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Analyzer struct {
	Description string `json:"description"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ping Pong!")
	})

	app.Get("/analyze", func(c *fiber.Ctx) error {
		analyzer := new(Analyzer)
		if err := c.BodyParser(analyzer); err != nil {
			return err
		}

		if len(analyzer.Description) < 1 {
			return c.Status(fasthttp.StatusBadRequest).SendString("Description is empty")
		}
		analyzer.Description = strings.Replace(analyzer.Description, " ", "%20", -1)
		apiUrl := config.Config("API_URL")
		apiKey := config.Config("API_KEY")

		client := fasthttp.Client{}
		req := fasthttp.AcquireRequest()
		req.SetRequestURI(apiUrl + "?replace=******&key=" + apiKey + "&msg=" + analyzer.Description)
		req.Header.SetMethod("GET")
		req.Header.Set("Content-Type", "application/json")
		resp := fasthttp.AcquireResponse()
		if err := client.Do(req, resp); err != nil {
			return err
		}
		return c.SendString(string(resp.Body()))
	})

	log.Fatal(app.Listen(":3008"))
}

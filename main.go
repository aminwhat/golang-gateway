package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const (
	NESTJS_V0 = "http://localhost:3700"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStream(bytes.NewBufferString("Hello, World!"))
	})

	app.All("/v0/*", func(c *fiber.Ctx) error {
		var method string = c.Method()

		var req *http.Request
		var resp *http.Response
		var err error

		req, err = http.NewRequest(method, NESTJS_V0+c.Path(), bytes.NewBuffer(c.Body()))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		req.Header.Set("Content-Type", c.Get("Content-Type"))
		req.Header.Set("Authorization", c.Get("Authorization"))
		// req.Header.Set("X-API-KEY", c.Get("X-API-KEY"))
		// req.Header.Set("X-API-SECRET", c.Get("X-API-SECRET"))
		// req.Header.Set("X-API-TOKEN", c.Get("X-API-TOKEN"))
		// req.Header.Set("X-API-TOKEN-SECRET", c.Get("X-API-TOKEN-SECRET"))
		// req.Header.Set("X-API-TOKEN-PUBLIC", c.Get("X-API-TOKEN-PUBLIC"))
		// req.Header.Set("X-API-TOKEN-PRIVATE", c.Get("X-API-TOKEN-PRIVATE"))
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Status(resp.StatusCode).Send(body)
	})

	log.Fatal(app.Listen(":3000"))
}

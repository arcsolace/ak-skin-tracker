package main

import (
    "log"
    "os"

    f "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"
    "github.com/arcsolace/ak-skin-tracker/config"
    r "github.com/arcsolace/ak-skin-tracker/routes"
)

func setupRoutes(app *f.App) {
    app.Get("/", func(c *f.Ctx) error {
        return c.Status(f.StatusOK).JSON(f.Map{
            "success":     true,
            "message":     "Hello, World!",
            "github_repo": "<https://github.com/arcsolace/ak-skin-tracker>",
        })
    })

    api := app.Group("/api")

    r.SkinsRoute(api.Group("/skin"))
	r.UserRoute(api.Group("/user"))
	r.ShareRoute(api.Group("/share"))
}

func main() {
    if os.Getenv("APP_ENV") != "production" {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
    }

    app := f.New()

    app.Use(cors.New())
    app.Use(logger.New())

    config.ConnectDB()

    setupRoutes(app)

    port := os.Getenv("PORT")
    err := app.Listen(":" + port)

    if err != nil {
        log.Fatal("Error app failed to start")
        panic(err)
    }
}
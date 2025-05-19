package main

import (
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "backend/router"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using default port 3000")
    }
    port := os.Getenv("PORT")
    if port == "" {
        port = "3333"
    }

    app := fiber.New()
    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000, http://127.0.0.1:3000, http://localhost:5173, http://127.0.0.1:5173",
        AllowCredentials: true,
    }))
    router.SetupRoutes(app)
    log.Fatal(app.Listen(":" + port))
}
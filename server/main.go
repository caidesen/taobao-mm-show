package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"tbmm/pkg/configs"
	"tbmm/pkg/middleware"
	"tbmm/pkg/routes"
	"tbmm/pkg/utils"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)
	app.Static("/", "./public")
	group := app.Group("api")
	routes.TbMMRoute(&group)
	middleware.FiberMiddleware(app)
	utils.StartServerWithGracefulShutdown(app)
}

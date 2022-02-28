package routes

import (
	"github.com/gofiber/fiber/v2"
	"tbmm/app/controllers"
)

func TbMMRoute(a *fiber.Router) {
	router := *a
	router.Get("/list", controllers.GetMMPicList)
}

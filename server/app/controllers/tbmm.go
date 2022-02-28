package controllers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"tbmm/app/service"
)

func GetMMPicList(c *fiber.Ctx) error {
	listLenQuery := c.Query("len", "1")
	listLen, err := strconv.Atoi(listLenQuery)
	if err != nil {
		return err
	}
	list := service.GetTbMMPicList(listLen)
	return c.JSON(fiber.Map{
		"code": 0,
		"data": list,
	})
}

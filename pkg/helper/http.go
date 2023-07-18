package helper

import "github.com/gofiber/fiber/v2"

// 200-299
func Ok(msg string, data any) error {
	return Ctx.JSON(fiber.Map{
		"data":    data,
		"msg":     msg,
		"success": true,
	})
}

// 400-499
func BadRequest(msg string, err error) error {
	return Ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":   err.Error(),
		"msg":     msg,
		"success": false,
	})
}

// 500-599
func InternalServerError(msg string, err error) error {
	return Ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   err.Error(),
		"msg":     msg,
		"success": false,
	})
}

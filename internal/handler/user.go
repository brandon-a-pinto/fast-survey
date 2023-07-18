package handler

import (
	"github.com/brandon-a-pinto/fast-survey/internal/schemas"
	"github.com/brandon-a-pinto/fast-survey/pkg/mongodb"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type user struct{}

func newUserHandler() *user {
	return &user{}
}

// Create User
func (h *user) HandlePostUser(c *fiber.Ctx) error {
	usersCollection := mongodb.MI.Collection("users")
	params := new(schemas.CreateUserParams)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"msg":     "BodyParser error",
			"success": false,
		})
	}

	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"msg":     "Bcrypt error",
			"success": false,
		})
	}

	user := schemas.User{
		Email:    params.Email,
		Name:     params.Name,
		Password: string(encpw),
	}

	res, err := usersCollection.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"msg":     "Database error",
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"data":    res,
		"msg":     "User created successfully",
		"success": true,
	})
}

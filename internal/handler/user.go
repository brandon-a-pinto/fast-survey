package handler

import (
	"github.com/brandon-a-pinto/fast-survey/internal/schemas"
	"github.com/brandon-a-pinto/fast-survey/pkg/helper"
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
	helper.Ctx = c
	usersCollection := mongodb.MI.Collection("users")
	params := new(schemas.CreateUserParams)

	if err := c.BodyParser(params); err != nil {
		return helper.BadRequest("BodyParser error", err)
	}
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return helper.InternalServerError("Bcrypt error", err)
	}
	user := schemas.User{
		Email:    params.Email,
		Name:     params.Name,
		Password: string(encpw),
	}
	res, err := usersCollection.InsertOne(c.Context(), user)
	if err != nil {
		return helper.InternalServerError("Database error", err)
	}

	return helper.Ok("User created successfully", res)
}

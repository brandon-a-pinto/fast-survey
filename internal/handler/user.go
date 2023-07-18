package handler

import (
	"fmt"

	"github.com/brandon-a-pinto/fast-survey/internal/schemas"
	"github.com/brandon-a-pinto/fast-survey/internal/validation"
	"github.com/brandon-a-pinto/fast-survey/pkg/helper"
	"github.com/brandon-a-pinto/fast-survey/pkg/mongodb"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
	if err := validation.CreateUserValidation(params); err != nil {
		return helper.BadRequest("Validation error", err)
	}
	emailExists, err := usersCollection.FindOne(c.Context(), bson.M{"email": params.Email}).DecodeBytes()
	if err != nil && err.Error() != "mongo: no documents in result" {
		return helper.InternalServerError("Database error", err)
	}
	if emailExists != nil {
		return helper.BadRequest("Email error", fmt.Errorf("email already in use"))
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

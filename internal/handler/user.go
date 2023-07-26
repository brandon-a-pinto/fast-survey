package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/brandon-a-pinto/fast-survey/internal/schemas"
	"github.com/brandon-a-pinto/fast-survey/internal/validation"
	"github.com/brandon-a-pinto/fast-survey/pkg/helper"
	"github.com/brandon-a-pinto/fast-survey/pkg/mongodb"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	UserHandler = newUserHandler()
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
	if err != nil && err != mongo.ErrNoDocuments {
		return helper.InternalServerError("Database error", err)
	}
	if emailExists != nil {
		return helper.Forbidden("Email error", fmt.Errorf("email already in use"))
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

// Authenticate User
func (h *user) HandleAuthUser(c *fiber.Ctx) error {
	helper.Ctx = c
	var user schemas.User
	usersCollection := mongodb.MI.Collection("users")
	params := new(schemas.AuthParams)

	if err := c.BodyParser(params); err != nil {
		return helper.BadRequest("BodyParser error", err)
	}
	err := usersCollection.FindOne(c.Context(), bson.M{"email": params.Email}).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		return helper.InternalServerError("Database error", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return helper.Unauthorized("Authentication error", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return helper.Unauthorized("JWT error", err)
	}
	res, err := usersCollection.UpdateOne(c.Context(), bson.M{"_id": user.ID}, bson.M{"$set": bson.M{"accessToken": tokenStr}})

	return helper.Ok("User authenticated successfully", res)
}

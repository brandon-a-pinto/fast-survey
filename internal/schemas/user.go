package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name" validate:"required"`
	Email    string             `json:"email" bson:"email" validate:"required"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	IsAdmin  bool               `bson:"isAdmin" json:"isAdmin" validate:"required"`
}

type CreateUserParams struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

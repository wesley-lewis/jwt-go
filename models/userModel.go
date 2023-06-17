package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	firstName    *string            `json:"first_name" validate:"required, min=2, max=100`
	lastName     *string            `json:"last_name" validate:"required, min=2, max=100"`
	password     *string            `json:"password" validate:"required, min=10, max=100"`
	email        *string            `json:"email" validate:"email, required"`
	phone        *string            `json:"phone" validate:"required"`
	token        *string            `json:"token"`
	user_type    *string            `json:"user_type" validate:"required, eq=ADMIN|wq=USER"`
	refreshToken *string            `json:"refresh_token"`
	createdAt    time.Time          `json:"created_at"`
	updatedAt    time.Time          `json:"updated_at"`
	userID       *string            `json:"user_id"`
}

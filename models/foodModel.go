package models

import (
	"go.mongodb.org/mongo-drive/bson/primitive"
)

type Food struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  *string            `json:"name" validate:"required,min=2,max=100"`
	Price *float64
}

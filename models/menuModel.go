package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Price      *float64           `json:"price" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Start_Date *time.time         `json:"start_date"`
	End_Date   *time.time         `json:"end_date"`
	Create_at  time.time          `json:"create_at"`
	Update_at  time.time          `json:"update_at"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}

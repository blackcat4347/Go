package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id"`
	Order_id         string             `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string            `json:"Payment_status" validate:"required.eq=PENDING|eq=PAID"`
	Payment_due_date time.time          `json:"Payment_due_date"`
	Create_at        time.time          `json:"create_at"`
	Update_at        time.time          `json:"update_at"`
}

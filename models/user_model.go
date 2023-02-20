package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	//Name     string             `json:"name,omitempty" validate:"required"`
	//Price    int    `json:"price,omitempty" validate:"required"`
	//Quantity string `json:"quantity,omitempty" validate:"required"`
	//UserID string `json:"userid,omitempty" validate:"required"`
}
type Products struct {
	Idcart  string `json:"idcart,omitempty"`
	Product []struct {
		Name     string `json:"name,omitempty" validate:"required"`
		Price    int    `json:"price,omitempty" validate:"required"`
		Quantity int    `json:"quantity,omitempty" validate:"required"`
	} `json:"product,omitempty" validate:"required"`
}

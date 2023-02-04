package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Class struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string               `json:"title" bson:"title"`
	Sks              int32                `json:"sks" bson:"sks"`
	Participants     []primitive.ObjectID `json:"participants" bson:"participants"`
}

func NewClass(title string, sks int32) *Class {
	return &Class{
		Title:        title,
		Sks:          sks,
		Participants: []primitive.ObjectID{},
	}
}

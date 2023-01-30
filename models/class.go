package models

import (
	"github.com/kamva/mgm/v3"
)

type Class struct {
	mgm.DefaultModel `bson:",inline"`
	Title       		string 						`json:"title" bson:"title"`
	Sks					int32    					`json:"sks" bson:"sks"`
	Participants		map[string]bool				`json:"participants" bson:"participants"`
}

func NewClass(title string, sks int32) *Class {
	return &Class{
		Title: title,
		Sks: 3,
		Participants: map[string]bool{},
	}
}
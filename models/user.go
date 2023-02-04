package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Username         string               `json:"username" bson:"username"`
	Email            string               `json:"email" bson:"email"`
	Password         string               `json:"password" bson:"password"`
	Groups           string               `json:"groups" bson:"groups"`
	FirstName        string               `json:"first_name" bson:"first_name"`
	LastName         string               `json:"last_name" bson:"last_name"`
	ClassesEnrolled  []primitive.ObjectID `json:"classes_enrolled" bson:"classes_enrolled"`
	RemSks           int32                `json:"rem_sks" bson:"rem_sks"`
}

type UserDetails struct {
	Username        string               `json:"username,omitempty"`
	Email           string               `json:"email,omitempty"`
	FirstName       string               `json:"first_name,omitempty"`
	LastName        string               `json:"last_name,omitempty"`
	ClassesEnrolled []primitive.ObjectID `json:"classes_enrolled,omitempty"`
	RemSks          int32                `json:"rem_sks,omitempty"`
}

func NewUser(username string, email string, password string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		RemSks:   24,
	}
}

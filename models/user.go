package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Username         		string 							`json:"username" bson:"username"`
	Email            		string 							`json:"email" bson:"email"`
	Password         		string 							`json:"password" bson:"password"`
	Groups					string 							`json:"groups" bson:"groups"`
	First_Name		 		string 							`json:"first_name" bson:"first_name"`
	Last_Name		 		string 							`json:"last_name" bson:"last_name"`
	Classes_Enrolled 		map[string]int32				`json:"classes_enrolled" bson:"classes_enrolled"`
	Rem_sks					int32							`json:"rem_sks" bson:"rem_sks"`
}

type UserDetails struct {
	Username       	 		string 							`json:"username,omitempty"`
	Email     		 		string 							`json:"email,omitempty"`
	First_Name		 		string 							`json:"first_name,omitempty"`
	Last_Name		 		string 							`json:"last_name,omitempty"`
	Classes_Enrolled 		map[string]int32				`json:"classes_enrolled,omitempty"`
	Rem_sks					int32							`json:"rem_sks,omitempty"`
}

func NewUser(username string, email string, password string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Rem_sks: 24,
	}
}
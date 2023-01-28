package models

type UsersResponse struct {
	ID int `json:"id"` 
	Nim string `json:"nim"` 
	Email string `json:"email"`
	Username string `json:"username"`
}

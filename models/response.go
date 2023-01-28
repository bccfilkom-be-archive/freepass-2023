package models

type CoursesOffersResponse struct {
	ID int `json:"id"` 
	Title string `json:"title"`
	Description string `json:"description"`
	InstructorName string `json:"instructorName"`
}

type AdminResponse struct {
	ID int `json:"id"` 
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type InstructorsResponse struct {
	ID int `json:"id"` 
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersResponse struct {
	ID int `json:"id"` 
	Nim string `json:"nim"` 
	Email string `json:"email"`
	Username string `json:"username"`
}

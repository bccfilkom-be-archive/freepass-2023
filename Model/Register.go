package Model

type Register struct {
	NIM       string `json:"nim"`
	StudentID uint   `json:"student_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
}

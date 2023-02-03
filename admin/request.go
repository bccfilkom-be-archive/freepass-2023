package admin

type AdminRequest struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	Position string `json:"position"`
	Password string `json:"password"`
}

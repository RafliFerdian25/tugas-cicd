package dto

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string
}
package dto

type RegisterDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

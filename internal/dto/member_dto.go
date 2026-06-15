package dto

type MemberRequestDTO struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Saldo     int    `json:"saldo"`
}

type MemberResponseDTO struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Saldo     int    `json:"saldo"`
}
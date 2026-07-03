package auth

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}
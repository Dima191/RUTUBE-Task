package models

type SignUp struct {
	Employee
	Password string `json:"password" validate:"required,min=8,max=32"`
}

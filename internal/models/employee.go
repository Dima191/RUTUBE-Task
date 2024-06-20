package models

type Employee struct {
	ID        uint32     `json:"id"`
	FullName  string     `json:"fullName"`
	BirthDate CustomDate `json:"birthDate"`
	Email     string     `json:"email" validate:"email"`
}

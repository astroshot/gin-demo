package model

type User struct {
	Name        *string
	Phone       *string
	Email       *string
	Description *string
	Status      *int
	BaseModel
}

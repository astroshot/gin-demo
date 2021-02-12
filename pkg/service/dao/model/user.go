package model

type User struct {
	Name        *string `json:"name"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	Description *string `json:"description"`
	Status      *int    `json:"status"`
	BaseModel
}

package model

import "encoding/json"

type User struct {
	Name        *string `json:"name"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	Description *string `json:"description"`
	Status      *int    `json:"status"`
	BaseModel
}

// TableName Defines name in database
func (instance User) TableName() string {
	return "user"
}

func (instance User) String() string {
	bytes, err := json.Marshal(instance)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

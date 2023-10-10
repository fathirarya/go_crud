package models

import (
	"time"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name" `
	Email     string    `json:"email"  gorm:"unique"`
	Password  string    `json:"password" `
}

func (updateUser User) ResponseConvertUser() UserResponse {
	Response := UserResponse{}
	Response.ID = updateUser.ID
	Response.Name = updateUser.Name
	Response.Email = updateUser.Email
	Response.Password = updateUser.Password

	return Response
}

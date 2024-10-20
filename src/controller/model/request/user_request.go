package request

import "time"

type UserRequest struct {
	Name      string    `json:"name" binding:"required,min=5,max=100"`
	Nick      string    `json:"nick" binding:"required,min=5,max=50"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=5,containsany=!@#$%&*"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package response

import "time"

type UserResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

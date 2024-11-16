package response

import "time"

type CategoryResponse struct {
	ID           uint64    `json:"id"`
	UserID       uint64    `json:"user_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

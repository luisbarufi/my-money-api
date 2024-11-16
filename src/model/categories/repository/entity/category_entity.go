package entity

import "time"

type CategoryEntity struct {
	ID           uint64    `json:"id,omitempty"`
	UserID       uint64    `json:"user_id,omitempty"`
	CategoryName string    `json:"category_name,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

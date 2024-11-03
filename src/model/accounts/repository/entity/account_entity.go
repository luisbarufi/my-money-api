package entity

import "time"

type AccountEntity struct {
	ID          uint64    `json:"id,omitempty"`
	UserID      uint64    `json:"user_id,omitempty"`
	AccountName string    `json:"account_name,omitempty"`
	Balance     float64   `json:"balance,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

package entity

import "time"

type TransactionEntity struct {
	ID              uint64    `json:"id,omitempty"`
	UserID          uint64    `json:"user_id,omitempty"`
	AccountID       uint64    `json:"account_id,omitempty"`
	CategoryID      uint64    `json:"category_id,omitempty"`
	Amount          float64   `json:"balance,omitempty"`
	TransactionType string    `json:"transaction_type,omitempty"`
	Description     string    `json:"description,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

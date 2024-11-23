package response

import "time"

type TransactionResponse struct {
	ID              uint64    `json:"id"`
	UserID          uint64    `json:"user_id"`
	AccountID       uint64    `json:"account_id"`
	CategoryID      uint64    `json:"category_id"`
	Amount          float64   `json:"balance"`
	TransactionType string    `json:"transaction_type" binding:"required,min=3,max=50"`
	Description     string    `json:"description" binding:",min=3,max=50"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

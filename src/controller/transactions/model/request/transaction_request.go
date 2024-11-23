package request

type TransactionRequest struct {
	AccountID       uint64  `json:"account_id"`
	CategoryID      uint64  `json:"category_id"`
	Amount          float64 `json:"balance"`
	TransactionType string  `json:"transaction_type" binding:"required,min=3,max=50"`
	Description     string  `json:"description" binding:"min=3,max=50"`
}

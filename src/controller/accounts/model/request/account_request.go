package request

type AccountRequest struct {
	AccountName string  `json:"account_name" binding:"required,min=3,max=50"`
	Balance     float64 `json:"balance"`
}

type UpdateAccountRequest struct {
	AccountName string `json:"account_name" binding:"required,min=3,max=50"`
}

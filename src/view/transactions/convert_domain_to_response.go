package view

import (
	"github.com/luisbarufi/my-money-api/src/controller/transactions/model/response"
	model "github.com/luisbarufi/my-money-api/src/model/transactions"
)

func ConvertDomainToResponse(
	transactionDomain model.TransactionDomainInterface,
) response.TransactionResponse {
	return response.TransactionResponse{
		ID:              transactionDomain.GetID(),
		UserID:          transactionDomain.GetUserID(),
		AccountID:       transactionDomain.GetAccountID(),
		CategoryID:      transactionDomain.GetCategoryID(),
		Amount:          transactionDomain.GetAmount(),
		Description:     transactionDomain.GetDescription(),
		TransactionType: transactionDomain.GetTransactionType(),
		CreatedAt:       transactionDomain.GetCreatedAt(),
		UpdatedAt:       transactionDomain.GetUpdatedAt(),
	}
}

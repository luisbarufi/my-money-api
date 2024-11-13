package view

import (
	response "github.com/luisbarufi/my-money-api/src/controller/accounts/model/response"
	model "github.com/luisbarufi/my-money-api/src/model/accounts"
)

func ConvertDomainToResponse(
	accountDomain model.AccountDomainInterface,
) response.AccountResponse {
	return response.AccountResponse{
		ID:          accountDomain.GetID(),
		UserID:      accountDomain.GetUserID(),
		AccountName: accountDomain.GetAccountName(),
		Balance:     accountDomain.GetBalance(),
		CreatedAt:   accountDomain.GetCreatedAt(),
		UpdatedAt:   accountDomain.GetUpdatedAt(),
	}
}

func ConvertDomainsToResponses(
	accountDomains []model.AccountDomainInterface,
) []response.AccountResponse {
	var responses []response.AccountResponse
	for _, accountDomain := range accountDomains {
		responses = append(responses, response.AccountResponse{
			ID:          accountDomain.GetID(),
			UserID:      accountDomain.GetUserID(),
			AccountName: accountDomain.GetAccountName(),
			Balance:     accountDomain.GetBalance(),
			CreatedAt:   accountDomain.GetCreatedAt(),
			UpdatedAt:   accountDomain.GetUpdatedAt(),
		})
	}
	return responses
}

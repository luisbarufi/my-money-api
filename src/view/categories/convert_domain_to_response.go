package view

import (
	response "github.com/luisbarufi/my-money-api/src/controller/categories/model/response"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
)

func ConvertDomainToResponse(
	categoryDomain model.CategoryDomainInterface,
) response.CategoryResponse {
	return response.CategoryResponse{
		ID:           categoryDomain.GetID(),
		UserID:       categoryDomain.GetUserID(),
		CategoryName: categoryDomain.GetCategoryName(),
		CreatedAt:    categoryDomain.GetCreatedAt(),
		UpdatedAt:    categoryDomain.GetUpdatedAt(),
	}
}

func ConvertDomainsToResponses(
	categoriesDomain []model.CategoryDomainInterface,
) []response.CategoryResponse {
	var responses []response.CategoryResponse
	for _, categoryDomain := range categoriesDomain {
		responses = append(responses, response.CategoryResponse{
			ID:           categoryDomain.GetID(),
			UserID:       categoryDomain.GetUserID(),
			CategoryName: categoryDomain.GetCategoryName(),
			CreatedAt:    categoryDomain.GetCreatedAt(),
			UpdatedAt:    categoryDomain.GetUpdatedAt(),
		})
	}
	return responses
}
